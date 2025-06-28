package handlers_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	"github.com/gatsu420/marianne/app/handlers"
	"github.com/gatsu420/marianne/app/usecases/food"
	mockfood "github.com/gatsu420/marianne/mocks/app/usecases/food"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
	mockFoodUsecase *mockfood.MockUsecase
	handler         handlers.Handler

	dummyPGText        pgtype.Text
	dummyPGTimestamptz pgtype.Timestamptz
	dummyGetFoodRow    *food.GetFoodRow
	dummyGetFoodRowStr string
}

func (s *testSuite) SetupSuite() {
	s.dummyPGText = pgtype.Text{
		String: "dummy",
		Valid:  true,
	}
	s.dummyPGTimestamptz = pgtype.Timestamptz{
		Time:  time.Now(),
		Valid: true,
	}

	s.dummyGetFoodRow = &food.GetFoodRow{
		ID:           99,
		Name:         "test",
		Type:         s.dummyPGText,
		IntakeStatus: s.dummyPGText,
		Feeder:       s.dummyPGText,
		Location:     s.dummyPGText,
		Remarks:      s.dummyPGText,
		CreatedAt:    s.dummyPGTimestamptz,
		UpdatedAt:    s.dummyPGTimestamptz,
	}
	gfrow, err := json.Marshal(s.dummyGetFoodRow)
	if err != nil {
		log.Fatal(err)
	}
	s.dummyGetFoodRowStr = string(gfrow)
}

func (s *testSuite) SetupTest() {
	s.mockFoodUsecase = mockfood.NewMockUsecase(s.T())
	s.handler = handlers.NewHandler(s.mockFoodUsecase)
}

func Test(t *testing.T) {
	suite.Run(t, &testSuite{})
}
