package golang_generic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestLength(t *testing.T) {
	result := Length[string]("Albarra")

	assert.Equal(t, "Albarra", result)

	resultNumber := Length[int](1405)
	assert.Equal(t, resultNumber, 1405)
}

func TestSample(t *testing.T) {
	assert.True(t, true)
}
