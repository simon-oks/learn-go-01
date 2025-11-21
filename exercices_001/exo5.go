package main

import "fmt"

func main() {
	var n int
	fmt.Print("Entrer un nombre: ")
	fmt.Scan(&n)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%v * %v = %v\n", n, i, n*i)
	}
}
