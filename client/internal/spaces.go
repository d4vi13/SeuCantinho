package internal

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type RequestSpace struct {
	ID         int     `json:"Id"`
	Location   string  `json:"location"`
	Substation string  `json:"substation"`
	Price      float64 `json:"price"`
	Capacity   int     `json:"capacity"`
	PNGBytes   []byte  `json:"Img"`
}

func CreateSpace(username string, password string) {
	reader := bufio.NewReader(os.Stdin)
	var space RequestSpace

	fmt.Printf("Localização do Espaço: ")
	location, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler localização")
		return
	}

	fmt.Printf("Filial: ")
	substation, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler filial")
		return
	}

	fmt.Printf("Custo de reserva por dia: ")
	input, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Println("Erro ao ler entrada: ", err)
		return
	}

	input = strings.TrimSpace(input)
	price, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Erro ao converter entrada")
		return
	}

	fmt.Printf("Capacidade Máxima: ")
	input, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Println("Erro ao ler entrada: ", err)
		return
	}

	input = strings.TrimSpace(input)
	capacity, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Erro ao converter entrada")
		return
	}

	location = location[:len(location)-1]
	substation = substation[:len(substation)-1]

	payload := map[string]interface{}{
		"username":   username,
		"password":   password,
		"location":   location,
		"substation": substation,
		"price":      price,
		"capacity":   capacity,
	}

	jsonData, _ := json.Marshal(payload)
	resp, err := http.Post("http://server:8080/space", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Não foi possível obter a resposta")
			return
		}

		var data map[string]string
		if err := json.Unmarshal(body, &data); err != nil {
			fmt.Printf("Não foi possível obter a resposta")
			return
		}

		fmt.Println("Erro:", data["error"])

		return
	}

	if resp.StatusCode == http.StatusConflict {
		fmt.Printf("Já existe um espaço com essa localização nessa filial")
		return
	}

	if resp.StatusCode == http.StatusInternalServerError {
		fmt.Printf("Houve um erro desconhecido no servidor\n")
		return
	}

	if resp.StatusCode == http.StatusCreated {
		if err := json.NewDecoder(resp.Body).Decode(&space); err != nil {
			panic(err)
		}

		fmt.Println()
		fmt.Println("Espaço criado com sucesso!")
		fmt.Println("ID: ", space.ID)
		fmt.Println("Localização: ", space.Location)
		fmt.Println("Filial: ", space.Substation)
		fmt.Println("Preço (R$): ", space.Price)
		fmt.Println("Capacidade (Pessoas): ", space.Capacity)
		fmt.Println()

		return
	}

	fmt.Println("Erro desconhecido")
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
