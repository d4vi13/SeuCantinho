package users

import (
	svc "github.com/d4vi13/SeuCantinho/internal/services/users"
)

type CreateRequestUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UsersController struct {
	usersService svc.UsersService
}

func (controller *UsersController) Init() {
	controller.usersService.Init()
}
