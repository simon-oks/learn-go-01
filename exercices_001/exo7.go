package main

import "fmt"

func main() {
	var n int
	fmt.Print("Entrer un nombre: ")
	fmt.Scan(&n)

	count := 0
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			count++
		}
	}

	fmt.Printf("Entre 1 et %v, il y %v nombrer pairs\n", n, count)

}
