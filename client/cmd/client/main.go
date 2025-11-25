package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	login "github.com/d4vi13/SeuCantinho/client/internal"
)

var reader *bufio.Reader

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

		username = username[:len(username)-1]
		password = password[:len(password)-1]

		session := login.Login(username, password)

		if session.Status == login.UserNotFound {
			fmt.Printf("O usuário %s não existe\n", username)
			return
		}

		if session.Status == login.WrongPassword {
			fmt.Printf("Senha incorreta\n")
			return
		}

		if session.Status == login.Unknown {
			fmt.Printf("Houve um erro desconhecido no servidor\n")
			return
		}

		fmt.Printf("Usuário %s conectado\n", username)
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

		username = username[:len(username)-1]
		password = password[:len(password)-1]
		password_repeat = password_repeat[:len(password_repeat)-1]

		if password != password_repeat {
			fmt.Printf("As senhas são diferentes\n")
			return
		}

		session := login.CreateUser(username, password)

		if session.Status == login.Conflict {
			fmt.Printf("Já existe um usuário com esse nome\n")
			return
		}

		if session.Status == login.Unknown {
			fmt.Printf("Houve um erro desconhecido no servidor\n")
			return
		}
		fmt.Printf("Usuário %s conectado\n", username)

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
