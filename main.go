package main

import "fmt"

func main() {
	var (
		x int
		y string
		z bool
	)
	//var x int
	//x = 15
	//y := 16

	x = 23
	y = "Simon"
	z = true

	//fmt.Printf("Mon age est: %v!", x)
	//fmt.Printf("\nMon prénom est: %v!", y)
	//fmt.Printf("\nJ'ai plus de 18 ans: %v!\n", z)
	fmt.Printf("Mon age est: %v, mon nom est: %v, j'ai plus de 18 ans: %v\n", x, y, z)
	fmt.Printf("Bonjour à tous, je suis %v, j'ai %vans et donc plus de 18 ans, c'est %v!\n", y, x, z)
}
