package cmd_crm

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/yumeei/go-tp/internal/app"
	"github.com/yumeei/go-tp/internal/storage"
)

var sqlitePath string

var rootCmd = &cobra.Command{
	Use:   "contact",
	Short: "contact est un outil qui permet de gérer et modifier sa liste de contacts facilement",
	Long:  `contact est un outil qui permet de gérer et modifier sa liste de contacts facilement`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running cobra command")
		var st storage.Storer
		if sqlitePath != "" {
			gormStore, err := storage.NewGORMStore(sqlitePath)
			if err != nil {
				log.Printf("Erreur à l'ouverture de la base SQLite (%s): %v. Utilisation du stockage en mémoire.", sqlitePath, err)
				st = storage.NewMemoryStorage()
			} else {
				st = gormStore
			}
		} else {
			st = storage.NewMemoryStorage()
		}

		app.Run(st)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&sqlitePath, "db", "d", "contacts.db", "Chemin vers la base SQLite (par défaut contacts.db). Si vide, utilise le stockage en mémoire.")
}
