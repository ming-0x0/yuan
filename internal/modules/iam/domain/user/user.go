package user

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/godruoyi/go-snowflake"
)

type status uint8

const (
	active   status = 1
	inactive status = 2
)

type User = *user

type user struct {
	ID             uint64
	FullName       string
	Email          string
	Username       string
	HashedPassword string
	Status         status
	IsReceiveEmail bool
}

func New(
	fullName string,
	email string,
	username string,
	hashedPassword string,
	isReceiveEmail bool,
) (User, error) {
	user := &user{
		ID:             snowflake.ID(),
		FullName:       fullName,
		Email:          email,
		Username:       username,
		HashedPassword: hashedPassword,
		Status:         active,
		IsReceiveEmail: isReceiveEmail,
	}
	if err := user.validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *user) validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.FullName, validation.Required),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Username, validation.Required),
		validation.Field(&u.HashedPassword, validation.Required),
		validation.Field(&u.Status, validation.In(active, inactive)),
	)
}

func (u *user) CanLogin() bool {
	return u.Status == active
}

func (u *user) Deactivate() {
	u.Status = inactive
}
