package command

import (
	"context"

	"github.com/ming-0x0/yuan/internal/common/decorator/command"
	"github.com/ming-0x0/yuan/internal/user/domain/user"
	"github.com/ming-0x0/yuan/pkg/logger"
)

type CreateUser struct {
	user *user.User
}

type CreateUserCommand command.Command[CreateUser]

type createUserCommand struct {
	userRepo user.UserRepositoryInterface
}

func NewCreateUserCommand(
	userRepo user.UserRepositoryInterface,
	logger logger.Logger,
) CreateUserCommand {
	return command.Execute(
		createUserCommand{userRepo: userRepo},
		logger,
	)
}

func (c createUserCommand) Handle(ctx context.Context, cmd CreateUser) error {
	return c.userRepo.Create(ctx, cmd.user, cmd.user.ID())
}
