package main

import "fmt"

func main() {
	fmt.Println("Bubble Sort")
	arrayNumber := [20]int{64, 34, 25, 12, 22, 11, 90, 67, 56, 45, 99, 88, 77, 66, 55, 44, 33, 22, 11, 10}
	n := len(arrayNumber)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arrayNumber[j] > arrayNumber[j+1] {
				temp := arrayNumber[j]
				arrayNumber[j] = arrayNumber[j+1]
				arrayNumber[j+1] = temp
			}
		}
	}

	fmt.Println("Setelah dilakukan pengurutan:")
	fmt.Println(arrayNumber)
}

