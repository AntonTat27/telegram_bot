package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-telegram/bot"
	"log"
	"os"
	"os/signal"
	"telegramBotTask/internal"
)

func main() {
	// Taking all env variables
	host := getenvStr("DATABASE_HOST")
	port := getenvInt("DATABASE_PORT")
	user := getenvStr("DATABASE_USER")
	password := getenvStr("DATABASE_PASSWORD")
	dbname := getenvStr("DATABASE_NAME")
	token := getenvStr("TELEGRAM_BOT_TOKEN")

	// Connecting to a database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	// Creating a bot
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(internal.Handler),
	}

	b, err := bot.New(token, opts...)
	if err != nil {
		panic(err)
	}

	// Registering handlers and starting the bot
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, internal.MyStartHandler)

	b.Start(ctx)
}
