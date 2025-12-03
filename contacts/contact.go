package contacts

import "fmt"

type Contact struct {
	ID     uint   `json:"id"`
	Nom    string `json:"nom"`
	Prenom string `json:"prenom"`
	Email  string `json:"email"`
}

func NewContact(nom string, prenom string, email string) (*Contact, error) {
	if nom == "" || prenom == "" || email == "" {
		return nil, fmt.Errorf("tous les champs doivent être remplis")
	}
	return &Contact{
		Nom:    nom,
		Prenom: prenom,
		Email:  email,
	}, nil
}

type Store interface {
	AjouterContact(c Contact) (Contact, error)
	SupprimerContact(id uint) error
	ModifierContact(id uint, c Contact) (Contact, error)
	GetContactsList() []Contact
	GetContactByID(id uint) (Contact, error)
}
