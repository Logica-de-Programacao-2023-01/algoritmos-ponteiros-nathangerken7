package main

import (
	"fmt"
	"strconv"
	"strings"
)

func somaDoisUltimosAlgarismos(n *int) {

	nString := strconv.Itoa(*n)
	nStringAlgarismos := strings.Split(nString, "")
	a1 := 0
	a2 := 0
	soma := 0

	i := len(nStringAlgarismos) - 1

	a1, _ = strconv.Atoi(nStringAlgarismos[i])
	a2, _ = strconv.Atoi(nStringAlgarismos[i-1])
	soma = a1 + a2

	*n = soma

}

func main() {

	x := 21324123421389

	somaDoisUltimosAlgarismos(&x)

	fmt.Println(x)

}
