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
	"time"
)

type RequestBooking struct {
	Id        int    `json:"id"`
	UserId    int    `json:"UserId"`
	SpaceId   int    `json:"SpaceId"`
	StartDate string `json:"StartDate"`
	EndDate   string `json:"EndDate"`
	Days      int    `json:"Days"`
}

type RequestPayment struct {
	Id    int   `json:"Id"`
	Total int64 `json:"TotalValue"`
	Payed int64 `json:"PayedValue"`
}

func getPayment(id int) (float64, float64) {
	var req RequestPayment
	url := fmt.Sprintf("http://server:8080/payments/%d", id)

	// Faz a requisição ao backend
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Trata valores de retorno
	if resp.StatusCode == http.StatusNotFound {
		fmt.Printf("Esse pagamento não existe\n")
		return -1, -1
	}

	if resp.StatusCode == http.StatusInternalServerError {
		fmt.Printf("Houve um erro interno no servidor\n")
		return -1, -1
	}

	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(&req); err != nil {
			panic(err)
		}

		return (float64(req.Total) / 100), (float64(req.Payed) / 100)
	}

	fmt.Println("Erro desconhecido")

	return -1, -1
}

func BookSpace(username string, password string) {
	reader := bufio.NewReader(os.Stdin)
	var req RequestBooking

	fmt.Println("Reserva de Espaço")
	fmt.Printf("ID do Espaço: ")
	input, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Println("Erro ao ler entrada: ", err)
		return
	}

	input = strings.TrimSpace(input)
	spaceId, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Erro ao converter entrada")
		return
	}

	fmt.Print("Data inicial (AAAA-MM-DD): ")
	startDate, _ := reader.ReadString('\n')
	startDate = strings.TrimSpace(startDate)

	fmt.Print("Quantidade de dias: ")
	daysStr, _ := reader.ReadString('\n')
	daysStr = strings.TrimSpace(daysStr)
	days, err := strconv.Atoi(daysStr)
	if err != nil {
		fmt.Println("Dias inválidos:", err)
		return
	}

	if days <= 0 {
		fmt.Printf("O número de dias deve ser maior que zero\n")
		return
	}

	payload := map[string]interface{}{
		"username":    username,
		"password":    password,
		"space":       spaceId,
		"startDate":   startDate,
		"bookingTime": days,
	}

	// Faz a requisição para o backend
	jsonData, _ := json.Marshal(payload)
	resp, err := http.Post("http://server:8080/bookings", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Trata valores de retorno
	if resp.StatusCode == http.StatusNotFound {
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
		fmt.Printf("Conflito ao realizar reserva\n")
		return
	}

	if resp.StatusCode == http.StatusBadRequest {
		fmt.Printf("Tentativa de reserva invalida\n")
		return
	}

	if resp.StatusCode == http.StatusUnauthorized {
		fmt.Printf("Senha do usuário incorreta\n")
		return
	}

	if resp.StatusCode == http.StatusCreated {
		err = json.NewDecoder(resp.Body).Decode(&req)
		if err != nil {
			panic(err)
		}

		total, _ := getPayment(req.Id)

		if total == -1 {
			return
		}

		startParsed, err := time.Parse("2006-01-02", startDate)
		if err != nil {
			fmt.Printf("Falha ao converter datas\n")
			return
		}

		endParsed := startParsed.AddDate(0, 0, days)
		endDate := endParsed.Format("2006-01-02")

		fmt.Println()
		fmt.Println("========================")
		fmt.Println("ID:", req.Id)
		fmt.Println("ID do Espaço: ", spaceId)
		fmt.Println("Data de Início: ", startDate)
		fmt.Println("Data de Fim: ", endDate)
		fmt.Println("Dias Reservados: ", days)
		fmt.Println("Custo Total (R$): ", total)
		fmt.Println("========================")
		fmt.Println()

		return
	}

	fmt.Println("Erro desconhecido")
}

func GetBooking() {
	var booking RequestBooking
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("ID da Reserva: ")
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

	url := fmt.Sprintf("http://server:8080/bookings/%d", id)

	// Faz a requisição ao backend
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Trata valores de retorno
	if resp.StatusCode == http.StatusNotFound {
		fmt.Printf("Essa reserva não existe\n")
		return
	}

	if resp.StatusCode == http.StatusInternalServerError {
		fmt.Printf("Houve um erro interno no servidor\n")
		return
	}

	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(&booking); err != nil {
			panic(err)
		}

		total, payed := getPayment(booking.Id)

		if total == -1 {
			return
		}

		fmt.Println()
		fmt.Println("========================")
		fmt.Println("ID:", booking.Id)
		fmt.Println("ID do Usuário: ", booking.UserId)
		fmt.Println("ID do Espaço: ", booking.SpaceId)
		fmt.Println("Data de Início: ", booking.StartDate)
		fmt.Println("Data de Fim: ", booking.EndDate)
		fmt.Println("Dias Reservados: ", booking.Days)
		fmt.Println("Valor Pago (R$): ", payed)
		fmt.Println("Custo Total (R$): ", total)
		fmt.Println("========================")
		fmt.Println()

		return
	}

	fmt.Println("Erro desconhecido")
}

func GetMyBookings(id int) {
	var bookings []RequestBooking

	url := fmt.Sprintf("http://server:8080/users/%d/bookings", id)

	// Faz a requisição ao backend
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Trata valores de retorno
	if resp.StatusCode == http.StatusNotFound {
		fmt.Printf("Não existe nenhuma reserva\n")
		return
	}

	if resp.StatusCode == http.StatusInternalServerError {
		fmt.Printf("Houve um erro interno no servidor\n")
		return
	}

	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(&bookings); err != nil {
			panic(err)
		}

		for _, b := range bookings {

			total, payed := getPayment(b.Id)
			if total == -1 {
				return
			}

			fmt.Println("========================")
			fmt.Println("ID:", b.Id)
			fmt.Println("ID do Usuário: ", b.UserId)
			fmt.Println("ID do Espaço: ", b.SpaceId)
			fmt.Println("Data de Início: ", b.StartDate)
			fmt.Println("Data de Fim: ", b.EndDate)
			fmt.Println("Dias Reservados: ", b.Days)
			fmt.Println("Valor Pago (R$): ", payed)
			fmt.Println("Custo Total (R$): ", total)
			fmt.Println("========================")
		}
		return
	}

	fmt.Println("Erro desconhecido")
}

func GetAllBookings() {
	var bookings []RequestBooking

	// Faz a requisição ao backend
	resp, err := http.Get("http://server:8080/bookings")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Trata valores de retorno
	if resp.StatusCode == http.StatusNotFound {
		fmt.Printf("Não existe nenhuma reserva\n")
		return
	}

	if resp.StatusCode == http.StatusInternalServerError {
		fmt.Printf("Houve um erro interno no servidor\n")
		return
	}

	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(&bookings); err != nil {
			panic(err)
		}

		for _, b := range bookings {

			total, payed := getPayment(b.Id)
			if total == -1 {
				return
			}

			fmt.Println("========================")
			fmt.Println("ID:", b.Id)
			fmt.Println("ID do Usuário: ", b.UserId)
			fmt.Println("ID do Espaço: ", b.SpaceId)
			fmt.Println("Data de Início: ", b.StartDate)
			fmt.Println("Data de Fim: ", b.EndDate)
			fmt.Println("Dias Reservados: ", b.Days)
			fmt.Println("Valor Pago (R$): ", payed)
			fmt.Println("Custo Total (R$): ", total)
			fmt.Println("========================")
		}
		return
	}

	fmt.Println("Erro desconhecido")
}
