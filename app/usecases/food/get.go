package food

import (
	stderr "errors"

	"github.com/gatsu420/marianne/common/errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type GetFoodRow struct {
	ID           int                `json:"id"`
	Name         string             `json:"name"`
	Type         pgtype.Text        `json:"type"`
	IntakeStatus pgtype.Text        `json:"intakeStatus"`
	Feeder       pgtype.Text        `json:"feeder"`
	Location     pgtype.Text        `json:"location"`
	Remarks      pgtype.Text        `json:"remarks"`
	CreatedAt    pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt    pgtype.Timestamptz `json:"updatedAt"`
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
