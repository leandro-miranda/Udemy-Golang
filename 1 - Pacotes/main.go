package main

import (
		"fmt"
		"modulo/auxiliar"
		"github.com/badoux/checkmail"
)

func main(){
		fmt.Println("Escrevendo do arquivo main")
		auxiliar.Escrever()

		erro := checkmail.ValidateFormat("devbook@gmail.com")
		fmt.Println(erro)

		erro2 := checkmail.ValidateFormat("123")
		fmt.Println(erro2)
}