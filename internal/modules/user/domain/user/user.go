package user

import "github.com/godruoyi/go-snowflake"

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
) (User, error) {
	return &user{
		ID:             snowflake.ID(),
		FullName:       fullName,
		Email:          email,
		Username:       username,
		HashedPassword: hashedPassword,
		Status:         active,
		IsReceiveEmail: false,
	}, nil
}

func (u *user) CanLogin() bool {
	return u.Status == active
}

func (u *user) Deactivate() {
	u.Status = inactive
}
