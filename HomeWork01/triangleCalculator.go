package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func HypotenuseCalculate(a float64, b float64) float64 {
	C := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	return C
}

func SquareCalculate(a float64, b float64) float64 {
	S := (a * b) / 2
	return S
}

func PerimeterCalculate(a float64, b float64, c float64) float64 {
	P := a + b + c
	return P
}

func main() {
	fmt.Println("Программа находит площадь, периметр, и гипотенузу прямоугольного треугольника исходя из его катетов.")
	var (
		answerA string
		answerB string
	)

	fmt.Println("Введите длину катета a: ")
	fmt.Scanln(&answerA)
	CathetA, err := strconv.Atoi(answerA)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Введите длину катета b: ")
	fmt.Scanln(&answerB)
	CathetB, err := strconv.Atoi(answerB)
	if err != nil {
		log.Fatalln(err)
	}

	hypotenuse := HypotenuseCalculate(float64(CathetA), float64(CathetB))
	square := SquareCalculate(float64(CathetA), float64(CathetB))
	perimeter := PerimeterCalculate(float64(CathetA), float64(CathetB), hypotenuse)

	fmt.Println("Длина гипотенузы: ", hypotenuse)
	fmt.Println("Площадь треугольника: ", square)
	fmt.Println("Длина периметра: ", perimeter)

}
