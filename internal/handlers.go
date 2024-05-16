package internal

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"strings"
	"telegramBotTask/storage"
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
	err := h.messagesDB.AddNewMessage(update.Message.Date, update.Message.Text, update.Message.From.ID, update.Message.ID)

	if err != nil {
		errMessage := fmt.Errorf("couldn't save the message in database, err: %s", err)
		log.Println(errMessage)
		resp = "The message was not saved as the bot had encountered an error."
	}

	if strings.Contains(update.Message.Text, h.filterWord) {
		err := h.messagesDB.AddFilteredMessage(update.Message.Date, update.Message.Text, update.Message.From.ID, update.Message.ID, h.filterWord)

		if err != nil {
			errMessage := fmt.Errorf("couldn't save the message in database, err: %s", err)
			log.Println(errMessage)
			resp = "The message has matched the filter but was not saved as the bot had encountered an error."
		} else {
			resp += fmt.Sprintf("\n \nThe message has macthed the filter as it contains the word '%s'", h.filterWord)
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
	h.filterWord = strings.TrimLeft(update.Message.Text, "/filter ")
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   fmt.Sprintf("All messages are now filtered by the word '%s'", h.filterWord),
	})
}

func (h *MessagesHandler) MyStartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Welcome to Message Filtering Bot!",
	})
}
