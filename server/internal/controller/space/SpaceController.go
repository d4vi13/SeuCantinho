package space

import (
	svc "github.com/d4vi13/SeuCantinho/server/internal/services/space"
)

type RequestSpace struct {
	Username   string  `json:"username"`
	Password   string  `json:"password"`
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
