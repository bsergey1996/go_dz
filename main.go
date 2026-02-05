package main

import "fmt"

const USD_TO_EUR = 0.92
const USD_TO_RUB = 90
const EUR_TO_RUB = USD_TO_RUB / USD_TO_EUR

func main() {
	fmt.Println(EUR_TO_RUB)
}
