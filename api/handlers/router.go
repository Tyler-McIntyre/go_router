package handlers

import (
	"github.com/go-chi/chi"
	api "github.com/tyler-mcintyre/go_service_layer/api"
)

const URL_CURRENT_WEATHER = "/current-weather"

type router struct {
	Mux     chi.Router
	Service api.Service
}

func NewRouter(
	mux chi.Router,
	service api.Service) router {

	return router{
		Mux:     mux,
		Service: service,
	}
}

func (router *router) AddRoutes() {
	router.Mux.Group(func(r chi.Router) {
		r.Get(URL_CURRENT_WEATHER, router.GetCurrentWeatherHandler)
	})
}
