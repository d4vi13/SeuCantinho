package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "github.com/d4vi13/SeuCantinho/internal/models/users"
	service "github.com/d4vi13/SeuCantinho/internal/services/users"
)

type CreateRequestUser struct {
	Username string `json:"username"`
	Passhash string `json:"password"`
	IsAdmin  bool   `json:"isAdmin"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var userReq CreateRequestUser
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	user, ret := service.CreateUser(userReq.Username, userReq.Passhash, userReq.IsAdmin)

	switch ret {
	case 0:
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "New User: %+v\n", user)
	case 1:
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "User already exists")
	}
}
