package main

import (
	"errors"
	"fmt"
)

func main() {
		var numero int32 = 100000000
		// numero := 100000000000000
		fmt.Println(numero)	

		var numero2 uint32 = 100000000
		fmt.Println(numero2)

		// alias
		// int32 = rune
		var numero3 rune = 123456
		fmt.Println(numero3)

		// uint8 = byte
		var numero4 byte = 123
		fmt.Println(numero4)

		// números reais
		var numeroReal1 float32 = 123.45
		fmt.Println(numeroReal1)

		var numeroReal2 float64 = 1230000000.45
		fmt.Println(numeroReal2)

		numeroReal3 := 1230000000.45
		fmt.Println(numeroReal3)

		// fim números reais

		// strings
		var str string = "Texto"
		fmt.Println(str)

		str2 := "Texto2"
		fmt.Println(str2)

		char := 'B'
		fmt.Println(char)

		// fim strings

		var texto string
		fmt.Println(texto) // valor padrão string vazia

		var texto2 int16
		fmt.Println(texto2) // valor padrão 0

		texto3 := 5
		fmt.Println(texto3)

		// booleanos
		var booleano1 bool
		fmt.Println(booleano1) // valor padrão false

		// erro
		var erro error
		fmt.Println(erro) // valor padrão nil

		var erro2 error = errors.New("Erro interno")
		fmt.Println(erro2)
}