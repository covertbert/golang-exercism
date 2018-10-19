package space

// Planet type
type Planet string

// Age returns an age based on the given planet
func Age(seconds float64, planet Planet) float64 {

	secondsInAnEarthYear := 31557600.0
	inputInEarthYears := seconds / secondsInAnEarthYear

	return inputInEarthYears / getPlanetYearInEarthYears(planet)
}

func getPlanetYearInEarthYears(planet Planet) float64 {
	var planetYearInEarthYears float64

	switch {
	case planet == "Earth":
		planetYearInEarthYears = 1
	case planet == "Mercury":
		planetYearInEarthYears = 0.2408467
	case planet == "Venus":
		planetYearInEarthYears = 0.61519726
	case planet == "Mars":
		planetYearInEarthYears = 1.8808158
	case planet == "Jupiter":
		planetYearInEarthYears = 11.862615
	case planet == "Saturn":
		planetYearInEarthYears = 29.447498
	case planet == "Uranus":
		planetYearInEarthYears = 84.016846
	case planet == "Neptune":
		planetYearInEarthYears = 164.79132
	}

	return planetYearInEarthYears
}
