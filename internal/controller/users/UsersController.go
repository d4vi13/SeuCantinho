package users

import (
	svc "github.com/d4vi13/SeuCantinho/internal/services/users"
)

type CreateRequestUser struct {
	Username string `json:"username"`
	Passhash string `json:"password"`
	IsAdmin  bool   `json:"isAdmin"`
}

type UsersController struct {
	usersService svc.UsersService
}

func (controller *UsersController) Init() {
	controller.usersService.Init()
}
