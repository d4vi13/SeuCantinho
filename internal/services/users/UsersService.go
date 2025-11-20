package users

import (
	models "github.com/d4vi13/SeuCantinho/internal/models/users"
	"github.com/d4vi13/SeuCantinho/internal/repository/users"
)

type UsersService struct {
	usersRepository users.UsersRepository
}

func (service *UsersService) Init() {
	service.usersRepository.Init()
}

func (service *UsersService) CreateUser(username string, passHash string, isAdmin bool) (*models.User, int) {

	user, _ := service.usersRepository.GetUserByName(username)
	if user != nil {
		return user, 1
	}

	user = &models.User{
		Username: username,
		PassHash: passHash,
		IsAdmin:  isAdmin,
	}

	id, err := service.usersRepository.Insert(user)

	if err != nil {
		return nil, 2
	}

	return &models.User{
		Id:       id,
		Username: username,
		PassHash: passHash,
		IsAdmin:  isAdmin,
	}, 0
}
