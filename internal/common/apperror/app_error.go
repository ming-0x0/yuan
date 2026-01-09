package apperror

import "errors"

//go:generate go tool stringer -type=ErrorCode

// ErrorCode represents a classified application-level error code.
// It is transport-agnostic and can be mapped to HTTP / gRPC status later.
type ErrorCode int

const (
	// OK indicates no error.
	OK ErrorCode = iota

	// Canceled indicates the operation was canceled.
	Canceled

	// Unknown indicates an unknown error occurred.
	Unknown

	// InvalidArgument indicates invalid input was provided.
	InvalidArgument

	// DeadlineExceeded indicates the operation exceeded its deadline.
	DeadlineExceeded

	// NotFound indicates the requested resource was not found.
	NotFound

	// AlreadyExists indicates the resource already exists.
	AlreadyExists

	// PermissionDenied indicates permission was denied.
	PermissionDenied

	// ResourceExhausted indicates resource exhaustion.
	ResourceExhausted

	// FailedPrecondition indicates the system is in an invalid state.
	FailedPrecondition

	// Aborted indicates the operation was aborted.
	Aborted

	// OutOfRange indicates a value is out of range.
	OutOfRange

	// Unimplemented indicates the operation is not implemented.
	Unimplemented

	// Internal indicates an internal system error.
	Internal

	// Unavailable indicates the service is unavailable.
	Unavailable

	// DataLoss indicates unrecoverable data loss.
	DataLoss

	// Unauthenticated indicates authentication is required.
	Unauthenticated

	// TokenExpired indicates an authentication token has expired.
	TokenExpired

	// InvalidToken indicates an authentication token is invalid.
	InvalidToken
)

// AppError represents an application-level error with a classified error code.
// It wraps an optional underlying cause and is safe to propagate across layers.
type AppError struct {
	errCode ErrorCode
	msg     string
	err     error
}

// WithCause creates a new AppError with the given error code
// and wraps the provided underlying cause.
func WithCause(code ErrorCode, cause error) *AppError {
	return &AppError{
		errCode: code,
		err:     cause,
	}
}

// WithMessage creates a new AppError with a custom message
// and an optional underlying.
func WithMessage(code ErrorCode, msg string, cause error) *AppError {
	return &AppError{
		errCode: code,
		msg:     msg,
		err:     cause,
	}
}

// WithCode creates a new AppError with the given error code only,
// without an underlying cause.
func WithCode(code ErrorCode) *AppError {
	return &AppError{
		errCode: code,
	}
}

// Error implements the error interface.
// It returns the underlying cause's error message if present.
func (e *AppError) Error() string {
	if e == nil {
		return ""
	}
	if e.err != nil {
		return e.err.Error()
	}
	return e.Message()
}

// Unwrap returns the underlying cause, allowing errors.Is / errors.As.
func (e *AppError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.err
}

// Code returns the classified error code.
func (e *AppError) Code() ErrorCode {
	if e == nil {
		return OK
	}
	return e.errCode
}

// Message returns the human-readable error message.
// If no custom message is set, the error code name is used.
func (e *AppError) Message() string {
	if e == nil {
		return ""
	}
	if e.msg != "" {
		return e.msg
	}
	return e.errCode.String()
}

func Is(code ErrorCode, err error) bool {
	if err == nil {
		return false
	}

	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr.Code() == code
	}

	return false
}

func As(err error, target any) bool {
	return errors.As(err, target)
}
