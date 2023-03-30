package model

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
)

var (
	ErrInvalidName  = errors.New("invalid name")
	ErrInvalidEmail = errors.New("invalid email")
)

type User struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
}

func NewUser(name, email string) (*User, error) {
	u := &User{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}

	if err := validateUser(u); err != nil {
		return nil, err
	}

	return u, nil
}

func validateUser(user *User) error {
	nameLen := utf8.RuneCountInString(user.Name)
	if nameLen < 1 || nameLen > 50 {
		return ErrInvalidName
	}
	if !strings.Contains(user.Email, "@") {
		return ErrInvalidEmail
	}

	return nil
}

func (u *User) UpdateName(name string) error {
	u.Name = name

	return validateUser(u)
}

func (u *User) UpdateEmail(email string) error {
	u.Email = email

	return validateUser(u)
}
