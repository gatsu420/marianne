package food_test

import (
	"errors"
	"time"

	"github.com/gatsu420/marianne/app/repository"
	"github.com/gatsu420/marianne/app/usecases/food"
	"github.com/stretchr/testify/mock"
)

func (s *testSuite) Test_ListFood() {
	testCases := []struct {
		testName     string
		args         *food.ListFoodArgs
		repoFood     []repository.ListFoodRow
		repoErr      error
		expectedFood []food.ListFoodRow
		expectedErr  error
	}{
		{
			testName: "repo error",
			args: &food.ListFoodArgs{
				StartTimestamp: time.Now(),
				EndTimestamp:   time.Now(),
				Type:           "test",
				IntakeStatus:   "test",
				Feeder:         "test",
				Location:       "test",
			},
			repoFood:     nil,
			repoErr:      errors.New("some error"),
			expectedFood: nil,
			expectedErr:  errors.New("some error"),
		},
		{
			testName: "success",
			args: &food.ListFoodArgs{
				StartTimestamp: time.Now(),
				EndTimestamp:   time.Now(),
				Type:           "test",
				IntakeStatus:   "test",
				Feeder:         "test",
				Location:       "test",
			},
			repoFood: []repository.ListFoodRow{
				{
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
			},
			repoErr: nil,
			expectedFood: []food.ListFoodRow{
				{
					ID:           99,
					Name:         "test",
					Type:         "dummy",
					IntakeStatus: "dummy",
					Feeder:       "dummy",
					Location:     "dummy",
					Remarks:      "dummy",
					CreatedAt:    s.dummyTimestamp,
					UpdatedAt:    s.dummyTimestamp,
				},
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.testName, func() {
			s.mockPGRepo.EXPECT().ListFood(
				mock.Anything,
				mock.AnythingOfType("repository.ListFoodArgs"),
			).Return(tc.repoFood, tc.repoErr).Once()

			rows, err := s.usecase.ListFood(s.ctx, tc.args)
			s.Equal(tc.expectedErr, err)
			s.Equal(tc.expectedFood, rows)
		})
	}
}
