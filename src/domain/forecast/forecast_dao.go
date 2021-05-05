package forecast

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gitlab.com/chertokdmitry/surfweatherbot/src/domain/weather"
	"gitlab.com/chertokdmitry/surfweatherbot/src/message"
	"gitlab.com/chertokdmitry/surfweatherbot/src/tg"
	"gitlab.com/chertokdmitry/surfweatherbot/src/utils/logger"
	"strconv"
	"time"
)

var currentHour, _ = strconv.Atoi(time.Now().Format("15"))
var currentData = time.Now().Format("02/01/06")

func Send(hours []weather.Hour, title string) {

	bot := tg.GetBot()
	chatId := tg.GetChatId()

	forecastMessage := message.Forecast24 + "\n" + title + " " + currentData
	for i, hour := range hours {
		if i%2 != 0 {
			continue
		}

		forecastMessage = forecastMessage + "\n" + getHour() + ":00 " + getIcon(hour.Wind_deg) + " " + getTemp(hour.Temp) + " " + getWind(hour.Wind_speed)

		if i == 24 {
			break
		}
	}

	msgInline := tgbotapi.NewMessage(chatId, forecastMessage)
	currentHour, _ = strconv.Atoi(time.Now().Format("15"))

	if _, err := bot.Send(msgInline); err != nil {
		logger.Error(message.ErrSendInline, err)
	}
}

func getTemp(temp float32) string {
	var sign string
	if temp > 1 {
		sign = " +"
	} else if temp > 0 && temp < 1 {
		sign = message.TripleSpace
	} else {
		sign = " "
	}

	return sign + strconv.Itoa(int(temp)) + "C "
}

func getWind(wind float32) string {
	var icon string
	switch {
	case wind > 10:
		icon = message.TripleStar
	case wind > 7:
		icon = message.DoubleStar
	case wind > 4:
		icon = message.Star
	default:
		icon = ""
	}

	return strconv.Itoa(int(wind)) + "m/s " + icon
}

func getIcon(degree int) string {
	var icon string
	switch {
	case degree < 23 || degree > 337:
		icon = message.Down
	case degree < 68 && degree > 22:
		icon = message.RightUp
	case degree < 113 && degree > 67:
		icon = message.Right
	case degree < 158 && degree > 112:
		icon = message.RightDown
	case degree < 203 && degree > 157:
		icon = message.Up
	case degree < 248 && degree > 202:
		icon = message.LeftDown
	case degree < 293 && degree > 247:
		icon = message.Left
	case degree < 338 && degree > 292:
		icon = message.LeftUp
	default:
		icon = ""
	}

	return icon
}

func getHour() string {
	zeroHour := ""
	if currentHour > 22 {
		currentHour = 0
	}

	if currentHour < 10 {
		zeroHour = "0"
	} else {
		zeroHour = ""
	}

	result := zeroHour + strconv.Itoa(currentHour)
	currentHour += 2

	return result
}
