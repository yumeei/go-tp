package cmd_crm

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yumeei/go-tp/internal/app"
	"github.com/yumeei/go-tp/internal/storage"
)

var rootCmd = &cobra.Command{
	Use:   "contact",
	Short: "contact est un outil qui permet de gérer et modifier sa liste de contacts facilement",
	Long:  `contact est un outil qui permet de gérer et modifier sa liste de contacts facilement`,
	Run: func(cmd *cobra.Command, args []string) {
		var store = storage.NewMemoryStorage()
		app.Run(*store)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func init() {

}
