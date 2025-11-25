package main

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

type User struct {
	id       int
	Username string
	Password string
}

var session User
var reader *bufio.Reader

func login() {
	fmt.Printf("Usuário: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler entrada: ", err)
		return
	}

	fmt.Printf("Senha: ")
	password, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler entrada: ", err)
		return
	}

	username = username[:len(username)-1] // remove o `\n`
	password = password[:len(password)-1] // remove o `\n`

	session.Username = username
	session.Password = password

	payload := map[string]interface{}{
		"username": session.Username,
		"password": session.Password,
	}

	jsonData, _ := json.Marshal(payload)
	resp, err := http.Post("http://server:8080/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Resposta:", string(body))
}

func main() {
	fmt.Printf("==== Cliente Seu Cantinho ===\n")
	reader = bufio.NewReader(os.Stdin)

	fmt.Printf("1- Login\n")
	fmt.Printf("2- Criar Usuário\n")
	fmt.Printf("Selecione uma opção: ")

	input, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Println("Erro ao ler entrada: ", err)
		return
	}

	fmt.Printf("\n")

	trimmedInput := strings.TrimSpace(input)

	num, _ := strconv.Atoi(trimmedInput)
	switch num {
	case 1:
		login()
	case 2:
		fmt.Printf("Criou...\n")
	default:
		return
	}

	// for {
	// 	fmt.Print("Digite algo: ")
	// 	text, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		fmt.Println("Erro na leitura:", err)
	// 		return
	// 	}
	// 	text = text[:len(text)-1] // remove o `\n`
	// 	if text == "exit" {
	// 		fmt.Println("Saindo...")
	// 		break
	// 	}
	// 	fmt.Println("Você digitou:", text)
	// }
}
