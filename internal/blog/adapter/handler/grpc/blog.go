package grpc

import (
	"context"

	blogv1 "github.com/ming-0x0/yuan/pkg/proto/blog/v1"
	"github.com/ming-0x0/yuan/internal/blog/application"
	"google.golang.org/protobuf/types/known/emptypb"
)

type BlogServer struct {
	blogv1.UnimplementedBlogServiceServer
	blogSvc application.BlogService
}

func NewBlogServer(blogSvc application.BlogService) blogv1.BlogServiceServer {
	return &BlogServer{
		blogSvc: blogSvc,
	}
}

func (s *BlogServer) GetBlog(ctx context.Context, req *blogv1.GetBlogRequest) (*blogv1.Blog, error) {
	blog, err := s.blogSvc.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &blogv1.Blog{
		Id:         blog.ID.String(),
		Title:      blog.Title,
		Content:    blog.Content,
		AuthorId:   blog.AuthorID.String(),
		CreatedAt:  blog.CreatedAt.String(),
		UpdatedAt:  blog.UpdatedAt.String(),
	}, nil
}

func (s *BlogServer) ListBlogs(ctx context.Context, req *blogv1.ListBlogsRequest) (*blogv1.ListBlogsResponse, error) {
	blogs, totalPages, totalItems, currentPage, err := s.blogSvc.List(ctx, int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, err
	}

	var pbBlogs []*blogv1.Blog
	for _, b := range blogs {
		pbBlogs = append(pbBlogs, &blogv1.Blog{
			Id:         b.ID.String(),
			Title:      b.Title,
			Content:    b.Content,
			AuthorId:   b.AuthorID.String(),
			CreatedAt:  b.CreatedAt.String(),
			UpdatedAt:  b.UpdatedAt.String(),
		})
	}

	return &blogv1.ListBlogsResponse{
		Blogs:        pbBlogs,
		TotalPages:   int32(totalPages),
		TotalItems:   int32(totalItems),
		CurrentPage:  int32(currentPage),
	}, nil
}

func (s *BlogServer) CreateBlog(ctx context.Context, req *blogv1.CreateBlogRequest) (*blogv1.Blog, error) {
	blog, err := s.blogSvc.Create(ctx, req.Title, req.Content)
	if err != nil {
		return nil, err
	}

	return &blogv1.Blog{
		Id:         blog.ID.String(),
		Title:      blog.Title,
		Content:    blog.Content,
		AuthorId:   blog.AuthorID.String(),
		CreatedAt:  blog.CreatedAt.String(),
		UpdatedAt:  blog.UpdatedAt.String(),
	}, nil
}

func (s *BlogServer) UpdateBlog(ctx context.Context, req *blogv1.UpdateBlogRequest) (*blogv1.Blog, error) {
	blog, err := s.blogSvc.Update(ctx, req.Id, req.Title, req.Content)
	if err != nil {
		return nil, err
	}

	return &blogv1.Blog{
		Id:         blog.ID.String(),
		Title:      blog.Title,
		Content:    blog.Content,
		AuthorId:   blog.AuthorID.String(),
		CreatedAt:  blog.CreatedAt.String(),
		UpdatedAt:  blog.UpdatedAt.String(),
	}, nil
}

func (s *BlogServer) DeleteBlog(ctx context.Context, req *blogv1.DeleteBlogRequest) (*emptypb.Empty, error) {
	err := s.blogSvc.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
