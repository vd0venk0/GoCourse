package main

import "fmt"

const limit = 100

func fibi(n int) uint64 {
	var a, b uint64 = 1, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

func main() {
	for j := 0; j < limit; j++ {
		fmt.Println(j+1, " = ", fibi(j))
	}
}
