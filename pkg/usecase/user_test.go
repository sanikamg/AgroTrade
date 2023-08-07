package usecase

import (
	"context"
	"golang_project_ecommerce/pkg/domain"
	mock "golang_project_ecommerce/pkg/mock/userRepoMock"
	interfaces "golang_project_ecommerce/pkg/usecase/interface"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func CreatingMock(t *testing.T) (interfaces.UserUsecase, *mock.MockUserRepository) {
	//creating a ctrl for mocking
	newMock := gomock.NewController(t)
	mockRepo := mock.NewMockUserRepository(newMock)
	UserUsecase := NewUserUsecase(mockRepo)

	return UserUsecase, mockRepo
}

func TestRegister(t *testing.T) {
	//creating a mock repo and Usecase
	UserUsecase, mockRepo := CreatingMock(t)

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
	hash, _ := bcrypt.GenerateFromPassword([]byte(usr.Password), 14)
	usr.Password = string(hash)

	//creating testcase
	testdata := []struct {
		name        string
		user        domain.Users
		beforeTest  func(userRepo *mock.MockUserRepository)
		expectedErr error
	}{
		{
			name: "register",
			user: domain.Users{
				User_Id:      1,
				Username:     "sanikamg",
				Name:         "sanika",
				Phone:        "6282099388",
				Email:        "sanika@gmail.com",
				BlockStatus:  false,
				Password:     "sanika",
				Verification: true,
			},
			beforeTest: func(userRepo *mock.MockUserRepository) {
				userRepo.EXPECT().FindStatus(gomock.Any(), gomock.Any()).
					Return(usr, nil)
				userRepo.EXPECT().UpdateUserDetails(gomock.Any(), gomock.Any()).
					Return(usr, nil)

			},
			expectedErr: nil,
		},
	}

	for _, testcase := range testdata {
		t.Run(testcase.name, func(t *testing.T) {
			testcase.beforeTest(mockRepo)
			usrResult, err := UserUsecase.Register(context.Background(), testcase.user)
			assert.Equal(t, testcase.expectedErr, err)
			assert.Equal(t, usr, usrResult)
		})
	}

}

func TestLogin(t *testing.T) {
	UserUsecase, mockRepo := CreatingMock(t)

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

	hash, _ := bcrypt.GenerateFromPassword([]byte(usr.Password), 14)
	usr.Password = string(hash)

	testdata := []struct {
		name        string
		user        domain.Users
		beforeTest  func(userRepo *mock.MockUserRepository)
		expectedErr error
	}{
		{
			name: "Login",
			user: domain.Users{

				Username: "sanikamg",

				Password: "sanika",
			},
			beforeTest: func(userRepo *mock.MockUserRepository) {
				userRepo.EXPECT().FindUser(gomock.Any(), gomock.Any()).
					Return(usr, nil)
			},
			expectedErr: nil,
		},
	}

	for _, testcase := range testdata {
		t.Run(testcase.name, func(t *testing.T) {
			testcase.beforeTest(mockRepo)
			_, err := UserUsecase.Login(context.Background(), testcase.user)
			assert.Equal(t, testcase.expectedErr, err)

		})
	}

}
