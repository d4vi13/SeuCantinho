package users

import (
	"encoding/json"
	"net/http"

	svc "github.com/d4vi13/SeuCantinho/internal/services/users"
)

func (controller *UsersController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userReq CreateRequestUser

	// Faz o parsing na requisição
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Chama o serviço para criar o usuário
	user, ret := controller.usersService.CreateUser(userReq.Username, userReq.Passhash)

	// Trata retornos
	w.Header().Set("Content-Type", "application/json")
	switch ret {
	case svc.UserCreated:
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	case svc.UserFound:
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(map[string]string{"error": "user already exists"})
	case svc.InternalError:
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
	default:
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "unknown status"})
	}
}
