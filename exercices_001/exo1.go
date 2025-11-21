package main

import "fmt"

func main() {
	var (
		longueur int
		largeur  int
	)

	fmt.Print("Entrer la largeur: ")
	fmt.Scan(&longueur)

	fmt.Print("Entrer la largeur: ")
	fmt.Scan(&largeur)

	fmt.Printf("\nPour une logueur de: %v, et un largeur de %v, Le périmètre est: %v\n", longueur, largeur, (longueur+largeur)*2)
}
