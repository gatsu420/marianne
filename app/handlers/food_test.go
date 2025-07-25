package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gatsu420/marianne/app/handlers"
	"github.com/gatsu420/marianne/app/usecases/food"
	commonerr "github.com/gatsu420/marianne/common/errors"
	"github.com/gatsu420/marianne/common/tests"
	mockusecases "github.com/gatsu420/marianne/mocks/app/usecases"
)

func Test_GetFood(t *testing.T) {
	mockRow, err := json.Marshal(&food.GetFoodRow{
		ID:           99,
		Name:         "mock",
		Type:         tests.MockPGText().String,
		IntakeStatus: tests.MockPGText().String,
		Feeder:       tests.MockPGText().String,
		Location:     tests.MockPGText().String,
		Remarks:      tests.MockPGText().String,
		CreatedAt:    tests.MockPGTimestamptz().Time,
		UpdatedAt:    tests.MockPGTimestamptz().Time,
	})
	if err != nil {
		t.Errorf("failed to serialize: %v", err)
	}

	testCases := []struct {
		testName           string
		usecaseErr         error
		w                  http.ResponseWriter
		r                  *http.Request
		expectedRespBody   string
		expectedStatusCode int
	}{
		{
			testName: "id is not supplied",
			w:        httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodGet,
				"http://localhost:8080/v1/food",
				nil),
			expectedRespBody:   "id must be supplied",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			testName: "id is not integer",
			w:        httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodGet,
				"http://localhost:8080/v1/food?id=pp",
				nil),
			expectedRespBody:   "id must be integer",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			testName:   "unable to get row from usecase",
			usecaseErr: commonerr.New(commonerr.ErrMsgInternal, commonerr.ErrInternal),
			w:          httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodGet,
				"http://localhost:8080/v1/food?id=99",
				nil),
			expectedRespBody:   commonerr.ErrMsgInternal,
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			testName:   "row is not found",
			usecaseErr: commonerr.New(commonerr.ErrMsgFoodNotFound, commonerr.ErrFoodNotFound),
			w:          httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodGet,
				"http://localhost:8080/v1/food?id=99",
				nil),
			expectedRespBody:   commonerr.ErrMsgFoodNotFound,
			expectedStatusCode: http.StatusNotFound,
		},
		{
			testName:   "handler is invoked successfully",
			usecaseErr: nil,
			w:          httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodGet,
				"http://localhost:8080/v1/food?id=99",
				nil),
			expectedRespBody:   string(mockRow),
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			mockUsecase := mockusecases.NewMockUsecase(
				mockusecases.WithExpectedErr(tc.usecaseErr),
			)
			handler := handlers.NewHandler(mockUsecase)

			handler.GetFood(tc.w, tc.r)
			resp := tc.w.(*httptest.ResponseRecorder)
			tests.AssertEqual(t, strings.TrimSuffix(resp.Body.String(), "\n"), tc.expectedRespBody)
			tests.AssertEqual(t, resp.Code, tc.expectedStatusCode)
		})
	}
}

func Test_ListFood(t *testing.T) {
	mockRows, err := json.Marshal([]food.ListFoodRow{
		{
			ID:           99,
			Name:         tests.MockPGText().String,
			Type:         tests.MockPGText().String,
			IntakeStatus: tests.MockPGText().String,
			Feeder:       tests.MockPGText().String,
			Location:     tests.MockPGText().String,
			Remarks:      tests.MockPGText().String,
			CreatedAt:    tests.MockPGTimestamptz().Time,
			UpdatedAt:    tests.MockPGTimestamptz().Time,
		},
	})
	if err != nil {
		t.Errorf("failed to serialize %v", err)
	}

	testCases := []struct {
		testName           string
		usecaseErr         error
		w                  http.ResponseWriter
		r                  *http.Request
		expectedRespBody   string
		expectedStatusCode int
	}{
		{
			testName: "start and end timestamps are not supplied in URL param",
			w:        httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodGet,
				"http://localhost:8080/v1/foodlist",
				nil),
			expectedRespBody:   "start and end timestamps must be supplied",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			testName: "only one of start and end timestamps are supplied",
			w:        httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodGet,
				"http://localhost:8080/v1/foodlist?startTimestamp=2021-02-20T08:04:05Z",
				nil),
			expectedRespBody:   "start and end timestamps must be supplied",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			testName: "start and end timestamps are not in time.Time type",
			w:        httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodGet,
				"http://localhost:8080/v1/foodlist?startTimestamp=2021-02-20&endTimestamp=2021-02-25",
				nil),
			expectedRespBody:   "start and end timestamps must be in time format",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			testName:   "unable to get rows from usecase",
			usecaseErr: commonerr.New(commonerr.ErrMsgInternal, commonerr.ErrInternal),
			w:          httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodGet,
				"http://localhost:8080/v1/foodlist?startTimestamp=2021-02-20T08:04:05Z&endTimestamp=2025-02-25T08:20:56Z",
				nil),
			expectedRespBody:   commonerr.ErrMsgInternal,
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			testName:   "handler is invoked successfully",
			usecaseErr: nil,
			w:          httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodGet,
				"http://localhost:8080/v1/foodlist?startTimestamp=2021-02-20T08:04:05Z&endTimestamp=2025-02-25T08:20:56Z",
				nil),
			expectedRespBody:   string(mockRows),
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			mockUsecase := mockusecases.NewMockUsecase(
				mockusecases.WithExpectedErr(tc.usecaseErr),
			)
			handler := handlers.NewHandler(mockUsecase)

			handler.ListFood(tc.w, tc.r)
			resp := tc.w.(*httptest.ResponseRecorder)
			tests.AssertEqual(t, strings.TrimSuffix(resp.Body.String(), "\n"), tc.expectedRespBody)
			tests.AssertEqual(t, resp.Code, tc.expectedStatusCode)
		})
	}
}
