package main

import "fmt"

func main() {
	var n int
	fmt.Print("Entrer un nombre: ")
	fmt.Scan(&n)

	if n > 0 {
		fmt.Println("Positif")
	} else if n < 0 {
		fmt.Println("Négatif")
	} else {
		fmt.Println("Zéro")
	}
}
