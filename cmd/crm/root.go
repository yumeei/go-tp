package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yumeei/go-tp/contacts"
	"github.com/yumeei/go-tp/internal/app"
	"github.com/yumeei/go-tp/internal/storage"
)

var (
	cfgFile string
	store   contacts.Store
)

var rootCmd = &cobra.Command{
	Use:   "crm",
	Short: "Un simple CRM en ligne de commande pour gérer les contacts.",
	// PersistentPreRun est exécuté avant la fonction Run de la commande.
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Lire la configuration depuis Viper
		storageType := viper.GetString("storage.type")
		fmt.Printf("Configuration du stockage : '%s'\n", storageType)

		var err error
		switch storageType {
		case "json":
			store, err = storage.NewJSONStore("contacts.json")
			if err != nil {
				log.Fatalf("Erreur lors de la création du JSON store: %v", err)
			}
			fmt.Println("Utilisation du stockage JSON.")
		case "memory":
			store = storage.NewMemoryStorage()
			fmt.Println("Utilisation du stockage en mémoire.")
		default:
			log.Fatalf("Type de stockage non supporté dans la configuration : %s", storageType)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		app.Run(store)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// La fonction init est appelée avant main().
	// On y attache la configuration de Viper.
	cobra.OnInitialize(initConfig)

	// On ajoute un flag persistant pour spécifier un fichier de config.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "fichier de configuration (par défaut: ./config.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		// Utiliser le fichier de config spécifié par le flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Chercher le fichier config.yaml dans le répertoire courant.
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	// Définir une valeur par défaut pour le type de stockage
	viper.SetDefault("storage.type", "json")

	viper.AutomaticEnv() // read in environment variables that match

	// Si un fichier de config est trouvé, le lire.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Utilisation du fichier de configuration:", viper.ConfigFileUsed())
	} else {
		// Si le fichier n'est pas trouvé, ce n'est pas forcément une erreur,on peut utiliser les valeurs par défaut.
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Fichier de configuration non trouvé, utilisation des valeurs par défaut/flags.")
		} else {
			// Une autre erreur est survenue lors de la lecture du fichier.
			log.Fatalf("Erreur de lecture du fichier de configuration: %s \n", err)
		}
	}
}
