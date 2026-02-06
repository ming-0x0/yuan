package grpc

import (
	"context"

	iamv1 "github.com/ming-0x0/yuan/pkg/proto/iam/v1"
	"github.com/ming-0x0/yuan/internal/iam/application/user"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserServer struct {
	iamv1.UnimplementedUserServiceServer
	userSvc user.UserService
}

func NewUserServer(userSvc user.UserService) iamv1.UserServiceServer {
	return &UserServer{
		userSvc: userSvc,
	}
}

func (s *UserServer) GetProfile(ctx context.Context, req *emptypb.Empty) (*iamv1.GetProfileResponse, error) {
	profile, err := s.userSvc.GetProfile(ctx)
	if err != nil {
		return nil, err
	}

	return &iamv1.GetProfileResponse{
		Id:        profile.ID.String(),
		Email:     profile.Email,
		FullName:  profile.FullName,
	}, nil
}

func (s *UserServer) UpdateProfile(ctx context.Context, req *iamv1.UpdateProfileRequest) (*iamv1.UpdateProfileResponse, error) {
	err := s.userSvc.UpdateProfile(ctx, req.FullName)
	if err != nil {
		return nil, err
	}

	return &iamv1.UpdateProfileResponse{
		Message: "Profile updated successfully",
	}, nil
}
