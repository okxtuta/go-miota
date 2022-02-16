package math_test

import (
	"testing"

	"github.com/okxtuta/go-miota/math"
	"github.com/stretchr/testify/assert"
)

func TestAbsInt64(t *testing.T) {
	assert.EqualValues(t, int64(9223372036854775807), math.AbsInt64(-9223372036854775807))
}
