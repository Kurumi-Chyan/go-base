package main

import "fmt"

func main() {
	var sli = []int{1, 2, 3, 4, 5}
	s := Delete(sli, 2)
	for i := 0; i < len(s); i++ {
		fmt.Printf("a[%d]:%d ", i+1, s[i])
	}
	println()
	s1 := Insert(s, 2, 7)
	for i := 0; i < len(s1); i++ {
		fmt.Printf("a[%d]:%d ", i+1, s1[i])
	}
}

func Delete(sli []int, index int) []int {
	sli = append(sli[:index], sli[index+1:]...)
	return sli
}

func Insert(sli []int, index int, n int) []int {
	tmp := append([]int{}, sli[index:]...)
	sli = append(sli[:index], n)
	sli = append(sli, tmp...)
	return sli
}
