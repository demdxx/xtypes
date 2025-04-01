package xtypes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAny(t *testing.T) {
	t.Run("general", func(t *testing.T) {
		v := &Any{nil}

		assert.True(t, v.IsNil(), "Expected IsNil to be true, got false")
		assert.True(t, v.IsEmpty(), "Expected IsEmpty to be true, got false")
		assert.Nil(t, v.Get(), "Expected Get to return nil, got non-nil value")
		assert.Equal(t, "default", v.GetOr("default"), "Expected GetOr to return default value, got non-default value")

		assert.NoError(t, v.Set("test"), "Expected Set to not return an error")
		assert.False(t, v.IsNil(), "Expected IsNil to be false, got true")
		assert.False(t, v.IsEmpty(), "Expected IsEmpty to be false, got true")
		assert.Equal(t, "test", v.Get(), "Expected Get to return 'test', got non-matching value")
		assert.Equal(t, "test", v.GetOr("default"), "Expected GetOr to return 'test', got non-matching value")

		assert.NoError(t, v.Set(100), "Expected Set to not return an error")
		assert.Equal(t, 100, v.Get(), "Expected Get to return 100, got non-matching value")
		assert.Equal(t, 100, v.GetOr("default"), "Expected GetOr to return 100, got non-matching value")
		assert.True(t, v.Bool(), "Expected Bool to be true, got false")
		assert.Equal(t, 100, v.Int(), "Expected Int to return 100, got non-matching value")
		assert.Equal(t, int64(100), v.Int64(), "Expected IntOr to return 100, got non-matching value")
		assert.Equal(t, uint(100), v.Uint(), "Expected Uint to return 100, got non-matching value")
		assert.Equal(t, uint64(100), v.Uint64(), "Expected Uint64 to return 100, got non-matching value")
		assert.Equal(t, float64(100), v.Float64(), "Expected Float64 to return 100, got non-matching value")
		assert.Equal(t, "100", v.String(), "Expected String to return '100', got non-matching value")

		if assert.NoError(t, v.Set(3.14), "Expected Set to not return an error") {
			// TryBool
			{
				vl, err := v.TryBool()
				assert.NoError(t, err, "Expected TryBool to not return an error")
				assert.True(t, vl, "Expected TryBool to be true, got false")
			}
			// TryInt
			{
				vl, err := v.TryInt()
				assert.NoError(t, err, "Expected TryInt to not return an error")
				assert.Equal(t, 3, vl, "Expected TryInt to return 3, got non-matching value")
			}
			// TryInt64
			{
				vl, err := v.TryInt64()
				assert.NoError(t, err, "Expected TryInt64 to not return an error")
				assert.Equal(t, int64(3), vl, "Expected TryInt64 to return 3, got non-matching value")
			}
			// TryUint
			{
				vl, err := v.TryUint()
				assert.NoError(t, err, "Expected TryUint to not return an error")
				assert.Equal(t, uint(3), vl, "Expected TryUint to return 3, got non-matching value")
			}
			// TryUint64
			{
				vl, err := v.TryUint64()
				assert.NoError(t, err, "Expected TryUint64 to not return an error")
				assert.Equal(t, uint64(3), vl, "Expected TryUint64 to return 3, got non-matching value")
			}
			// TryFloat64
			{
				vl, err := v.TryFloat64()
				assert.NoError(t, err, "Expected TryFloat64 to not return an error")
				assert.Equal(t, 3.14, vl, "Expected TryFloat64 to return 3.14, got non-matching value")
			}
			// TryString
			{
				vl, err := v.TryString()
				assert.NoError(t, err, "Expected TryString to not return an error")
				assert.Equal(t, "3.14", vl, "Expected TryString to return '3.14', got non-matching value")
			}
		}
	})

	t.Run("nil", func(t *testing.T) {
		assert.Nil(t, (*Any)(nil).Get(), "Expected nil Any instance to be nil")

		assert.EqualError(t, (*Any)(nil).Set("test"), ErrBoxContainerIsNil.Error(), "Expected Set to return an error when called on nil Any instance")

		_, err := (*Any)(nil).TryBool()
		assert.EqualError(t, err, ErrBoxContainerIsNil.Error(), "Expected TryBool to return an error when called on nil Any instance")

		_, err = (*Any)(nil).TryInt()
		assert.EqualError(t, err, ErrBoxContainerIsNil.Error(), "Expected TryInt to return an error when called on nil Any instance")

		_, err = (*Any)(nil).TryString()
		assert.EqualError(t, err, ErrBoxContainerIsNil.Error(), "Expected TryString to return an error when called on nil Any instance")
	})
}
