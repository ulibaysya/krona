package storage

import (
	"errors"
	"fmt"
)

var (
	ErrBadEngine = errors.New("unsupported database engine")
	ErrBadType = errors.New("unsupported storage type")
	ErrAff = errors.New("error when affecting row(s)")
)

func NewErrBadEngine(engine string) error {
	return fmt.Errorf("%w: %s", ErrBadEngine, engine)
}

func NewErrBadType(method string) error {
	return fmt.Errorf("%w: %s", ErrBadType, method)
}

func NewErrAff(message string) error {
	return fmt.Errorf("%w: %s", ErrAff, message)
}
