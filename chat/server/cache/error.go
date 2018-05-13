package cache

import (
	"errors"
)

var (
	ErrUserNotExists = errors.New("user not exists")
	ErrInvalidPasswd = errors.New("Passwd or username not right")
	ErrInvalidParams = errors.New("Invalid params")
	ErrUserExist     = errors.New("user exist")
)
