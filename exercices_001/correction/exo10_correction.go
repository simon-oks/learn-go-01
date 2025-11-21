package main

import "fmt"

func main() {
	var a, b float64
	var op string

	fmt.Print("Nombre 1: ")
	fmt.Scan(&a)

	fmt.Print("Opérateur (+ - * /): ")
	fmt.Scan(&op)

	fmt.Print("Nombre 2: ")
	fmt.Scan(&b)

	switch op {
	case "+":
		fmt.Println("Résultat =", a+b)
	case "-":
		fmt.Println("Résultat =", a-b)
	case "*":
		fmt.Println("Résultat =", a*b)
	case "/":
		if b != 0 {
			fmt.Println("Résultat =", a/b)
		} else {
			fmt.Println("Erreur: division par zéro")
		}
	default:
		fmt.Println("Opérateur invalide.")
	}
}
