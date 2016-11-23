package users

import (
	infra "github.com/g0dgarden/ddd-api/infrastructures"
)

type Repository interface {
	GetUser(exec infra.Executor, userId int64) (*User, error)
	GetUsers(exec infra.Executor) ([]User, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetUser(exec infra.Executor, id int64) (*User, error) {
	user := &User{}
	if err := exec.SelectOne(user, "select * from users where id = ?", id); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) GetUsers(exec infra.Executor) ([]User, error) {
	users := []User{}
	_, err := exec.Select(&users, "select * from users")
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repository) Create(exec infra.Executor, user *User) error {
	return exec.Insert(user)
}
