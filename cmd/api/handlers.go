package main

import (
	"logger/data"
	"net/http"
)

type jsonPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	var requestPayload jsonPayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}

	err = app.Models.LogEntry.Insert(event)

	if err != nil {
		app.errorJson(w, err, http.StatusInternalServerError)
		return
	}

	rsp := jsonResponse{
		Error:   false,
		Message: "logged",
	}

	app.writeJSON(w, http.StatusAccepted, rsp)

}
