package main

import (
	"fmt"
	"bufio"
  "os"
	"strings"
)

func main() {
	reader:= bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Voici tous les menus disponibles:")
		fmt.Println("Ajouter un contact")
		fmt.Println("Lister tous les contacts")
		fmt.Println("Supprimer un contact")
		fmt.Println("Mettre à jour un contact")
		fmt.Println("Quitter l'application")
		fmt.Print("Choisissez votre menu: ")

		input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)

		switch input {
			case "1":
				fmt.Println("-> Ajouter un contact")
			case "2":
				fmt.Println("-> Lister tous les contacts")
			case "3":
				fmt.Println("-> Supprimer un contact")
			case "4":
				fmt.Println("-> Mettre à jour un contact")
			case "5":
				fmt.Println("-> Quitter l'application")
			default:
        fmt.Println("Option invalide, essayez encore.")
			}
    fmt.Println()
	}
}