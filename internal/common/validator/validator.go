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
}

// Assert validates the given rule immediately.
//
// If the validator already contains an error, the rule is skipped.
// This method follows a fail-fast strategy and supports fluent chaining.
func (v *Validator) Assert(rule Rule) *Validator {
	if v.err != nil {
		return v
	}

	v.err = rule.Validate()
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

// Message overrides the error message when validation has failed.
//
// It wraps the error with a domain InvalidArgument error.
func (v *Validator) Message(msg string) *Validator {
	if v.err == nil {
		return v
	}
	v.err = domainerrors.Wrap(domainerrors.InvalidArgument, errors.New(msg))
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
