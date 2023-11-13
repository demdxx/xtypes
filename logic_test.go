package xtypes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstVal(t *testing.T) {
	assert.Equal(t, 1, *FirstVal[int](nil, nil, nil, &[]int{1}[0]))
	assert.Nil(t, FirstVal[int]())
}
