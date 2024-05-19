package handlers

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"strings"
	"telegram_bot/storage"
)

type MessagesHandler struct {
	messagesDB storage.MessagesDB
	filterWord string
}

// InitMessageHandler initializes a MessagesHandler with the given messagesDB and sets the filterWord to an empty string.
func InitMessageHandler(messagesDB storage.MessagesDB) MessagesHandler {
	res := MessagesHandler{messagesDB: messagesDB, filterWord: ""}
	return res
}

// DefaultHandler handles the incoming message by checking if it contains the filter word and saving it accordingly in the database.
func (h *MessagesHandler) DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	resp := "Message is saved successfully"
	var err error

	// Checking if the message contains the filter word
	if (strings.Contains(update.Message.Text, h.filterWord)) && (h.filterWord != "") {
		// If there is a filter word, save the message to the table with filtered messages
		err := h.messagesDB.AddFilteredMessage(update.Message.Date, update.Message.Text, update.Message.From.ID, update.Message.ID, h.filterWord)

		if err != nil {
			errMessage := fmt.Errorf("couldn't save the message in database, err: %s", err)
			log.Println(errMessage)
			resp = "The message has matched the filter but was not saved as the bot had encountered an error."
		} else {
			resp += fmt.Sprintf("\n \nThe message has macthed the filter as it contains the word '%s'", h.filterWord)
		}
	} else {
		// If there is no filter word, save the message to the table with unfiltered messages
		err := h.messagesDB.AddNewMessage(update.Message.Date, update.Message.Text, update.Message.From.ID, update.Message.ID)

		if err != nil {
			errMessage := fmt.Errorf("couldn't save the message in database, err: %s", err)
			log.Println(errMessage)
			resp = "The message was not saved as the bot had encountered an error."
		}
	}

	// Sending response
	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:          update.Message.Chat.ID,
		Text:            resp,
		ReplyParameters: &models.ReplyParameters{ChatID: update.Message.Chat.ID, MessageID: update.Message.ID},
	})

	if err != nil {
		log.Println(err)
	}
}

// SetFilterHandler sets the filter word for message filtering..
func (h *MessagesHandler) SetFilterHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	// Extracting the filter word from the message text
	word, _ := strings.CutPrefix(update.Message.Text, "/filter")
	word = strings.TrimSpace(word)
	resp := ""

	// If filter word is not given, return an error. Otherwise, set filterWord to the word given
	if word == "" {
		resp = "*Error*\nNo filter word specified"
	} else {
		h.filterWord = word
		resp = fmt.Sprintf("All messages are now filtered by the word '%s'", h.filterWord)
	}

	// Sending success message if the word is set and error otherwise
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      resp,
		ParseMode: models.ParseModeMarkdown,
	})
	if err != nil {
		log.Println(err)
	}
}

func (h *MessagesHandler) MyStartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	text := "Hello\\! Welcome to Message Filtering Bot that saves and filters all messages \n" +
		"Initially, it filters by the word 'filter' \n\n" +
		"*Commands:* \n" +
		"'/filter _option_' sets the word, to be filtered by, to '_option_'"

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      text,
		ParseMode: models.ParseModeMarkdown,
	})
	if err != nil {
		log.Println(err)
	}
}
