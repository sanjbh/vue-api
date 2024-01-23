package main

import (
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	type credentials struct {
		UserName string `json:"email"`
		Password string `json:"password"`
	}

	var creds credentials
	var payload jsonResponse

	err := app.readJson(w, r, &creds)
	if err != nil {
		app.errorLog.Printf("Invalid json: %v\n", err)
		payload.Error = true
		payload.Message = "Invalid JSON"

		err := app.writeJson(w, http.StatusBadRequest, payload)
		if err != nil {
			app.errorLog.Println(err)
		}
		return
	}

	//authenticate
	app.infoLog.Printf("Username: %s, Password: %s\n", creds.UserName, creds.Password)
	//send back a response
	payload.Error = false
	payload.Message = "Signed In"

	err = app.writeJson(w, http.StatusOK, payload)

	if err != nil {
		app.errorLog.Println(err)
	}
}
