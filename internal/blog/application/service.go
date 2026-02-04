package application

import (
	"context"
	"errors"
	"time"

	"github.com/ming-0x0/yuan/internal/blog/domain"
	"github.com/ming-0x0/yuan/internal/common/domain/id"
)

type BlogService interface {
	Create(ctx context.Context, title, content string, authorID id.ID) error
	Update(ctx context.Context, id id.ID, title, content string, authorID id.ID) error
	Delete(ctx context.Context, id id.ID, authorID id.ID) error
	Get(ctx context.Context, id id.ID) (*domain.Blog, error)
	List(ctx context.Context, page, limit int) ([]*domain.Blog, error)
}

type blogService struct {
	repo domain.Repository
}

func NewBlogService(repo domain.Repository) BlogService {
	return &blogService{
		repo: repo,
	}
}

func (s *blogService) Create(ctx context.Context, title, content string, authorID id.ID) error {
	blog := &domain.Blog{
		ID:        id.MustNew(),
		Title:     title,
		Content:   content,
		AuthorID:  authorID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return s.repo.Create(ctx, blog)
}

func (s *blogService) Update(ctx context.Context, id id.ID, title, content string, authorID id.ID) error {
	blog, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if blog.AuthorID != authorID {
		return errors.New("unauthorized")
	}

	if title != "" {
		blog.Title = title
	}
	if content != "" {
		blog.Content = content
	}

	return s.repo.Update(ctx, blog)
}

func (s *blogService) Delete(ctx context.Context, id id.ID, authorID id.ID) error {
	blog, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if blog.AuthorID != authorID {
		return errors.New("unauthorized")
	}

	return s.repo.Delete(ctx, id)
}

func (s *blogService) Get(ctx context.Context, id id.ID) (*domain.Blog, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *blogService) List(ctx context.Context, page, limit int) ([]*domain.Blog, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit
	return s.repo.List(ctx, limit, offset)
}
