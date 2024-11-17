package usecase

import "github.com/omnia-core/go-echo-template/domain"

type userUsecase struct {
	userStore domain.UserStore
}

func NewUserUsecase(userStore domain.UserStore) domain.UserUsecase {
	return &userUsecase{
		userStore: userStore,
	}
}

func (u userUsecase) CreateUser(request domain.CreateUserRequest) error {
	return u.userStore.CreateUser(request.ToUser())
}

func (u userUsecase) GetUser(request domain.GetUserRequest) (domain.GetUserResponse, error) {
	user, err := u.userStore.GetUser(request.ID)
	if err != nil {
		return domain.GetUserResponse{}, err
	}
	return user.ToGetUserResponse(), nil
}

func (u userUsecase) UpdateUser(request domain.UpdateUserRequest) error {
	return u.userStore.UpdateUser(request.ToUser())
}

func (u userUsecase) DeleteUser(request domain.DeleteUserRequest) error {
	return u.userStore.DeleteUser(request.ID)
}
