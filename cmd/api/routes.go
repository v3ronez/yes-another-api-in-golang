package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/api/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/api/v1/movies", app.storeMovie)
	router.HandlerFunc(http.MethodGet, "/api/v1/movies/:id", app.showMovie)
	return router
}
