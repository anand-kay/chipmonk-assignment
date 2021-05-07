package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/anand-kay/chipmonk-assignment/libs"
	"github.com/anand-kay/chipmonk-assignment/models"
)

func LoginHandler(w http.ResponseWriter, req *http.Request) {
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

	if !user.CheckUserExists() {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	storedHash := user.RetreiveStoredHash()

	isPwdMatch := libs.CheckHash(user.Password, storedHash)
	if !isPwdMatch {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Incorrect password"))
		return
	}

	sessionToken := user.CreateNewSession()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(sessionToken))
}
