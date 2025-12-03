package storage

import (
	"fmt"

	"github.com/yumeei/go-tp/contacts"
)

type Store struct {
	list   map[uint]contacts.Contact
	nextID uint
}

func NewMemoryStorage() *Store {
	return &Store{
		list:   make(map[uint]contacts.Contact),
		nextID: 0,
	}
}

func (m *Store) AjouterContact(c contacts.Contact) (contacts.Contact, error) {
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

func (m *Store) ModifierContact(id uint, c contacts.Contact) (contacts.Contact, error) {
	fmt.Printf("Tentative de modification Id : %v\n", id)

	_, exists := m.list[id]

	if !exists {
		return contacts.Contact{}, fmt.Errorf("contact avec ID %d non trouvé", id)
	}

	m.list[id] = c
	return c, nil

}

func (m *Store) GetContactsList() []contacts.Contact {
	var contactList []contacts.Contact

	for _, contact := range m.list {
		contactList = append(contactList, contact)
	}

	return contactList
}

func (m *Store) GetContactByID(id uint) (contacts.Contact, error) {
	_, exists := m.list[id]

	if !exists {
		return contacts.Contact{}, fmt.Errorf("contact avec ID %d non trouvé", id)
	}
	return m.list[id], nil
}