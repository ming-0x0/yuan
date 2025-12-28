package user

import (
	"github.com/ming-0x0/yuan/adapter/repository/postgres/bun/internal"
	"github.com/ming-0x0/yuan/internal/user/domain/user"
	userDTO "github.com/ming-0x0/yuan/internal/user/dto/user"
)

func ToDomain(model *internal.User) (*user.User, error) {
	return user.New(model.ID, model.Email, model.Username, model.Password)
}

func ToDTO(model *internal.User) (userDTO.User, error) {
	return userDTO.User{
		ID:       model.ID,
		Email:    model.Email,
		Username: model.Username,
	}, nil
}

func ToModel(domain *user.User) (*internal.User, error) {
	return &internal.User{
		ID:       domain.ID(),
		Email:    domain.Email(),
		Username: domain.Username(),
		Password: domain.HashPassword(),
	}, nil
}
