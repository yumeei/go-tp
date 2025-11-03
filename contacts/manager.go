package contacts

import "fmt"

type Manager struct {
	list   map[uint]Contact
	nextID uint
}

func NewManager() *Manager {
	return &Manager{
		list:   make(map[uint]Contact),
		nextID: 1,
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

func (m *Manager) ModifierContact(c Contact) (Contact, error) {
	id := c.ID
	fmt.Printf("Id : %v", id)

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
