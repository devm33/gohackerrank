package main

import "fmt"

func solve(values []int, weights []int) {
	fmt.Printf("TODO solve for: %v, %v\n", &values, &weights)
}

func readIntLine(count int) []int {
	var slice = make([]int, count)
	var v int
	for i := range slice {
		fmt.Scanf("%d", &v)
		slice[i] = v
	}
	return slice
}

func main() {
	var num_test_cases int
	var N int
	fmt.Scanf("%d\n", &num_test_cases)
	for i := 0; i < num_test_cases; i++ {
		fmt.Scanf("%d\n", &N)
		solve(readIntLine(N), readIntLine(N))
	}
}
