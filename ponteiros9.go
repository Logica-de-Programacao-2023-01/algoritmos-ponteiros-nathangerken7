package main

import "fmt"

type Livro2 struct {
	Titulo string
	Autor  string
	Preco  float64
}

func desconto(l *Livro2) {

	l.Preco -= l.Preco * 0.1

}

func main() {

	l := Livro2{
		Titulo: "123456",
		Autor:  "1234545677686",
		Preco:  500,
	}

	desconto(&l)

	fmt.Println(l.Preco)

}
