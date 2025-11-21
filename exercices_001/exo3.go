package main

import "fmt"

func main() {
	var (
		n1 int
		n2 int
	)

	fmt.Print("Entrer le premier nombre: ")
	fmt.Scan(&n1)

	fmt.Print("Entrer le deuxième nombre: ")
	fmt.Scan(&n2)

	if n1 > n2 {
		fmt.Println("Le plus grand nombre est:", n1)
	} else {
		fmt.Println("Le plus grand nombre est:", n2)
	}
}
