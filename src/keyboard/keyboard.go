package keyboard

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/chertokdmitry/surfweatherbot/src/domain/cameras"
	"gitlab.com/chertokdmitry/surfweatherbot/src/domain/spots"
	"gitlab.com/chertokdmitry/surfweatherbot/src/domain/subscription"
	"gitlab.com/chertokdmitry/surfweatherbot/src/message"
	"gitlab.com/chertokdmitry/surfweatherbot/src/tg"
	"gitlab.com/chertokdmitry/surfweatherbot/src/utils/logger"
	"strconv"
)

type Request struct {
	Action string `json:"action"`
	Id     string `json:"id"`
}

var hours = []string{"06:00", "07:00", "08:00", "09:00", "10:00", "11:00", "12:00", "13:00", "14:00", "15:00", "16:00", "17:00", "18:00", "19:00", "20:00", "21:00", "22:00", "23:00"}

var MainPage = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(message.Russia),
		tgbotapi.NewKeyboardButton(message.Worldwide),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(message.Subscribe),
		tgbotapi.NewKeyboardButton(message.SubList),
	),
)

func GetHoursMenu() {
	var rows = [][]tgbotapi.InlineKeyboardButton{
		[]tgbotapi.InlineKeyboardButton{},
		[]tgbotapi.InlineKeyboardButton{},
		[]tgbotapi.InlineKeyboardButton{}}

	bot := tg.GetBot()
	msgInline := tgbotapi.NewMessage(tg.GetChatId(), message.GetHour)
	r := 0

	for i, hour := range hours {
		request := &Request{
			Action: "hour",
			Id:     hour,
		}
		var msg, _ = json.Marshal(request)
		row := tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(hour, string(msg)))

		if i%6 == 0 && i != 0 {
			r++
		}

		rows[r] = append(rows[r], row[0])
	}
	msgInline.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(rows...)
	if _, err := bot.Send(msgInline); err != nil {
		logger.Error(message.ErrSendInline, err)
	}
}

func GetMainMenu() {
	if subscription.IsSpotId == true || subscription.IsHour {
		return
	}

	update := tg.GetUpdate()
	bot := tg.GetBot()
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyMarkup = MainPage

	if _, err := bot.Send(msg); err != nil {
		logger.Error(message.ErrSendInline, err)
	}
}

func GetSpot(regionId int64) {
	var rows [][]tgbotapi.InlineKeyboardButton
	update := tg.GetUpdate()
	bot := tg.GetBot()
	msgInline := tgbotapi.NewMessage(tg.GetChatId(), update.Message.Text)

	for _, spot := range spots.Get(regionId) {
		request := &Request{
			Action: "spot",
			Id:     strconv.Itoa(spot.Id),
		}
		var msg, _ = json.Marshal(request)

		var cams []tgbotapi.InlineKeyboardButton
		spotButton := tgbotapi.NewInlineKeyboardButtonData(spot.Title, string(msg))
		cams = append(cams, spotButton)

		camerasIds := cameras.GetBySpot(int64(spot.Id))
		if len(camerasIds) > 0 {
			for _, camId := range camerasIds {
				cameraRequest := &Request{
					Action: "camera",
					Id:     strconv.FormatInt(camId, 10),
				}

				var cameraMsg, _ = json.Marshal(cameraRequest)
				cam := tgbotapi.NewInlineKeyboardButtonData(message.Frame, string(cameraMsg))
				cams = append(cams, cam)

				aviRequest := &Request{
					Action: "avi",
					Id:     strconv.FormatInt(camId, 10),
				}

				var aviMsg, _ = json.Marshal(aviRequest)
				avi := tgbotapi.NewInlineKeyboardButtonData(message.Camera, string(aviMsg))
				cams = append(cams, avi)
			}
		}

		row := tgbotapi.NewInlineKeyboardRow(cams...)
		rows = append(rows, row)
	}

	msgInline.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(rows...)
	if _, err := bot.Send(msgInline); err != nil {
		logger.Error(message.ErrSendInline, err)
	}
}

func GetSubs() {
	var rows [][]tgbotapi.InlineKeyboardButton
	update := tg.GetUpdate()
	bot := tg.GetBot()
	msgInline := tgbotapi.NewMessage(tg.GetChatId(), update.Message.Text)

	for _, sub := range subscription.List() {
		request := &Request{
			Action: "delete",
			Id:     strconv.Itoa(sub.Id),
		}
		var msg, _ = json.Marshal(request)
		title := sub.Title + " " + sub.Hour + message.Delete

		row := tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(title, string(msg)))

		rows = append(rows, row)
	}

	msgInline.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(rows...)
	if _, err := bot.Send(msgInline); err != nil {
		logger.Error(message.ErrSendInline, err)
	}
}

func BackCatalog() tgbotapi.InlineKeyboardMarkup {
	requestData := &Request{
		Action: "back_catalog",
		Id:     "",
	}

	msgData, _ := json.Marshal(requestData)

	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(message.Back, string(msgData)),
		),
	)
}
