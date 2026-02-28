package storage

import (
	"errors"
	"fmt"
)

var (
	ErrBadEngine = errors.New("unsupported database engine")
	ErrBadMethod = errors.New("unsupported storage method")
	ErrAff = errors.New("error when affecting row(s)")
)

func NewErrBadEngine(engine string) error {
	return fmt.Errorf("%w: %s", ErrBadEngine, engine)
}

func NewErrBadMethod(method string) error {
	return fmt.Errorf("%w: %s", ErrBadMethod, method)
}

func NewErrAff(message string) error {
	return fmt.Errorf("%w: %s", ErrAff, message)
}
