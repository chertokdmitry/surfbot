package routers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/chertokdmitry/surfweatherbot/src/controllers"
	"gitlab.com/chertokdmitry/surfweatherbot/src/domain/spots"
	"gitlab.com/chertokdmitry/surfweatherbot/src/message"
)

type KeyboardRouter struct {
	Update tgbotapi.Update
}

func (r *KeyboardRouter) Route() {
	switch r.Update.Message.Text {
	case message.Russia:
		message.Send(message.Get(r.Update.Message.Text))
		regionId := int64(message.RussiaRegionId)
		controllers.GetSpot(regionId)

	case message.Worldwide:
		message.Send(message.Get(r.Update.Message.Text))
		regionId := int64(message.WorldRegionId)
		controllers.GetSpot(regionId)

	case message.Subscribe:
		message.Send(message.Get(spots.GetSubcriptionList()))
		controllers.Subscribe()

	case message.SubList:
		message.Send(message.Get(r.Update.Message.Text))
		controllers.GetSubList()

	default:
		controllers.GetMainMenu()
	}
}
