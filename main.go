package main // Obrigatoriamente sempre preciso declarar o pacote main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	res, rem := dividir(10, 3) // Aqui eu estou iniciando a variavel, então sempre sera com :=
	fmt.Println(res, rem)

	f := somar(2)
	x := f(1)
	fmt.Println(x)

	foo := func() { // Aqui eu estou criando uma funcao anonima
		fmt.Println("Executando a funcao foo")
	}

	foo()

	fmt.Println(somar2(10, 10, 10))
}

func dividir(a, b int) (res, rem int) { // Aqui eu estou criando uma função para dividir os valores e retornar o resultado e o resto
	res = a / b // Aqui eu estou apenas atualizando os valores de res e rem
	rem = a % b
	return res, rem // Aqui eu estou retornando os valores
}

// Função Higher Order
func somar(a int) func(int) int { // Função que retorna outra função
	return func(b int) int { // Função closure
		return a + b
	}
}

func somar2(nums ...int) int { // Função variatica
	var out int
	for _, n := range nums { // Aqui eu estou percorrendo os valores
		out += n
	}
	return out
}
