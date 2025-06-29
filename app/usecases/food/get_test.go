package food_test

import (
	"errors"

	"github.com/gatsu420/marianne/app/repository"
	"github.com/gatsu420/marianne/app/usecases/food"
	commonerr "github.com/gatsu420/marianne/common/errors"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/mock"
)

func (s *testSuite) Test_GetFood() {
	testCases := []struct {
		testName    string
		id          int
		repoFood    repository.GetFoodRow
		repoErr     error
		expectedRow *food.GetFoodRow
		expectedErr error
	}{
		{
			testName:    "repo error",
			id:          99,
			repoFood:    repository.GetFoodRow{},
			repoErr:     errors.New("some repo error"),
			expectedRow: nil,
			expectedErr: commonerr.New(commonerr.ErrMsgInternal, commonerr.ErrInternal),
		},
		{
			testName:    "ID is not found",
			id:          99,
			repoFood:    repository.GetFoodRow{},
			repoErr:     pgx.ErrNoRows,
			expectedRow: nil,
			expectedErr: commonerr.New(commonerr.ErrMsgFoodNotFound, commonerr.ErrFoodNotFound),
		},
		{
			testName: "success",
			id:       99,
			repoFood: repository.GetFoodRow{
				ID:           99,
				Name:         "test",
				Type:         s.dummyPGText,
				IntakeStatus: s.dummyPGText,
				Feeder:       s.dummyPGText,
				Location:     s.dummyPGText,
				Remarks:      s.dummyPGText,
				CreatedAt:    s.dummyPGTimestamptz,
				UpdatedAt:    s.dummyPGTimestamptz,
			},
			repoErr: nil,
			expectedRow: &food.GetFoodRow{
				ID:           99,
				Name:         "test",
				Type:         s.dummyPGText,
				IntakeStatus: s.dummyPGText,
				Feeder:       s.dummyPGText,
				Location:     s.dummyPGText,
				Remarks:      s.dummyPGText,
				CreatedAt:    s.dummyPGTimestamptz,
				UpdatedAt:    s.dummyPGTimestamptz,
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.testName, func() {
			s.mockPGRepo.EXPECT().GetFood(
				mock.AnythingOfType("int"),
			).Return(tc.repoFood, tc.repoErr).Once()
		})

		row, err := s.usecase.GetFood(tc.id)
		s.Equal(tc.expectedRow, row)
		s.Equal(tc.expectedErr, err)
	}
}
