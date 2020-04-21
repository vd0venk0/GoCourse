package main

import (
	"fmt"
)

const years = 5

func main() {

	var deposit float64
	var interest float64

	fmt.Println("Введите сумму вклада: ")
	fmt.Scanln(&deposit)

	fmt.Println("Введите процент по вкладу: ")
	fmt.Scanln(&interest)

	fmt.Println("Доходность по вкладу: ")
	for i := 1; i <= years; i++ {
		deposit += (deposit * interest / 100)
		fmt.Println("Сумма за ", i, " год равна ", deposit)
	}
}
