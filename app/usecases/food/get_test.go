package food_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/gatsu420/marianne/app/usecases/food"
	commonerr "github.com/gatsu420/marianne/common/errors"
	"github.com/gatsu420/marianne/common/testassert"
	mockrepository "github.com/gatsu420/marianne/mocks/app/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func Test_GetFood(t *testing.T) {
	testCases := []struct {
		testName    string
		id          int
		repoErr     error
		expectedRow *food.GetFoodRow
		expectedErr *commonerr.Err
	}{
		{
			testName:    "repo error",
			id:          99,
			repoErr:     errors.New("some repo error"),
			expectedRow: nil,
			expectedErr: commonerr.New(commonerr.ErrMsgInternal, commonerr.ErrInternal),
		},
		{
			testName:    "food is not found",
			id:          99,
			repoErr:     pgx.ErrNoRows,
			expectedRow: nil,
			expectedErr: commonerr.New(commonerr.ErrMsgFoodNotFound, commonerr.ErrFoodNotFound),
		},
		{
			testName: "success",
			id:       99,
			repoErr:  nil,
			expectedRow: &food.GetFoodRow{
				ID:           99,
				Name:         "mock",
				Type:         pgtype.Text{String: "mock", Valid: true},
				IntakeStatus: pgtype.Text{String: "mock", Valid: true},
				Feeder:       pgtype.Text{String: "mock", Valid: true},
				Location:     pgtype.Text{String: "mock", Valid: true},
				Remarks:      pgtype.Text{String: "mock", Valid: true},
				CreatedAt:    pgtype.Timestamptz{Time: time.Date(2025, time.July, 4, 20, 47, 0, 0, time.UTC), Valid: true},
				UpdatedAt:    pgtype.Timestamptz{Time: time.Date(2025, time.July, 4, 20, 47, 0, 0, time.UTC), Valid: true},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			mockPGRepo := mockrepository.NewMockPGRepo(
				mockrepository.WithExpectedErr(tc.repoErr),
			)
			usecase := food.NewUsecase(mockPGRepo)
			ctx := context.Background()

			row, err := usecase.GetFood(ctx, tc.id)
			cerr, ok := err.(*commonerr.Err)
			if err != nil && !ok {
				t.Error("unable to cast err as *commonerr.Err")
			}

			testassert.Equal(t, row, tc.expectedRow)
			testassert.Equal(t, cerr, tc.expectedErr)
		})
	}
}
