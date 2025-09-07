package main

import (
	"fmt"
	"math"
)

func main() {
	const BMIPow = 2

	var userHeightMeters float64
	var userWeightKg float64

	fmt.Println("BMI calculator")

	fmt.Println("Enter your user height in meters")
	fmt.Scan(&userHeightMeters)

	fmt.Println("Enter your user weight in kg")
	fmt.Scan(&userWeightKg)

	BMI := userWeightKg / math.Pow(userHeightMeters, BMIPow)
	fmt.Printf("Your BMI is: %.0f", BMI)
}
