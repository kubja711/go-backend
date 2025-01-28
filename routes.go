package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /links", app.getLinks)
	router.HandleFunc("GET /temperature", app.getTemperature)
	router.HandleFunc("GET /meteo", app.getMeteo)

	return router
}
