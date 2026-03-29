package postgres

import (
	"context"
	"learning_golang/user/internal/domain"
	"learning_golang/user/internal/repository/postgres/queries"
	"learning_golang/user/pkg/db"
	"logger"
)

type UserRepo struct {
	db      *db.PostgresDB
	queries *queries.UserQueries
	logger  *logger.Logger
}

func NewUserRepo(db *db.PostgresDB, log *logger.Logger) domain.UserRepo {
	return &UserRepo{
		db:      db,
		queries: queries.GetUserQueries(),
		logger:  log,
	}
}

func (r *UserRepo) GetByID(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User

	err := r.db.QueryRowContext(ctx, r.queries.GetUserByID, id).Scan(&user)

	if err != nil {
		r.logger.Error("Error to get user by id %v", err)
	}

	return &user, nil
}

func (r *UserRepo) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User

	err := r.db.QueryRowContext(ctx, r.queries.GetUserByEmail, email).Scan(&user)

	if err != nil {
		r.logger.Error("Error to get user by email %v", err)
	}

	return &user, nil
}

func (r *UserRepo) Update(ctx context.Context, user *domain.User) error {
	r.logger.Info("Update user %v", user.Email)

	err := r.db.QueryRowContext(ctx, r.queries.UpdateUser, user.ID, user.Name, user.Email).Scan(&user)

	if err != nil {
		r.logger.Error("Failed update user %v: %v", user.Email, err)
		return err
	}

	return nil
}

func (r *UserRepo) UpdateStatus(ctx context.Context, id string, status int) error {
	r.logger.Info("Update status user %v", id)

	result, err := r.db.ExecContext(ctx, r.queries.UpdateUserStatus, id, status)

	if err != nil {
		r.logger.Error("Failed update status user %v: %v", id, err)
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		r.logger.Info("User not found %v", id)
	}

	return nil
}

func (r *UserRepo) CheckEmailExists(ctx context.Context, email string) (bool, error) {
	var exists bool
	err := r.db.QueryRowContext(ctx, r.queries.CheckEmailExist, email).Scan(&exists)
	return exists, err
}

func (r *UserRepo) Search(ctx context.Context, query string, limit int, offset int) ([]*domain.User, error) {
	rows, err := r.db.QueryContext(ctx, r.queries.SearchUsers, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*domain.User, 0)
	for rows.Next() {
		var user domain.User
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Name,
			&user.Status,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (r *UserRepo) Create(ctx context.Context, user *domain.User) error {

	r.logger.Info("Creating user with email %v", user.Email)

	err := r.db.QueryRowContext(ctx, r.queries.CreateUser,
		user.ID,
		user.Email,
		user.Password,
		user.Status,
		user.CreatedAt,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.Status,
		&user.CreatedAt,
	)

	if err != nil {
		r.logger.Error("Error creating user with email %v: %v", user.Email, err)
	}

	r.logger.Info("Created user with email %v in database ", user.Email)

	return nil
}
