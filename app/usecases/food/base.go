package food

import (
	"context"

	"github.com/gatsu420/marianne/app/repository"
)

type Usecase interface {
	GetFood(ctx context.Context, id int) (*GetFoodRow, error)
	// ListFood(ctx context.Context, args *ListFoodArgs) ([]ListFoodRow, error)
}

type usecaseImpl struct {
	pgRepo repository.PGRepo
}

func NewUsecase(pgRepo repository.PGRepo) Usecase {
	return &usecaseImpl{
		pgRepo: pgRepo,
	}
}
