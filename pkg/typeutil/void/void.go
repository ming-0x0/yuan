package void

import (
	"database/sql"
	"database/sql/driver"
	"encoding"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

// Verify interface implementations
var (
	_ json.Marshaler           = (*Void[any])(nil)
	_ json.Unmarshaler         = (*Void[any])(nil)
	_ encoding.TextUnmarshaler = (*Void[any])(nil)
	_ driver.Valuer            = (*Void[any])(nil)
	_ sql.Scanner              = (*Void[any])(nil)
)

// Void is a generic wrapper type that can represent a value that may be explicitly
// void or unset, which is particularly useful for:
//   - JSON marshaling/unmarshaling where fields can be omitted
//   - Database operations where NULL values need to be distinguished from zero values
//
// Supported types for V include:
//   - Basic types
//   - time.Time for timestamp handling
//   - []byte for binary data
//   - nil for explicit NULL values in databases
//
// The zero value of Void[V] is considered void (valid = false).
type Void[V any] struct {
	value V
	valid bool
}

func New[V any](value V) Void[V] {
	return Void[V]{
		value: value,
		valid: true,
	}
}

func (v Void[V]) Get() V {
	return v.value
}

func (v *Void[V]) Set(value V) {
	v.value = value
	v.valid = true
}

func (v *Void[V]) Unset() {
	var val V
	v.value = val
	v.valid = false
}

// UnmarshalJSON implements json.Unmarshaler
func (v *Void[V]) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &v.value); err != nil {
		return err
	}

	v.valid = true
	return nil
}

// MarshalJSON implements json.Marshaler
func (v Void[V]) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(v.value)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UnmarshalText implements encoding.TextUnmarshaler
func (v *Void[V]) UnmarshalText(text []byte) error {
	v.valid = len(text) > 0
	if textUnmarshaler, ok := any(&v.value).(encoding.TextUnmarshaler); ok {
		if err := textUnmarshaler.UnmarshalText(text); err != nil {
			return err
		}
		v.valid = true
		return nil
	}

	return errors.New("Void: cannot unmarshal text: underlying value doesn't implement encoding.TextUnmarshaler")
}

// Value implements driver.Valuer
func (v Void[V]) Value() (driver.Value, error) {
	if !v.valid {
		return nil, nil
	}

	if valuer, ok := any(v.value).(driver.Valuer); ok {
		v, err := valuer.Value()
		return v, err
	}
	return v.value, nil
}

// Scan implements sql.Scanner
func (v *Void[V]) Scan(src any) error {
	v.valid = true

	if src == nil {
		var zero V
		v.value = zero
		return nil
	}

	if val, ok := src.(V); ok {
		v.value = val
		return nil
	}

	typ := reflect.TypeFor[V]()
	val := reflect.ValueOf(src)

	if val.Type().ConvertibleTo(typ) {
		v.value = val.Convert(typ).Interface().(V)
		return nil
	}

	v.valid = false
	return fmt.Errorf("Void: Scan() incompatible types (src: %T, dst: %T)", src, v.value)
}

func (v Void[V]) IsVoid() bool {
	return !v.valid
}

func (v Void[V]) IsZero() bool {
	return v.valid && reflect.ValueOf(v.value).IsZero()
}

func (v Void[V]) Ptr() *V {
	if v.valid {
		return &v.value
	}

	return nil
}

func (v Void[V]) Equal(other Void[V]) bool {
	if v.valid != other.valid {
		return false
	}
	if !v.valid {
		return true
	}

	return reflect.DeepEqual(v.value, other.value)
}
