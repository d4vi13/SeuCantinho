package space

import (
	"encoding/json"
	"fmt"
	"net/http"

	svc "github.com/d4vi13/SeuCantinho/internal/services/space"
)

func (controller *SpaceController) GetAllSpaces(w http.ResponseWriter, r *http.Request) {
	spaces, ret := controller.spaceService.GetAllSpaces()

	// Trata valores de retorno
	w.Header().Set("Content-Type", "application/json")
	switch ret {
	case svc.SpacesFound:
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(spaces)
		fmt.Printf("INFO: Spaces found\n")
	case svc.SpacesNotFound:
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "spaces not found"})
		fmt.Printf("INFO: Spaces not found\n")
	default:
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "unknown status"})
		fmt.Printf("ERROR: Unknown Status\n")
	}

}
