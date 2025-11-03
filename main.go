package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/yumeei/go-tp/contacts"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var contactList []contacts.Contact

	var contactManager = contacts.NewManager()

	for {
		fmt.Println("Voici tous les menus disponibles:")
		fmt.Println("1- Ajouter un contact")
		fmt.Println("2- Lister tous les contacts")
		fmt.Println("3- Supprimer un contact")
		fmt.Println("4- Mettre à jour un contact")
		fmt.Println("5- Quitter l'application")
		fmt.Print("Choisissez votre menu: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Println("-> Ajouter un contact")
			fmt.Println("-> Rentrez le nom")
			inputNom, _ := reader.ReadString('\n')
			inputNom = strings.TrimSpace(inputNom)
			fmt.Println("-> Rentrez le prenom")
			inputPrenom, _ := reader.ReadString('\n')
			inputPrenom = strings.TrimSpace(inputPrenom)
			fmt.Println("-> Rentrez l'email")
			inputMail, _ := reader.ReadString('\n')
			inputMail = strings.TrimSpace(inputMail)

			addContact := contacts.Contact{
				ID:     0,
				Nom:    inputNom,
				Prenom: inputPrenom,
				Email:  inputMail,
			}
			addContact, error := contactManager.AjouterContact(addContact)
			if error != nil {
				fmt.Printf("Une erreur est survenue: %v\n\n", error)
			} else {
				fmt.Printf("Utilisateur ajouté: %v \n\n", addContact)
			}

			updateContact := contacts.Contact{
				ID:     0,
				Nom:    "test",
				Prenom: "test",
				Email:  "test",
			}
			updateContact, error2 := contactManager.ModifierContact(updateContact)
			if error != nil {
				fmt.Printf("Une erreur est survenue: %v\n\n", error2)
			} else {
				fmt.Printf("Utilisateur modifié: %v \n\n", updateContact)
			}

		case "2":
			fmt.Println("-> Lister tous les contacts")
			contactList = contactManager.GetContactsList()
			if len(contactList) != 0 {
				for i := 0; i < len(contactList); i++ {
					fmt.Printf("Utilisateur n° %v : %v, %v, %v \n\n", i, contactList[i].Prenom, contactList[i].Nom, contactList[i].Email)
				}
			} else {
				fmt.Println("Pas de contact enregistrés")
			}
		case "3":
			fmt.Println("-> Supprimer un contact")
		case "4":
			fmt.Println("-> Mettre à jour un contact")
		case "5":
			fmt.Println("-> Quitter l'application")
			os.Exit(0)
		default:
			fmt.Println("Option invalide, essayez encore.")
		}
		fmt.Println()
	}
}
