package main

import (
	"fmt"
	"sort"
)

type Meeting struct {
	start int
	end   int
	index int
}

func main() {
	var N int
	fmt.Print("Введите количество заявок: ")
	fmt.Scan(&N)

	meetings := make([]Meeting, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&meetings[i].start, &meetings[i].end)
		meetings[i].index = i + 1
	}

	sort.Slice(meetings, func(i, j int) bool {
		if meetings[i].end == meetings[j].end {
			return meetings[i].start < meetings[j].start
		}
		return meetings[i].end < meetings[j].end
	})

	selectedIndices := []int{}
	lastEnd := -1

	for _, m := range meetings {
		if m.start >= lastEnd {
			selectedIndices = append(selectedIndices, m.index)
			lastEnd = m.end
		}
	}

	sort.Ints(selectedIndices)

	fmt.Println("Результат:", len(selectedIndices))
	for _, idx := range selectedIndices {
		fmt.Printf("%d ", idx)
	}
	fmt.Println()
}
