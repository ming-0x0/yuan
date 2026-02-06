package application

import (
	"context"
	"time"

	"github.com/ming-0x0/yuan/internal/blog/domain"
	"github.com/ming-0x0/yuan/internal/common/domain/id"
)

type BlogService interface {
	Create(ctx context.Context, title, content string) (*domain.Blog, error)
	Update(ctx context.Context, id string, title, content string) (*domain.Blog, error)
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*domain.Blog, error)
	List(ctx context.Context, page, pageSize int) ([]*domain.Blog, int, int, int, error)
}

type blogService struct {
	repo domain.Repository
}

func NewBlogService(repo domain.Repository) BlogService {
	return &blogService{
		repo: repo,
	}
}

func (s *blogService) Create(ctx context.Context, title, content string) (*domain.Blog, error) {
	// TODO: Get authorID from context
	authorID := id.MustNew() // Temporary

	blog := &domain.Blog{
		ID:        id.MustNew(),
		Title:     title,
		Content:   content,
		AuthorID:  authorID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.repo.Create(ctx, blog)
	if err != nil {
		return nil, err
	}

	return blog, nil
}

func (s *blogService) Update(ctx context.Context, idStr string, title, content string) (*domain.Blog, error) {
	blogID, err := id.Parse(idStr)
	if err != nil {
		return nil, err
	}

	blog, err := s.repo.FindByID(ctx, blogID)
	if err != nil {
		return nil, err
	}

	// TODO: Get authorID from context
	// authorID := id.MustNew() // Temporary

	if title != "" {
		blog.Title = title
	}
	if content != "" {
		blog.Content = content
	}
	blog.UpdatedAt = time.Now()

	err = s.repo.Update(ctx, blog)
	if err != nil {
		return nil, err
	}

	return blog, nil
}

func (s *blogService) Delete(ctx context.Context, idStr string) error {
	blogID, err := id.Parse(idStr)
	if err != nil {
		return err
	}

	// TODO: Get authorID from context
	// authorID := id.MustNew() // Temporary

	return s.repo.Delete(ctx, blogID)
}

func (s *blogService) Get(ctx context.Context, idStr string) (*domain.Blog, error) {
	blogID, err := id.Parse(idStr)
	if err != nil {
		return nil, err
	}

	return s.repo.FindByID(ctx, blogID)
}

func (s *blogService) List(ctx context.Context, page, pageSize int) ([]*domain.Blog, int, int, int, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	blogs, err := s.repo.List(ctx, pageSize, offset)
	if err != nil {
		return nil, 0, 0, 0, err
	}

	totalItems := len(blogs)
	totalPages := (totalItems + pageSize - 1) / pageSize

	return blogs, totalPages, totalItems, page, nil
}
