package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	const years = 5
	var answerSum string
	var answerProcent string
	var summa float64

	fmt.Println("Введите сумму вклада: ")
	fmt.Scanln(&answerSum)
	deposit, err := strconv.Atoi(answerSum)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Введите процент по вкладу: ")
	fmt.Scanln(&answerProcent)
	interest, err := strconv.Atoi(answerProcent)
	if err != nil {
		log.Fatalln(err)
	}

	summa = float64(deposit)
	fmt.Println("Доходность по вкладу: ")
	for i := 1; i <= years; i++ {
		summa = (summa * float64(interest) / 100) + summa
		fmt.Println("Сумма за ", i, " год равна ", summa)
	}
}
