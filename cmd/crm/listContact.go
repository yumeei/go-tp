package cmd_crm

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yumeei/go-tp/internal/storage"
)

var listContactCmd = &cobra.Command{
	Use:   "listContact",
	Short: "list contact command",
	Long:  `list contact command`,
	Run: func(cmd *cobra.Command, args []string) {
		var store = storage.NewMemoryStorage()

		fmt.Println("-> Lister tous les contacts")
		contactList := store.GetContactsList()
		if len(contactList) != 0 {
			for i := 0; i < len(contactList); i++ {
				fmt.Printf("Utilisateur n° %v : %v, %v, %v \n\n", i, contactList[i].Prenom, contactList[i].Nom, contactList[i].Email)
			}
		} else {
			fmt.Println("Pas de contact enregistrés")
		}
	},
}

func init() {
	rootCmd.AddCommand(listContactCmd)
}
