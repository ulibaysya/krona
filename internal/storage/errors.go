package storage

import (
	"errors"
)

var (
	ErrNoneAffected = errors.New("no objects affected")
	ErrUniqueViolation = errors.New("unique constraint violated")
)
