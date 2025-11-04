package contacts

import "fmt"

type Contact struct {
	Nom    string
	Prenom string
	Email  string
}

func NewContact(nom string, prenom string, email string) (*Contact, error) {
	if nom == "" || prenom == "" || email == "" {
		return nil, fmt.Errorf("Tous les champs doivent être remplis")
	}
	return &Contact{
		Nom:    nom,
		Prenom: prenom,
		Email:  email,
	}, nil
}
