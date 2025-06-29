package food

import (
	"errors"

	commonerr "github.com/gatsu420/marianne/common/errors"
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
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, commonerr.New(commonerr.ErrMsgFoodNotFound, commonerr.ErrFoodNotFound)
		}
		return nil, commonerr.New(commonerr.ErrMsgInternal, commonerr.ErrInternal)
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
