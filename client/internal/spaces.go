package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestSpace struct {
	ID         int     `json:"Id"`
	Location   string  `json:"location"`
	Substation string  `json:"substation"`
	Price      float64 `json:"price"`
	Capacity   int     `json:"capacity"`
	PNGBytes   []byte  `json:"Img"`
}

func GetAllSpaces() {
	var spaces []RequestSpace
	resp, err := http.Get("http://server:8080/space")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		fmt.Printf("Não existe nenhum espaço\n")
		return
	}

	if resp.StatusCode == http.StatusInternalServerError {
		fmt.Printf("Houve um erro interno no servidor\n")
		return
	}

	if err := json.NewDecoder(resp.Body).Decode(&spaces); err != nil {
		panic(err)
	}

	for _, s := range spaces {
		fmt.Println("========================")
		fmt.Println("ID: ", s.ID)
		fmt.Println("Localização: ", s.Location)
		fmt.Println("Filial: ", s.Substation)
		fmt.Println("Preço (R$): ", s.Price)
		fmt.Println("Capacidade (Pessoas): ", s.Capacity)
		fmt.Println("========================")

	}

}
