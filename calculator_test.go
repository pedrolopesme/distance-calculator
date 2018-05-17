package distance_calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcMetersWithNoDistanceBetweenTwoPoints(test *testing.T) {
	dist := CalcMeters(0.0, 0.0, 0.0, 0.0)
	assert.Equal(test, 0.0, dist)
}

func TestCalcMetersWithVerySmallDistanceBetweenTwoPoints(test *testing.T) {
	dist := CalcMeters(0.00001, 0.0, 00002.0, 0.0)
	assert.Equal(test, 0.22237860120277003, dist)
}

func TestCalcMetersWithVeryLongDistanceBetweenTwoPoints(test *testing.T) {
	dist := CalcMeters(90.0, 90.0, 0.0, 0.0)
	assert.Equal(test, 10.007037053999998, dist)
}

func TestCalcKilometersWithNoDistanceBetweenTwoPoints(test *testing.T) {
	dist := CalcKilometers(0.0, 0.0, 0.0, 0.0)
	assert.Equal(test, 0.0, dist)
}

func TestCalcKilometersWithVerySmallDistanceBetweenTwoPoints(test *testing.T) {
	dist := CalcKilometers(0.00001, 0.0, 0.00002, 0.0)
	assert.Equal(test, 0.002486381704489686, dist)
}

func TestCalcKilometersWithVeryLongDistanceBetweenTwoPoints(test *testing.T) {
	dist := CalcKilometers(90.0, 90.0, 0.0, 0.0)
	assert.Equal(test, 10007.037053999999, dist)
}

func TestCalcMilesWithNoDistanceBetweenTwoPoints(test *testing.T) {
	dist := CalcMiles(0.0, 0.0, 0.0, 0.0)
	assert.Equal(test, 0.0, dist)
}

func TestCalcMilesWithVerySmallDistanceBetweenTwoPoints(test *testing.T) {
	dist := CalcMiles(0.00001, 0.0, 0.00002, 0.0)
	assert.Equal(test, 0.0015449698040747673, dist)
}

func TestCalcMilesWithVeryLongDistanceBetweenTwoPoints(test *testing.T) {
	dist := CalcMiles(90.0, 90.0, 0.0, 0.0)
	assert.Equal(test, 6218.099999999999, dist)
}

func TestCalcNauticMilesWithNoDistanceBetweenTwoPoints(test *testing.T) {
	dist := CalcNauticalMiles(0.0, 0.0, 0.0, 0.0)
	assert.Equal(test, 0.0, dist)
}

func TestCalcNauticMilesWithVerySmallDistanceBetweenTwoPoints(test *testing.T) {
	dist := CalcNauticalMiles(0.00001, 0.0, 0.00002, 0.0)
	assert.Equal(test, 0.0013416517778585277, dist)
}

func TestCalcNauticMilesWithVeryLongDistanceBetweenTwoPoints(test *testing.T) {
	dist := CalcMiles(90.0, 90.0, 0.0, 0.0)
	assert.Equal(test, 6218.099999999999, dist)
}
