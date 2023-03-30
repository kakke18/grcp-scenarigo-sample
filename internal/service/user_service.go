package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/kakke18/grcp-scenarigo-sample/internal/model"
	"github.com/kakke18/grcp-scenarigo-sample/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	_ pb.UserServiceServer = (*UserService)(nil)
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	userDao userDao
}

type userDao interface {
	GetByID(context.Context, string) (*model.User, error)
	List(context.Context, int, int) ([]*model.User, error)
	Insert(context.Context, *model.User) error
	Update(context.Context, *model.User) error
	Delete(context.Context, string) error
}

func NewUserService(userDao userDao) *UserService {
	return &UserService{
		UnimplementedUserServiceServer: pb.UnimplementedUserServiceServer{},
		userDao:                        userDao,
	}
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := s.userDao.GetByID(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, "failed to get user")
	}

	return userModelToPb(user), nil
}

func (s *UserService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	users, err := s.userDao.List(ctx, int(req.GetLimit()), int(req.GetOffset()))
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to list users")
	}

	res := make([]*pb.GetUserResponse, len(users))
	for i, user := range users {
		res[i] = userModelToPb(user)
	}

	return &pb.ListUsersResponse{
		Users: res,
	}, nil
}

func userModelToPb(user *model.User) *pb.GetUserResponse {
	return &pb.GetUserResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user, err := model.NewUser(req.GetName(), req.GetEmail())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid user")
	}

	if err := s.userDao.Insert(ctx, user); err != nil {
		return nil, status.Error(codes.Internal, "failed to create user")
	}

	return &pb.CreateUserResponse{
		Id: user.ID,
	}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	user, err := s.userDao.GetByID(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, "failed to get user")
	}

	if err := user.UpdateName(req.Name); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid name")
	}

	if err := user.UpdateEmail(req.Email); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid email")
	}

	if err := s.userDao.Update(ctx, user); err != nil {
		return nil, status.Error(codes.Internal, "failed to update user")
	}

	return &pb.UpdateUserResponse{
		Id: user.ID,
	}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	if err := s.userDao.Delete(ctx, req.GetId()); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, "failed to delete user")
	}

	return &pb.DeleteUserResponse{}, nil
}
