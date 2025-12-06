package cmd_crm

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yumeei/go-tp/internal/storage"
)

var (
	inputUpdateId   string
	updateLastName  string
	updateFirstName string
	updateMail      string
)

var updateContactCmd = &cobra.Command{
	Use:   "updateContact",
	Short: "update contact command",
	Long:  `update contact command`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		var store = storage.NewMemoryStorage()

		var inputIdInt int64
		var err error

		// ID
		if inputUpdateId == "" {
			fmt.Println("-> Mettre à jour un contact")
			fmt.Println("-> Rentrez l'ID")
			inputId, _ := reader.ReadString('\n')
			inputId = strings.TrimSpace(inputId)
			inputIdInt, err = strconv.ParseInt(inputId, 10, 64)
			if err != nil {
				fmt.Printf("Erreur dans la conversion de l'ID %v\n", err)
				return
			}
		} else {
			inputId := strings.TrimSpace(inputUpdateId)
			inputIdInt, err = strconv.ParseInt(inputId, 10, 64)
			if err != nil {
				fmt.Printf("Erreur dans la conversion de l'ID %v\n", err)
				return
			}
		}

		// Nom
		if updateLastName == "" {
			fmt.Println("-> Rentrez le nouveau nom")
			n, _ := reader.ReadString('\n')
			updateLastName = strings.TrimSpace(n)
		}

		// Prenom
		if updateFirstName == "" {
			fmt.Println("-> Rentrez le nouveau prenom")
			p, _ := reader.ReadString('\n')
			updateFirstName = strings.TrimSpace(p)
		}

		// Email
		if updateMail == "" {
			fmt.Println("-> Rentrez le nouvel email")
			e, _ := reader.ReadString('\n')
			updateMail = strings.TrimSpace(e)
		}

		updateContact, err := storage.NewContact(updateLastName, updateFirstName, updateMail)
		if err != nil {
			fmt.Printf("Erreur à la création du contact: %v\n", err)
			return
		}

		err = store.ModifierContact(uint(inputIdInt), updateContact)
		if err != nil {
			fmt.Printf("Une erreur est survenue: %v\n\n", err)
		} else {
			fmt.Printf("Utilisateur modifié \n\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(updateContactCmd)

	updateContactCmd.Flags().StringVarP(&inputUpdateId, "contactId", "c", "", "Id du contact à modifier")
	updateContactCmd.Flags().StringVarP(&updateLastName, "lastName", "l", "", "Nouveau nom de famille du contact")
	updateContactCmd.Flags().StringVarP(&updateFirstName, "firstName", "f", "", "Nouveau prénom du contact")
	updateContactCmd.Flags().StringVarP(&updateMail, "mail", "m", "", "Nouveau mail du contact")
}
