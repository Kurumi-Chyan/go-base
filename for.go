package main

import "fmt"

func main() {
	s := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(s); i++ {
		fmt.Printf("s%d: %d ", i, s[i])
	}
	println()
	for i, value := range s {
		fmt.Printf("i: %d, value: %d", i, value)
	}
}
