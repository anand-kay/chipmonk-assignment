package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/anand-kay/chipmonk-assignment/libs"
	"github.com/anand-kay/chipmonk-assignment/models"
)

func ActiveUsersHandler(w http.ResponseWriter, req *http.Request) {
	sessionToken := req.URL.Query().Get("token")

	err := libs.ValidateSessionToken(sessionToken)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if !models.IsUserActive(sessionToken) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("User not authorized to make this request"))
		return
	}

	activeUsers := models.FetchActiveUsers()

	type response struct {
		ActiveUsers []string `json:"active-users"`
	}

	res, err := json.Marshal(&response{
		ActiveUsers: activeUsers,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
