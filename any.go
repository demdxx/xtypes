package xtypes

import (
	"errors"

	"github.com/demdxx/gocast/v2"
)

var (
	ErrBoxContainerIsNil = errors.New("container is nil")
)

// Any is a type alias for any.
type Any struct {
	Val any
}

// IsNil checks if the value is nil.
// It returns true if the value is nil, otherwise false.
func (a *Any) IsNil() bool {
	return a == nil || a.Val == nil
}

// IsEmpty checks if the value is empty.
// It uses the gocast.IsEmpty function to determine if the value is empty.
// It returns true if the value is empty, otherwise false.
// It also checks if the Any instance itself is nil.
func (a *Any) IsEmpty() bool {
	return a == nil || gocast.IsEmpty(a.Val)
}

// Set sets the value of the Any instance.
// It returns an error if the Any instance is nil.
// Otherwise, it sets the value and returns nil.
func (a *Any) Set(val any) error {
	if a == nil {
		return ErrBoxContainerIsNil
	}
	a.Val = val
	return nil
}

// Get retrieves the value of the Any instance.
// It returns nil if the Any instance is nil.
// Otherwise, it returns the value.
func (a *Any) Get() any {
	if a == nil {
		return nil
	}
	return a.Val
}

// GetOr retrieves the value of the Any instance or returns a default value.
// It returns the value if the Any instance is not nil, otherwise it returns the default value.
func (a *Any) GetOr(def any) any {
	if a == nil || a.Val == nil {
		return def
	}
	return a.Val
}

// Bool retrieves the value of the Any instance as a boolean.
// It uses the gocast.Bool function to convert the value to a boolean.
// It returns false if the Any instance is nil or if the value cannot be converted to a boolean.
// Otherwise, it returns the converted boolean value.
func (a *Any) Bool() bool {
	return gocast.Bool(a.Get())
}

// TryBool attempts to retrieve the value of the Any instance as a boolean.
// It uses the gocast.TryCast function to convert the value to a boolean.
// It returns an error if the Any instance is nil or if the value cannot be converted to a boolean.
// Otherwise, it returns the converted boolean value.
// It also returns an error if the conversion fails.
func (a *Any) TryBool() (bool, error) {
	if a == nil {
		return false, ErrBoxContainerIsNil
	}
	return gocast.TryCast[bool](a.Val)
}

// Int retrieves the value of the Any instance as an integer.
// It uses the gocast.Number function to convert the value to an integer.
// It returns 0 if the Any instance is nil or if the value cannot be converted to an integer.
// Otherwise, it returns the converted integer value.
// It also returns an error if the conversion fails.
func (a *Any) Int() int {
	return gocast.Number[int](a.Get())
}

// TryInt attempts to retrieve the value of the Any instance as an integer.
// It uses the gocast.TryNumber function to convert the value to an integer.
// It returns an error if the Any instance is nil or if the value cannot be converted to an integer.
// Otherwise, it returns the converted integer value.
// It also returns an error if the conversion fails.
func (a *Any) TryInt() (int, error) {
	return tryNumber[int](a)
}

// Int64 retrieves the value of the Any instance as a 64-bit integer.
// It uses the gocast.Number function to convert the value to a 64-bit integer.
// It returns 0 if the Any instance is nil or if the value cannot be converted to a 64-bit integer.
// Otherwise, it returns the converted 64-bit integer value.
func (a *Any) Int64() int64 {
	return gocast.Number[int64](a.Get())
}

// TryInt64 attempts to retrieve the value of the Any instance as a 64-bit integer.
// It uses the gocast.TryNumber function to convert the value to a 64-bit integer.
// It returns an error if the Any instance is nil or if the value cannot be converted to a 64-bit integer.
// Otherwise, it returns the converted 64-bit integer value.
// It also returns an error if the conversion fails.
func (a *Any) TryInt64() (int64, error) {
	return tryNumber[int64](a)
}

// Uint retrieves the value of the Any instance as an unsigned integer.
// It uses the gocast.Number function to convert the value to an unsigned integer.
// It returns 0 if the Any instance is nil or if the value cannot be converted to an unsigned integer.
// Otherwise, it returns the converted unsigned integer value.
func (a *Any) Uint() uint {
	return gocast.Number[uint](a.Get())
}

// TryUint attempts to retrieve the value of the Any instance as an unsigned integer.
func (a *Any) TryUint() (uint, error) {
	return tryNumber[uint](a)
}

// Uint64 retrieves the value of the Any instance as a 64-bit unsigned integer.
// It uses the gocast.Number function to convert the value to a 64-bit unsigned integer.
// It returns 0 if the Any instance is nil or if the value cannot be converted to a 64-bit unsigned integer.
// Otherwise, it returns the converted 64-bit unsigned integer value.
func (a *Any) Uint64() uint64 {
	return gocast.Number[uint64](a.Get())
}

// TryUint64 attempts to retrieve the value of the Any instance as a 64-bit unsigned integer.
// It uses the gocast.TryNumber function to convert the value to a 64-bit unsigned integer.
// It returns an error if the Any instance is nil or if the value cannot be converted to a 64-bit unsigned integer.
// Otherwise, it returns the converted 64-bit unsigned integer value.
// It also returns an error if the conversion fails.
func (a *Any) TryUint64() (uint64, error) {
	return tryNumber[uint64](a)
}

// Float64 retrieves the value of the Any instance as a 64-bit floating-point number.
// It uses the gocast.Number function to convert the value to a 64-bit floating-point number.
// It returns 0 if the Any instance is nil or if the value cannot be converted to a 64-bit floating-point number.
// Otherwise, it returns the converted 64-bit floating-point number.
func (a *Any) Float64() float64 {
	return gocast.Number[float64](a.Get())
}

// TryFloat64 attempts to retrieve the value of the Any instance as a 64-bit floating-point number.
// It uses the gocast.TryNumber function to convert the value to a 64-bit floating-point number.
// It returns an error if the Any instance is nil or if the value cannot be converted to a 64-bit floating-point number.
// Otherwise, it returns the converted 64-bit floating-point number.
// It also returns an error if the conversion fails.
func (a *Any) TryFloat64() (float64, error) {
	return tryNumber[float64](a)
}

// String retrieves the value of the Any instance as a string.
// It uses the gocast.String function to convert the value to a string.
// It returns an empty string if the Any instance is nil or if the value cannot be converted to a string.
// Otherwise, it returns the converted string value.
func (a *Any) String() string {
	return gocast.Str(a.Get())
}

// TryString attempts to retrieve the value of the Any instance as a string.
// It uses the gocast.TryStr function to convert the value to a string.
// It returns an error if the Any instance is nil or if the value cannot be converted to a string.
// Otherwise, it returns the converted string value.
// It also returns an error if the conversion fails.
func (a *Any) TryString() (string, error) {
	if a == nil {
		return "", ErrBoxContainerIsNil
	}
	return gocast.TryStr(a.Val)
}

func tryNumber[T gocast.Numeric](a *Any) (T, error) {
	if a == nil {
		return 0, ErrBoxContainerIsNil
	}
	return gocast.TryNumber[T](a.Val)
}
