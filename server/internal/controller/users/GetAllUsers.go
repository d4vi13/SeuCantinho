package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	svc "github.com/d4vi13/SeuCantinho/server/internal/services/users"
)

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
