package main

import "fmt"

func main() {
	const password = "azerty"
	var input string

	for {
		fmt.Print("Mot de passe: ")
		fmt.Scan(&input)

		if input == password {
			fmt.Println("Accès autorisé")
			break
		} else {
			fmt.Println("Incorrect, réessayez.")
		}
	}
}
