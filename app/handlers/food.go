package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gatsu420/marianne/common/errors"
)

func (h *handlerImpl) GetFood(w http.ResponseWriter, r *http.Request) {
	qid := r.URL.Query().Get("id")
	if qid == "" {
		http.Error(w, "id must be supplied", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(qid)
	if err != nil {
		http.Error(w, "id must be integer", http.StatusBadRequest)
		return
	}

	food, err := h.foodUsecases.GetFood(id)
	if err != nil {
		if ferr, ok := err.(*errors.Err); ok && ferr.Code == errors.ErrFoodNotFound {
			http.Error(w, ferr.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(&food); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
