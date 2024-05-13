package service

import (
	"context"
	"time"

	"github.com/mystpen/Pet-API/internal/dto"
	"github.com/mystpen/Pet-API/pkg"
)

type UserStorage interface {
	CreatUser(context.Context, *dto.RegistrationRequest, []byte) error
}

type UserService struct {
	userStorage UserStorage
}

func NewUserService(userStorage UserStorage) *UserService {
	return &UserService{
		userStorage: userStorage,
	}
}

func (us *UserService) RegisterUser(request *dto.RegistrationRequest) error {
	hash, err := pkg.SetPassword(request.Password)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = us.userStorage.CreatUser(ctx, request, hash)
	if err != nil {
		return err
	}
	return nil
}
