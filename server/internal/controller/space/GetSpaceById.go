package space

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	svc "github.com/d4vi13/SeuCantinho/server/internal/services/space"
)

// GetSpaceById godoc
// @Summary Retorna um espaço específico
// @Description Retorna um JSON com os dados do espaço especificado pelo ID
// @Tags Spaces
// @Produce json
// @Param id path int true "ID do espaço"
// @Success 200 {object} models.Space "Espaço encontrado com sucesso"
// @Failure 404 {object} map[string]string "Espaço não encontrado"
// @Failure 500 {object} map[string]string "Erro interno do servidor"
// @Router /space/{id} [get]
func (controller *SpaceController) GetSpaceById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if (err != nil) || (id < 1) {
		http.NotFound(w, r)
		fmt.Printf("Failed to parsing req\n")
		return
	}

	space, ret := controller.spaceService.GetSpaceById(id)

	// Trata valores de retorno
	w.Header().Set("Content-Type", "application/json")
	switch ret {
	case svc.SpaceFound:
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(space)
		fmt.Printf("INFO: Space %d found\n", id)
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
