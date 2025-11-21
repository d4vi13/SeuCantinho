package space

import (
	svc "github.com/d4vi13/SeuCantinho/internal/services/space"
)

type SpaceController struct {
	spaceService svc.SpaceService
}

func (controller *SpaceController) Init() {
	controller.spaceService.Init()
}
