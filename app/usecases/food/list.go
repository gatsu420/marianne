package food

import (
	"context"
	"time"

	"github.com/gatsu420/marianne/app/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type ListFoodRow struct {
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

type ListFoodArgs struct {
	StartTimestamp time.Time
	EndTimestamp   time.Time
	Type           string
	IntakeStatus   string
	Feeder         string
	Location       string
}

func (u *usecaseImpl) ListFood(ctx context.Context, args *ListFoodArgs) ([]ListFoodRow, error) {
	food, err := u.pgRepo.ListFood(ctx, repository.ListFoodArgs{
		StartTimestamp: pgtype.Timestamptz{
			Time:  args.StartTimestamp,
			Valid: !args.StartTimestamp.IsZero(),
		},
		EndTimestamp: pgtype.Timestamptz{
			Time:  args.EndTimestamp,
			Valid: !args.EndTimestamp.IsZero(),
		},
		Type: pgtype.Text{
			String: args.Type,
			Valid:  args.Type != "",
		},
		IntakeStatus: pgtype.Text{
			String: args.IntakeStatus,
			Valid:  args.IntakeStatus != "",
		},
		Feeder: pgtype.Text{
			String: args.Feeder,
			Valid:  args.Feeder != "",
		},
		Location: pgtype.Text{
			String: args.Location,
			Valid:  args.Location != "",
		},
	})
	if err != nil {
		return nil, err
	}

	var rows []ListFoodRow
	for _, f := range food {
		rows = append(rows, ListFoodRow{
			ID:           f.ID,
			Name:         f.Name,
			Type:         f.Type.String,
			IntakeStatus: f.IntakeStatus.String,
			Feeder:       f.Feeder.String,
			Location:     f.Location.String,
			Remarks:      f.Remarks.String,
			CreatedAt:    f.CreatedAt.Time,
			UpdatedAt:    f.UpdatedAt.Time,
		})
	}
	return rows, nil
}
