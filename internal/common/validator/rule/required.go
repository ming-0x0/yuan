package rule

import (
	"errors"
	"strings"
	"unicode/utf8"
)

type required struct {
	value string
}

func (r *required) Validate() error {
	if utf8.RuneCountInString(strings.TrimSpace(r.value)) == 0 {
		return errors.New("field is required")
	}
	return nil
}

func Required(value string) *required {
	return &required{
		value: value,
	}
}
