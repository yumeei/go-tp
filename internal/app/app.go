package app

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/yumeei/go-tp/internal/storage"
)

func Run(store storage.Storer) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Voici tous les menus disponibles:")
		fmt.Println("1- Ajouter un contact")
		fmt.Println("2- Lister tous les contacts")
		fmt.Println("3- Supprimer un contact")
		fmt.Println("4- Mettre à jour un contact")
		fmt.Println("5- Chercher un contact par ID")
		fmt.Println("6- Quitter l'application")
		fmt.Print("Choisissez votre menu: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			handleAddContact(reader, store)
		case "2":
			handleListContacts(store)
		case "3":
			handleDeleteContact(reader, store)
		case "4":
			handleUpdateContact(reader, store)
		case "5":
			handleFindByID(reader, store)
		case "6":
			fmt.Println("-> Quitter l'application")
			os.Exit(0)
		default:
			fmt.Println("Option invalide, essayez encore.")
		}
		fmt.Println()
	}
}

func handleAddContact(reader *bufio.Reader, store storage.Storer) {
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

	addedContact, err := storage.NewContact(inputNom, inputPrenom, inputMail)
	if err != nil {
		fmt.Printf("Erreur de validation: %v\n\n", err)
		return
	}
	addedContactValue, err := store.AjouterContact(addedContact)
	if err != nil {
		fmt.Printf("Erreur d'ajout au manager: %v\n\n", err)
	} else {
		fmt.Printf("Utilisateur ajouté: %v \n\n", addedContactValue)
	}
}

func handleListContacts(store storage.Storer) {
	fmt.Println("-> Lister tous les contacts")
	contactList := store.GetContactsList()
	if len(contactList) != 0 {
		for i, contact := range contactList {
			fmt.Printf("Utilisateur n° %v : %v, %v, %v \n\n", i, contact.Prenom, contact.Nom, contact.Email)
		}
	} else {
		fmt.Println("Pas de contact enregistrés")
	}
}

func handleDeleteContact(reader *bufio.Reader, store storage.Storer) {
	fmt.Println("-> Supprimer un contact")
	fmt.Println("-> Entrez l'ID du contact à supprimer :")
	inputId, _ := reader.ReadString('\n')
	inputId = strings.TrimSpace(inputId)
	inputIdInt, err := strconv.ParseInt(inputId, 10, 64)
	if err != nil {
		fmt.Printf("Erreur dans la conversion de l'ID : %v", err)
		return
	}

	err = store.SupprimerContact(uint(inputIdInt))
	if err != nil {
		fmt.Printf("Une erreur est survenue : %v\n\n", err)
	} else {
		fmt.Println("Contact supprimé avec succès")
	}
}

func handleUpdateContact(reader *bufio.Reader, store storage.Storer) {
	fmt.Println("-> Mettre à jour un contact")
	fmt.Println("-> Rentrez l'ID")
	inputId, _ := reader.ReadString('\n')
	inputId = strings.TrimSpace(inputId)
	inputIdInt, err := strconv.ParseInt(inputId, 10, 64)
	if err != nil {
		fmt.Printf("Erreur dans la conversion de l'ID %v", err)
		return
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
	updateContact, err := storage.NewContact(inputNom, inputPrenom, inputMail)
	if err != nil {
		fmt.Printf("Erreur à la création du contact: %v", err)
		return
	}
	err = store.ModifierContact(uint(inputIdInt), updateContact)
	if err != nil {
		fmt.Printf("Une erreur est survenue: %v\n\n", err)
	} else {
		fmt.Printf("Utilisateur modifié \n\n")
	}
}

func handleFindByID(reader *bufio.Reader, store storage.Storer) {
	fmt.Println("-> Chercher par ID")
	fmt.Println("-> Rentrez l'ID")
	inputId, _ := reader.ReadString('\n')
	inputId = strings.TrimSpace(inputId)
	inputIdInt, err := strconv.ParseInt(inputId, 10, 64)
	if err != nil {
		fmt.Printf("Erreur dans la conversion de l'ID %v", err)
		return
	}
	contactFound, err := store.GetContactByID(uint(inputIdInt))
	if err != nil {
		fmt.Printf("Une erreur est survenue : %v", err)
	} else {
		fmt.Printf("Utilisateur trouvé : %v, %v, %v \n\n", contactFound.Prenom, contactFound.Nom, contactFound.Email)
	}
}
