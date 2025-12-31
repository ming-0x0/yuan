package rule

import (
	"fmt"
	"unicode/utf8"
)

type maxLength struct {
	str string
	max int
}

func (l *maxLength) Validate() error {
	if utf8.RuneCountInString(l.str) > l.max {
		return fmt.Errorf("length must be less than or equal to %d", l.max)
	}
	return nil
}

// MaxLength returns a validator that ensures the string length is less than or equal to the specified maximum.
func MaxLength(str string, max int) *maxLength {
	return &maxLength{
		str: str,
		max: max,
	}
}

type minLength struct {
	str string
	min int
}

func (l *minLength) Validate() error {
	if utf8.RuneCountInString(l.str) < l.min {
		return fmt.Errorf("length must be greater than or equal to %d", l.min)
	}
	return nil
}

// MinLength returns a validator that ensures the string length is greater than or equal to the specified minimum.
func MinLength(str string, min int) *minLength {
	return &minLength{
		str: str,
		min: min,
	}
}
