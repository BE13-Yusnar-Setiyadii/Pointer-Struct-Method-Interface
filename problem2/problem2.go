package main

import (
	"fmt"
)

func GetMinMax(numbers ...*int) (min int, max int) {
	// your code here
	min = *numbers[0]
	for i := 0; i < len(numbers); i++ {
		if *numbers[i] > max {
			max = *numbers[i]
		} else if *numbers[i] < min {
			min = *numbers[i]
		}
	}
	return min, max
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)
	min, max = GetMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai min ", min)
	fmt.Println("Nilai max ", max)
}
