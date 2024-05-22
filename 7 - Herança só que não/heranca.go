package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura uint8
}

type estudante struct {
	pessoa
	curso string
	faculdade string
}

func main() {
	fmt.Println("Herança é uma forma de reutilizar código em Go")

	p1 := pessoa{"João", "Pedro", 20, 178}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1)

	e2 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e2.nome)
	fmt.Println(e2.faculdade)
}