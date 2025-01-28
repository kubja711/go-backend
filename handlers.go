package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) getLinks(w http.ResponseWriter, r *http.Request) {
	// pathid := r.PathValue("id")
	// id, err := strconv.Atoi(pathid)
	// if err != nil {
	// 	id = 0
	// }
	link, err := app.links.Get()
	if err != nil {
		app.errorLog.Fatal(err)
	}
	json, err := json.MarshalIndent(link, "", "\t")
	if err != nil {
		app.errorLog.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(json)
}

func (app *application) getTemperature(w http.ResponseWriter, r *http.Request) {
	// pathid := r.PathValue("id")
	// id, err := strconv.Atoi(pathid)
	// if err != nil {
	// 	id = 0
	// }
	temperature, err := app.temperature.Get()
	if err != nil {
		app.errorLog.Fatal(err)
	}
	json, err := json.MarshalIndent(temperature, "", "\t")
	if err != nil {
		app.errorLog.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(json)
}

func (app *application) getMeteo(w http.ResponseWriter, r *http.Request) {
	// pathid := r.PathValue("id")
	// id, err := strconv.Atoi(pathid)
	// if err != nil {
	// 	id = 0
	// }
	temperature, err := app.meteo.Get()
	if err != nil {
		app.errorLog.Fatal(err)
	}
	json, err := json.MarshalIndent(temperature, "", "\t")
	if err != nil {
		app.errorLog.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(json)
}
