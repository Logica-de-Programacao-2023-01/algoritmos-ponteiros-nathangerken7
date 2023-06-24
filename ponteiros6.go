package main

import "fmt"

// Escreva uma função em Go que receba um ponteiro para um struct Livro com campos título e autor,
// e altere o título do livro para "Desconhecido" se o autor for "Anônimo".

type Livro struct {
	Titulo string
	Autor  string
}

func anonimo(l *Livro) {

	if l.Autor == "Desconhecido" {

		l.Titulo = "Anônimo"

	}

}

func main() {

	l := Livro{

		Titulo: "Granny, a obra",
		Autor:  "Felipão",
	}

	l2 := Livro{
		Titulo: "123456",
		Autor:  "Desconhecido",
	}

	anonimo(&l)
	anonimo(&l2)

	fmt.Println(l)
	fmt.Println(l2)

}
