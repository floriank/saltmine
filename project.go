package main

import (
	"net/http"
	"time"
)

// Project is the struct for projects
type Project struct {
	ID          int
	Identifier  string `sql:"size:30"`
	Title       string
	Description string `sql:"type:text'"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// ProjectIndex represents the index handler on GET /projects
func ProjectIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json, charset=UTF8")
	w.Write([]byte("foo"))
}

// ProjectCreate represents the create handler on POST /projects
func ProjectCreate(w http.ResponseWriter, r *http.Request) {
	project := Project{
		Identifier:  "florian",
		Title:       "a demo project",
		Description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Tempora vitae, praesentium iure natus nobis dicta, expedita tempore, deserunt iusto ut accusantium deleniti, enim beatae rem quae facilis ex sapiente quod.",
	}

	db.NewRecord(project)
	_, err := db.Create(&project)

	if err == nil {
		w.Write([]byte("ok"))
	} else {
		w.Write([]byte("not ok"))
	}
}
