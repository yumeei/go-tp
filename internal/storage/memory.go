package storage

import (
	"fmt"
	"sync"
)

// Store is an in-memory implementation of the Storer interface.
type Store struct {
	mu     sync.RWMutex
	list   map[uint]Contact
	nextID uint
}

func NewMemoryStorage() *Store {
	return &Store{
		list:   make(map[uint]Contact),
		nextID: 1,
	}
}

// AjouterContact ajoute un nouveau contact et retourne le contact créé.
func (m *Store) AjouterContact(c *Contact) (*Contact, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	c.ID = m.nextID
	m.list[c.ID] = *c
	m.nextID++

	return c, nil
}

// SupprimerContact supprime un contact par son ID.
func (m *Store) SupprimerContact(id uint) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.list[id]; !ok {
		return fmt.Errorf("le contact avec l'ID %d n'existe pas", id)
	}

	delete(m.list, id)
	return nil
}

// ModifierContact met à jour un contact existant.
func (m *Store) ModifierContact(id uint, c *Contact) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.list[id]; !exists {
		return fmt.Errorf("contact avec ID %d non trouvé", id)
	}

	c.ID = id
	m.list[id] = *c
	return nil
}

// GetContactsList retourne la liste des contacts sous forme de slice de pointeurs.
func (m *Store) GetContactsList() []*Contact {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var contactList []*Contact
	for _, contact := range m.list {
		c := contact
		contactList = append(contactList, &c)
	}
	return contactList
}

// GetContactByID retourne un contact par son ID.
func (m *Store) GetContactByID(id uint) (*Contact, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	contact, exists := m.list[id]
	if !exists {
		return nil, fmt.Errorf("contact avec ID %d non trouvé", id)
	}
	return &contact, nil
}
