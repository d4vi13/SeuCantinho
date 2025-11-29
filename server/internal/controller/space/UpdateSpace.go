package space

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	svc "github.com/d4vi13/SeuCantinho/server/internal/services/space"
)

// UpdateSpace godoc
// @Summary Atualiza um espaço existente
// @Description Atualiza os dados de um espaço existente com base no ID informado. Apenas administradores podem atualizar espaços.
// @Tags Spaces
// @Accept json
// @Produce json
// @Param id path int true "ID do espaço"
// @Param space body RequestSpace true "Dados atualizados do espaço"
// @Success 201 {object} models.Space "Espaço atualizado com sucesso"
// @Failure 400 {object} models.ErrorResponse "Usuário não encontrado, não é admin, senha incorreta ou espaço não existe"
// @Failure 500 {object} models.ErrorResponse "Erro interno do servidor"
// @Router /space/{id} [put]
func (controller *SpaceController) UpdateSpace(w http.ResponseWriter, r *http.Request) {
	var spaceReq RequestSpace

	// Parsing do id do espaço
	id, err := strconv.Atoi(r.PathValue("id"))
	if (err != nil) || (id < 1) {
		http.NotFound(w, r)
		fmt.Printf("Failed to parsing req\n")
		return
	}

	spaceReq.Location = ""
	spaceReq.Substation = ""
	spaceReq.Price = -1
	spaceReq.Capacity = -1
	spaceReq.PNGBytes = nil

	// Faz o parsing da requisição
	err = json.NewDecoder(r.Body).Decode(&spaceReq)
	if err != nil {
		fmt.Printf("ERROR: %+v\n", err)
		return
	}

	// Chama o serviço para criar o espaço
	space, ret := controller.spaceService.UpdateSpace(id, spaceReq.Username, spaceReq.Password, spaceReq.Location, spaceReq.Substation, spaceReq.Price, spaceReq.Capacity, spaceReq.PNGBytes)

	// Trata valores de retorno
	w.Header().Set("Content-Type", "application/json")
	switch ret {
	case svc.SpaceUpdated:
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(space)
		fmt.Printf("INFO: Space %s [%s] updated succesfuly\n", spaceReq.Location, space.Substation)
	case svc.UserNotFound:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "user not found"})
		fmt.Printf("INFO: User %s not found\n", spaceReq.Username)
	case svc.InvalidAdmin:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "user is not an admin"})
		fmt.Printf("INFO: User %s is not an Admin\n", spaceReq.Username)
	case svc.WrongPassword:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "wrong password"})
		fmt.Printf("ERROR: The password for User %s is incorrect.\n", spaceReq.Username)
	case svc.SpaceNotFound:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "space not found"})
		fmt.Printf("ERROR: Space %d not found\n", id)
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
