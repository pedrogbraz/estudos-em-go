package main

import "fmt"

func main() {
	// Maneiras de declarar variaveis

	var x int = 8                                   // Declaração comprida de variaveis
	var nome, sobrenome string = "Pedro", "Gontijo" // Declaração dupla de variaveis
	nome2 := "Pedro"                                // Declaração simples, eu só posso declarar uma variavel assim dentro de uma função
	teste, teste2 := "Pedro", 18                    // Aqui eu estou declarando duas variaveis em uma linha de tipos diferentes
	// Mas sempre que eu for declarar uma variavel dupla de tipos diferentes, eu ja preciso atribuir um valor a elas

	fmt.Println(teste, teste2)
	fmt.Println(nome, sobrenome, nome2, x)
}
