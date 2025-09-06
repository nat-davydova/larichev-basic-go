package main

import "fmt"

func main() {
	const USDtoEUR = 0.85
	const USDtoRUB = 81.18
	const EURtoRUB = USDtoRUB / USDtoEUR

	fmt.Printf("EURtoRUB", EURtoRUB)
}
