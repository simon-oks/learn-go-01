package main

import "fmt"

func main() {

	const correctPwd = "azerty"
	var pwd string

	fmt.Print("Entrer le mot de passe: ")
	fmt.Scan(&pwd)

	for pwd != correctPwd {
		fmt.Print("Entrer le mot de passe: ")
		fmt.Scan(&pwd)
	}

	fmt.Println("Accès autorisé")
}
