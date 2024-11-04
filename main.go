package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hello World !")
	//var somme int = add(14, 27) ou bien:
	somme := add(14, 77)
	fmt.Println(somme)
	//f string
	fmt.Printf("La somme est de : %v\n", somme)
}
