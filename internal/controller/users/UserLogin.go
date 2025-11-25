package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	svc "github.com/d4vi13/SeuCantinho/internal/services/users"
)

func (controller *UsersController) UserLogin(w http.ResponseWriter, r *http.Request) {
	var userReq RequestUser

	// Faz o parsing na requisição
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Chama o serviço para criar o usuário
	user, ret := controller.usersService.Login(userReq.Username, userReq.Password)

	// Trata retornos
	w.Header().Set("Content-Type", "application/json")
	switch ret {
	case svc.Logged:
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
		fmt.Printf("INFO: User %s login succesfuly\n", userReq.Username)
	case svc.UserNotFound:
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "user not found"})
		fmt.Printf("INFO: User %s not found\n", userReq.Username)
	case svc.WrongPassword:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "wrong password"})
		fmt.Printf("ERROR: The password for User %s is incorrect.\n", userReq.Username)
	default:
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "unknown status"})
		fmt.Printf("ERROR: Internal Server Error\n")
	}

}
