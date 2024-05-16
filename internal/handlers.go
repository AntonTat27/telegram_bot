package internal

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"telegramBotTask/storage"
)

type MessagesHandler struct {
	messagesDB storage.MessagesDB
	filterWord string
}

func InitMessageHandler(messagesDB storage.MessagesDB) MessagesHandler {
	res := MessagesHandler{messagesDB: messagesDB, filterWord: ""}
	return res
}

func (h *MessagesHandler) DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	resp := "Message is saved successfully"
	err := h.messagesDB.AddNewMessage(update.Message.Date, update.Message.Text, update.Message.From.ID, update.Message.ID)

	if err != nil {
		errMessage := fmt.Errorf("couldn't save the message in database, err: %s", err)
		log.Println(errMessage)
		resp = "The bot encountered the following error: \n" + errMessage.Error()
	}

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:          update.Message.Chat.ID,
		Text:            resp,
		ReplyParameters: &models.ReplyParameters{ChatID: update.Message.Chat.ID, MessageID: update.Message.ID},
	})
}

func (h *MessagesHandler) MyStartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Welcome to Goland Projects Bot!",
	})

	if err != nil {
		return
	}
}
