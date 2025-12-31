package validator

import (
	"errors"
)

type Rule interface {
	// Validate checks the rule and returns an error if it fails.
	Validate() error
}

type Validator struct {
	err error

	// lastAssertFailed indicates whether the most recent Assert call failed.
	// It is used to decide if Yield() is allowed to override the error.
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
		// A previous assertion already failed; do not allow Yield() to override anymore.
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

// Yield overrides the error message of the most recently failed assertion.
func (v *Validator) Yield(msg string) *Validator {
	if !v.lastAssertFailed || v.err == nil {
		return v
	}

	v.err = errors.New(msg)
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
