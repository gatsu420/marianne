package food

import (
	stderr "errors"

	"github.com/gatsu420/marianne/common/errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type GetFoodRow struct {
	ID           int
	Name         string
	Type         pgtype.Text
	IntakeStatus pgtype.Text
	Feeder       pgtype.Text
	Location     pgtype.Text
	Remarks      pgtype.Text
	CreatedAt    pgtype.Timestamptz
	UpdatedAt    pgtype.Timestamptz
}

func (u *usecaseImpl) GetFood(id int) (*GetFoodRow, error) {
	food, err := u.pgRepo.GetFood(id)
	if err != nil {
		if stderr.Is(err, pgx.ErrNoRows) {
			return nil, errors.New(errors.ErrMsgFoodNotFound, errors.ErrFoodNotFound)
		}
		return nil, errors.New(errors.ErrMsgInternal, errors.ErrInternal)
	}

	return &GetFoodRow{
		ID:           food.ID,
		Name:         food.Name,
		Type:         food.Type,
		IntakeStatus: food.IntakeStatus,
		Feeder:       food.Feeder,
		Location:     food.Location,
		Remarks:      food.Remarks,
		CreatedAt:    food.CreatedAt,
		UpdatedAt:    food.UpdatedAt,
	}, nil
}
