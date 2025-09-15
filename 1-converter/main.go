package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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
	var targetCurrencyErr error

	for {
		inputCurrency, inputCurrencyErr = getInputCurrency()

		if inputCurrencyErr == nil {
			break
		} else {
			fmt.Printf("Error getting input currency - %v\n", inputCurrencyErr)
		}
	}

	for {
		moneyAmount, moneyAmountErr = getMoneyAmount()

		if moneyAmountErr == nil {
			break
		} else {
			fmt.Printf("Error getting input currency - %v\n", moneyAmountErr)
		}
	}

	for {
		targetCurrency, targetCurrencyErr = getTargetCurrency(inputCurrency)

		if targetCurrencyErr == nil {
			break
		} else {
			fmt.Printf("Error getting target currency - %v\n", targetCurrencyErr)
		}
	}

	return inputCurrency, targetCurrency, moneyAmount
}

func getInputCurrency() (string, error) {
	var initCurrency string

	fmt.Println("Enter your init currency (RUB, USD, EUR):")
	fmt.Scan(&initCurrency)

	initCurrencyFormatted := strings.ToLower(initCurrency)

	if initCurrencyFormatted != "rub" && initCurrencyFormatted != "usd" && initCurrencyFormatted != "eur" {
		return "", errors.New("invalid init currency, you're allowed to enter only one of (RUB, USD, EUR)")
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
		return -1, errors.New("invalid money amount, you're allowed to enter only numbers")
	}
}

func getTargetCurrency(inputCurrency string) (string, error) {
	var targetCurrency string
	var description string

	inputCurrencyFormatted := strings.ToLower(inputCurrency)

	switch inputCurrencyFormatted {
	case "eur":
		description = "(USD, RUB)"
	case "usd":
		description = "(EUR, RUB)"
	case "rub":
		description = "(EUR, USD)"
	default:
		return "", errors.New("invalid init currency")
	}

	fmt.Printf("Enter your target currency %v:", description)
	fmt.Scan(&targetCurrency)

	targetCurrencyFormatted := strings.ToLower(targetCurrency)

	switch {
	case inputCurrencyFormatted == "eur" && (targetCurrencyFormatted == "usd" || targetCurrencyFormatted == "rub"):
	case inputCurrencyFormatted == "usd" && (targetCurrencyFormatted == "eur" || targetCurrencyFormatted == "rub"):
	case inputCurrencyFormatted == "rub" && (targetCurrencyFormatted == "eur" || targetCurrencyFormatted == "usd"):
		return targetCurrency, nil
	default:
		errorText := fmt.Sprintf("invalid target currency, you're allowed to enter only one of %v", description)
		return "", errors.New(errorText)
	}

	return targetCurrency, nil
}

func convertCurrencies(initCurrency string, targetCurrency string, moneyAmount float64) float64 {
	fmt.Printf("init currency - %v, target currency - %v, money amount - %v", initCurrency, targetCurrency, moneyAmount)
	return moneyAmount
}
