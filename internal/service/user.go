package service

type UserStorage interface{

}

type UserService struct{
	userStorage UserStorage
}

func NewUserService(userStorage UserStorage) *UserService{
	return &UserService{
		userStorage: userStorage,
	}
}