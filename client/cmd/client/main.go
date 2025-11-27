package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	internal "github.com/d4vi13/SeuCantinho/client/internal"
)

var data *internal.SessionData

func main() {
	var opt int
	fmt.Printf("==== Cliente Seu Cantinho ===\n")
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("0- Sair\n")
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
	case 0:
		return
	case 1:
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

		username = strings.TrimSpace(username)
		password = strings.TrimSpace(password)

		data = internal.Login(username, password)

		if data.Status == internal.UserNotFound {
			fmt.Printf("O usuário %s não existe\n", username)
			return
		}

		if data.Status == internal.WrongPassword {
			fmt.Printf("Senha incorreta\n")
			return
		}

		if data.Status == internal.Unknown {
			fmt.Printf("Houve um erro desconhecido no servidor\n")
			return
		}

	case 2:
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

		fmt.Printf("Confirme sua senha: ")
		password_repeat, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao ler entrada: ", err)
			return
		}

		username = strings.TrimSpace(username)
		password = strings.TrimSpace(password)
		password_repeat = strings.TrimSpace(password_repeat)

		if password != password_repeat {
			fmt.Printf("As senhas são diferentes\n")
			return
		}

		data = internal.CreateUser(username, password)

		if data.Status == internal.Conflict {
			fmt.Printf("Já existe um usuário com esse nome\n")
			return
		}

		if data.Status == internal.Unknown {
			fmt.Printf("Houve um erro desconhecido no servidor\n")
			return
		}

	default:
		return
	}

	fmt.Printf("Usuário %s conectado\n\n", data.User.Username)

	var session internal.Session

	if data.IsAdmin {
		session = &internal.AdminSession{Data: data}
	} else {
		session = &internal.ClientSession{Data: data}
	}

	opt = 1
	for opt != 0 {
		session.ShowOptions()
		fmt.Printf("Selecione uma opção: ")

		input, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println("Erro ao ler entrada: ", err)
			return
		}

		fmt.Println()
		trimmedInput := strings.TrimSpace(input)
		num, _ = strconv.Atoi(trimmedInput)
		opt = session.Handler(num)
		fmt.Println()
	}
}
