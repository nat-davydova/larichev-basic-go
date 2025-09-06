package main

import (
	"fmt"
	"math"
)

func main() {
	userHeightMeters := 1.8
	var userWeightKg float64 = 100
	BMI := userWeightKg / math.Pow(userHeightMeters, 2)
	fmt.Println(BMI)
}
