package user

import (
	"encoding/json"
	"golayout/internal/api/helper"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
)

type user struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var u user
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	claims := make(map[string]interface{})

	jwtauth.SetExpiry(claims, time.Now().Add(time.Hour*24))
	//#TODO: add user_id to claims
	claims["user_id"] = 0

	auth := helper.GetJWTAuth()
	_, token, err := auth.Encode(claims)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(token))
	w.WriteHeader(http.StatusOK)
}
