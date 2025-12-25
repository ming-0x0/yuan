package void_test

import (
	"database/sql/driver"
	"encoding/json"
	"testing"
	"time"

	"github.com/ming-0x0/yuan/pkg/typeutil/void"
	"github.com/stretchr/testify/assert"
)

func TestVoid_New_Set_Get_Unset(t *testing.T) {
	t.Parallel()

	type args struct {
		v void.Void[string]
	}

	testCases := []struct {
		name  string
		given func(t *testing.T, args *args)
		when  func(args *args)
		then  func(t *testing.T, args *args)
	}{
		{
			name: "should create new valid void",
			given: func(t *testing.T, args *args) {
				// No setup needed
			},
			when: func(args *args) {
				args.v = void.New("test")
			},
			then: func(t *testing.T, args *args) {
				assert.False(t, args.v.IsVoid())
				assert.Equal(t, "test", args.v.Get())
				assert.NotNil(t, args.v.Ptr())
				assert.Equal(t, "test", *args.v.Ptr())
			},
		},
		{
			name: "should set value and become valid",
			given: func(t *testing.T, args *args) {
				args.v = void.Void[string]{} // Undefined initially
			},
			when: func(args *args) {
				args.v.Set("updated")
			},
			then: func(t *testing.T, args *args) {
				assert.False(t, args.v.IsVoid())
				assert.Equal(t, "updated", args.v.Get())
			},
		},
		{
			name: "should unset value and become undefined",
			given: func(t *testing.T, args *args) {
				args.v = void.New("initial")
			},
			when: func(args *args) {
				args.v.Unset()
			},
			then: func(t *testing.T, args *args) {
				assert.True(t, args.v.IsVoid())
				assert.Zero(t, args.v.Get())
				assert.Nil(t, args.v.Ptr())
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var args args
			if tc.given != nil {
				tc.given(t, &args)
			}
			if tc.when != nil {
				tc.when(&args)
			}
			if tc.then != nil {
				tc.then(t, &args)
			}
		})
	}
}

func TestVoid_MarshalJSON(t *testing.T) {
	t.Parallel()

	type args struct {
		v    void.Void[string]
		data []byte
		err  error
	}

	testCases := []struct {
		name  string
		given func(t *testing.T, args *args)
		when  func(args *args)
		then  func(t *testing.T, args *args)
	}{
		{
			name: "should marshal valid value",
			given: func(t *testing.T, args *args) {
				args.v = void.New("123")
			},
			when: func(args *args) {
				args.data, args.err = json.Marshal(args.v)
			},
			then: func(t *testing.T, args *args) {
				assert.NoError(t, args.err)
				assert.Equal(t, "\"123\"", string(args.data))
			},
		},
		{
			name: "should marshal undefined value (as zero value)",
			given: func(t *testing.T, args *args) {
				args.v = void.Void[string]{}
			},
			when: func(args *args) {
				args.data, args.err = json.Marshal(args.v)
			},
			then: func(t *testing.T, args *args) {
				assert.NoError(t, args.err)
				// Based on implementation: v.value is "", so it marshals to "\"\""
				assert.Equal(t, "\"\"", string(args.data))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var args args
			if tc.given != nil {
				tc.given(t, &args)
			}
			if tc.when != nil {
				tc.when(&args)
			}
			if tc.then != nil {
				tc.then(t, &args)
			}
		})
	}
}

func TestVoid_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	type args struct {
		jsonStr string
		v       void.Void[string]
		err     error
	}

	testCases := []struct {
		name  string
		given func(t *testing.T, args *args)
		when  func(args *args)
		then  func(t *testing.T, args *args)
	}{
		{
			name: "should unmarshal valid string",
			given: func(t *testing.T, args *args) {
				args.jsonStr = `"hello"`
			},
			when: func(args *args) {
				args.err = json.Unmarshal([]byte(args.jsonStr), &args.v)
			},
			then: func(t *testing.T, args *args) {
				assert.NoError(t, args.err)
				assert.False(t, args.v.IsVoid())
				assert.Equal(t, "hello", args.v.Get())
			},
		},
		{
			name: "should unmarshal null",
			given: func(t *testing.T, args *args) {
				args.jsonStr = `null`
			},
			when: func(args *args) {
				args.err = json.Unmarshal([]byte(args.jsonStr), &args.v)
			},
			then: func(t *testing.T, args *args) {
				assert.NoError(t, args.err)
				assert.False(t, args.v.IsVoid()) // Explicit null sets it to valid(true) with zero value
				assert.Equal(t, "", args.v.Get())
			},
		},
		{
			name: "should error on invalid json",
			given: func(t *testing.T, args *args) {
				args.jsonStr = `invalid`
			},
			when: func(args *args) {
				args.err = json.Unmarshal([]byte(args.jsonStr), &args.v)
			},
			then: func(t *testing.T, args *args) {
				assert.Error(t, args.err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var args args
			if tc.given != nil {
				tc.given(t, &args)
			}
			if tc.when != nil {
				tc.when(&args)
			}
			if tc.then != nil {
				tc.then(t, &args)
			}
		})
	}
}

