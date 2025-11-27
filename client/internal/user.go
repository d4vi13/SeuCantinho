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

type RequestUser struct {
	ID       int    `json:"Id"`
	Username string `json:"Username"`
	PassHash string `json:"PassHash"`
	IsAdmin  bool   `json:"IsAdmin"`
}

func GetUser() {
	var user RequestUser
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("ID do Usuário: ")
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

	url := fmt.Sprintf("http://server:8080/users/%d", id)

	// Faz a requisição ao backend
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Trata valores de retorno
	if resp.StatusCode == http.StatusNotFound {
		fmt.Printf("Esse usuário não existe\n")
		return
	}

	if resp.StatusCode == http.StatusInternalServerError {
		fmt.Printf("Houve um erro interno no servidor\n")
		return
	}

	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
			panic(err)
		}

		fmt.Println()
		fmt.Println("========================")
		fmt.Println("ID: ", user.ID)
		fmt.Println("Username: ", user.Username)
		fmt.Println("PassHash: ", user.PassHash)
		fmt.Println("Is Admin?: ", user.IsAdmin)
		fmt.Println("========================")
		fmt.Println()

		return
	}

	fmt.Println("Erro desconhecido")
}

func GetAllUsers() {
	var users []RequestUser

	// Faz a requisição ao backend
	resp, err := http.Get("http://server:8080/users")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Trata valores de retorno
	if resp.StatusCode == http.StatusNotFound {
		fmt.Printf("Não existe nenhum usuário\n")
		return
	}

	if resp.StatusCode == http.StatusInternalServerError {
		fmt.Printf("Houve um erro interno no servidor\n")
		return
	}

	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
			panic(err)
		}

		for _, user := range users {
			fmt.Println("========================")
			fmt.Println("ID: ", user.ID)
			fmt.Println("Username: ", user.Username)
			fmt.Println("PassHash: ", user.PassHash)
			fmt.Println("Is Admin?: ", user.IsAdmin)
			fmt.Println("========================")
		}
		return
	}

	fmt.Println("Erro desconhecido")
}

func DeleteUser(username string, password string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Deleção de Usuário")
	fmt.Printf("ID do Usuário: ")
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

	url := fmt.Sprintf("http://server:8080/users/%d", id)

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
		fmt.Printf("O usuário não foi encontrado\n")
		return
	}

	if resp.StatusCode == http.StatusBadRequest {
		fmt.Printf("A senha do administrador está incorreta\n")
		return
	}

	if resp.StatusCode == http.StatusUnauthorized {
		fmt.Printf("O usuário não é um admin\n")
		return
	}

	if resp.StatusCode == http.StatusLocked {
		fmt.Printf("Não é possível deletar um admin\n")
		return
	}

	if resp.StatusCode == http.StatusInternalServerError {
		fmt.Printf("Houve um erro interno no servidor\n")
		return
	}

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("O usuário foi deletado com sucesso\n")
		return
	}

	fmt.Println("Erro desconhecido")
}
