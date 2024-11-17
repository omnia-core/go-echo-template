package domain

type User struct {
	ID   uint
	Name string
	Age  uint
}

func (u User) ToGetUserResponse() GetUserResponse {
	return GetUserResponse{
		ID:   u.ID,
		Name: u.Name,
		Age:  u.Age,
	}
}
