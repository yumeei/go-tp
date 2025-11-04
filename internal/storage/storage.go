package storage

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

type Storer interface {
	AjouterContact(contact *Contact) (*Contact, error)
	SupprimerContact(id uint) error
	ModifierContact(id uint, c *Contact) error
	GetContactsList() ([]*Contact, error)
	GetContactByID(id int) (*Contact, error)
}

var ErrContactNotFound = func(id int) error {
	return fmt.Errorf("Contact avec l'ID %d non trouvé", id)
}
