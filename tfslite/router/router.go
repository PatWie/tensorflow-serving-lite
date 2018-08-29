// Patrick Wieschollek <mail@patwie.com>

package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/patwie/tensorflow-serving-lite/tfslite/router/api"
)

func GetRouter() *chi.Mux {

	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.NoCache)

	r.Get("/", api.HomeHandler)
	r.Post("/v{version}/models/{modelname}:{signature}", api.PredictHandler)

	return r
}
