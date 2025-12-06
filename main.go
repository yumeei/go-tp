package main

import (
	"fmt"

	cmd_crm "github.com/yumeei/go-tp/cmd/crm"
)

func main() {
	fmt.Println("DEBUG: démarrage de main")
	cmd_crm.Execute()
}
