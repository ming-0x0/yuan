package validator

import (
	"errors"

	domainerrors "github.com/ming-0x0/yuan/internal/common/domain/errors"
)

type Rule interface {
	// Validate checks the rule and returns an error if it fails.
	Validate() error
}

type Validator struct {
	err error

	// lastAssertFailed indicates whether the most recent Assert call failed.
	// It is used to decide if Message() is allowed to override the error.
	lastAssertFailed bool
}

// Assert validates the given rule immediately.
//
// It follows a fail-fast strategy:
//   - If a previous assertion has already failed, the rule is skipped.
//   - If the rule fails, the validator stores the error and marks the last assertion as failed.
//
// This method supports fluent chaining.
func (v *Validator) Assert(rule Rule) *Validator {
	if v.err != nil {
		// A previous assertion already failed; do not allow Message() to override anymore.
		v.lastAssertFailed = false
		return v
	}

	v.err = rule.Validate()
	v.lastAssertFailed = v.err != nil
	return v
}

// Error implements the error interface.
//
// It returns an empty string if there is no validation error.
func (v *Validator) Error() string {
	if v == nil || v.err == nil {
		return ""
	}
	return v.err.Error()
}

// Message overrides the error message of the most recently failed assertion.
//
// It only takes effect when:
//   - The last Assert call failed
//   - And the current error has not already been wrapped as a DomainError
//
// The error is wrapped as a domain InvalidArgument error.
func (v *Validator) Message(msg string) *Validator {
	if !v.lastAssertFailed || v.err == nil {
		return v
	}

	// Avoid double-wrapping if Message() is called multiple times
	// or the rule already returned a DomainError.
	var domainErr *domainerrors.DomainError
	if errors.As(v.err, &domainErr) {
		return v
	}

	v.err = domainerrors.Wrap(domainerrors.InvalidArgument, errors.New(msg))
	v.lastAssertFailed = true
	return v
}

// Err returns the underlying validation error, if any.
func (v *Validator) Err() error {
	if v == nil {
		return nil
	}
	return v.err
}

// New creates a new Validator instance.
func New() *Validator {
	return &Validator{}
}
