package user

import (
	userCmd "github.com/ming-0x0/yuan/internal/user/command/user"
	userQuery "github.com/ming-0x0/yuan/internal/user/query/user"
)

type Module struct {
	Command Command
	Query   Query
}

type Command struct {
	CreateUser userCmd.CreateUserCommand
}

type Query struct {
	GetUserByID userQuery.GetUserByIDQuery
}
