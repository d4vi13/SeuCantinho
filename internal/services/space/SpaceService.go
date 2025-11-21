package space

import "github.com/d4vi13/SeuCantinho/internal/repository/space"

type SpaceService struct {
	spaceRepository space.SpaceRepository
}

func (service *SpaceService) Init() {
	service.spaceRepository.Init()
}
