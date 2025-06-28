package food

import "github.com/gatsu420/marianne/app/repository"

type Usecase interface {
	GetFood(id int) (*GetFoodRow, error)
}

type usecaseImpl struct {
	pgRepo repository.PGRepo
}

func NewUsecase(pgRepo repository.PGRepo) Usecase {
	return &usecaseImpl{
		pgRepo: pgRepo,
	}
}
