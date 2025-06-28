package handlers_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"

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
