package service

import (
	"context"
	"learning_golang/user/internal/domain"
	"logger"
)

type UserService struct {
	repo   domain.UserRepo
	logger *logger.Logger
}

func NewUserService(repo domain.UserRepo, logger *logger.Logger) domain.UserService {
	return &UserService{
		repo:   repo,
		logger: logger,
	}
}

func (u *UserService) CreateUser(ctx context.Context, req *domain.CreateUserRequest) (*domain.User, error) {
	return nil, nil
}
