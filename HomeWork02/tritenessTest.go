package main

import "fmt"

func TritenessTest(digit int) bool {
	var resultOfChecking bool
	if digit%3 != 0 {
		resultOfChecking = false
	} else {
		resultOfChecking = true
	}
	return resultOfChecking
}

func main() {
	fmt.Println(TritenessTest(12))
	fmt.Println(TritenessTest(11))
}
