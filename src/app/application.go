package app

import (
	"gitlab.com/chertokdmitry/surfweatherbot/src/routers"
	"gitlab.com/chertokdmitry/surfweatherbot/src/tg"
)

func StartApplication() {
	for update := range tg.GetUpdatesChannel() {
		tg.SetUpdate(update)

		router := &routers.SubscriptionRouter{Update: update}
		router.Route()

		if update.Message == nil && update.CallbackQuery == nil {
			continue
		}

		Map()
	}
}

