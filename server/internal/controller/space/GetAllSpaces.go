package space

import (
	"encoding/json"
	"fmt"
	"net/http"

	svc "github.com/d4vi13/SeuCantinho/server/internal/services/space"
)

// GetAllSpaces godoc
// @Summary Lista todos os espaços
// @Description Retorna um array com todos os espaço existentes
// @Tags Spaces
// @Produce json
// @Success 200 {array} models.Space "Lista de espaços existentes"
// @Failure 404 {object} map[string]string "Nenhum espaço encontrado"
// @Failure 500 {object} map[string]string "Erro interno do servidor"
// @Router /space [get]
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
