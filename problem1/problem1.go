package main

import "fmt"

func Swap(a, b *int) (int, int) {
	//your code here
	c := *a
	*a = *b
	*b = c
	return *a, *b
}

func main() {
	a := 10
	b := 20

	Swap(&a, &b)
	fmt.Println(a, b)
}
