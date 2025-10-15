package favorapi

import (
	"encoding/json"
	"net/http"

	"github.com/nanagoboiler/internal/favors"
	"github.com/nanagoboiler/models"
)

func Create(s favors.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.FavorRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid Request Json", http.StatusBadRequest)
			return
		}
		err = s.CreateFavor(r.Context(), &req)
		if err != nil {
			http.Error(w, "Service Failure", http.StatusInternalServerError)
		}

	}

}
