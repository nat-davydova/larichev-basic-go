package main

import (
	"errors"
	"fmt"
	"strconv"
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

	var moneyAmount float64
	var moneyAmountErr error

	var targetCurrency string

	for {
		inputCurrency, inputCurrencyErr = getInputCurrency()

		if inputCurrencyErr == nil {
			break
		}
	}

	for {
		moneyAmount, moneyAmountErr = getMoneyAmount()

		if moneyAmountErr == nil {
			break
		}
	}

	fmt.Println("Enter your target currency:")
	fmt.Scan(&targetCurrency)

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

func getMoneyAmount() (float64, error) {
	var moneyAmount string

	fmt.Println("Enter your money amount:")
	fmt.Scan(&moneyAmount)

	parsed, err := strconv.ParseFloat(moneyAmount, 64)

	if err == nil {
		return parsed, nil
	} else {
		return -1, errors.New("invalid money amount")
	}
}

func convertCurrencies(initCurrency string, targetCurrency string, moneyAmount float64) float64 {
	fmt.Printf("init currency - %v, target currency - %v, money amount - %v", initCurrency, targetCurrency, moneyAmount)
	return moneyAmount
}
