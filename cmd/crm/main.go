package main

import (
	"log"

	"github.com/yumeei/go-tp/internal/app"
	"github.com/yumeei/go-tp/internal/storage"
)

func main() {
	// On remplace le stockage en mémoire par le stockage JSON.
	// On lui donne le nom du fichier à utiliser pour la sauvegarde.
	store, err := storage.NewJSONStore("contacts.json")
	if err != nil {
		log.Fatalf("Erreur lors de l'initialisation du stockage JSON : %v", err)
	}
	app.Run(store)
}