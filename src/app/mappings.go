package app

import (
	"gitlab.com/chertokdmitry/surfweatherbot/src/routers"
	"gitlab.com/chertokdmitry/surfweatherbot/src/tg"
)

// Inline or keyboard selection
func Map() {
	update := tg.GetUpdate()
	if update.CallbackQuery != nil {
		router := &routers.InlineRouter{Update: update}
		router.Route()
	} else {
		router := &routers.KeyboardRouter{Update: update}
		router.Route()
	}
}
