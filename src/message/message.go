package message

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/chertokdmitry/surfweatherbot/src/tg"
	"gitlab.com/chertokdmitry/surfweatherbot/src/utils/logger"
)

var botMessage string

func Get(updateText string) string {
	if botMessage != "" {
		answer := botMessage
		botMessage = ""

		return answer
	} else {
		return updateText
	}
}

func Set(message string) {
	botMessage = message
}

func Send(message string) {
	bot := tg.GetBot()
	msgMain := tgbotapi.NewMessage(tg.GetChatId(), message)

	if _, err := bot.Send(msgMain); err != nil {
		logger.Error("error when trying to send main response message", err)
	}
}
