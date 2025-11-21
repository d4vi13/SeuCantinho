package space

import (
	"github.com/d4vi13/SeuCantinho/internal/repository/space"
	"github.com/d4vi13/SeuCantinho/internal/services/users"
)

const (
	SpaceCreated = iota
	SpaceFound
	SpaceNotFound
	InvalidAdmin
)

type SpaceService struct {
	spaceRepository space.SpaceRepository
	userService     users.UsersService
}

func (service *SpaceService) Init() {
	service.spaceRepository.Init()
}

func (service *SpaceService) CreateSpace(username, passHash, location, substation string, price float64, capacity int, img []byte) int {

	_, ret := service.userService.AuthenticateUser(username, passHash)
	if ret != 0 {
		return InvalidAdmin
	}

	return SpaceCreated
}
