package space

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	svc "github.com/d4vi13/SeuCantinho/internal/services/space"
)

func (controller *SpaceController) DeleteSpace(w http.ResponseWriter, r *http.Request) {
	// Parsing do id do espaço
	id, err := strconv.Atoi(r.PathValue("id"))
	if (err != nil) || (id < 1) {
		http.NotFound(w, r)
		fmt.Printf("Failed to parsing req\n")
		return
	}

	var ret int = controller.spaceService.DeleteSpace(id)

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
	default:
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "unknown status"})
		fmt.Printf("ERROR: Unknown Status\n")

	}
}
