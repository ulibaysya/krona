package storage

import (
	"errors"
	"fmt"
)

var (
	ErrBadEngine = errors.New("unsupported database engine")
	ErrBadMethod = errors.New("unsupported storage method")
)

func NewBadEngine(engine string) error {
	return fmt.Errorf("%w: %s", ErrBadEngine, engine)
}

func NewBadMethod(method string) error {
	return fmt.Errorf("%w: %s", ErrBadMethod, method)
}
