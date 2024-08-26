package service

import (
	"fmt"
)

var (
	ErrCannotParseToken    = fmt.Errorf("cannot parse token")
	ErrCannotSignToken     = fmt.Errorf("cannot sign token")
	ErrCannotGenerateToken = fmt.Errorf("cannot generate token")

	ErrCannotCreateUser  = fmt.Errorf("cannot create user")
	ErrCannotGetUser     = fmt.Errorf("cannot get user")
	ErrUserAlreadyExists = fmt.Errorf("user already exists")
	ErrPasswordRequired  = fmt.Errorf("password is required")
	ErrUsernameRequired  = fmt.Errorf("username is required")

	ErrCannotCheckMistakes = fmt.Errorf("cannot check mistakes")
	ErrTextRequired        = fmt.Errorf("text is required")
)
