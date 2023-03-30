package dao

import (
	"context"
	"database/sql"
	"time"

	"github.com/kakke18/grcp-scenarigo-sample/internal/model"
)

type UserDao struct{}

func NewUserDao() *UserDao {
	return &UserDao{}
}

var userList = []*model.User{
	{
		ID:        "1",
		Name:      "Alice",
		Email:     "alice@example.com",
		CreatedAt: time.Now(),
	},
	{
		ID:        "2",
		Name:      "Bob",
		Email:     "bob@example.com",
		CreatedAt: time.Now(),
	},
	{
		ID:        "3",
		Name:      "Carol",
		Email:     "carol@example.com",
		CreatedAt: time.Now(),
	},
	{
		ID:        "4",
		Name:      "Dave",
		Email:     "dave@example.com",
		CreatedAt: time.Now(),
	},
	{
		ID:        "5",
		Name:      "Eve",
		Email:     "eve@example.com",
		CreatedAt: time.Now(),
	},
}

func (d *UserDao) GetByID(_ context.Context, id string) (*model.User, error) {
	for _, user := range userList {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, sql.ErrNoRows
}

func (d *UserDao) List(_ context.Context, limit, offset int) ([]*model.User, error) {
	if len(userList) < offset {
		return nil, nil
	}

	if len(userList) < limit+offset {
		return userList[offset:], nil
	}

	return userList[offset : limit+offset], nil
}

func (d *UserDao) Insert(_ context.Context, user *model.User) error {
	userList = append(userList, user)

	return nil
}

func (d *UserDao) Update(_ context.Context, user *model.User) error {
	for i, u := range userList {
		if u.ID == user.ID {
			userList[i] = user
			return nil
		}
	}

	return sql.ErrNoRows
}

func (d *UserDao) Delete(_ context.Context, id string) error {
	for i, user := range userList {
		if user.ID == id {
			userList = append(userList[:i], userList[i+1:]...)
			return nil
		}
	}

	return sql.ErrNoRows
}
