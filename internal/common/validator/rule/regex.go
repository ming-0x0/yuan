package rule

import (
	"errors"
	"regexp"
)

var (
	emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
)

type email struct {
	value string
}

func (e *email) Validate() error {
	if !emailRegex.MatchString(e.value) {
		return errors.New("invalid email format")
	}
	return nil
}

// IsEmail returns a rule that validates if a string is a valid email address.
func IsEmail(value string) *email {
	return &email{value: value}
}
