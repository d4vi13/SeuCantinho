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

		login.Login(username, password)
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
