package main

import (
	"fmt"
	"math"
)

func main() {
	var isOneMoreTime bool = true

	fmt.Println("BMI calculator")

	for isOneMoreTime {
		if isOneMoreTime {
			handleBMI()
		} else {
			break
		}
	}

}

func handleBMI() {
	var description string

	userHeightCm, userWeightKg := getUserInput()

	var userHeightMeters = userHeightCm / 100

	BMI := calcBMI(userHeightMeters, userWeightKg)

	switch {
	case BMI < 16:
		description = "BMI is too small"
	case BMI > 16 && BMI < 18.5:
		description = "BMI is small"
	case BMI < 25:
		description = "BMI is normal"
	case BMI < 35:
		description = "BMI is large"
	case BMI >= 35:
		description = "BMI is too large"
	default:
		description = "Sorry, are you a human?"
	}

	outputResult(BMI, description)
}

func outputResult(result float64, description string) {
	fmt.Printf("Your BMI is: %.0f, %v \n ______________ \n", result, description)
}

func calcBMI(height float64, weight float64) float64 {
	const BMIPow = 2
	return weight / math.Pow(height, BMIPow)
}

func getUserInput() (float64, float64) {
	var userHeightCm float64
	var userWeightKg float64

	fmt.Println("Enter your user height in cm")
	fmt.Scan(&userHeightCm)

	fmt.Println("Enter your user weight in kg")
	fmt.Scan(&userWeightKg)

	return userHeightCm, userWeightKg
}
