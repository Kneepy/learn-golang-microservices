package domain

import "context"

type User struct {
    ID        string `json:"id"`
    Name      string `json:"name"`
    Email     string `json:"email"`
    CreatedAt int64  `json:"created_at"`
    Password  string `json:"-"`
}

type UserRepo interface {
    Create(ctx context.Context, user *User) (*User, error)
}

type UserService interface {
    CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error)
}

type CreateUserRequest struct {
    Password string
    Name     string
    Email    string
}
