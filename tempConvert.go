package main

import (
	"fmt"
)

func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func main() {
	var value, result float64
	var inputType string

	fmt.Print("Enter the temperature value: ")
	_, err := fmt.Scan(&value)
	if err != nil {
		fmt.Println("Invalid input for value.")
		return
	}

	fmt.Print("Enter the temperature type (C for Celsius, F for Fahrenheit): ")
	_, err = fmt.Scan(&inputType)
	if err != nil || (inputType != "C" && inputType != "F" && inputType != "c" && inputType != "f") {
		fmt.Println("Invalid input for temperature type. Use 'C' or 'F'.")
		return
	}

	if inputType == "C" || inputType == "c" {
		result = celsiusToFahrenheit(value)
		fmt.Printf("%.2f°C is %.2f°F\n", value, result)
	} else {
		result = fahrenheitToCelsius(value)
		fmt.Printf("%.2f°F is %.2f°C\n", value, result)
	}
}
