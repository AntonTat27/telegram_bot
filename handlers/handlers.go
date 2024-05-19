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

func InitMessageHandler(messagesDB storage.MessagesDB) MessagesHandler {
	res := MessagesHandler{messagesDB: messagesDB, filterWord: "filter"}
	return res
}

func (h *MessagesHandler) DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	resp := "Message is saved successfully"
	var err error

	if strings.Contains(update.Message.Text, h.filterWord) {
		err := h.messagesDB.AddFilteredMessage(update.Message.Date, update.Message.Text, update.Message.From.ID, update.Message.ID, h.filterWord)

		if err != nil {
			errMessage := fmt.Errorf("couldn't save the message in database, err: %s", err)
			log.Println(errMessage)
			resp = "The message has matched the filter but was not saved as the bot had encountered an error."
		} else {
			resp += fmt.Sprintf("\n \nThe message has macthed the filter as it contains the word '%s'", h.filterWord)
		}
	} else {
		err := h.messagesDB.AddNewMessage(update.Message.Date, update.Message.Text, update.Message.From.ID, update.Message.ID)

		if err != nil {
			errMessage := fmt.Errorf("couldn't save the message in database, err: %s", err)
			log.Println(errMessage)
			resp = "The message was not saved as the bot had encountered an error."
		}
	}

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:          update.Message.Chat.ID,
		Text:            resp,
		ReplyParameters: &models.ReplyParameters{ChatID: update.Message.Chat.ID, MessageID: update.Message.ID},
	})

	if err != nil {
		return
	}
}

func (h *MessagesHandler) SetFilterHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	word, _ := strings.CutPrefix(update.Message.Text, "/filter")
	h.filterWord = strings.Trim(word, " ")
	resp := fmt.Sprintf("All messages are now filtered by the word '_%s_'", h.filterWord)
	if h.filterWord == "" {
		resp = "*Error*\nNo filter word specified"
	}

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
