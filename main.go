package main

import "fmt"

const (
	USDToEUR = 0.84 // 1 USD ≈ 0.84 EUR
	USDToRUB = 77.0 // 1 USD ≈ 77 RUB
	EURToRUB = USDToRUB / USDToEUR
)

func inputCurrency(promt string) string {
	for {
		fmt.Println(promt)
		fmt.Println("Доступные валюты: USD, EUR, RUB")

		var currency string
		fmt.Scan(&currency)

		switch currency {
		case "USD", "EUR", "RUB":
			return currency
		default:
			fmt.Println("Неверная валюта попробуйте снова")
		}
	}
}

func inputAmount() float64 {
	for {
		fmt.Println("Введите сумму:")

		var amount float64
		_, err := fmt.Scan(&amount)

		if err != nil {
			fmt.Println("Ошибка ввода, попробуйте снова.")
			continue
		}

		if amount <= 0 {
			fmt.Println("Сумма должна быть больше 0.")
			continue
		}

		return amount
	}
}

func userInput() float64 {
	var money float64
	fmt.Scan(&money)
	return money
}

func calculateMoney(money float64, from string, to string) float64 {
	if from == to {
		return money
	}

	if from == "USD" && to == "EUR" {
		return money * USDToEUR
	}
	if from == "USD" && to == "RUB" {
		return money * USDToRUB
	}

	if from == "EUR" && to == "USD" {
		return money / USDToEUR
	}
	if from == "EUR" && to == "RUB" {
		return money * EURToRUB
	}

	if from == "RUB" && to == "USD" {
		return money / USDToRUB
	}
	if from == "RUB" && to == "EUR" {
		return money / EURToRUB
	}

	return 0
}

func main() {
	from := inputCurrency("Введите исходную валюту:")
	amount := inputAmount()
	to := inputCurrency("Введите целевую валюту:")

	result := calculateMoney(amount, from, to)

	fmt.Printf("Результаты %.2f %s\n", result, to)
}
