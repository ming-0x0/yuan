package user

import (
	"github.com/ming-0x0/yuan/internal/common/validator"
	"github.com/ming-0x0/yuan/internal/common/validator/rule"
)

type User struct {
	id           int64
	email        string
	username     string
	hashPassword string
}

func New(
	id int64,
	email string,
	username string,
	hashPassword string,
) *User {
	user := &User{
		id:           id,
		email:        email,
		username:     username,
		hashPassword: hashPassword,
	}

	if err := user.validate(); err != nil {
		return nil
	}

	return user
}

func (u *User) ID() int64 {
	return u.id
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Username() string {
	return u.username
}

func (u *User) validate() error {
	return validator.New().
		Assert(rule.Required(u.email)).Message("invalid email").
		Assert(rule.Required(u.username)).Message("invalid username").
		Assert(rule.Required(u.hashPassword)).Message("invalid hash password").
		Err()
}
