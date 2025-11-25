package space

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	svc "github.com/d4vi13/SeuCantinho/server/internal/services/space"
)

func (controller *SpaceController) DeleteSpace(w http.ResponseWriter, r *http.Request) {
	var spaceReq RequestSpace

	// Parsing do id do espaço
	id, err := strconv.Atoi(r.PathValue("id"))
	if (err != nil) || (id < 1) {
		http.NotFound(w, r)
		fmt.Printf("Failed to parsing req\n")
		return
	}

	// Faz o parsing da requisição
	err = json.NewDecoder(r.Body).Decode(&spaceReq)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var ret int = controller.spaceService.DeleteSpace(id, spaceReq.Username, spaceReq.Password)

	// Trata valores de retorno
	w.Header().Set("Content-Type", "application/json")
	switch ret {
	case svc.SpaceDeleted:
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"info": "deleted"})
		fmt.Printf("INFO: Space %d was deleted\n", id)
	case svc.SpaceNotFound:
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "space not found"})
		fmt.Printf("INFO: Space %d not found\n", id)
	case svc.WrongPassword:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "wrong password"})
		fmt.Printf("ERROR: The password for User %s is incorrect.\n", spaceReq.Username)
	default:
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "unknown status"})
		fmt.Printf("ERROR: Unknown Status\n")

	}
}
