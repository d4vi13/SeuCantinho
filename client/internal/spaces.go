package internal

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type RequestSpace struct {
	ID         int    `json:"Id"`
	Location   string `json:"location"`
	Substation string `json:"substation"`
	Price      int64  `json:"price"`
	Capacity   int    `json:"capacity"`
}

type SpaceUpdateRequest struct {
	Username   *string `json:"username"`
	Password   *string `json:"password"`
	Location   *string `json:"location,omitempty"`
	Substation *string `json:"substation,omitempty"`
	Price      *int64  `json:"price,omitempty"`
	Capacity   *int    `json:"capacity,omitempty"`
}

func CreateSpace(username string, password string) {
	reader := bufio.NewReader(os.Stdin)
	var space RequestSpace

	fmt.Println("Criação de Espaço")
	fmt.Printf("Localização do Espaço: ")
	location, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler localização")
		return
	}
	location = strings.TrimSpace(location)

	fmt.Printf("Filial: ")
	substation, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler filial")
		return
	}
	substation = strings.TrimSpace(substation)

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
	var intPrice int64 = int64(math.Round(price * 100))

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

	payload := map[string]interface{}{
		"username":   username,
		"password":   password,
		"location":   location,
		"substation": substation,
		"price":      intPrice,
		"capacity":   capacity,
	}

	// Faz a requisição para o backend
	jsonData, _ := json.Marshal(payload)
	resp, err := http.Post("http://server:8080/space", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Trata valores de retorno
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
		fmt.Println("Preço (R$): ", (float64(space.Price) / 100))
		fmt.Println("Capacidade (Pessoas): ", space.Capacity)
		fmt.Println()

		return
	}

	fmt.Println("Erro desconhecido")
}

func UpdateSpace(username string, password string) {
	req := SpaceUpdateRequest{}
	space := &RequestSpace{}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Atualização de Espaço")
	fmt.Printf("ID do Espaço: ")
	input, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Println("Erro ao ler entrada: ", err)
		return
	}

	input = strings.TrimSpace(input)
	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Erro ao converter entrada")
		return
	}

	req.Username = &username
	req.Password = &password

	fmt.Printf("Para atualizar o campo, insira o novo valor, caso contrário, aperte 'Enter'\n")

	fmt.Print("Nova Localização: ")
	location, _ := reader.ReadString('\n')
	location = strings.TrimSpace(location)
	if location != "" {
		req.Location = &location
	}

	fmt.Print("Nova Filial: ")
	substation, _ := reader.ReadString('\n')
	substation = strings.TrimSpace(substation)
	if substation != "" {
		req.Substation = &substation
	}

	fmt.Print("Novo Custo: ")
	priceInput, _ := reader.ReadString('\n')
	priceInput = strings.TrimSpace(priceInput)
	if priceInput != "" {
		price, err := strconv.ParseFloat(priceInput, 64)
		if err != nil {
			fmt.Println("Preço inválido")
			return
		}
		var intPrice int64 = int64(math.Round(price * 100))
		req.Price = &intPrice
	}

	fmt.Print("Nova Capacidade: ")
	capInput, _ := reader.ReadString('\n')
	capInput = strings.TrimSpace(capInput)
	if capInput != "" {
		capacity, err := strconv.Atoi(capInput)
		if err != nil {
			fmt.Println("Capacidade inválida")
			return
		}
		req.Capacity = &capacity
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("http://server:8080/space/%d", id)

	requisition, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	// Faz a requisição para o backend
	client := &http.Client{}
	resp, err := client.Do(requisition)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Trata valores de retorno
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

	if resp.StatusCode == http.StatusInternalServerError {
		fmt.Printf("Houve um erro interno no servidor\n")
		return
	}

	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(space); err != nil {
			panic(err)
		}

		fmt.Println()
		fmt.Println("Espaço atualizado com sucesso!")
		fmt.Println("ID: ", space.ID)
		fmt.Println("Localização: ", space.Location)
		fmt.Println("Filial: ", space.Substation)
		fmt.Println("Preço (R$): ", (float64(space.Price) / 100))
		fmt.Println("Capacidade (Pessoas): ", space.Capacity)
		fmt.Println()

		return
	}

	fmt.Println("Erro desconhecido")
}

func GetSpace() {
	var space RequestSpace
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("ID do Espaço: ")
	input, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Println("Erro ao ler entrada: ", err)
		return
	}

	input = strings.TrimSpace(input)
	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Erro ao converter entrada")
		return
	}

	url := fmt.Sprintf("http://server:8080/space/%d", id)

	// Faz a requisição ao backend
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Trata valores de retorno
	if resp.StatusCode == http.StatusNotFound {
		fmt.Printf("Esse espaço não existe\n")
		return
	}

	if resp.StatusCode == http.StatusInternalServerError {
		fmt.Printf("Houve um erro interno no servidor\n")
		return
	}

	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(&space); err != nil {
			panic(err)
		}

		fmt.Println()
		fmt.Println("========================")
		fmt.Println("ID: ", space.ID)
		fmt.Println("Localização: ", space.Location)
		fmt.Println("Filial: ", space.Substation)
		fmt.Println("Preço (R$): ", (float64(space.Price) / 100))
		fmt.Println("Capacidade (Pessoas): ", space.Capacity)
		fmt.Println("========================")
		fmt.Println()

		return
	}

	fmt.Println("Erro desconhecido")
}

func GetAllSpaces() {
	var spaces []RequestSpace

	// Faz a requisição ao backend
	resp, err := http.Get("http://server:8080/space")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Trata valores de retorno
	if resp.StatusCode == http.StatusNotFound {
		fmt.Printf("Não existe nenhum espaço\n")
		return
	}

	if resp.StatusCode == http.StatusInternalServerError {
		fmt.Printf("Houve um erro interno no servidor\n")
		return
	}

	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(&spaces); err != nil {
			panic(err)
		}

		for _, s := range spaces {
			fmt.Println("========================")
			fmt.Println("ID: ", s.ID)
			fmt.Println("Localização: ", s.Location)
			fmt.Println("Filial: ", s.Substation)
			fmt.Println("Preço (R$): ", (float64(s.Price) / 100))
			fmt.Println("Capacidade (Pessoas): ", s.Capacity)
			fmt.Println("========================")
		}
		return
	}

	fmt.Println("Erro desconhecido")
}

func DeleteSpace(username string, password string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Deleção de Espaço")
	fmt.Printf("ID do Espaço: ")
	input, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Println("Erro ao ler entrada: ", err)
		return
	}

	input = strings.TrimSpace(input)
	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Erro ao converter entrada")
		return
	}

	url := fmt.Sprintf("http://server:8080/space/%d", id)

	payload := map[string]interface{}{
		"username": username,
		"password": password,
	}

	jsonData, _ := json.Marshal(payload)
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	// Faz a requisição para o backend
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Trata valores de retorno
	if resp.StatusCode == http.StatusNotFound {
		fmt.Printf("O Espaço não foi encontrado\n")
		return
	}

	if resp.StatusCode == http.StatusBadRequest {
		fmt.Printf("A senha do administrador está incorreta\n")
		return
	}

	if resp.StatusCode == http.StatusInternalServerError {
		fmt.Printf("Houve um erro interno no servidor\n")
		return
	}

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("O Espaço foi deletado com sucesso\n")
		return
	}

	fmt.Println("Erro desconhecido")
}
