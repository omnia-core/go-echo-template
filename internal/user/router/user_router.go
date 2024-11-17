package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/omnia-core/go-echo-template/domain"
)

type UserRouter struct {
	userUsecase domain.UserUsecase
}

func NewUserRouter(echo *echo.Echo, userUsecase domain.UserUsecase) *UserRouter {
	userRouter := &UserRouter{
		userUsecase: userUsecase,
	}

	echo.GET("/users/:id", userRouter.GetUser)
	echo.POST("/users", userRouter.CreateUser)
	echo.PUT("/users/:id", userRouter.UpdateUser)
	echo.DELETE("/users/:id", userRouter.DeleteUser)

	return userRouter
}

// CreateUser
// @Tags Users
// @Summary Create User
// @Description This API is used to create a user.
// @Param request body domain.CreateUserRequest true "request"
// @Success 201 "User created successfully"
// @Security BearerAuth
// @Router /users [post]
func (u UserRouter) CreateUser(c echo.Context) error {
	var request domain.CreateUserRequest
	if err := c.Bind(&request); err != nil {
		return err
	}

	if err := u.userUsecase.CreateUser(request); err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

// GetUser
// @Tags Users
// @Summary Get User
// @Description This API is used to get a user.
// @Param id path int true "User ID"
// @Success 200 {object} domain.GetUserResponse
// @Security BearerAuth
// @Router /users/{id} [get]
func (u UserRouter) GetUser(c echo.Context) error {
	var request domain.GetUserRequest
	if err := c.Bind(&request); err != nil {
		return err
	}

	user, err := u.userUsecase.GetUser(request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

// UpdateUser
// @Tags Users
// @Summary Update User
// @Description This API is used to update a user.
// @Param id path int true "User ID"
// @Param request body domain.UpdateUserRequest true "request"
// @Success 202 "User updated successfully"
// @Security BearerAuth
// @Router /users/{id} [put]
func (u UserRouter) UpdateUser(c echo.Context) error {
	var request domain.UpdateUserRequest
	if err := c.Bind(&request); err != nil {
		return err
	}

	if err := u.userUsecase.UpdateUser(request); err != nil {
		return err
	}

	return c.NoContent(http.StatusAccepted)
}

// DeleteUser
// @Tags Users
// @Summary Delete User
// @Description This API is used to delete a user.
// @Param id path int true "User ID"
// @Success 202 "User deleted successfully"
// @Security BearerAuth
// @Router /users/{id} [delete]
func (u UserRouter) DeleteUser(c echo.Context) error {
	var request domain.DeleteUserRequest
	if err := c.Bind(&request); err != nil {
		return err
	}

	if err := u.userUsecase.DeleteUser(request); err != nil {
		return err
	}

	return c.NoContent(http.StatusAccepted)
}
