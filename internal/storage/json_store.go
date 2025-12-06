package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
)

// JSONStore implémente l'interface Store pour la persistance des données dans un fichier JSON.
type JSONStore struct {
	mu     sync.Mutex
	path   string
	list   map[uint]Contact
	nextID uint
}

// GetContactByID implements Storer.
func (s *JSONStore) GetContactByID(id uint) (*Contact, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	contact, ok := s.list[id]
	if !ok {
		return nil, fmt.Errorf("le contact avec l'ID %d n'existe pas", id)
	}
	c := contact
	return &c, nil
}

// NewJSONStore crée une nouvelle instance de JSONStore.
func NewJSONStore(path string) (*JSONStore, error) {
	store := &JSONStore{
		path:   path,
		list:   make(map[uint]Contact),
		nextID: 1, // On commence à 1
	}

	if err := store.load(); err != nil {
		return nil, err
	}

	return store, nil
}

// load charge les contacts depuis le fichier JSON.
func (s *JSONStore) load() error {
	file, err := os.Open(s.path)
	if err != nil {
		// Si le fichier n'existe pas, c'est normal au premier lancement.
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	// Si le fichier est vide, pas besoin de désérialiser.
	if len(bytes) == 0 {
		return nil
	}

	if err := json.Unmarshal(bytes, &s.list); err != nil {
		return err
	}

	// Mettre à jour nextID pour éviter les conflits
	for id := range s.list {
		if id >= s.nextID {
			s.nextID = id + 1
		}
	}

	return nil
}

// save sauvegarde les contacts dans le fichier JSON.
func (s *JSONStore) save() error {
	bytes, err := json.MarshalIndent(s.list, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.path, bytes, 0644)
}

// AjouterContact ajoute un nouveau contact et sauvegarde.
func (s *JSONStore) AjouterContact(c *Contact) (*Contact, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	c.ID = s.nextID
	s.list[c.ID] = *c
	s.nextID++

	return c, s.save()
}

// SupprimerContact supprime un contact par son ID et sauvegarde.
func (s *JSONStore) SupprimerContact(id uint) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.list[id]; !ok {
		return fmt.Errorf("le contact avec l'ID %d n'existe pas", id)
	}

	delete(s.list, id)
	return s.save()
}

// ModifierContact modifie un contact existant et sauvegarde.
func (s *JSONStore) ModifierContact(id uint, c *Contact) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.list[id]; !ok {
		return fmt.Errorf("contact avec ID %d non trouvé", id)
	}

	c.ID = id // Assure que l'ID est conservé
	s.list[id] = *c
	return s.save()
}

// GetContactsList retourne la liste de tous les contacts
func (s *JSONStore) GetContactsList() []*Contact {
	s.mu.Lock()
	defer s.mu.Unlock()

	var contactList []*Contact
	for _, contact := range s.list {
		c := contact
		contactList = append(contactList, &c)
	}
	return contactList
}
