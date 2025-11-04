package storage

import (
	"fmt"
)

type Store struct {
	list   map[uint]*Contact
	nextID uint
}

func NewMemoryStorage() *Store {
	return &Store{
		list:   make(map[uint]*Contact),
		nextID: 0,
	}
}

func (m *Store) AjouterContact(c *Contact) (*Contact, error) {
	id := m.nextID

	m.list[id] = c

	m.nextID++

	return c, nil
}

func (m *Store) SupprimerContact(id uint) error {
	if _, ok := m.list[id]; !ok {
		return fmt.Errorf("le contact avec l'ID %d n'existe pas", id)
	}

	delete(m.list, id)

	return nil
}

func (m *Store) ModifierContact(id uint, c *Contact) error {
	fmt.Printf("Tentative de modification Id : %v\n", id)

	_, exists := m.list[id]

	if !exists {
		return fmt.Errorf("Contact avec ID %d non trouvé", id)
	}

	m.list[id] = c
	return nil

}

func (m *Store) GetContactsList() []*Contact {
	var contactList []*Contact

	for _, contact := range m.list {
		contactList = append(contactList, contact)
	}

	return contactList
}

func (m *Store) GetContactByID(id uint) (*Contact, error) {
	_, exists := m.list[id]

	if !exists {
		return nil, fmt.Errorf("Contact avec ID %d non trouvé", id)
	}
	return m.list[id], nil
}