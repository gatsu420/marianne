package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gatsu420/marianne/app/usecases/food"
	commonerr "github.com/gatsu420/marianne/common/errors"
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
		if ferr, ok := err.(*commonerr.Err); ok && ferr.Code == commonerr.ErrFoodNotFound {
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

type CreateFoodReqBody struct {
	Name           string `json:"name"`
	TypeID         int    `json:"type_id"`
	IntakeStatusID int    `json:"intake_status_id"`
	FeederID       int    `json:"feeder_id"`
	LocationID     int    `json:"location_id"`
	Remarks        string `json:"remarks"`
}

func (h *handlerImpl) CreateFood(w http.ResponseWriter, r *http.Request) {
	reqBody := CreateFoodReqBody{}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if err := h.foodUsecases.CreateFood(r.Context(), &food.CreateFoodArgs{
		Name:           reqBody.Name,
		TypeID:         reqBody.TypeID,
		IntakeStatusID: reqBody.IntakeStatusID,
		FeederID:       reqBody.FeederID,
		LocationID:     reqBody.LocationID,
		Remarks:        reqBody.Remarks,
	}); err != nil {
		if errors.Is(err, commonerr.New(commonerr.ErrMsgBadRequest, commonerr.ErrBadRequest)) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("food is created"))
}
