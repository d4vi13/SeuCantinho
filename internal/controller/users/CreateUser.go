package users

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (controller *UsersController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userReq CreateRequestUser

	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	user, ret := controller.usersService.CreateUser(userReq.Username, userReq.Passhash, userReq.IsAdmin)

	switch ret {
	case 0:
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "New User: %+v\n", user)
	case 1:
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "User already exists")
	case 2:
		w.WriteHeader(http.StatusInternalServerError)
	}
}
