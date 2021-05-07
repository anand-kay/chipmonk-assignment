package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/anand-kay/chipmonk-assignment/db"
	"github.com/anand-kay/chipmonk-assignment/libs"
	"github.com/anand-kay/chipmonk-assignment/models"
)

func RegisterHandler(w http.ResponseWriter, req *http.Request) {
	var user models.User

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.Unmarshal(reqBody, &user)

	err = libs.ValidateForm(user.Username, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if _, ok := db.Users[user.Username]; ok {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Username exists already"))
		return
	}

	hashedPwd, err := libs.HashPassword(user.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	user.Password = hashedPwd

	err = user.RegisterUser()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))
}
