package golang_generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsSame[T comparable](value1, value2 T) bool {
	return value1 == value2
}

func TestISame(t *testing.T) {
	assert.True(t, IsSame[string]("Albarra", "Albarra"))
	assert.True(t, IsSame[int](1405, 1405))
}
