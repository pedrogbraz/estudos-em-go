package main // Obrigatoriamente sempre preciso declarar o pacote main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	res, rem := dividir(10, 3) // Aqui eu estou iniciando a variavel, então sempre sera com :=
	fmt.Println(res, rem)
}

func dividir(a, b int) (res, rem int) { // Aqui eu estou criando uma função para dividir os valores e retornar o resultado e o resto
	res = a / b // Aqui eu estou apenas atualizando os valores de res e rem
	rem = a % b
	return res, rem // Aqui eu estou retornando os valores
}
