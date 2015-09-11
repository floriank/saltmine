package datastore

import (
	"github.com/jinzhu/gorm"
	"time"
)

// TicketStorer defines an interface for TicketDataStorers
type TicketStorer interface {
	Get(int) (*Ticket, error)
	List() ([]*Ticket, error)
	Create(*Ticket) (*Ticket, error)
	Delete(*Ticket) (*Ticket, error)
	Update(*Ticket) (*Ticket, error)
}

// TicketStore provides the repository for Tickets
type TicketStore struct {
	db *gorm.DB
}

// Ticket is the struct for Tickets
type Ticket struct {
	ID          int
	Subject     string
	Description string
	Project     Project

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewTicketStore(db *gorm.DB) *TicketStore {
	return &TicketStore{
		db: db,
	}
}

// Get will return a single Ticket based on an id given
func (p *TicketStore) Get(id int) (*Ticket, error) {
	ticket := Ticket{}
	assoc := p.db.Find(&ticket, id)
	return &ticket, assoc.Error
}

// List will provide all Tickets the store can find
func (p *TicketStore) List() ([]*Ticket, error) {
	var tickets []*Ticket
	assoc := p.db.Find(&tickets)
	return tickets, assoc.Error
}

// Create will attempt to create a new Ticket
func (p *TicketStore) Create(ticket *Ticket) (*Ticket, error) {
	p.db.NewRecord(&ticket)
	assoc := p.db.Create(&ticket)
	return ticket, assoc.Error
}

// Delete will destroy an existing Ticket based on its given id
func (p *TicketStore) Delete(ticket *Ticket) (*Ticket, error) {
	assoc := p.db.Delete(ticket)
	return ticket, assoc.Error
}

// Update will provide new Data for the fields of a Ticket
func (p *TicketStore) Update(ticket *Ticket) (*Ticket, error) {
	assoc := p.db.Save(ticket)
	return ticket, assoc.Error
}
