package internal

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramBotTask/storage"
)

type MessagesHandler struct {
	messagesDB storage.MessagesDB
}

func InitMessageHandler(messagesDB storage.MessagesDB) MessagesHandler {
	res := MessagesHandler{messagesDB: messagesDB}
	return res
}

func (h *MessagesHandler) DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})

	//@ToDo handle error
	if err != nil {
		return
	}
}

func (h *MessagesHandler) MyStartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Welcome to Goland Projects Bot!",
	})
	//@ToDo handle error
	if err != nil {
		return
	}
}
