package main

import "fmt"

func ParityCheck(digit int) bool {
	var resultOfChecking bool
	if digit%2 != 0 {
		resultOfChecking = false
	} else {
		resultOfChecking = true
	}
	return resultOfChecking
}

func main() {
	fmt.Println(ParityCheck(10))
	fmt.Println(ParityCheck(11))
}
