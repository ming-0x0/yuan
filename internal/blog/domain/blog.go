package domain

import (
	"context"
	"time"

	"github.com/ming-0x0/yuan/internal/common/domain/id"
)

type Blog struct {
	ID        id.ID
	Title     string
	Content   string
	AuthorID  id.ID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Repository interface {
	Create(ctx context.Context, blog *Blog) error
	Update(ctx context.Context, blog *Blog) error
	Delete(ctx context.Context, id id.ID) error
	FindByID(ctx context.Context, id id.ID) (*Blog, error)
	List(ctx context.Context, limit, offset int) ([]*Blog, error)
}
