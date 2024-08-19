package service

import (
	pb "auth_service/genproto/user"
	"auth_service/storage/postgres"
	"context"
	"fmt"
	"log/slog"
)

type UserManagementService interface {
	Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error)
	Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error)
	Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.Message, error)
	GetUserProfile(ctx context.Context, req *pb.Void) (*pb.ProfileResponse, error)
	UpdateUserProfile(ctx context.Context, req *pb.ProfileUpdateRequest) (*pb.ProfileUpdateResponse, error)
	ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.Message, error)
	ForgotPassword(ctx context.Context, req *pb.ForgotPasswordRequest) (*pb.Message, error)
	ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.Message, error)
	GetByUserId(ctx context.Context, req *pb.GetByUserIdRequest) (*pb.GetByUserIdResponse, error)
	DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error)
}

type UserService struct {
	pb.UnimplementedUsersServer
	authRepo postgres.UserManagementRepository
	Logger   *slog.Logger
}

func NewUserService(authRepo postgres.UserManagementRepository, logger *slog.Logger) pb.UsersServer {
	return &UserService{authRepo: authRepo, Logger: logger}
}

func (s *UserService) GetByUserId(ctx context.Context, req *pb.GetByUserIdRequest) (*pb.GetByUserIdResponse, error) {
	return s.authRepo.GetByUserId(ctx, req)
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return s.authRepo.DeleteUser(ctx, req)
}

func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	resp, err := s.authRepo.Register(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return resp, nil
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	fmt.Println("...................", req.Email, req.Password)
	resp, err := s.authRepo.Login(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return resp, nil
}

func (s *UserService) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.Message, error) {
	resp, err := s.authRepo.Logout(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}

	return resp, nil
}

func (s *UserService) GetUserProfile(ctx context.Context, req *pb.Void) (*pb.ProfileResponse, error) {
	resp, err := s.authRepo.GetUserProfile(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return resp, nil
}

func (s *UserService) UpdateUserProfile(ctx context.Context, req *pb.ProfileUpdateRequest,) (*pb.ProfileUpdateResponse, error) {
	resp, err := s.authRepo.UpdateUserProfile(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return resp, nil
}

func (s *UserService) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.Message, error) {
	resp, err := s.authRepo.ChangePassword(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return resp, nil
}

func (s *UserService) ForgotPassword(ctx context.Context, req *pb.ForgotPasswordRequest) (*pb.Message, error) {
	resp, err := s.authRepo.ForgotPassword(ctx, req)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return resp, nil
}

func (s *UserService) ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.Message, error) {
	resp, err := s.authRepo.ResetPassword(ctx, req)
	fmt.Println(req.Email, req.NewPassword)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return resp, nil
}
