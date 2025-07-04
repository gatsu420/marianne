package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gatsu420/marianne/app/usecases/food"
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

	food, err := h.foodUsecases.GetFood(r.Context(), id)
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

func (h *handlerImpl) ListFood(w http.ResponseWriter, r *http.Request) {
	qst := r.URL.Query().Get("startTimestamp")
	qet := r.URL.Query().Get("endTimestamp")
	if qst == "" || qet == "" {
		http.Error(w, "start and end timestamps must be supplied", http.StatusBadRequest)
		return
	}

	st, sterr := time.Parse(time.RFC3339, qst)
	et, eterr := time.Parse(time.RFC3339, qet)
	if sterr != nil || eterr != nil {
		http.Error(w, "start and end timestamps must be in time format", http.StatusBadRequest)
		return
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	llst, llet := st.In(loc), et.In(loc)

	food, err := h.foodUsecases.ListFood(r.Context(), &food.ListFoodArgs{
		StartTimestamp: llst,
		EndTimestamp:   llet,
		Type:           r.URL.Query().Get("type"),
		IntakeStatus:   r.URL.Query().Get("intakeStatus"),
		Feeder:         r.URL.Query().Get("feeder"),
		Location:       r.URL.Query().Get("location"),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&food); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
