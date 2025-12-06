package main

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
	inputContactId string
)

var delContactCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		var store = storage.NewMemoryStorage()

		var inputIdInt int64
		var err error

		if inputContactId == "" {
			fmt.Println("-> Supprimer un contact")
			inputId, _ := reader.ReadString('\n')
			inputId = strings.TrimSpace(inputId)
			inputIdInt, err = strconv.ParseInt(inputId, 10, 64)
			if err != nil {
				fmt.Printf("Erreur dans la conversion de l'ID : %v", err)
				return
			}

		} else {
			inputId := inputContactId
			inputId = strings.TrimSpace(inputId)
			inputIdInt, err = strconv.ParseInt(inputId, 10, 64)
			if err != nil {
				fmt.Printf("Erreur dans la conversion de l'ID : %v", err)
				return
			}
		}

		err = store.SupprimerContact(uint(inputIdInt))
		if err != nil {
			fmt.Printf("Une erreur est survenue : %v\n\n", err)
		} else {
			fmt.Println("Contact supprimé avec succès")
		}
	},
}

func init() {
	rootCmd.AddCommand(delContactCmd)

	addContactCmd.Flags().StringVarP(&inputContactId, "contactId", "c", "", "Id du contact à supprimer")
}