func TestVoid_UnmarshalText(t *testing.T) {
	t.Parallel()

	type args struct {
		text  []byte
		vTime void.Void[time.Time]
		err   error
	}

	testCases := []struct {
		name  string
		given func(t *testing.T, args *args)
		when  func(args *args)
		then  func(t *testing.T, args *args)
	}{
		{
			name: "should unmarshal text for supported type (time.Time)",
			given: func(t *testing.T, args *args) {
				args.text = []byte("2023-01-01T12:00:00Z")
			},
			when: func(args *args) {
				args.err = args.vTime.UnmarshalText(args.text)
			},
			then: func(t *testing.T, args *args) {
				assert.NoError(t, args.err)
				assert.False(t, args.vTime.IsVoid())
				expected, _ := time.Parse(time.RFC3339, "2023-01-01T12:00:00Z")
				assert.Equal(t, expected.Unix(), args.vTime.Get().Unix())
			},
		},
		{
			name: "should handle empty text",
			given: func(t *testing.T, args *args) {
				args.text = []byte("")
			},
			when: func(args *args) {
				args.err = args.vTime.UnmarshalText(args.text)
			},
			then: func(t *testing.T, args *args) {
				assert.Error(t, args.err)
			},
		},
		{
			name: "should return error when handling invalid text",
			given: func(t *testing.T, args *args) {
				args.text = []byte("invalid")
			},
			when: func(args *args) {
				args.err = args.vTime.UnmarshalText(args.text)
			},
			then: func(t *testing.T, args *args) {
				assert.Error(t, args.err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var args args
			if tc.given != nil {
				tc.given(t, &args)
			}
			if tc.when != nil {
				tc.when(&args)
			}
			if tc.then != nil {
				tc.then(t, &args)
			}
		})
	}
}

func TestVoid_Value(t *testing.T) {
	t.Parallel()

	type args struct {
		v   void.Void[string]
		val driver.Value
		err error
	}

	testCases := []struct {
		name  string
		given func(t *testing.T, args *args)
		when  func(args *args)
		then  func(t *testing.T, args *args)
	}{
		{
			name: "should return nil for undefined",
			given: func(t *testing.T, args *args) {
				args.v = void.Void[string]{}
			},
			when: func(args *args) {
				args.val, args.err = args.v.Value()
			},
			then: func(t *testing.T, args *args) {
				assert.NoError(t, args.err)
				assert.Nil(t, args.val)
			},
		},
		{
			name: "should return value for valid",
			given: func(t *testing.T, args *args) {
				args.v = void.New("foo")
			},
			when: func(args *args) {
				args.val, args.err = args.v.Value()
			},
			then: func(t *testing.T, args *args) {
				assert.NoError(t, args.err)
				assert.Equal(t, "foo", args.val)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var args args
			if tc.given != nil {
				tc.given(t, &args)
			}
			if tc.when != nil {
				tc.when(&args)
			}
			if tc.then != nil {
				tc.then(t, &args)
			}
		})
	}
}

func TestVoid_Scan(t *testing.T) {
	t.Parallel()

	type args struct {
		v   void.Void[string]
		src any
		err error
	}

	testCases := []struct {
		name  string
		given func(t *testing.T, args *args)
		when  func(args *args)
		then  func(t *testing.T, args *args)
	}{
		{
			name: "should scan matching type",
			given: func(t *testing.T, args *args) {
				args.src = "scanned"
			},
			when: func(args *args) {
				args.err = args.v.Scan(args.src)
			},
			then: func(t *testing.T, args *args) {
				assert.NoError(t, args.err)
				assert.False(t, args.v.IsVoid())
				assert.Equal(t, "scanned", args.v.Get())
			},
		},
		{
			name: "should scan nil",
			given: func(t *testing.T, args *args) {
				args.src = nil
			},
			when: func(args *args) {
				args.err = args.v.Scan(args.src)
			},
			then: func(t *testing.T, args *args) {
				assert.NoError(t, args.err)
				assert.False(t, args.v.IsVoid()) // Scan(nil) sets valid=true, value=zero
				assert.Equal(t, "", args.v.Get())
			},
		},
		{
			name: "should scan convertible type (bytes to string)",
			given: func(t *testing.T, args *args) {
				args.src = []byte("bytes")
			},
			when: func(args *args) {
				args.err = args.v.Scan(args.src)
			},
			then: func(t *testing.T, args *args) {
				assert.NoError(t, args.err)
				assert.Equal(t, "bytes", args.v.Get())
			},
		},
		{
			name: "should error on incompatible type",
			given: func(t *testing.T, args *args) {
				args.src = struct{}{}
			},
			when: func(args *args) {
				args.err = args.v.Scan(args.src)
			},
			then: func(t *testing.T, args *args) {
				assert.Error(t, args.err)
				assert.True(t, args.v.IsVoid())
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var args args
			if tc.given != nil {
				tc.given(t, &args)
			}
			if tc.when != nil {
				tc.when(&args)
			}
			if tc.then != nil {
				tc.then(t, &args)
			}
		})
	}
}

func TestVoid_IsZero(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		given any
		then  bool
	}{
		{
			name:  "should return true for valid zero int",
			given: void.New(0),
			then:  true,
		},
		{
			name:  "should return false for valid non-zero int",
			given: void.New(42),
			then:  false,
		},
		{
			name:  "should return false for undefined",
			given: void.Void[int]{},
			then:  false,
		},
		{
			name:  "should return true for valid zero string",
			given: void.New(""),
			then:  true,
		},
		{
			name:  "should return false for valid non-zero string",
			given: void.New("hello"),
			then:  false,
		},
		{
			name:  "should return true for valid nil slice",
			given: void.New([]int(nil)),
			then:  true,
		},
		{
			name:  "should return false for valid empty slice",
			given: void.New([]int{}),
			then:  false,
		},
		{
			name:  "should return true for valid nil map",
			given: void.New(map[string]int(nil)),
			then:  true,
		},
		{
			name:  "should return false for valid empty map",
			given: void.New(map[string]int{}),
			then:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var result bool
			switch v := tc.given.(type) {
			case void.Void[int]:
				result = v.IsZero()
			case void.Void[string]:
				result = v.IsZero()
			case void.Void[[]int]:
				result = v.IsZero()
			case void.Void[map[string]int]:
				result = v.IsZero()
			}
			assert.Equal(t, tc.then, result)
		})
	}
}

