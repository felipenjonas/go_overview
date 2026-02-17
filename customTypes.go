package main

import "fmt"

//Definição de tipo personalizado.
type hotdog int

var b hotdog = 42

func main() {

	x := 10
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	//Exemplo Sem conversão de tipo.
	// x = 101.1 //Erro de compilação.
	//Conversão de tipo.
	x = int(b)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
