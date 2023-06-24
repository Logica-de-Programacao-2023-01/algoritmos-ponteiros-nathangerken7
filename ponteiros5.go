package main

import (
	"fmt"
	"math"
)

func mediaPi(n *float64) {

	*n = (math.Pi + *n) / 2

}

func main() {

	x := 34.32

	mediaPi(&x)

	fmt.Println(x)

}