func TestVoid_Equal(t *testing.T) {
	t.Parallel()

	type args struct {
		v1    void.Void[int]
		v2    void.Void[int]
		equal bool
	}

	testCases := []struct {
		name  string
		given func(t *testing.T, args *args)
		when  func(args *args)
		then  func(t *testing.T, args *args)
	}{
		{
			name: "should be equal when both undefined",
			given: func(t *testing.T, args *args) {
				args.v1 = void.Void[int]{}
				args.v2 = void.Void[int]{}
			},
			when: func(args *args) {
				args.equal = args.v1.Equal(args.v2)
			},
			then: func(t *testing.T, args *args) {
				assert.True(t, args.equal)
			},
		},
		{
			name: "should not be equal when one valid one invalid",
			given: func(t *testing.T, args *args) {
				args.v1 = void.New(1)
				args.v2 = void.Void[int]{}
			},
			when: func(args *args) {
				args.equal = args.v1.Equal(args.v2)
			},
			then: func(t *testing.T, args *args) {
				assert.False(t, args.equal)
			},
		},
		{
			name: "should be equal when values equal",
			given: func(t *testing.T, args *args) {
				args.v1 = void.New(10)
				args.v2 = void.New(10)
			},
			when: func(args *args) {
				args.equal = args.v1.Equal(args.v2)
			},
			then: func(t *testing.T, args *args) {
				assert.True(t, args.equal)
			},
		},
		{
			name: "should not be equal when values differ",
			given: func(t *testing.T, args *args) {
				args.v1 = void.New(10)
				args.v2 = void.New(20)
			},
			when: func(args *args) {
				args.equal = args.v1.Equal(args.v2)
			},
			then: func(t *testing.T, args *args) {
				assert.False(t, args.equal)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var args args
			if tc.given != nil {
				tc.given(t, &args)
			}
			if tc.when != nil {
				tc.when(&args)
			}
			if tc.then != nil {
				tc.then(t, &args)
			}
		})
	}
}
