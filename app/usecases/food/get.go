package food

import (
	"context"
	"errors"
	"time"

	commonerr "github.com/gatsu420/marianne/common/errors"
	"github.com/jackc/pgx/v5"
)

type GetFoodRow struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	IntakeStatus string    `json:"intakeStatus"`
	Feeder       string    `json:"feeder"`
	Location     string    `json:"location"`
	Remarks      string    `json:"remarks"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func (u *usecaseImpl) GetFood(ctx context.Context, id int) (*GetFoodRow, error) {
	food, err := u.pgRepo.GetFood(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, commonerr.New(commonerr.ErrMsgFoodNotFound, commonerr.ErrFoodNotFound)
		}
		return nil, commonerr.New(commonerr.ErrMsgInternal, commonerr.ErrInternal)
	}

	return &GetFoodRow{
		ID:           food.ID,
		Name:         food.Name,
		Type:         food.Type.String,
		IntakeStatus: food.IntakeStatus.String,
		Feeder:       food.Feeder.String,
		Location:     food.Location.String,
		Remarks:      food.Remarks.String,
		CreatedAt:    food.CreatedAt.Time,
		UpdatedAt:    food.UpdatedAt.Time,
	}, nil
}
