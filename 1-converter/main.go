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
	initCurrency, targetCurrency, moneyAmount := getUserInput()

	convertedMoneyAmount, err := convertCurrencies(initCurrency, targetCurrency, moneyAmount)

	if err == nil {
		targetCurrencyFormatted := strings.ToUpper(targetCurrency)
		fmt.Printf("Your converted money amount: %.2f %v \n", convertedMoneyAmount, targetCurrencyFormatted)
	} else {
		fmt.Printf("Convertion error: %v \n", err)
	}

}

func getUserInput() (string, string, float64) {
	inputCurrency, _ := getInputCurrency()
	moneyAmount, _ := getMoneyAnount()
	targetCurrency, _ := getTargetCurrency(inputCurrency)

	return inputCurrency, targetCurrency, moneyAmount
}

func getInputCurrencyUserInput() (string, error) {
	var inputCurrency string

	fmt.Println("Enter your init currency (RUB, USD, EUR):")
	fmt.Scan(&inputCurrency)

	inputCurrencyTrimmed := strings.TrimSpace(inputCurrency)
	inputCurrencyFormatted := strings.ToLower(inputCurrencyTrimmed)

	if inputCurrencyFormatted != rub && inputCurrencyFormatted != usd && inputCurrencyFormatted != eur {
		return "", errors.New("invalid init currency, you're allowed to enter only one of (RUB, USD, EUR)")
	}

	return inputCurrency, nil
}

func getMoneyAmountUserInput() (float64, error) {
	var moneyAmount string

	fmt.Println("Enter your money amount:")
	fmt.Scan(&moneyAmount)

	moneyAmountTrimmed := strings.TrimSpace(moneyAmount)
	moneyAmountFormatted := strings.ReplaceAll(moneyAmountTrimmed, ",", ".")

	parsed, err := strconv.ParseFloat(moneyAmountFormatted, 64)

	if err == nil && parsed > 0 {
		return parsed, nil
	} else {
		return -1, errors.New("invalid money amount, you're allowed to enter only positive numbers")
	}
}

func getTargetCurrencyInput(inputCurrency string) (string, error) {
	var targetCurrency string
	var description string

	inputCurrencyTrimmed := strings.TrimSpace(inputCurrency)
	inputCurrencyFormatted := strings.ToLower(inputCurrencyTrimmed)

	switch inputCurrencyFormatted {
	case eur:
		description = "(USD, RUB)"
	case usd:
		description = "(EUR, RUB)"
	case rub:
		description = "(EUR, USD)"
	default:
		return "", errors.New("invalid init currency")
	}

	fmt.Printf("Enter your target currency %v:", description)
	fmt.Scan(&targetCurrency)

	targetCurrencyTrimmed := strings.TrimSpace(targetCurrency)
	targetCurrencyFormatted := strings.ToLower(targetCurrencyTrimmed)

	switch {
	case inputCurrencyFormatted == eur && (targetCurrencyFormatted == usd || targetCurrencyFormatted == rub):
	case inputCurrencyFormatted == usd && (targetCurrencyFormatted == eur || targetCurrencyFormatted == rub):
	case inputCurrencyFormatted == rub && (targetCurrencyFormatted == eur || targetCurrencyFormatted == usd):
		return targetCurrency, nil
	default:
		errorText := fmt.Sprintf("invalid target currency, you're allowed to enter only one of %v", description)
		return "", errors.New(errorText)
	}

	return targetCurrency, nil
}

func convertCurrencies(initCurrency string, targetCurrency string, moneyAmount float64) (float64, error) {
	var convertedMoneyAmount float64

	initCurrencyFormatted := strings.ToLower(initCurrency)
	targetCurrencyFormatted := strings.ToLower(targetCurrency)

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

func getInputCurrency() (string, error) {
	inputCurrency, inputCurrencyErr := getInputCurrencyUserInput()

	if inputCurrencyErr == nil {
		return inputCurrency, nil
	} else {
		fmt.Printf("Error getting input currency - %v\n", inputCurrencyErr)
		val, err := getInputCurrency()
		return val, err
	}
}

func getMoneyAnount() (float64, error) {
	moneyAmount, moneyAmountErr := getMoneyAmountUserInput()

	if moneyAmountErr == nil {
		return moneyAmount, nil
	} else {
		fmt.Printf("Error getting input currency - %v\n", moneyAmountErr)
		val, err := getMoneyAnount()
		return val, err
	}
}

func getTargetCurrency(inputCurrency string) (string, error) {
	targetCurrency, targetCurrencyErr := getTargetCurrencyInput(inputCurrency)

	if targetCurrencyErr == nil {
		return targetCurrency, nil
	} else {
		fmt.Printf("Error getting target currency - %v\n", targetCurrencyErr)
		val, err := getTargetCurrency(inputCurrency)
		return val, err
	}
}
