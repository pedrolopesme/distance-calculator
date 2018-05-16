package distance_calculator

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCalcKilometersWithNoDistanceBetweenTwoPoints(test *testing.T) {
	dist := CalcKilometers(1.0,2.0,1.0,2.0)
	assert.Equal(test, 0, dist)
}