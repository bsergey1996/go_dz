package main

import "fmt"

// Вместо отдельных констант создаем единый источник истины — map.
// Базовой валютой делаем USD (ее курс равен 1.0).
var exchangeRates = map[string]float64{
	"USD":  1.0,
	"EUR":  0.84,
	"RUB":  77.0,
	"USDT": 0.9,
}

func outputCurrency() {
	fmt.Print("Доступные валюты: ")
	for currency := range exchangeRates {
		fmt.Printf("%s ", currency)
	}
	fmt.Println() // Перенос строки после списка
}

func inputCurrency(promt string) string {
	for {
		fmt.Println(promt)

		outputCurrency()

		var currency string
		fmt.Scan(&currency)

		// Проверяем, есть ли введенная валюта в нашей map ключей.
		// Заменяем громоздкий switch на изящную проверку.
		if _, exists := exchangeRates[currency]; exists {
			return currency
		}

		fmt.Println("Неверная валюта попробуйте снова")
	}
}

func inputAmount() float64 {
	for {
		fmt.Println("Введите сумму:")

		var amount float64
		_, err := fmt.Scan(&amount)

		if err != nil {
			fmt.Println("Ошибка ввода, попробуйте снова.")
			// Очищаем буфер ввода, чтобы избежать зацикливания при вводе букв
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

// Функция сохранена по твоей просьбе, хоть сейчас она и не используется в main
func userInput() float64 {
	var money float64
	fmt.Scan(&money)
	return money
}

func calculateMoney(money float64, from string, to string) float64 {
	if from == to {
		return money
	}

	// Достаем курсы валют из нашей map
	rateFrom := exchangeRates[from]
	rateTo := exchangeRates[to]

	// Математика конвертации через базовую валюту (USD):
	// 1. Делим сумму на курс исходной валюты (получаем эквивалент в USD)
	// 2. Умножаем на курс целевой валюты
	return (money / rateFrom) * rateTo
}

func main() {
	from := inputCurrency("Введите исходную валюту:")
	amount := inputAmount()
	to := inputCurrency("Введите целевую валюту:")

	result := calculateMoney(amount, from, to)

	fmt.Printf("Результаты %.2f %s\n", result, to)
}
