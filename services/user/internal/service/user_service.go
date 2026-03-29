package service

import (
    "context"
    "fmt"
    "learning_golang/user/internal/domain"
    "logger"
    "regexp"
    "strings"
    "time"

    "github.com/google/uuid"
    "golang.org/x/crypto/bcrypt"
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
    if err := u.validateCreateUserRequest(req); err != nil {
        return nil, domain.ErrCreateUser
    }

    exists, err := u.repo.CheckEmailExists(ctx, req.Email)
    if err != nil {
        return nil, fmt.Errorf("Failed to check email existence: %v", err)
    }

    if exists {
        return nil, domain.ErrUserAlreadyExists
    }

    hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, fmt.Errorf("failed to hash password: %w", err)
    }

    now := time.Now().Unix()
    user := &domain.User{
        ID:        uuid.New().String(),
        Email:     req.Email,
        Name:      req.Name,
        Password:  string(hashPassword),
        Status:    domain.USER_STATUS_INACTIVE,
        CreatedAt: now,
    }

    if err := u.repo.Create(ctx, user); err != nil {
        return nil, fmt.Errorf("failed to create user: %w", err)
    }

    u.logger.Info("User created successfully")

    return user, nil
}

func (u *UserService) validateCreateUserRequest(req *domain.CreateUserRequest) error {
    if req.Email == "" || !isValidEmail(req.Email) {
        return domain.ErrInvalidEmail
    }

    if req.Name == "" || len(req.Name) < 2 || len(req.Name) > 100 {
        return domain.ErrInvalidName
    }

    if req.Password == "" || len(req.Password) < 8 || len(req.Password) > 72 {
        return domain.ErrInvalidPassword
    }

    return nil
}

func isValidEmail(email string) bool {
    emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
    return emailRegex.MatchString(strings.ToLower(email))
}
