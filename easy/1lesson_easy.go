package main

import (
	"fmt"
)

func main() {
	easyLevel := [5]int{10, 2, 3, 4, 5}
	resEasy := easyHomeWork(easyLevel[:])
	fmt.Println("Результат простого задания:", resEasy)
}

func easyHomeWork(arr []int) int {
	var sum int
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}
