package domain

import "errors"

var (
    ErrCreateUser        = errors.New("create user error: invalid argument")
    ErrUserAlreadyExists = errors.New("user already exists")
    ErrInvalidEmail      = errors.New("invalid email")
    ErrInvalidName       = errors.New("the name must be at least two characters long")
    ErrInvalidPassword   = errors.New("the password must be longer than 8 and shorter than 72 characters")
)
