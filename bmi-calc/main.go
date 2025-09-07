package main

import (
	"fmt"
	"math"
)

func main() {
	const BMIPow = 2

	var userHeightCm float64
	var userWeightKg float64

	fmt.Println("BMI calculator")

	fmt.Println("Enter your user height in cm")
	fmt.Scan(&userHeightCm)

	fmt.Println("Enter your user weight in kg")
	fmt.Scan(&userWeightKg)

	var userHeightMeters = userHeightCm / 100

	BMI := userWeightKg / math.Pow(userHeightMeters, BMIPow)
	fmt.Printf("Your BMI is: %.0f", BMI)
}
