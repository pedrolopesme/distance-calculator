package distance_calculator

import "math"

const(

	// Unit in Meter
	UnitMeter = "METER"

	// Unit in Mile
	UnitMile = "MILE"

	// Unit in Kilometer
	UnitKilometer = "KILOMETER"

	// Unit in Nautical
	UnitNautical = "NAUTICAL"
)

// deg2rad converts decimal degrees to radians
func deg2rad(deg float64) float64 {
	return deg * math.Pi / 180.0
}

// rad2deg converts radians to decimal degrees
func rad2deg(rad float64) float64 {
	return (rad * 180.0) / math.Pi
}

// calculate
func Calculate(firstLatitude float64, firstLongitude float64, secondLongitude float64, secondLatitude float64, unit string) (dist float64) {

	theta := firstLongitude - secondLongitude
	dist = math.Sin((deg2rad(firstLatitude)) * math.Sin(deg2rad(secondLatitude)) + math.Cos(deg2rad(firstLatitude)) * math.Cos(deg2rad(secondLatitude)) * math.Cos(deg2rad(theta)))
	dist = math.Acos(dist)
	dist = rad2deg(dist)
	dist = dist * 60

	if unit == UnitMile {
		dist = dist *  1.1515
	} else if unit == UnitMeter {
		dist = dist * 0.609344
	} else if unit == UnitKilometer {
		dist = dist * 1.60934
	} else if unit == UnitNautical {
		dist = dist * 0.8684
	} else {
		dist = -1
	}
	return

}

func CalcMeters(firstLatitude float64, firstLongitude float64, secondLongitude float64, secondLatitude float64) {
	Calculate(firstLatitude, firstLongitude, secondLongitude, secondLatitude, UnitMeter)
}

func CalcMiles(firstLatitude float64, firstLongitude float64, secondLongitude float64, secondLatitude float64) {
	Calculate(firstLatitude, firstLongitude, secondLongitude, secondLatitude, UnitMile)
}

func CalcKilometers(firstLatitude float64, firstLongitude float64, secondLongitude float64, secondLatitude float64) {
	Calculate(firstLatitude, firstLongitude, secondLongitude, secondLatitude, UnitKilometer)
}

func CalcNauticals(firstLatitude float64, firstLongitude float64, secondLongitude float64, secondLatitude float64) {
	Calculate(firstLatitude, firstLongitude, secondLongitude, secondLatitude, UnitNautical)
}