package space

import (
	"fmt"

	models "github.com/d4vi13/SeuCantinho/internal/models/space"
	"github.com/d4vi13/SeuCantinho/internal/repository/space"
	"github.com/d4vi13/SeuCantinho/internal/services/users"
)

const (
	SpaceCreated = iota
	SpaceFound
	SpaceNotFound
	UserNotFound
	InvalidAdmin
	InternalError
)

type SpaceService struct {
	spaceRepository space.SpaceRepository
	userService     users.UsersService
}

func (service *SpaceService) Init() {
	service.spaceRepository.Init()
}

func (service *SpaceService) CreateSpace(username, password, location, substation string, price float64, capacity int, img []byte) (*models.Space, int) {

	// Verifica se o usuário existe
	var ret int = service.userService.AuthenticateUser(username, password)
	if ret == users.UserNotFound {
		fmt.Printf("SpaceService: User Not Found\n")
		return nil, UserNotFound
	}

	if ret == users.WrongPassword {
		fmt.Printf("SpaceService: Wrong Password\n")
		return nil, InvalidAdmin
	}

	// Verifica se o usuário é um administrador
	var adm bool = service.userService.UserIsAdmin(username)
	if !adm {
		fmt.Printf("SpaceService: User isn't an Admin\n")
		return nil, InvalidAdmin
	}

	// Verifica se o espaço já exsite
	space, _ := service.spaceRepository.GetSpace(location, substation)
	if space != nil {
		fmt.Printf("SpaceService: Space already exists\n")
		return space, SpaceFound
	}

	space = &models.Space{
		Location:   location,
		Substation: substation,
		Price:      price,
		Capacity:   capacity,
		Img:        img,
	}

	// Cria o novo espaço
	id, err := service.spaceRepository.Insert(space)

	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, InternalError
	}

	// Retorna o modelo do novo espaço
	space.Id = id

	fmt.Printf("SpaceService: Space created\n")
	return space, SpaceCreated
}
