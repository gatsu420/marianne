package handlers_test

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/gatsu420/marianne/app/handlers"
	"github.com/gatsu420/marianne/app/usecases/food"
	commonerr "github.com/gatsu420/marianne/common/errors"
	"github.com/stretchr/testify/mock"
)

func (s *testSuite) Test_GetFood() {
	testCases := []struct {
		testName           string
		w                  http.ResponseWriter
		r                  *http.Request
		urlQueryErr        error
		usecaseFood        *food.GetFoodRow
		usecaseErr         error
		expectedRespBody   string
		expectedStatusCode int
	}{
		{
			testName: "id is not supplied",
			w:        httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodGet,
				"http://localhost:8080/v1/food",
				nil),
			urlQueryErr:        errors.New("some error"),
			expectedRespBody:   "id must be supplied",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			testName: "id is not integer",
			w:        httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodGet,
				"http://localhost:8080/v1/food?id=pp",
				nil),
			urlQueryErr:        errors.New("some error"),
			expectedRespBody:   "id must be integer",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			testName: "unable to get food from usecase",
			w:        httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodGet,
				"http://localhost:8080/v1/food?id=99",
				nil),
			urlQueryErr:        nil,
			usecaseFood:        nil,
			usecaseErr:         errors.New("some error"),
			expectedRespBody:   "some error",
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			testName: "food is not found",
			w:        httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodGet,
				"http://localhost:8080/v1/food?id=99",
				nil),
			urlQueryErr:        nil,
			usecaseFood:        nil,
			usecaseErr:         commonerr.New(commonerr.ErrMsgFoodNotFound, commonerr.ErrFoodNotFound),
			expectedRespBody:   commonerr.ErrMsgFoodNotFound,
			expectedStatusCode: http.StatusNotFound,
		},
		{
			testName: "success",
			w:        httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodGet,
				"http://localhost:8080/v1/food?id=99",
				nil),
			urlQueryErr:        nil,
			usecaseFood:        s.dummyGetFoodRow,
			usecaseErr:         nil,
			expectedRespBody:   s.dummyGetFoodRowStr,
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.testName, func() {
			if tc.urlQueryErr == nil {
				s.mockFoodUsecase.EXPECT().GetFood(
					mock.Anything,
					mock.AnythingOfType("int"),
				).Return(tc.usecaseFood, tc.usecaseErr).Once()
				s.handler = handlers.NewHandler(s.mockFoodUsecase)
			}

			s.handler.GetFood(tc.w, tc.r)
			resp := tc.w.(*httptest.ResponseRecorder)
			s.Equal(tc.expectedRespBody, strings.TrimSuffix(resp.Body.String(), "\n"))
			s.Equal(tc.expectedStatusCode, resp.Code)
		})
	}
}

func (s *testSuite) Test_ListFood() {
	dummyFoodRows := []food.ListFoodRow{
		{
			ID:           99,
			Name:         "dummy",
			Type:         "dummy",
			IntakeStatus: "dummy",
			Feeder:       "dummy",
			Location:     "dummy",
			Remarks:      "dummy",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}
	frows, err := json.Marshal(dummyFoodRows)
	if err != nil {
		log.Fatal(err)
	}
	dummyFoodRowsStr := string(frows)

	testCases := []struct {
		testName             string
		w                    http.ResponseWriter
		r                    *http.Request
		timestampURLParamErr error
		usecaseFood          []food.ListFoodRow
		usecaseErr           error
		expectedRespBody     string
		expectedStatusCode   int
	}{
		{
			testName: "start and end timestamps are not supplied in URL param",
			w:        httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodGet,
				"http://localhost:8080/v1/foodlist",
				nil),
			timestampURLParamErr: errors.New("some error"),
			expectedRespBody:     "start and end timestamps must be supplied",
			expectedStatusCode:   http.StatusBadRequest,
		},
		{
			testName: "only one of both start and end timestamps are supplied in URL param",
			w:        httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodGet,
				"http://localhost:8080/v1/foodlist?startTimestamp=2021-02-20T08:04:05Z",
				nil),
			timestampURLParamErr: errors.New("some error"),
			expectedRespBody:     "start and end timestamps must be supplied",
			expectedStatusCode:   http.StatusBadRequest,
		},
		{
			testName: "start and end timestamps are not in time format",
			w:        httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodGet,
				"http://localhost:8080/v1/foodlist?startTimestamp=2021-02-20&endTimestamp=2021-02-25",
				nil),
			timestampURLParamErr: errors.New("some error"),
			expectedRespBody:     "start and end timestamps must be in time format",
			expectedStatusCode:   http.StatusBadRequest,
		},
		{
			testName: "unable to get food list from usecase",
			w:        httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodGet,
				"http://localhost:8080/v1/foodlist?startTimestamp=2021-02-20T08:04:05Z&endTimestamp=2025-02-25T08:20:56Z",
				nil),
			timestampURLParamErr: nil,
			usecaseFood:          nil,
			usecaseErr:           errors.New("something is wrong"),
			expectedRespBody:     "something is wrong",
			expectedStatusCode:   http.StatusInternalServerError,
		},
		{
			testName: "success",
			w:        httptest.NewRecorder(),
			r: httptest.NewRequest(http.MethodGet,
				"http://localhost:8080/v1/foodlist?startTimestamp=2021-02-20T08:04:05Z&endTimestamp=2025-02-25T08:20:56Z",
				nil),
			timestampURLParamErr: nil,
			usecaseFood:          dummyFoodRows,
			expectedRespBody:     dummyFoodRowsStr,
			expectedStatusCode:   http.StatusOK,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.testName, func() {
			if tc.timestampURLParamErr == nil {
				s.mockFoodUsecase.EXPECT().ListFood(
					mock.Anything,
					mock.AnythingOfType("*food.ListFoodArgs"),
				).Return(tc.usecaseFood, tc.usecaseErr).Once()
			}

			s.handler.ListFood(tc.w, tc.r)
			resp := tc.w.(*httptest.ResponseRecorder)
			s.Equal(tc.expectedRespBody, strings.TrimSuffix(resp.Body.String(), "\n"))
			s.Equal(tc.expectedStatusCode, resp.Code)
		})
	}
}
