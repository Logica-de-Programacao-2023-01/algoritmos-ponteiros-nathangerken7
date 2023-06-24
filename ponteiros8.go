package main

import "fmt"

func alterar(v *int) {

	x := 0

	fmt.Println("Digite um valor para a variável:")
	fmt.Scanln(&x)

	*v = x

}

func main() {

	x := 0

	alterar(&x)

	fmt.Println(x)

}
