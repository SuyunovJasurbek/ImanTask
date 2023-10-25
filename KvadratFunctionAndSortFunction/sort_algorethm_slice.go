package main

import (
	"fmt"
)

func main() {
	slice := []int{34, 5, 3, 2, 45, 6, 312, 56, 23, 1, 4, 23}
	fmt.Println(slice)
	fmt.Println("________________________________________________________")
	Selectionsort(slice)
	fmt.Println(slice)
}

func Selectionsort(items []int) {
	var n = len(items)
	for i := 0; i < n; i++ {
		var a = i
		for j := i; j < n; j++ {
			if items[j] < items[a] {
				a = j
			}
		}
		sw := items[a]
		items[a] = items[i]
		items[i] = sw
	}
}