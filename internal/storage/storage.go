package storage

import "fmt"

type Contact struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Nom    string `json:"nom"`
	Prenom string `json:"prenom"`
	Email  string `json:"email"`
}

func NewContact(nom string, prenom string, email string) (*Contact, error) {
	if nom == "" || prenom == "" || email == "" {
		return nil, fmt.Errorf("Tous les champs doivent être remplis")
	}
	return &Contact{
		ID:     0,
		Nom:    nom,
		Prenom: prenom,
		Email:  email,
	}, nil
}

type Storer interface {
	AjouterContact(contact *Contact) (*Contact, error)
	SupprimerContact(id uint) error
	ModifierContact(id uint, c *Contact) error
	GetContactsList() []*Contact
	GetContactByID(id uint) (*Contact, error)
}

var ErrContactNotFound = func(id uint) error {
	return fmt.Errorf("Contact avec l'ID %d non trouvé", id)
}
