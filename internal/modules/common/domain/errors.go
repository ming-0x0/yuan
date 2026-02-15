package domain

import "errors"

var (
	ErrRequiredField  = errors.New("required field is missing")
	ErrInvalidEmail    = errors.New("invalid email format")
	ErrInvalidInternal = errors.New("internal domain error")
	ErrInvalidFormat   = errors.New("invalid format")
	ErrNotFound        = errors.New("resource not found")
	ErrAlreadyExists   = errors.New("resource already exists")
)
