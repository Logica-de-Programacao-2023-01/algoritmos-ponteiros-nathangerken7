package main

import "fmt"

func main() {
	x := 1
	n := 4
	var ptr *int = &x
	var ptr2 *int = &n
	nnumeros(ptr, ptr2)
	fmt.Print(nnumeros(ptr, ptr2))
}

func nnumeros(ptr *int, ptr2 *int) int {
	soma := 0
	for i := *ptr; i <= *ptr2; i++ {
		soma += i
	}

	return soma
}
