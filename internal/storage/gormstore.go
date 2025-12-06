package storage

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type GORMStore struct {
	db *gorm.DB
}

func NewGORMStore(dsn string) (*GORMStore, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite database: %w", err)
	}

	if err := db.AutoMigrate(&Contact{}); err != nil {
		return nil, fmt.Errorf("auto migrate failed: %w", err)
	}

	return &GORMStore{
		db: db,
	}, nil
}

func (g *GORMStore) AjouterContact(contact *Contact) (*Contact, error) {
	if contact == nil {
		return nil, fmt.Errorf("nil contact provided")
	}

	result := g.db.Create(contact)
	if result.Error != nil {
		return nil, result.Error
	}
	return contact, nil
}

func (g *GORMStore) SupprimerContact(id uint) error {
	if id == 0 {
		return fmt.Errorf("invalid id")
	}
	result := g.db.Delete(&Contact{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrContactNotFound(id)
	}
	return nil
}

func (g *GORMStore) ModifierContact(id uint, c *Contact) error {
	if id == 0 {
		return fmt.Errorf("invalid id")
	}
	if c == nil {
		return fmt.Errorf("nil contact provided")
	}

	var existing Contact
	if err := g.db.First(&existing, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrContactNotFound(id)
		}
		return err
	}

	c.ID = id
	if err := g.db.Save(c).Error; err != nil {
		return err
	}
	return nil
}

func (g *GORMStore) GetContactsList() []*Contact {
	var contacts []*Contact
	if err := g.db.Find(&contacts).Error; err != nil {
		fmt.Printf("failed to list contacts: %v\n", err)
		return []*Contact{}
	}
	return contacts
}

func (g *GORMStore) GetContactByID(id uint) (*Contact, error) {
	if id == 0 {
		return nil, fmt.Errorf("invalid id")
	}
	var contact Contact
	if err := g.db.First(&contact, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrContactNotFound(id)
		}
		return nil, err
	}
	return &contact, nil
}
