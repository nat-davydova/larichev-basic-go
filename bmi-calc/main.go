package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	fmt.Println("BMI calculator")

	for {
		handleBMI()
		isOneMoreTime := getIsOneMoreTime()

		if !isOneMoreTime {
			break
		}
	}

}

func handleBMI() {
	var description string

	userHeightCm, userWeightKg := getUserBMIInput()

	var userHeightMeters = userHeightCm / 100

	BMI, err := calcBMI(userHeightMeters, userWeightKg)

	if err != nil {
		fmt.Printf("Error calculating BMI: %v\n", err)
		return
	}

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

func calcBMI(height float64, weight float64) (float64, error) {
	if weight <= 0 || height <= 0 {
		return 0, errors.New("invalid parameters")
	}

	const BMIPow = 2
	return weight / math.Pow(height, BMIPow), nil
}

func getUserBMIInput() (float64, float64) {
	var userHeightCm float64
	var userWeightKg float64

	fmt.Println("Enter your user height in cm")
	fmt.Scan(&userHeightCm)

	fmt.Println("Enter your user weight in kg")
	fmt.Scan(&userWeightKg)

	return userHeightCm, userWeightKg
}

func getIsOneMoreTime() bool {
	var isOneMoreTime string

	fmt.Println("Press Y if you want to calc once again")
	fmt.Scan(&isOneMoreTime)

	return isOneMoreTime == "Y" || isOneMoreTime == "y"
}
