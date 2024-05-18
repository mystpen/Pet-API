package service

import (
	"context"
	"encoding/base64"
	"time"

	"github.com/mystpen/Pet-API/internal/dto"
	"github.com/mystpen/Pet-API/internal/model"
	"github.com/mystpen/Pet-API/pkg"
	"golang.org/x/crypto/bcrypt"
)

type UserStorage interface {
	CreatUser(context.Context, *dto.RegistrationRequest, []byte) error
	GetUserByEmail(context.Context, string) (*model.User, error)
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

func (us *UserService) GetRegisteredUser(req *dto.LogInRequest) (*model.User, error) { //TODO 
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	user, err := us.userStorage.GetUserByEmail(ctx, req.Email)
	if err != nil{
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(*req.Password))
	if err != nil{
		return nil, model.ErrNoMatch
	}
	return user, nil
}

func (us *UserService) CreateToken(user *model.User) string {
	auth := user.UserName + ":" + user.PlainPassword
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
