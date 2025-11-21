package main

import (
	"fmt"
)

func main() {
	var (
		n1       int
		operator string
		n2       int
	)

	fmt.Print("Entrer le premier nombre: ")
	fmt.Scan(&n1)
	fmt.Print("Entrer l'opérateur: ")
	fmt.Scan(&operator)
	fmt.Print("Entrer le deuxième nombre: ")
	fmt.Scan(&n2)

	if operator == "+" {
		fmt.Printf("%v %v %v = %v\n", n1, operator, n2, n1+n2)
	} else if operator == "-" {
		fmt.Printf("%v %v %v = %v\n", n1, operator, n2, n1-n2)
	} else if operator == "*" {
		fmt.Printf("%v %v %v = %v\n", n1, operator, n2, n1*n2)
	} else if operator == "/" {
		fmt.Printf("%v %v %v = %v\n", n1, operator, n2, n1/n2)
	} else if operator == "%" {
		fmt.Printf("%v %v %v = %v\n", n1, operator, n2, n1%n2)
	} else {
		fmt.Println("Opérateur nom prise en charge!")
	}
}
