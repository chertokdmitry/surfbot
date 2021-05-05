package tg

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/chertokdmitry/surfweatherbot/src/env"
	"sync"
)

type MyBot struct {
	bot *tgbotapi.BotAPI
}

var singleton *MyBot
var doOnce sync.Once

func GetMyBot() *MyBot {
	doOnce.Do(func() {
		bot, err := tgbotapi.NewBotAPI(env.TOKEN)
		if err != nil {
			fmt.Println(err)
		}
		singleton = &MyBot{bot: bot}
	})
	return singleton
}

func (mb *MyBot) Get() *tgbotapi.BotAPI {
	return mb.bot
}

func GetBot() *tgbotapi.BotAPI {
	myBotStruct := GetMyBot()
	bot := myBotStruct.Get()
	return bot
}

func GetChatId() int64 {
	update := GetUpdate()
	if update.CallbackQuery != nil {
		return update.CallbackQuery.Message.Chat.ID
	} else {
		return update.Message.Chat.ID
	}
}
