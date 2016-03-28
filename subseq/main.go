package main

import "fmt"

type pair struct {
	value  int
	weight int
}

func solve(N int, list []pair) {
	// fmt.Printf("TODO solve for: %v\n", &list)
	var best = make([]pair, N)
	var max int
	copy(best, list)
	for i := 1; i < N; i++ {
		max = 0
		for j := i - 1; j >= 0; j-- {
			if best[j].value < best[i].value && best[j].weight > max {
				max = best[j].weight
			}
		}
		best[i].weight += max
	}
	max = 0
	for _, p := range best {
		if max < p.weight {
			max = p.weight
		}
	}
	fmt.Printf("%v\n", max)
}

func main() {
	var num_test_cases int
	var N int
	var v int
	fmt.Scanf("%d\n", &num_test_cases)
	for i := 0; i < num_test_cases; i++ {
		fmt.Scanf("%d\n", &N)
		var list = make([]pair, N)
		for i := range list {
			fmt.Scanf("%d", &v)
			list[i].value = v
		}
		for i := range list {
			fmt.Scanf("%d", &v)
			list[i].weight = v
		}
		solve(N, list)
	}
}
