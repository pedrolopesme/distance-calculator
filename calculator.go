package distance_calculator

import "math"

const (
	// Unit in Meter
	UnitMeter = "METER"

	// Unit in Mile
	UnitMile = "MILE"

	// Unit in Kilometer
	UnitKilometer = "KILOMETER"

	// Unit in Nautical
	UnitNauticalMile = "NAUTICAL_MILE"
)

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

// deg2rad converts decimal degrees to radians
func deg2rad(deg float64) float64 {
	return deg * math.Pi / 180.0
}

// rad2deg converts radians to decimal degrees
func rad2deg(rad float64) float64 {
	return (rad * 180.0) / math.Pi
}

// calculate two coordinates and return the distance according to a given dist unit
func Calculate(first Coordinate, second Coordinate, unit string) (dist float64) {

	theta := first.Longitude - second.Longitude
	dist = math.Sin((deg2rad(first.Latitude))*math.Sin(deg2rad(second.Latitude)) + math.Cos(deg2rad(first.Latitude))*math.Cos(deg2rad(second.Latitude))*math.Cos(deg2rad(theta)))
	dist = math.Acos(dist)
	dist = rad2deg(dist)
	dist = dist * 60

	if unit == UnitMile {
		dist = dist * 1.1515
	} else if unit == UnitMeter {
		dist = dist * 0.609344
	} else if unit == UnitKilometer {
		dist = dist * 1.60934
	} else if unit == UnitNauticalMile {
		dist = dist * 0.8684
	} else {
		dist = -1
	}
	return
}

// CalcMeters calculates the distance between two coordinates in Meters
func CalcMeters(firstLatitude float64, firstLongitude float64, secondLongitude float64, secondLatitude float64) float64 {
	first := Coordinate{Latitude: firstLatitude, Longitude: firstLongitude}
	second := Coordinate{Latitude: secondLatitude, Longitude: secondLongitude}
	return Calculate(first, second, UnitMeter)
}

// CalcMiles calculates the distance between two coordinates in Miles
func CalcMiles(firstLatitude float64, firstLongitude float64, secondLongitude float64, secondLatitude float64) float64 {
	first := Coordinate{Latitude: firstLatitude, Longitude: firstLongitude}
	second := Coordinate{Latitude: secondLatitude, Longitude: secondLongitude}
	return Calculate(first, second, UnitMile)
}

// CalcKilometers calculates the distance between two coordinates in Kilometers
func CalcKilometers(firstLatitude float64, firstLongitude float64, secondLongitude float64, secondLatitude float64) float64 {
	first := Coordinate{Latitude: firstLatitude, Longitude: firstLongitude}
	second := Coordinate{Latitude: secondLatitude, Longitude: secondLongitude}
	return Calculate(first, second, UnitKilometer)
}

// CalcNautical calculates the distance between two coordinates in Nautical Miles
func CalcNauticalMiles(firstLatitude float64, firstLongitude float64, secondLongitude float64, secondLatitude float64) float64 {
	first := Coordinate{Latitude: firstLatitude, Longitude: firstLongitude}
	second := Coordinate{Latitude: secondLatitude, Longitude: secondLongitude}
	return Calculate(first, second, UnitNauticalMile)
}
