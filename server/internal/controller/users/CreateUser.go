package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	svc "github.com/d4vi13/SeuCantinho/server/internal/services/users"
)

// CreateUser godoc
// @Summary Cria um novo usuário
// @Description Cria um usuário com username e password informados no corpo da requisição
// @Tags Users
// @Accept json
// @Produce json
// @Param user body RequestUser true "Dados do novo usuário"
// @Success 201 {object} models.User "Usuário criado com sucesso"
// @Failure 409 {object} map[string]string "Usuário já existe"
// @Failure 500 {object} map[string]string "Erro interno do servidor"
// @Router /users [post]
func (controller *UsersController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userReq RequestUser

	// Faz o parsing na requisição
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Chama o serviço para criar o usuário
	user, ret := controller.usersService.CreateUser(userReq.Username, userReq.Password)

	// Trata retornos
	w.Header().Set("Content-Type", "application/json")
	switch ret {
	case svc.UserCreated:
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
		fmt.Printf("INFO: User %s created succesfuly\n", userReq.Username)
	case svc.UserFound:
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(map[string]string{"error": "user already exists"})
		fmt.Printf("INFO: User %s already exists\n", userReq.Username)
	case svc.InternalError:
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
		fmt.Printf("ERROR: Internal Server Error\n")
	default:
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "unknown status"})
		fmt.Printf("ERROR: Internal Server Error\n")
	}

}
