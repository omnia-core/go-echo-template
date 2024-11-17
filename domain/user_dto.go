package domain

type CreateUserRequest struct {
	Name string `json:"name" validate:"required"`
	Age  uint   `json:"age" validate:"required"`
}

func (r CreateUserRequest) ToUser() User {
	return User{
		Name: r.Name,
		Age:  r.Age,
	}
}

type GetUserRequest struct {
	ID uint `swaggerignore:"true" param:"id"`
}

type GetUserResponse struct {
	ID   uint   `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
	Age  uint   `json:"age" validate:"required"`
}

type UpdateUserRequest struct {
	ID   uint   `swaggerignore:"true" param:"id"`
	Name string `json:"name" validate:"optional"`
	Age  uint   `json:"age" validate:"optional"`
}

func (r UpdateUserRequest) ToUser() User {
	return User{
		ID:   r.ID,
		Name: r.Name,
		Age:  r.Age,
	}
}

type DeleteUserRequest struct {
	ID uint `swaggerignore:"true" param:"id"`
}
