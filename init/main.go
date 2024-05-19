package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-telegram/bot"
	_ "github.com/lib/pq"
	"log"
	"os"
	"os/signal"
	"regexp"
	"telegram_bot/handlers"
	"telegram_bot/init/internal"
	"telegram_bot/storage"
)

const (
	messagesTableNamespace         = "messages"
	filteredMessagesTableNamespace = "filtered_messages"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	// Taking all env variables
	host := internal.GetEnvStr("DATABASE_HOST")
	port := internal.GetEnvInt("DATABASE_PORT")
	user := internal.GetEnvStr("DATABASE_USER")
	password := internal.GetEnvStr("DATABASE_PASSWORD")
	dbname := internal.GetEnvStr("DATABASE_NAME")
	token := internal.GetEnvStr("TELEGRAM_BOT_TOKEN")

	// Connecting to a database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Creating storage
	messageStorage := storage.InitMessagesDB(db, messagesTableNamespace, filteredMessagesTableNamespace)

	// Creating a handler
	messageHandler := handlers.InitMessageHandler(messageStorage)

	// Creating a bot
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(messageHandler.DefaultHandler),
	}

	b, err := bot.New(token, opts...)
	if err != nil {
		panic(err)
	}

	// Registering handlers and starting the bot
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, messageHandler.MyStartHandler)

	re := regexp.MustCompile(`^/filter`)
	b.RegisterHandlerRegexp(bot.HandlerTypeMessageText, re, messageHandler.SetFilterHandler)

	b.Start(ctx)
}
