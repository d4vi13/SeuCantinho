package users

import (
	models "github.com/d4vi13/SeuCantinho/internal/models/users"
)

var globalID int = 0

type UsersService struct {
}

func (service *UsersService) Init() {}

func (service *UsersService) CreateUser(username string, passHash string, isAdmin bool) (models.User, int) {

	user := models.User{
		Id:       globalID,
		Username: username,
		PassHash: passHash,
		IsAdmin:  isAdmin,
	}
	globalID++

	return user, 0
}
