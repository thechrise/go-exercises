// Package weather provides a simple weather forecast.
// It contains a function to get the current weather condition and location.
// It also contains variables to store the current condition and location.
package weather

// CurrentCondition is a string variable that stores the current weather condition.
var CurrentCondition string

// CurrentLocation is a string variable that stores the current location.
var CurrentLocation string

// Forecast returns the current weather condition and location.
// It takes two string arguments: city and condition.
// It returns a string that contains the current weather condition and location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
