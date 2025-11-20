package users

import (
	"github.com/d4vi13/SeuCantinho/internal/services/users"
)

type CreateRequestUser struct {
	Username string `json:"username"`
	Passhash string `json:"password"`
	IsAdmin  bool   `json:"isAdmin"`
}

type UsersController struct {
	usersService users.UsersService
}

func (controller *UsersController) Init() {
	controller.usersService.Init()
}
