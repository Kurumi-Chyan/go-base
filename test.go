package main

import "fmt"

func main() {
	var s string
	s = stringTest("hhh")
	println(s)

	var a = [2]int{1, 2}
	fmt.Printf("a1: %v, len: %d , cap: %d \n", a, len(a), cap(a))

	var a2 = []int{1, 2, 3}
	fmt.Printf("a2 : %v , len : %d , cap : %d", a2, len(a2), cap(a2))
}

func stringTest(name string) string {
	return name
}
