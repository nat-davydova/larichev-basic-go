package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	usdToEur = 0.85
	usdToRub = 81.18
	eurToRub = usdToRub / usdToEur
)

const (
	usd = "usd"
	eur = "eur"
	rub = "rub"
)

func main() {
	inputCurrency := getInputCurrency()
	moneyAmount := getMoneyAmount()
	targetCurrency := getTargetCurrency(inputCurrency)

	converted, err := convertCurrencies(inputCurrency, targetCurrency, moneyAmount)

	if err != nil {
		fmt.Printf("Convertion error: %v \n", err)
		return
	}

	targetCurrencyFormatted := strings.ToUpper(targetCurrency)
	fmt.Printf("Your converted money amount: %.2f %v \n", converted, targetCurrencyFormatted)
}

func getInputCurrencyUserInput() (string, error) {
	var inputCurrency string

	fmt.Println("Enter your init currency (RUB, USD, EUR):")
	fmt.Scan(&inputCurrency)

	inputCurrencyFormatted := formatInput(inputCurrency)

	if inputCurrencyFormatted != rub && inputCurrencyFormatted != usd && inputCurrencyFormatted != eur {
		return "", errors.New("invalid init currency, you're allowed to enter only one of (RUB, USD, EUR)")
	}

	return inputCurrency, nil
}

func getMoneyAmountUserInput() (float64, error) {
	var moneyAmount string

	fmt.Println("Enter your money amount:")
	fmt.Scan(&moneyAmount)

	moneyAmountFormatted := formatInput(moneyAmount)
	moneyAmountReplaced := strings.ReplaceAll(moneyAmountFormatted, ",", ".")

	parsed, err := strconv.ParseFloat(moneyAmountReplaced, 64)

	if err != nil || parsed <= 0 {
		return -1, errors.New("invalid money amount, you're allowed to enter only positive numbers")
	}

	return parsed, nil
}

func getTargetCurrencyInput(inputCurrency string) (string, error) {
	inputCurrencyFormatted := formatInput(inputCurrency)
	availableCurrency1, availableCurrency2, err := getAvailableCurrencies(inputCurrencyFormatted)

	if err != nil {
		return "", err
	}

	var targetCurrency string
	description := fmt.Sprintf("(%v, %v)", strings.ToUpper(availableCurrency1), strings.ToUpper(availableCurrency2))

	fmt.Printf("Enter your target currency %v:", description)
	fmt.Scan(&targetCurrency)

	targetCurrencyFormatted := formatInput(targetCurrency)

	if targetCurrencyFormatted == availableCurrency1 || targetCurrencyFormatted == availableCurrency2 {
		return targetCurrencyFormatted, nil
	} else {
		errorText := fmt.Sprintf("invalid target currency, you're allowed to enter only one of %v", description)
		return "", errors.New(errorText)
	}
}

func getAvailableCurrencies(inputCurrency string) (string, string, error) {
	switch inputCurrency {
	case eur:
		return usd, eur, nil
	case usd:
		return eur, rub, nil
	case rub:
		return eur, usd, nil
	default:
		return "", "", errors.New("invalid init currency")
	}
}

func convertCurrencies(initCurrency string, targetCurrency string, moneyAmount float64) (float64, error) {
	var convertedMoneyAmount float64

	initCurrencyFormatted := formatInput(initCurrency)
	targetCurrencyFormatted := formatInput(targetCurrency)

	switch {
	case initCurrencyFormatted == usd && targetCurrencyFormatted == eur:
		convertedMoneyAmount = moneyAmount * usdToEur
	case initCurrencyFormatted == usd && targetCurrencyFormatted == rub:
		convertedMoneyAmount = moneyAmount * usdToRub
	case initCurrencyFormatted == eur && targetCurrencyFormatted == usd:
		convertedMoneyAmount = moneyAmount / usdToEur
	case initCurrencyFormatted == eur && targetCurrencyFormatted == rub:
		convertedMoneyAmount = moneyAmount * eurToRub
	case initCurrencyFormatted == rub && targetCurrencyFormatted == usd:
		convertedMoneyAmount = moneyAmount / usdToRub
	case initCurrencyFormatted == rub && targetCurrencyFormatted == eur:
		convertedMoneyAmount = moneyAmount / eurToRub
	default:
		errorText := fmt.Sprintf("invalid converting %v %v in %v", initCurrency, moneyAmount, targetCurrency)
		return -1, errors.New(errorText)
	}

	return convertedMoneyAmount, nil
}

func getInputCurrency() string {
	inputCurrency, err := getInputCurrencyUserInput()

	if err != nil {
		fmt.Printf("Error getting input currency - %v\n", err)
		val := getInputCurrency()
		return val
	}

	return inputCurrency
}

func getMoneyAmount() float64 {
	moneyAmount, err := getMoneyAmountUserInput()

	if err != nil {
		fmt.Printf("Error getting input currency - %v\n", err)
		val := getMoneyAmount()
		return val
	}

	return moneyAmount
}

func getTargetCurrency(inputCurrency string) string {
	targetCurrency, err := getTargetCurrencyInput(inputCurrency)

	if err == nil {
		return targetCurrency
	} else {
		fmt.Printf("Error getting target currency - %v\n", err)
		val := getTargetCurrency(inputCurrency)
		return val
	}
}

func formatInput(input string) string {
	lowerCased := strings.ToLower(input)
	trimmed := strings.TrimSpace(lowerCased)

	return trimmed
}
