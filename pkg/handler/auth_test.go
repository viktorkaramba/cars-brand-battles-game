package handler

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	carsBrandsBattle "github.com/viktorkaramba/cars-brand-random-generator-app"
	"github.com/viktorkaramba/cars-brand-random-generator-app/pkg/service"
	mock_service "github.com/viktorkaramba/cars-brand-random-generator-app/pkg/service/mocks"
	"net/http/httptest"
	"testing"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, user carsBrandsBattle.User)

	testTable := []struct {
		name                 string
		inputBody            string
		inputUser            carsBrandsBattle.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"Test", "username":"test", "password": "qwerty"}`,
			inputUser: carsBrandsBattle.User{
				Name:     "Test",
				Username: "test",
				Password: "qwerty",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user carsBrandsBattle.User) {
				s.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "Empty Fields",
			inputBody:            `{"username":"test", "password": "qwerty"}`,
			mockBehavior:         func(s *mock_service.MockAuthorization, user carsBrandsBattle.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"Message":"invalid input body"}`,
		},
		{
			name:      "Service Failure",
			inputBody: `{"name":"Test", "username":"test", "password": "qwerty"}`,
			inputUser: carsBrandsBattle.User{
				Name:     "Test",
				Username: "test",
				Password: "qwerty",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user carsBrandsBattle.User) {
				s.EXPECT().CreateUser(user).Return(1, errors.New("service failure"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"Message":"service failure"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.inputUser)
			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
