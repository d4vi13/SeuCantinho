package internal

import "fmt"

type Session interface {
	ShowOptions()
	Handler(opt int) int
}

type AdminSession struct {
	Data *SessionData
}
type ClientSession struct {
	Data *SessionData
}

func (session *AdminSession) ShowOptions() {
	fmt.Printf("0  - Encerrar Execução\n")
	fmt.Printf("1  - Criar Espaço\n")
	fmt.Printf("2  - Atualizar Espaço\n")
	fmt.Printf("3  - Obter Espaço\n")
	fmt.Printf("4  - Obter todos os espaços\n")
	fmt.Printf("5  - Deletar Espaço\n")
	fmt.Printf("6  - Obter Usuário\n")
	fmt.Printf("7  - Obter todos os usuários\n")
	fmt.Printf("8  - Deletar Usuário\n")
	fmt.Printf("9  - Criar Reserva\n")
	fmt.Printf("10 - Obter todas as reservas\n")
	fmt.Printf("11 - Cancelar Reserva\n")
}

func (session *ClientSession) ShowOptions() {
	fmt.Printf("0 - Encerrar Execução\n")
	fmt.Printf("1 - Obter todos os espaços\n")
	fmt.Printf("2 - Criar Reserva\n")
	fmt.Printf("3 - Obter todas as reservas\n")
	fmt.Printf("4 - Cancelar Reserva\n")
}

func (session *AdminSession) Handler(opt int) int {
	if opt == 0 {
		return 0
	}

	if opt == 1 {
		CreateSpace(session.Data.User.Username, session.Data.User.Password)
		return 1
	}

	if opt == 2 {
		UpdateSpace(session.Data.User.Username, session.Data.User.Password)
		return 1
	}

	if opt == 3 {
		GetSpace()
		return 1
	}

	if opt == 4 {
		GetAllSpaces()
		return 1
	}

	if opt == 5 {
		DeleteSpace(session.Data.User.Username, session.Data.User.Password)
		return 1
	}

	if opt == 6 {
		GetUser()
		return 1
	}

	if opt == 7 {
		GetAllUsers()
		return 1
	}

	return 0
}

func (session *ClientSession) Handler(opt int) int {
	if opt == 0 {
		return 0
	}

	if opt == 1 {
		GetAllSpaces()
		return 1
	}

	return 0
}
