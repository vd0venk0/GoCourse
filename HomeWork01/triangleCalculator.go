package main

import (
	"fmt"
	"math"
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
	var CathetA, CathetB float64

	fmt.Println("Введите длину катета a: ")
	fmt.Scanln(&CathetA)
	fmt.Println("Введите длину катета b: ")
	fmt.Scanln(&CathetB)

	hypotenuse := HypotenuseCalculate(CathetA, CathetB)
	square := SquareCalculate(CathetA, CathetB)
	perimeter := PerimeterCalculate(CathetA, CathetB, hypotenuse)

	fmt.Println("Длина гипотенузы: ", hypotenuse)
	fmt.Println("Площадь треугольника: ", square)
	fmt.Println("Длина периметра: ", perimeter)

}
