package controllers

import (
	"gitlab.com/chertokdmitry/surfweatherbot/src/domain/subscription"
	"gitlab.com/chertokdmitry/surfweatherbot/src/keyboard"
)

func GetSpot(regionId int64) {
	keyboard.GetSpot(regionId)
}

func GetMainMenu() {
	keyboard.GetMainMenu()
}

func Subscribe() {
	subscription.VarsFlash()
	subscription.GetSpotId()
}

func GetSubList() {
	keyboard.GetSubs()
}

func GetHoursMenu() {
	keyboard.GetHoursMenu()
}
