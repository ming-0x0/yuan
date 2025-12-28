package query

import (
	"context"

	"github.com/ming-0x0/yuan/internal/common/decorator/query"
	userdto "github.com/ming-0x0/yuan/internal/user/dto/user"
	"github.com/ming-0x0/yuan/pkg/logger"
)

type GetUserByID struct {
	ID int64
}

type GetUserByIDQuery query.Query[GetUserByID, userdto.User]

type getUserByIDQuery struct {
	readModel GetUserByIDReadModel
}

type GetUserByIDReadModel interface {
	FindByID(ctx context.Context, id int64) (userdto.User, error)
}

func NewGetUserByIDQuery(
	readModel GetUserByIDReadModel,
	logger logger.Logger,
) query.Query[GetUserByID, userdto.User] {
	return query.Execute(
		getUserByIDQuery{readModel: readModel},
		logger,
	)
}

func (h getUserByIDQuery) Handle(ctx context.Context, query GetUserByID) (userdto.User, error) {
	return h.readModel.FindByID(ctx, query.ID)
}
