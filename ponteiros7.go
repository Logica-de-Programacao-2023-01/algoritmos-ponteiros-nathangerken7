package main

import "fmt"

type Conta struct {
	Saldo   float64
	Titular string
}

func addValor(c *Conta) {

	x := 0.0

	fmt.Println("Digite um valor para adicionar a sua conta: ")
	fmt.Scanln(&x)

	c.Saldo += x

}

func main() {

	c := Conta{
		Saldo:   500,
		Titular: "Felipe",
	}

	addValor(&c)

	fmt.Println(c.Saldo)

}
