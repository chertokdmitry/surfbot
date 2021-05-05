package controllers

import (
	"gitlab.com/chertokdmitry/surfweatherbot/src/domain/avi"
	"gitlab.com/chertokdmitry/surfweatherbot/src/domain/cameras"
	"gitlab.com/chertokdmitry/surfweatherbot/src/domain/forecast"
	"gitlab.com/chertokdmitry/surfweatherbot/src/domain/subscription"
	"gitlab.com/chertokdmitry/surfweatherbot/src/domain/weather"
)

// show image from camera
func SendAvi(cameraId string) {
	avi.SendAvi(cameraId)
}

func SendImage(cameraId string) {
	cameras.SendImage(cameraId)
}

// show weather by spot
func GetWeather(spotId int64) {
	forecast.Send(weather.Get(spotId))
}

// insert new sub
func NewSubscription(hour string) {
	subscription.Insert(hour)
}

// delete sub by id
func DelSubscription(subId int64) {
	subscription.Delete(subId)
}
