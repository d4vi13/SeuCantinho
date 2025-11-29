package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	svc "github.com/d4vi13/SeuCantinho/server/internal/services/users"
)

// GetAllUsers godoc
// @Summary Lista todos os usuários
// @Description Retorna um array com todos os usuários cadastrados
// @Tags Users
// @Produce json
// @Success 200 {array} models.User "Lista de usuários cadastrados"
// @Failure 404 {object} map[string]string "Nenhum usuário encontrado"
// @Failure 500 {object} map[string]string "Erro interno do servidor"
// @Router /users [get]
func (controller *UsersController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, ret := controller.usersService.GetAllUsers()

	// Trata valores de retorno
	w.Header().Set("Content-Type", "application/json")
	switch ret {
	case svc.UsersFound:
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
		fmt.Printf("INFO: Users found\n")
	case svc.UsersNotFound:
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "users not found"})
		fmt.Printf("INFO: Spaces not found\n")
	default:
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "unknown status"})
		fmt.Printf("ERROR: Unknown Status\n")
	}

}
