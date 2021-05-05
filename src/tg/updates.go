package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"sync"
)

type Update struct {
	val tgbotapi.Update
}

var singletonUpdate *Update
var doOnceUpdate sync.Once

func GetMyUpdate() *Update {
	doOnceUpdate.Do(func() {
		singletonUpdate = &Update{}
	})
	return singletonUpdate
}

func (u *Update) GetVal() tgbotapi.Update {
	return u.val
}

func GetUpdate() tgbotapi.Update {
	myUpdateStruct := GetMyUpdate()
	update := myUpdateStruct.GetVal()
	return update
}

func SetUpdate(update tgbotapi.Update) {
	myUpdateStruct := GetMyUpdate()
	myUpdateStruct.val = update
}

func GetUpdatesChannel() tgbotapi.UpdatesChannel {
	bot := GetBot()
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	tgUpdates, _ := bot.GetUpdatesChan(u)

	return tgUpdates
}
