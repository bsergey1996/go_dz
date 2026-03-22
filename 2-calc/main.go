package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseNumbers(input string) ([]float64, error) {
	parts := strings.Split(input, ",")

	numbers := make([]float64, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)

		num, err := strconv.ParseFloat(part, 64)
		if err != nil {
			return nil, fmt.Errorf("Не удалось преобразовать '%s'", part)
		}

		numbers = append(numbers, num)
	}

	return numbers, nil
}

func sum(numbers []float64) float64 {
	total := 0.0
	for _, n := range numbers {
		total += n
	}
	return total
}

func averadge(numbers []float64) float64 {
	return sum(numbers) / float64(len(numbers))
}

func median(numbers []float64) float64 {
	// делаем копию, чтобы не портить исходный порядок
	copySlice := append([]float64(nil), numbers...)

	sort.Float64s(copySlice)

	n := len(copySlice)

	if n%2 == 1 {
		return copySlice[n/2]
	}

	return (copySlice[n/2-1] + copySlice[n/2]) / 2

}

func main() {
	var operation string
	var input string

	fmt.Print("Введите операцию (SUM, AVG, MED): ")
	fmt.Scanln(&operation)

	fmt.Print("Введите числа через запятую: ")
	fmt.Scanln(&input)

	numbers, err := parseNumbers(input)
	if err != nil {
		fmt.Println("Ошибка:", err)
		os.Exit(1)
	}

	if len(numbers) == 0 {
		fmt.Println("Нет чисел для расчёта")
		return
	}

	switch strings.ToUpper(operation) {
	case "SUM":
		fmt.Println("Результат:", sum(numbers))
	case "AVG":
		fmt.Println("Результат:", averadge(numbers))
	case "MED":
		fmt.Println("Результат:", median(numbers))
	default:
		fmt.Println("Неизвестная операция")
	}
}
