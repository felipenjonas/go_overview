package main

import "fmt"

var x int
var y string
var z bool

func main() {

	fmt.Println(x, y, z)
	// Mostra o valor da string em formato de string.
	s := fmt.Sprintf("%v", y)
	fmt.Println(s)

}
