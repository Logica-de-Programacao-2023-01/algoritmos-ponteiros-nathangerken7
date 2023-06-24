package main

import (
	"fmt"
	"strings"
)

func inverterString(s *string) {

	newS := strings.Split(*s, "")
	*s = ""

	for i := len(newS) - 1; i >= 0; i-- {

		*s += newS[i]

	}

}

func main() {

	s := "Oi"

	inverterString(&s)

	fmt.Println(s)

}
