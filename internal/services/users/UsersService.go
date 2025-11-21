package users

import (
	models "github.com/d4vi13/SeuCantinho/internal/models/users"
	"github.com/d4vi13/SeuCantinho/internal/repository/users"
)

const (
	UserCreated = iota
	UserFound
	UserNotFound
	WrongPassword
	InternalError
)

type UsersService struct {
	usersRepository users.UsersRepository
}

func (service *UsersService) Init() {
	service.usersRepository.Init()
}

func (service *UsersService) GetUserId(username string) int {

	user, err := service.usersRepository.GetUserByName(username)
	if err != nil {
		return -1
	}

	return user.Id
}

func (service *UsersService) AuthenticateUser(username string, passHash string) (bool, int) {

	user, err := service.usersRepository.GetUserByName(username)
	if err != nil {
		return false, UserNotFound
	}

	if user.PassHash != passHash {
		return false, WrongPassword
	}

	return true, 0
}

func (service *UsersService) UserIsAdmin(username string) bool {

	user, err := service.usersRepository.GetUserByName(username)
	if err != nil {
		return false
	}

	return user.IsAdmin
}

func (service *UsersService) CreateUser(username string, passHash string, isAdmin bool) (*models.User, int) {

	// Verifica se o usuário já existe
	user, _ := service.usersRepository.GetUserByName(username)
	if user != nil {
		return user, UserFound
	}

	user = &models.User{
		Username: username,
		PassHash: passHash,
		IsAdmin:  isAdmin,
	}

	// Insere o novo usuário no banco de dados
	id, err := service.usersRepository.Insert(user)

	if err != nil {
		return nil, InternalError
	}

	// Retorna o modelo do novo usuário
	return &models.User{
		Id:       id,
		Username: username,
		PassHash: passHash,
		IsAdmin:  isAdmin,
	}, UserCreated
}
