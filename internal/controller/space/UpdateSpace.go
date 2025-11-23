package space

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	svc "github.com/d4vi13/SeuCantinho/internal/services/space"
)

func (controller *SpaceController) UpdateSpace(w http.ResponseWriter, r *http.Request) {
	var spaceReq RequestSpace

	// Parsing do id do espaço
	id, err := strconv.Atoi(r.PathValue("id"))
	if (err != nil) || (id < 1) {
		http.NotFound(w, r)
		fmt.Printf("Failed to parsing req\n")
		return
	}

	// Faz o parsing da requisição
	var body map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Printf("ERROR: %+v\n", err)
		return
	}

	username := body["username"].(string)
	password := body["password"].(string)

	location, hasLocation := body["location"]
	substation, hasSubstation := body["substation"]
	price, hasPrice := body["price"]
	capacity, hasCapacity := body["capacity"]

	img, hasImg := body["img"]

	if !hasLocation {
		location = ""
	}

	if !hasSubstation {
		substation = ""
	}

	if !hasPrice {
		price = -1
	}

	if !hasCapacity {
		capacity = -1
	}

	if hasImg {
		imgStr, ok := img.(string)
		if !ok {
			http.Error(w, "img must be a base64 string", http.StatusBadRequest)
			return
		}

		decoded, err := base64.StdEncoding.DecodeString(imgStr)
		if err != nil {
			http.Error(w, "invalid base64 image", http.StatusBadRequest)
			return
		}

		img = decoded
	} else {
		img = nil
	}

	// Chama o serviço para criar o espaço
	space, ret := controller.spaceService.UpdateSpace(id, username, password, location.(string),
		substation.(string), price.(float64), capacity.(int), img.([]byte))

	// Trata valores de retorno
	w.Header().Set("Content-Type", "application/json")
	switch ret {
	case svc.SpaceUpdated:
		w.WriteHeader(http.StatusCreated)
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
