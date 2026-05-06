package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return "Не найдено"
	}

	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	required := len(t)
	left, right := 0, 0
	minLen := math.MaxInt32
	start := 0

	for right < len(s) {
		rc := s[right]
		if need[rc] > 0 {
			required--
		}
		need[rc]--
		right++

		for required == 0 {
			if right-left < minLen {
				minLen = right - left
				start = left
			}
			lc := s[left]
			need[lc]++
			if need[lc] > 0 {
				required++
			}
			left++
		}
	}

	if minLen == math.MaxInt32 {
		return "Не найдено"
	}
	return s[start : start+minLen]
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите строку S: ")
	S, _ := reader.ReadString('\n')
	S = strings.TrimSpace(S)

	fmt.Print("Введите строку T: ")
	T, _ := reader.ReadString('\n')
	T = strings.TrimSpace(T)

	result := minWindow(S, T)
	fmt.Printf("Результат: %s\n", result)
}
