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

func (m *Manager) SupprimerContact(id uint) error {
	if _, ok := m.list[id]; !ok {
		return fmt.Errorf("le contact avec l'ID %d n'existe pas", id)
	}

	delete(m.list, id)

	return nil
}

func (m *Manager) ModifierContact(id uint, c Contact) (Contact, error) {
	fmt.Printf("Tentative de modification Id : %v\n", id)

	_, exists := m.list[id]

	if !exists {
		return Contact{}, errors.New(fmt.Sprintf("Contact avec ID %d non trouvé", id))
	}

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
