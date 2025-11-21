package main

import "fmt"

func main() {
	var n int
	fmt.Print("Entrer un nombre: ")
	fmt.Scan(&n)

	som := 0
	for i := 1; i <= n; i++ {
		som += i
	}

	fmt.Printf("La somme des nombres entiers entrer 1 et %v inclus est %v\n", n, som)
}
