package api

import (
	. "github.com/floriank/saltmine/datastore"
	"time"
)

var api *SaltmineAPI

func init() {
	api = &SaltmineAPI{
		projects: &MockProjectsStore{},
		tickets:  &MockTicketsStore{},
		version:  "0.0.0",
	}
}

type MockProjectsStore struct{}

type MockTicketsStore struct{}

func (p *MockProjectsStore) Get(id int) (*Project, error) {
	return &Project{
		12,
		"testIdentifier",
		"Title",
		"Lorem ipsum dolor sit amet, consectetur adipisicing elit. Minima, incidunt tempore, itaque magni totam quis ipsum atque vero, perferendis sequi ducimus dolores. Ducimus harum consequuntur, iste explicabo totam labore dolores!",
		time.Now(),
		time.Now(),
		nil,
	}, nil
}

func (p *MockProjectsStore) List() ([]*Project, error) {
	var projects []*Project

	proj1 := &Project{
		ID:          1,
		Identifier:  "foo",
		Title:       "Foo",
		Description: "lorem",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
	}

	proj2 := &Project{
		ID:          2,
		Identifier:  "bar",
		Title:       "Bar",
		Description: "ispum",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
	}

	projects = append(projects, proj1)
	projects = append(projects, proj2)

	return projects, nil
}

func (p *MockProjectsStore) Create(project *Project) (*Project, error) {
	return &Project{
		12,
		"testIdentifier",
		"Title",
		"Lorem ipsum dolor sit amet, consectetur adipisicing elit. Minima, incidunt tempore, itaque magni totam quis ipsum atque vero, perferendis sequi ducimus dolores. Ducimus harum consequuntur, iste explicabo totam labore dolores!",
		time.Now(),
		time.Now(),
		nil,
	}, nil
}

func (p *MockProjectsStore) Delete(project *Project) (*Project, error) {
	return &Project{
		12,
		"testIdentifier",
		"Title",
		"Lorem ipsum dolor sit amet, consectetur adipisicing elit. Minima, incidunt tempore, itaque magni totam quis ipsum atque vero, perferendis sequi ducimus dolores. Ducimus harum consequuntur, iste explicabo totam labore dolores!",
		time.Now(),
		time.Now(),
		nil,
	}, nil
}

func (p *MockProjectsStore) Update(project *Project) (*Project, error) {
	return &Project{
		12,
		"testIdentifier",
		"Title",
		"Lorem ipsum dolor sit amet, consectetur adipisicing elit. Minima, incidunt tempore, itaque magni totam quis ipsum atque vero, perferendis sequi ducimus dolores. Ducimus harum consequuntur, iste explicabo totam labore dolores!",
		time.Now(),
		time.Now(),
		nil,
	}, nil
}

var ticketProject = Project{
	ID:          42,
	Identifier:  "project-with-tickets",
	Title:       "Project with Tickets",
	Description: "A simple project with tickets",
}

func (t *MockTicketsStore) Get(id int) (*Ticket, error) {
	return &Ticket{
		ID:          1,
		Subject:     "A problem",
		Description: "A serious problem",
		Project:     ticketProject,
	}, nil
}

func (t *MockTicketsStore) List() ([]*Ticket, error) {
	var tickets []*Ticket

	ticket1 := &Ticket{
		ID:          1,
		Subject:     "A problem",
		Description: "A serious problem",
		Project:     ticketProject,
	}

	ticket2 := &Ticket{
		ID:          2,
		Subject:     "An even more serious problem",
		Description: "A seriously serious problem",
		Project:     ticketProject,
	}

	tickets = append(tickets, ticket1)
	tickets = append(tickets, ticket2)
	return tickets, nil
}

func (t *MockTicketsStore) Create(ticket *Ticket) (*Ticket, error) {
	return &Ticket{
		ID:          1,
		Subject:     "A problem",
		Description: "A serious problem",
		Project:     ticketProject,
	}, nil
}

func (t *MockTicketsStore) Delete(ticket *Ticket) (*Ticket, error) {
	return &Ticket{
		ID:          1,
		Subject:     "A problem",
		Description: "A serious problem",
		Project:     ticketProject,
	}, nil
}

func (t *MockTicketsStore) Update(ticket *Ticket) (*Ticket, error) {
	return &Ticket{
		ID:          1,
		Subject:     "A problem",
		Description: "A serious problem",
		Project:     ticketProject,
	}, nil
}
