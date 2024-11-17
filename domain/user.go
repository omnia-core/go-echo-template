package domain

type UserStore interface {
	CreateUser(user User) error
	GetUser(id uint) (User, error)
	UpdateUser(user User) error
	DeleteUser(id uint) error
}

type UserUsecase interface {
	CreateUser(request CreateUserRequest) error
	GetUser(request GetUserRequest) (GetUserResponse, error)
	UpdateUser(request UpdateUserRequest) error
	DeleteUser(request DeleteUserRequest) error
}
