package food

import (
	"context"

	"github.com/gatsu420/marianne/app/repository"
	commonerr "github.com/gatsu420/marianne/common/errors"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateFoodArgs struct {
	Name           string
	TypeID         int
	IntakeStatusID int
	FeederID       int
	LocationID     int
	Remarks        string
}

func (u *usecaseImpl) CreateFood(ctx context.Context, args *CreateFoodArgs) error {
	err := u.pgRepo.CreateFood(ctx, repository.CreateFoodArgs{
		Name:           args.Name,
		TypeID:         args.TypeID,
		IntakeStatusID: args.IntakeStatusID,
		FeederID:       args.FeederID,
		LocationID:     args.LocationID,
		Remarks: pgtype.Text{
			String: args.Remarks,
			Valid:  args.Remarks != "",
		},
	})
	if err != nil {
		return commonerr.New(commonerr.ErrMsgInternal, commonerr.ErrInternal)
	}
	return nil
}
