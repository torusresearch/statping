package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/torusresearch/statping/types/errors"
	"github.com/torusresearch/statping/types/users"
	"github.com/torusresearch/statping/utils"
	"net/http"
)

func findUser(r *http.Request) (*users.User, int64, error) {
	vars := mux.Vars(r)
	if utils.NotNumber(vars["id"]) {
		return nil, 0, errors.NotNumber
	}
	num := utils.ToInt(vars["id"])
	user, err := users.Find(num)
	if err != nil {
		return nil, num, errors.Missing(&users.User{}, num)
	}
	return user, num, nil
}

func apiUserHandler(w http.ResponseWriter, r *http.Request) {
	user, _, err := findUser(r)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	user.Password = ""
	returnJson(user, w, r)
}

func apiUserUpdateHandler(w http.ResponseWriter, r *http.Request) {
	user, _, err := findUser(r)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	err = DecodeJSON(r, &user)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	if user.Password != "" {
		user.Password = utils.HashPassword(user.Password)
	}

	err = user.Update()
	if err != nil {
		sendErrorJson(fmt.Errorf("issue updating user #%d: %s", user.Id, err), w, r)
		return
	}
	sendJsonAction(user, "update", w, r)
}

func apiUserDeleteHandler(w http.ResponseWriter, r *http.Request) {
	allUsers := users.All()
	if len(allUsers) == 1 {
		sendErrorJson(errors.New("cannot delete the last user"), w, r)
		return
	}
	user, _, err := findUser(r)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	if err := user.Delete(); err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(user, "delete", w, r)
}

func apiAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	allUsers := users.All()
	returnJson(allUsers, w, r)
}

func apiCreateUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user *users.User
	err := DecodeJSON(r, &user)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	err = user.Create()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(user, "create", w, r)
}
