package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	svc "github.com/d4vi13/SeuCantinho/server/internal/services/users"
)

func (controller *UsersController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var userReq RequestUser

	// Parsing do id do usuário
	id, err := strconv.Atoi(r.PathValue("id"))
	if (err != nil) || (id < 1) {
		http.NotFound(w, r)
		fmt.Printf("Failed to parsing req\n")
		return
	}

	// Faz o parsing da requisição
	err = json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var ret int = controller.usersService.DeleteUser(id, userReq.Username, userReq.Password)

	// Trata valores de retorno
	w.Header().Set("Content-Type", "application/json")
	switch ret {
	case svc.UserDeleted:
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"info": "deleted"})
		fmt.Printf("INFO: User %d was deleted\n", id)
	case svc.UserNotFound:
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "user not found"})
		fmt.Printf("INFO: User %d not found\n", id)
	case svc.WrongPassword:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "wrong password"})
		fmt.Printf("ERROR: The password for User %s is incorrect.\n", userReq.Username)
	case svc.InvalidAdmin:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "user is not an admin"})
		fmt.Printf("ERROR: User %s is not an admin\n", userReq.Username)
	case svc.InvalidDelete:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "this user is an admin"})
		fmt.Printf("ERROR: User %d is an admin.\n", id)
	case svc.InternalError:
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
		fmt.Printf("ERROR: Internal Server Error\n")
	default:
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "unknown status"})
		fmt.Printf("ERROR: Unknown Status\n")

	}
}
