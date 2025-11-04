package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

			addedContact, err := contacts.NewContact(inputNom, inputPrenom, inputMail)
			if err != nil {
				fmt.Printf("Erreur de validation: %v\n\n", err)
				break
			}

			// 2. AJOUT au Manager (La variable err est réutilisée)
			addedContact, err = contactManager.AjouterContact(addedContact) // NOTE: on utilise "="
			if err != nil {
				fmt.Printf("Erreur d'ajout au manager: %v\n\n", err)
			} else {
				fmt.Printf("Utilisateur ajouté: %v \n\n", addedContact)
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
			fmt.Println("-> Entrez l'ID du contact à supprimer :")
			inputId, _ := reader.ReadString('\n')
			inputId = strings.TrimSpace(inputId)
			inputIdInt, err := strconv.ParseInt(inputId, 10, 64)
			if err != nil {
				fmt.Printf("Erreur dans la conversion de l'ID : %v", err)
			} else {
				err = contactManager.SupprimerContact(uint(inputIdInt))
				if err != nil {
					fmt.Printf("Une erreur est survenue : %v\n\n", err)
				} else {
					fmt.Println("Contact supprimé avec succès\n")
				}
			}
		case "4":
			fmt.Println("-> Mettre à jour un contact")
			fmt.Println("-> Rentrez l'ID")
			inputId, _ := reader.ReadString('\n')
			inputId = strings.TrimSpace(inputId)
			inputIdInt, err := strconv.ParseInt(inputId, 10, 64)
			if err != nil {
				fmt.Printf("Erreur dans la conversion de l'ID %v", err)
				break
			}
			fmt.Println("-> Rentrez le nouveau nom")
			inputNom, _ := reader.ReadString('\n')
			inputNom = strings.TrimSpace(inputNom)
			fmt.Println("-> Rentrez le nouveau prenom")
			inputPrenom, _ := reader.ReadString('\n')
			inputPrenom = strings.TrimSpace(inputPrenom)
			fmt.Println("-> Rentrez le nouvel email")
			inputMail, _ := reader.ReadString('\n')
			inputMail = strings.TrimSpace(inputMail)
			updateContact, err := contacts.NewContact(inputNom, inputPrenom, inputMail)
			if err != nil {
				fmt.Printf("Erreur à la création du contact: %v", err)
				break
			}
			updateContact, err = contactManager.ModifierContact(uint(inputIdInt), updateContact)
			if err != nil {
				fmt.Printf("Une erreur est survenue: %v\n\n", err)
			} else {
				fmt.Printf("Utilisateur modifié: %v \n\n", updateContact)
			}
		case "5":
			fmt.Println("-> Quitter l'application")
			os.Exit(0)
		default:
			fmt.Println("Option invalide, essayez encore.")
		}
		fmt.Println()
	}
}
