package space

import (
	"encoding/json"
	"net/http"

	svc "github.com/d4vi13/SeuCantinho/internal/services/space"
)

func (controller *SpaceController) CreateSpace(w http.ResponseWriter, r *http.Request) {
	var spaceReq CreateRequestSpace

	err := json.NewDecoder(r.Body).Decode(&spaceReq)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	ret := controller.spaceService.CreateSpace(spaceReq.Username, spaceReq.Passhash, spaceReq.Location,
		spaceReq.Substation, spaceReq.Price, spaceReq.Capacity, spaceReq.PNGBytes)

	w.Header().Set("Content-Type", "application/json")
	if ret == svc.SpaceCreated {
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(spaceReq)

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}
