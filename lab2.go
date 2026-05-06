package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	var prev, curr int
	fmt.Scan(&prev)

	count := 0

	for i := 1; i < N; i++ {
		fmt.Scan(&curr)
		if curr == prev {
			count++
		}
		prev = curr
	}

	fmt.Printf("Результат: %d\n", count)
}
