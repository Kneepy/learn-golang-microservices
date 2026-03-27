package postgres

import (
	"context"
	"learning_golang/user/internal/domain"
	"learning_golang/user/pkg/db"
)

type UserRepo struct {
	db *db.PostgresDB
}

func NewUserRepo(db *db.PostgresDB) domain.UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	return nil, nil
}
