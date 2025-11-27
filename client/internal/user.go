package internal

import (
	"bufio"
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
		fmt.Println("========================")
		fmt.Println("ID: ", user.ID)
		fmt.Println("Username: ", user.Username)
		fmt.Println("PassHash: ", user.PassHash)
		fmt.Println("Is Admin?: ", user.IsAdmin)
		fmt.Println("========================")
		return
	}

	fmt.Println("Erro desconhecido")
}
