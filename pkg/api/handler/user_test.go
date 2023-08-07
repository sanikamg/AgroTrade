package handler

import (
	"encoding/json"
	"golang_project_ecommerce/pkg/common/response"
	"golang_project_ecommerce/pkg/domain"
	mock "golang_project_ecommerce/pkg/mock/userUsecaseMock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func CreatingMock(t *testing.T) (*UserHandler, *mock.MockUserUsecase) {
	//creating a ctrl for mocking
	newMock := gomock.NewController(t)
	mockUsecase := mock.NewMockUserUsecase(newMock)
	UserHandler := NewUserhandler(mockUsecase)

	return UserHandler, mockUsecase
}

func TestRegister(t *testing.T) {

	UserHandler, mockUsecase := CreatingMock(t)

	usr := domain.Users{
		User_Id:      1,
		Username:     "sanikamg",
		Name:         "sanika",
		Phone:        "6282099388",
		Email:        "sanika@gmail.com",
		BlockStatus:  false,
		Password:     "sanika",
		Verification: true,
	}

	//creating testcase
	testdata := []struct {
		name       string
		data       string
		beforeTest func(Usecase *mock.MockUserUsecase)
		response   response.Response
	}{
		{
			name: "register",
			beforeTest: func(Usecase *mock.MockUserUsecase) {
				Usecase.EXPECT().Register(gomock.Any(), gomock.Any()).
					Return(usr, nil)

			},

			data: `{
				"username":"jerin",
				"name":"jerin",
				"phone":"9447426879",
				"email":"jerin@gmail.com",
				"password":"jerin@123"
			 }`,

			response: response.Response{
				StatusCode: 200,
				Message:    "Registration completed please login",
				Errors:     nil,
				Data:       usr,
			},
		},
	}

	for _, testcase := range testdata {
		t.Run(testcase.name, func(t *testing.T) {
			router := gin.Default()
			router.POST("/signup/register", UserHandler.Register)

			req, _ := http.NewRequest("POST", "/signup/register", strings.NewReader(testcase.data))
			req.Header.Set("Content-Type", "application/json")

			respRecorder := httptest.NewRecorder()
			testcase.beforeTest(mockUsecase)

			router.ServeHTTP(respRecorder, req)

			var actual response.Response
			json.Unmarshal(respRecorder.Body.Bytes(), &actual)

			assert.Equal(t, testcase.response.StatusCode, actual.StatusCode)
			assert.Equal(t, testcase.response.Message, actual.Message)
			assert.Equal(t, testcase.response.Errors, actual.Errors)

		})
	}

}
