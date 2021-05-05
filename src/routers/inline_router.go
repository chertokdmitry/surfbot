package routers

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/chertokdmitry/surfweatherbot/src/controllers"
	"gitlab.com/chertokdmitry/surfweatherbot/src/keyboard"
	"gitlab.com/chertokdmitry/surfweatherbot/src/message"
	"gitlab.com/chertokdmitry/surfweatherbot/src/utils/logger"
	"strconv"
)

type InlineRouter struct {
	Update tgbotapi.Update
}

// switch inline actions
func (r *InlineRouter) Route() {
	data := keyboard.Request{}

	if err := json.Unmarshal([]byte(r.Update.CallbackQuery.Data), &data); err != nil {
		logger.Error(message.ErrUnmarshal, err)
	}

	switch data.Action {
	case "avi":
		controllers.SendAvi(data.Id)

	case "camera":
		controllers.SendImage(data.Id)

	case "spot":
		spotId, _ := strconv.ParseInt(data.Id, 10, 64)
		controllers.GetWeather(spotId)

	case "hour":
		controllers.NewSubscription(data.Id)

	case "delete":
		subId, _ := strconv.ParseInt(data.Id, 10, 64)
		controllers.DelSubscription(subId)
	}
}
