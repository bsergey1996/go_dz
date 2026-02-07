package main

import "fmt"

const (
	USDToEUR = 0.92
	USDToRUB = 90.0
	EURToRUB = USDToRUB / USDToEUR
)

func userInput() float64 {
	var money float64
	fmt.Scan(&money)
	return money
}

func calculateMoney(money float64, currencyFrom string, currencyTo string) {
}

func main() {
	money := userInput()
	calculateMoney(money, "USD", "RUB")
}
