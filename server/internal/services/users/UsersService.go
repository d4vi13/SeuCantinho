package users

import (
	"fmt"

	models "github.com/d4vi13/SeuCantinho/server/internal/models/users"
	"github.com/d4vi13/SeuCantinho/server/internal/repository/users"
	"golang.org/x/crypto/bcrypt"
)

const (
	UserCreated = iota
	UserFound
	UsersFound
	UserNotFound
	UsersNotFound
	UserAuthenticated
	InvalidAdmin
	InvalidDelete
	WrongPassword
	InternalError
	UserDeleted
	Logged
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

func (service *UsersService) GetUserById(userId int) (*models.User, int) {
	// Obtém o usuário atravês do Id

	users, err := service.usersRepository.GetUserById(userId)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return nil, UserNotFound
	}

	return users, UserFound
}

func (service *UsersService) AuthenticateUser(username string, password string) int {

	user, err := service.usersRepository.GetUserByName(username)
	if err != nil {
		return UserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PassHash), []byte(password))
	if err != nil {
		return WrongPassword
	}

	return UserAuthenticated
}

func (service *UsersService) UserIsAdmin(username string) bool {

	user, err := service.usersRepository.GetUserByName(username)
	if err != nil {
		return false
	}

	return user.IsAdmin
}

func (service *UsersService) IsThisAnAdminId(userId int) bool {

	user, err := service.usersRepository.GetUserById(userId)
	if err != nil {
		return false
	}

	return user.IsAdmin
}

func (service *UsersService) Login(username string, password string) (*models.User, int) {
	user := &models.User{}

	user, err := service.usersRepository.GetUserByName(username)
	if err != nil {
		fmt.Printf("UserService: User not found\n")
		return nil, UserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PassHash), []byte(password))
	if err != nil {
		fmt.Printf("UserService: Wrong password\n")
		return nil, WrongPassword
	}

	user.PassHash = "hashed_password"
	user.IsAdmin = service.UserIsAdmin(username)

	return user, Logged
}

func (service *UsersService) CreateUser(username string, password string) (*models.User, int) {

	// Verifica se o usuário já existe
	user, _ := service.usersRepository.GetUserByName(username)
	if user != nil {
		fmt.Printf("UserService: User Already Exists\n")
		return user, UserFound
	}

	// Gera hash da senha
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, InternalError
	}

	user = &models.User{
		Username: username,
		PassHash: string(hashedPass),
		IsAdmin:  false,
	}

	// Insere o novo usuário no banco de dados
	id, err := service.usersRepository.Insert(user)

	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, InternalError
	}

	// Retorna o modelo do novo usuário
	user.Id = id

	fmt.Printf("UserService: User created\n")
	return user, UserCreated
}

func (service *UsersService) GetAllUsers() ([]models.User, int) {
	users, err := service.usersRepository.GetAllUsers()
	if err != nil {
		fmt.Printf("%+v\n", err)
		return nil, InternalError
	}

	if len(users) == 0 {
		fmt.Printf("UsersService: Users not found\n")
		return nil, UsersNotFound
	}

	return users, UsersFound
}

func (service *UsersService) DeleteUser(userId int, username string, password string) int {
	// Verifica se o usuário existe
	var ret int = service.AuthenticateUser(username, password)
	if ret == UserNotFound {
		fmt.Printf("UsersService: User Not Found\n")
		return UserNotFound
	}

	if ret == WrongPassword {
		fmt.Printf("UsersService: Wrong Password\n")
		return WrongPassword
	}

	var adm bool = service.IsThisAnAdminId(userId)
	if adm {
		fmt.Printf("UsersService: It is not possible to delete an administrator user\n")
		return InvalidDelete
	}

	// Verifica se o usuário é um administrador
	adm = service.UserIsAdmin(username)
	if !adm {
		fmt.Printf("UsersService: User isn't an Admin\n")
		return InvalidAdmin
	}

	err := service.usersRepository.Delete(userId)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return InternalError
	}

	return UserDeleted

}
