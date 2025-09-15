package main

import (
	"errors"
	"fmt"
)

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
	var inputCurrency string
	var inputCurrencyErr error

	var targetCurrency string
	var moneyAmount float64

	for {
		inputCurrency, inputCurrencyErr = getInputCurrency()

		if inputCurrencyErr == nil {
			break
		}
	}

	fmt.Println("Enter your target currency:")
	fmt.Scan(&targetCurrency)

	fmt.Println("Enter your money amount:")
	fmt.Scan(&moneyAmount)

	return inputCurrency, targetCurrency, moneyAmount
}

func getInputCurrency() (string, error) {
	var initCurrency string

	fmt.Println("Enter your init currency (RUB, USD, EUR):")
	fmt.Scan(&initCurrency)

	if initCurrency != "RUB" && initCurrency != "USD" && initCurrency != "EUR" {
		return "", errors.New("invalid init currency")
	}

	return initCurrency, nil
}

func convertCurrencies(initCurrency string, targetCurrency string, moneyAmount float64) float64 {
	fmt.Printf("init currency - %v, target currency - %v, money amount - %v", initCurrency, targetCurrency, moneyAmount)
	return moneyAmount
}
