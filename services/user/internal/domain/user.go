package domain

import "context"

type User struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	CreatedAt int64      `json:"created_at"`
	Password  string     `json:"-"`
	Status    UserStatus `json:"status"`
}

type UserStatus int32

const (
	USER_STATUS_INACTIVE UserStatus = 0
	USER_STATUS_ACTIVE   UserStatus = 1
)

type UserRepo interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Update(ctx context.Context, user *User) error
	UpdateStatus(ctx context.Context, id string, status int) error
	CheckEmailExists(ctx context.Context, email string) (bool, error)
	Search(ctx context.Context, query string, limit int, offset int) ([]*User, error)
}

type UserService interface {
	CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error)
}

type CreateUserRequest struct {
	Password string
	Name     string
	Email    string
}
