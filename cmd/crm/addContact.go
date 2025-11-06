package cmd_crm

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yumeei/go-tp/internal/storage"
)

var (
	inputLastName  string
	inputFirstName string
	inputMail      string
)

var addContactCmd = &cobra.Command{
	Use:   "addContact",
	Short: "add contact command",
	Long:  `add contact command`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		var store = storage.NewMemoryStorage()

		if inputLastName == "" || inputFirstName == "" || inputMail == "" {
			fmt.Println("-> Ajouter un contact")
		}

		if inputLastName == "" {
			fmt.Println("-> Rentrez le nom")
			lastName, _ := reader.ReadString('\n')
			inputLastName = strings.TrimSpace(lastName)
		}

		if inputFirstName == "" {
			fmt.Println("-> Rentrez le prenom")
			firstName, _ := reader.ReadString('\n')
			inputFirstName = strings.TrimSpace(firstName)
		}

		if inputMail == "" {
			fmt.Println("-> Rentrez l'email")
			mail, _ := reader.ReadString('\n')
			inputMail = strings.TrimSpace(mail)
		}

		addedContact, err := storage.NewContact(inputLastName, inputFirstName, inputMail)
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
	},
}

func init() {
	rootCmd.AddCommand(addContactCmd)

	addContactCmd.Flags().StringVarP(&inputLastName, "lastName", "l", "", "Nom de famille du contact")
	addContactCmd.Flags().StringVarP(&inputFirstName, "firstName", "f", "", "Prénom du contact")
	addContactCmd.Flags().StringVarP(&inputMail, "mail", "m", "", "Mail du contact")
}
