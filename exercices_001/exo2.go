package main

import "fmt"

func main() {
	var n int
	fmt.Print("Entrer un nombre: ")
	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Println("Pair")
	} else {
		fmt.Println("Impair")
	}
}
