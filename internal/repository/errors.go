package repository

import "fmt"

var (
	ErrCannotCheckUserExist = fmt.Errorf("cannot check user exists")
	ErrUserAlreadyExists    = fmt.Errorf("user already exists")
	ErrCannotCreateUser     = fmt.Errorf("cannot create user")
	ErrCannotGetUser        = fmt.Errorf("cannot get user")

	ErrCannotCreateNote = fmt.Errorf("cannot create note")
	ErrCannotGetNotex   = fmt.Errorf("cannot get notex")
)
