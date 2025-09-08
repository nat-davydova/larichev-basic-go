package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("BMI calculator")

	userHeightCm, userWeightKg := getUserInput()

	var userHeightMeters = userHeightCm / 100

	BMI := calcBMI(userHeightMeters, userWeightKg)

	outputResult(BMI)
}

func outputResult(result float64) {
	fmt.Printf("Your BMI is: %.0f", result)
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
