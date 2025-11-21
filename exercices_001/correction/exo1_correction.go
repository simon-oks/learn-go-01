package main

import "fmt"

func main() {
	var longueur, largeur float64

	fmt.Print("Longeur: ")
	fmt.Scan(&longueur)

	fmt.Print("Largeur: ")
	fmt.Scan(&largeur)

	perimetre := 2 * (longueur + largeur)
	fmt.Println("Périmètre = ", perimetre)
}
