package avi

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/chertokdmitry/surfweatherbot/src/message"
	"gitlab.com/chertokdmitry/surfweatherbot/src/tg"
	"gitlab.com/chertokdmitry/surfweatherbot/src/utils/logger"
)

func SendAvi(id string) {
	url := message.AviUrl + id + ".mov"
	msgFile := tgbotapi.NewVideoUpload(tg.GetChatId(), url)

	if _, err := tg.GetBot().Send(msgFile); err != nil {
		logger.Error(message.ErrSendInline, err)
	}
}
