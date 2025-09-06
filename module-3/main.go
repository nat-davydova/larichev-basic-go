package main

import (
	"fmt"
	"math"
)

func main() {
	const BMIPow = 2

	userHeightMeters := 1.8
	var userWeightKg float64 = 100
	BMI := userWeightKg / math.Pow(userHeightMeters, BMIPow)
	fmt.Println(BMI)
}
