package handlers

import (
	"net/http"

	"github.com/gatsu420/marianne/app/usecases/food"
)

type Handler interface {
	GetFood(w http.ResponseWriter, r *http.Request)
}

type handlerImpl struct {
	foodUsecases food.Usecase
}

func NewHandler(foodUsecases food.Usecase) Handler {
	return &handlerImpl{
		foodUsecases: foodUsecases,
	}
}
