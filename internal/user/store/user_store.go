package store

import (
	"github.com/omnia-core/go-echo-template/domain"
	"gorm.io/gorm"
)

type userStore struct {
	conn *gorm.DB
}

func NewUserStore(conn *gorm.DB) domain.UserStore {
	return &userStore{conn: conn}
}

func (u userStore) CreateUser(user domain.User) error {
	return u.conn.Create(&user).Error
}

func (u userStore) GetUser(id uint) (domain.User, error) {
	var user domain.User
	err := u.conn.First(&user, id).Error
	return user, err
}

func (u userStore) UpdateUser(user domain.User) error {
	return u.conn.Save(&user).Error
}

func (u userStore) DeleteUser(id uint) error {
	return u.conn.Delete(&domain.User{}, id).Error
}
