package main

// доделаю чуть позже...
func fibonacciNumbers(limit int) []int {
	var numbers []int
	for i := 0; i <= limit; i++ {
		for k := 0; k <= limit; k++ {
			if k < 2 {
				numbers = append(numbers, k)

			} else if numbers[i] == numbers[i]+numbers[i-1] {

			}
		}

	}

}
