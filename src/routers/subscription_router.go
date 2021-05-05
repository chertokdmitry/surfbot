package routers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/chertokdmitry/surfweatherbot/src/controllers"
	"gitlab.com/chertokdmitry/surfweatherbot/src/domain/spots"
	"gitlab.com/chertokdmitry/surfweatherbot/src/domain/subscription"
	"gitlab.com/chertokdmitry/surfweatherbot/src/message"
	"gitlab.com/chertokdmitry/surfweatherbot/src/utils/generics"
	"strconv"
)

type SubscriptionRouter struct {
	Update tgbotapi.Update
}

func (r *SubscriptionRouter) Route() {
	if subscription.IsSpotId == true && r.Update.Message != nil {
		subscription.SpotId, _ = strconv.Atoi(r.Update.Message.Text)
		_, found := generics.Find(spots.SpotIds, subscription.SpotId)
		if found {
			subscription.GetHour()
			controllers.GetHoursMenu()
		} else {
			message.Send(message.WrongSpotId)
		}

	}
}
