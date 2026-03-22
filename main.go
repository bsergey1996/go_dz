package main

import "fmt"

func main() {
	// Выносим данные из глобальной области (Global State — это зло для тестов)
	rates := map[string]float64{
		"USD":  1.0,
		"EUR":  0.84,
		"RUB":  77.0,
		"USDT": 0.9,
	}

	from := inputCurrency("Введите исходную валюту:", rates)
	amount := inputAmount()
	to := inputCurrency("Введите целевую валюту:", rates)

	result := calculateMoney(amount, from, to, rates)

	fmt.Printf("Результаты %.2f %s\n", result, to)
}

// Передаем мапу как обычный параметр.
// Go скопирует только указатель, это очень быстро (8 байт).
func outputCurrency(rates map[string]float64) {
	fmt.Print("Доступные валюты: ")
	for currency := range rates {
		fmt.Printf("%s ", currency)
	}
	fmt.Println()
}

func inputCurrency(promt string, rates map[string]float64) string {
	for {
		fmt.Println(promt)
		outputCurrency(rates)

		var currency string
		fmt.Scan(&currency)

		if _, exists := rates[currency]; exists {
			return currency
		}

		fmt.Println("Неверная валюта, попробуйте снова")
	}
}

func calculateMoney(money float64, from string, to string, rates map[string]float64) float64 {
	if from == to {
		return money
	}

	// Математика остается прежней, но теперь данные приходят из аргумента
	return (money / rates[from]) * rates[to]
}

func inputAmount() float64 {
	for {
		fmt.Println("Введите сумму:")
		var amount float64
		_, err := fmt.Scan(&amount)

		if err != nil {
			fmt.Println("Ошибка ввода, попробуйте снова.")
			var trash string
			fmt.Scanln(&trash)
			continue
		}

		if amount <= 0 {
			fmt.Println("Сумма должна быть больше 0.")
			continue
		}
		return amount
	}
}
