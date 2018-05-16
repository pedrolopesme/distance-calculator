package distance_calculator

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCalcKilometersWithNoDistanceBetweenTwoPoints(test *testing.T) {
	dist := CalcKilometers(0.0,0.0,0.0,0.0)
	assert.Equal(test, 0.0, dist)
}

func TestCalcKilometersWithVerySmallDistanceBetweenTwoPoints(test *testing.T) {
	dist := CalcKilometers(0.00001,0.0,0.0,0.0)
	assert.Equal(test, 0.0011111329509752147, dist)
}

func TestCalcKilometersWithVeryLongDistanceBetweenTwoPoints(test *testing.T) {
	dist := CalcKilometers(90.0,90.0,0.0,0.0)
	assert.Equal(test, 10007.037053999999, dist)
}