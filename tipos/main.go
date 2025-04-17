package main

import "fmt"

func main() {
	// Tipos de variaveis
	// Números Inteiros
	var x int8 = 127                  // 8 bits (de -128 a 127)
	var y int16 = 32767               // 16 bits (de -32.768 a 32.767)
	var z int32 = 2147483647          // 32 bits (de -2.147.483.648 a 2.147.483.647)
	var w int64 = 9223372036854775807 // 64 bits (de -9.223.372.036.854.775.808 a 9.223.372.036.854.775.807)

	// Números com ponto flutuante
	var a float32 = 3.14              // 32 bits
	var b float64 = 3.141592653589793 // 64 bits

	// Booleanos
	var c bool = true
	var d bool = false

	// Strings
	var e string = "Ola, Mundo!"

	// Ponteiros
	var f *int = nil

	fmt.Println(x, y, z, w, a, b, c, d, e, f)
}
