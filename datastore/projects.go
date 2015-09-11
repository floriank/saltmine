package datastore

import (
	"github.com/jinzhu/gorm"
	"time"
)

// ProjectStorer defines an interface for ProjectDataStorers
type ProjectStorer interface {
	Get(int) (*Project, error)
	List() ([]*Project, error)
	Create(*Project) (*Project, error)
	Delete(*Project) (*Project, error)
	Update(*Project) (*Project, error)
}

// ProjectStore provides the repository for Projects
type ProjectStore struct {
	db *gorm.DB
}

// Project is the struct for projects
type Project struct {
	ID          int
	Identifier  string
	Title       string
	Description string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewProjectStore(db *gorm.DB) *ProjectStore {
	return &ProjectStore{
		db: db,
	}
}

// Get will return a single Project based on an id given
func (p *ProjectStore) Get(id int) (*Project, error) {
	project := Project{}
	assoc := p.db.Find(&project, id)
	return &project, assoc.Error
}

// List will provide all projects the sotre can find
func (p *ProjectStore) List() ([]*Project, error) {
	var projects []*Project
	assoc := p.db.Find(&projects)
	return projects, assoc.Error
}

// Create will attempt to create a new Project
func (p *ProjectStore) Create(project *Project) (*Project, error) {
	p.db.NewRecord(&project)
	assoc := p.db.Create(&project)
	return project, assoc.Error
}

// Delete will destroy an existing Project based on its given id
func (p *ProjectStore) Delete(project *Project) (*Project, error) {
	assoc := p.db.Delete(project)
	return project, assoc.Error
}

// Update will provide new Data for the fields of a Project
func (p ProjectStore) Update(project *Project) (*Project, error) {
	assoc := p.db.Save(&project)
	return project, assoc.Error
}
