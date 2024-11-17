package golang_generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Age int

type Number interface {
	~int | int8 | int16 | int32 | int64 | float32 | float64
}

func Min[T Number](first, second T) T {
	if first < second {
		return first
	}

	return second
}

func TestMin(t *testing.T) {
	assert.Equal(t, 100, Min(100, 200))
	assert.Equal(t, Age(30), Min[Age](Age(40), Age(30)))
}
