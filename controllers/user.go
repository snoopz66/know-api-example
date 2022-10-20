package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/snoopz66/know-api-example/pkg/errors"

	"github.com/snoopz66/know-api-example/models"

	"github.com/gorilla/mux"

	"github.com/snoopz66/know-api-example/services"
)

type User struct {
	UserService *services.User
}

func (u *User) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, err := u.UserService.GetUser(vars["uuid"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(errors.Internal{
			Type:    "not found",
			Message: "user was not found",
			Code:    http.StatusNotFound,
			Err:     err,
		})
		return
	}
	_ = json.NewEncoder(w).Encode(user)
}

func (u *User) CreateUser(w http.ResponseWriter, r *http.Request) {
	userBody := &models.User{}
	err := json.NewDecoder(r.Body).Decode(userBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(errors.Internal{
			Type:    "bad request",
			Message: "could not decode json",
			Code:    http.StatusBadRequest,
			Err:     err,
		})
	}
	newUser, err := u.UserService.CreateUser(userBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(errors.Internal{
			Type:    "internal",
			Message: "could not persist user",
			Code:    http.StatusInternalServerError,
			Err:     err,
		})
	}
	_ = json.NewEncoder(w).Encode(newUser)
}
