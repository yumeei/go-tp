package contacts

import (
	"errors"
	"fmt"
)

type Manager struct {
	list   map[uint]Contact
	nextID uint
}

func NewManager() *Manager {
	return &Manager{
		list:   make(map[uint]Contact),
		nextID: 0,
	}
}

func (m *Manager) AjouterContact(c Contact) (Contact, error) {
	id := m.nextID

	m.list[id] = c

	m.nextID++

	return c, nil
}

func (m *Manager) SupprimerContact() {

}

func (m *Manager) ModifierContact(id uint, c Contact) (Contact, error) {
	fmt.Printf("Tentative de modification Id : %v\n", id)

	// 1. Vérifier si l'ID existe dans la map
	// La syntaxe "_, exists" est la façon idiomatique de vérifier
	// l'existence d'une clé dans une map.
	_, exists := m.list[id]

	if !exists {
		// Si l'ID n'existe pas, on ne peut pas modifier, on retourne une erreur
		return Contact{}, errors.New(fmt.Sprintf("Contact avec ID %d non trouvé", id))
	}

	// 2. Si l'ID existe, on le modifie
	m.list[id] = c
	return c, nil

}

func (m *Manager) GetContactsList() []Contact {
	var contactList []Contact

	for _, contact := range m.list {
		contactList = append(contactList, contact)
	}

	return contactList
}
