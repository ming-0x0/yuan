package postgres

import (
	"context"
	"time"

	"github.com/ming-0x0/yuan/internal/blog/domain"
	"github.com/ming-0x0/yuan/internal/common/domain/id"
	"github.com/ming-0x0/yuan/pkg/logger"
	"github.com/ming-0x0/yuan/pkg/orm/bun/client"
	"github.com/uptrace/bun"
)

type blogRepository struct {
	client *client.Client
	logger logger.Logger
}

func New(client *client.Client, logger logger.Logger) *blogRepository {
	return &blogRepository{
		client: client,
		logger: logger,
	}
}

type Blog struct {
	bun.BaseModel `bun:"table:blogs,alias:b"`

	ID        int64     `bun:"id,pk"`
	Title     string    `bun:"title"`
	Content   string    `bun:"content"`
	AuthorID  int64     `bun:"author_id"`
	CreatedAt time.Time `bun:"created_at"`
	UpdatedAt time.Time `bun:"updated_at"`
}

func toInternal(b *domain.Blog) *Blog {
	return &Blog{
		ID:        b.ID.Int64(),
		Title:     b.Title,
		Content:   b.Content,
		AuthorID:  b.AuthorID.Int64(),
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}
}

func toDomain(b *Blog) *domain.Blog {
	return &domain.Blog{
		ID:        id.FromInt64(b.ID),
		Title:     b.Title,
		Content:   b.Content,
		AuthorID:  id.FromInt64(b.AuthorID),
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}
}

func (r *blogRepository) Create(ctx context.Context, blog *domain.Blog) error {
	b := toInternal(blog)
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()

	_, err := r.client.DB(ctx).NewInsert().Model(b).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *blogRepository) Update(ctx context.Context, blog *domain.Blog) error {
	b := toInternal(blog)
	b.UpdatedAt = time.Now()

	_, err := r.client.DB(ctx).NewUpdate().Model(b).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *blogRepository) Delete(ctx context.Context, id id.ID) error {
	_, err := r.client.DB(ctx).NewDelete().Model((*Blog)(nil)).Where("id = ?", id.Int64()).Exec(ctx)
	return err
}

func (r *blogRepository) FindByID(ctx context.Context, id id.ID) (*domain.Blog, error) {
	b := new(Blog)
	err := r.client.DB(ctx).NewSelect().Model(b).Where("id = ?", id.Int64()).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return toDomain(b), nil
}

func (r *blogRepository) List(ctx context.Context, limit, offset int) ([]*domain.Blog, error) {
	var blogs []*Blog
	err := r.client.DB(ctx).NewSelect().Model(&blogs).Limit(limit).Offset(offset).Order("created_at DESC").Scan(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]*domain.Blog, len(blogs))
	for i, b := range blogs {
		res[i] = toDomain(b)
	}
	return res, nil
}
