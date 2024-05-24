package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mystpen/Pet-API/internal/dto"
	mock_service "github.com/mystpen/Pet-API/internal/service/mock"
	"github.com/stretchr/testify/assert"
)

var password string = "1234556@!"

func TestUserService_RegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)

	userStorageMock := mock_service.NewMockUserStorage(ctrl)

	userService := NewUserService(userStorageMock)

	testCases := []struct {
		name     string
		req      *dto.RegistrationRequest
		mockFunc func()
		err      error
	}{
		{
			name: "Correct data",
			req: &dto.RegistrationRequest{
				UserName: "testUser",
				Email:    "test@mail.com",
				Password: &password,
			},
			mockFunc: func() {
				req := &dto.RegistrationRequest{
					UserName: "testUser",
					Email:    "test@mail.com",
					Password: &password,
				}
				//generatedPass, _ := bcrypt.GenerateFromPassword([]byte(*&password), 12)
				userStorageMock.EXPECT().CreatUser(gomock.Any(), req, gomock.Any()).Return(nil)
			},
			err: nil,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.mockFunc()
			err := userService.RegisterUser(test.req)
			assert.Equal(t, test.err, err)
		})
	}

}
