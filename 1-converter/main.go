package main

import "fmt"

func main() {
	var initCurrency string
	var targetCurrency string
	var moneyAmount float64

	const USDtoEUR = 0.85
	const USDtoRUB = 81.18
	const EURtoRUB = USDtoRUB / USDtoEUR

	initCurrency, targetCurrency, moneyAmount = getUserInput()

	fmt.Println("EURtoRUB", EURtoRUB)
	fmt.Printf("init currency - %v, target currency - %v, money amount - %v", initCurrency, targetCurrency, moneyAmount)
}

func getUserInput() (string, string, float64) {
	var initCurrency string
	var targetCurrency string
	var moneyAmount float64

	fmt.Println("Enter your init currency:")
	fmt.Scan(&initCurrency)

	fmt.Println("Enter your target currency:")
	fmt.Scan(&targetCurrency)

	fmt.Println("Enter your money amount:")
	fmt.Scan(&moneyAmount)

	return initCurrency, targetCurrency, moneyAmount
}

func convertCurrencies(initCurrency string, targetCurrency string, moneyAmount float64) float64 {
	fmt.Printf("init currency - %v, target currency - %v, money amount - %v", initCurrency, targetCurrency, moneyAmount)
	return moneyAmount
}
