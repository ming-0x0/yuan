package rule

import (
	"errors"
	"strings"
	"unicode/utf8"
)

type required[V comparable] struct {
	value V
}

func (r *required[V]) Validate() error {
	var (
		zero V
		err  = errors.New("field is required")
	)
	switch v := any(r.value).(type) {
	case string:
		if utf8.RuneCountInString(strings.TrimSpace(v)) == 0 {
			return err
		}
	default:
		if r.value == zero {
			return err
		}
	}
	return nil
}

// Required returns a validator that ensures the value is neither empty nor the zero value.
func Required[V comparable](value V) *required[V] {
	return &required[V]{
		value: value,
	}
}
