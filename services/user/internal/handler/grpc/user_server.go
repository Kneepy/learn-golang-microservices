package grpc

import (
	"context"
	gen "gen/user"
	"learning_golang/user/internal/domain"
	"logger"
)

type UserServer struct {
	gen.UnimplementedUserServiceServer
	service domain.UserService
	logger  *logger.Logger
}

func NewUserServer(service domain.UserService, log *logger.Logger) *UserServer {
	return &UserServer{
		logger:  log,
		service: service,
	}
}

func (s *UserServer) CreateUser(ctx context.Context, req *gen.CreateUserRequest) (*gen.CreateUserResponse, error) {
	s.logger.Info("CreateUser called with: %v", req)

	domainReq := &domain.CreateUserRequest{
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
	}

	user, err := s.service.CreateUser(ctx, domainReq)

	if err != nil {
		s.logger.Error("CreateUser returned error: %v", err)

		return nil, err
	}

	return &gen.CreateUserResponse{
		User: &gen.User{
			Id:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			Status:    1, // это надо на статус поменять как сделаю чета
		},
	}, nil
}
