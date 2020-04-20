package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	const course int = 75
	var answer string
	fmt.Println("Программа конвертирует доллары в рубли по фиксированному курсу.")
	fmt.Println("Введите сумму в долларах для конвертации : ")
	fmt.Scanln(&answer)
	dollars, err := strconv.Atoi(answer)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("После конвертации Вы получите", dollars*course, "руб.")
}
