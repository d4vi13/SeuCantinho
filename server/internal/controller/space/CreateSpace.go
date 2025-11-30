package space

import (
	"encoding/json"
	"fmt"
	"net/http"

	svc "github.com/d4vi13/SeuCantinho/server/internal/services/space"
)

// CreateSpace godoc
// @Summary Cria um novo espaço
// @Description Cria um espaço com os dados enviados no corpo da requisição. Apenas administradores podem criar espaços.
// @Tags Spaces
// @Accept json
// @Produce json
// @Param space body RequestSpace true "Dados do novo espaço"
// @Success 201 {object} models.Space "Espaço criado com sucesso"
// @Failure 400 {object} models.ErrorResponse "Usuário não encontrado ou não é administrador"
// @Failure 409 {object} models.ErrorResponse "Espaço já existe"
// @Failure 500 {object} models.ErrorResponse "Erro interno do servidor"
// @Router /space [post]
func (controller *SpaceController) CreateSpace(w http.ResponseWriter, r *http.Request) {
	var spaceReq RequestSpace

	// Faz o parsing da requisição
	err := json.NewDecoder(r.Body).Decode(&spaceReq)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Chama o serviço para criar o espaço
	space, ret := controller.spaceService.CreateSpace(spaceReq.Username, spaceReq.Password, spaceReq.Location,
		spaceReq.Substation, spaceReq.Price, spaceReq.Capacity, spaceReq.PNGBytes)

	// Trata valores de retorno
	w.Header().Set("Content-Type", "application/json")
	switch ret {
	case svc.SpaceCreated:
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(space)
		fmt.Printf("INFO: Space %s [%s] created succesfuly\n", spaceReq.Location, space.Substation)
	case svc.UserNotFound:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "user not found"})
		fmt.Printf("INFO: User %s not found\n", spaceReq.Username)
	case svc.InvalidAdmin:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "user is not an admin"})
		fmt.Printf("INFO: User %s is not an Admin\n", spaceReq.Username)
	case svc.SpaceFound:
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(map[string]string{"error": "space already exists"})
		fmt.Printf("INFO: Space %s [%s] already exists\n", spaceReq.Location, spaceReq.Substation)
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
