package main

import (
	"fmt"
)

func kelvinToCelsius(temp float32) float32 {
	return temp - 273
}

func main() {
	tempInKelvin := float32(373.2)
	tempInCelsius := kelvinToCelsius(tempInKelvin)
	fmt.Printf("Water boiling temperature is %.2f C\n", tempInCelsius)
}
