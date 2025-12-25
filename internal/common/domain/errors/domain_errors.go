package domainerrors

//go:generate go tool stringer -type=ErrorCode
type ErrorCode int32

const (
	OK ErrorCode = iota
	Canceled
	Unknown
	InvalidArgument
	DeadlineExceeded
	NotFound
	AlreadyExists
	PermissionDenied
	ResourceExhausted
	FailedPrecondition
	Aborted
	OutOfRange
	Unimplemented
	Internal
	Unavailable
	DataLoss
	Unauthenticated
	TokenExpired
	InvalidToken
)

type DomainError struct {
	errCode ErrorCode
	msg     string
	err     error
}

func Wrap(
	errCode ErrorCode,
	err error,
) *DomainError {
	return &DomainError{
		errCode: errCode,
		err:     err,
	}
}

func Wrapf(
	errCode ErrorCode,
	msg string,
	err error,
) *DomainError {
	return &DomainError{
		errCode: errCode,
		msg:     msg,
		err:     err,
	}
}

func Wrapc(
	errCode ErrorCode,
) *DomainError {
	return Wrap(errCode, nil)
}

func (e *DomainError) Error() string {
	if e == nil || e.err == nil {
		return ""
	}

	return e.err.Error()
}

func (e *DomainError) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.err
}

func (e *DomainError) ErrorCode() ErrorCode {
	if e == nil {
		return OK
	}

	return e.errCode
}

func (e *DomainError) Message() string {
	if e == nil {
		return ""
	}

	if e.msg == "" {
		return e.errCode.String()
	}

	return e.msg
}
