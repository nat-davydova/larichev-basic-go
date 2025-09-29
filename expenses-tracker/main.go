package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var transactions []float64
}

func getTransactionInput() (float64, error) {
	var transaction string

	fmt.Println("Enter your transaction amount:")
	fmt.Scan(&transaction)

	transactionTrimmed := strings.TrimSpace(transaction)
	transactionFormatted := strings.ReplaceAll(transactionTrimmed, ",", ".")

	parsed, err := strconv.ParseFloat(transactionFormatted, 64)

	if err != nil || parsed == 0 {
		return 0, errors.New("invalid transaction amount: should be a number and not equal to 0")
	}

	return parsed, nil
}
