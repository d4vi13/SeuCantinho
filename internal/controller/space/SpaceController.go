package space

import (
	svc "github.com/d4vi13/SeuCantinho/internal/services/space"
)

type CreateRequestSpace struct {
	Location   string  `json:"location"`
	Substation string  `json:"substation"`
	Price      float64 `json:"price"`
	Capacity   int     `json:"capacity"`
	PNGBytes   []byte  `json:"png"`
}

type SpaceController struct {
	spaceService svc.SpaceService
}

func (controller *SpaceController) Init() {
	controller.spaceService.Init()
}
