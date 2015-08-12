package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

// Project is the struct for projects
type Project struct {
	ID          int
	Identifier  string `sql:"size:30"`
	Title       string
	Description string `sql:"type:text"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// ProjectIndex represents the index handler on GET /projects
func ProjectIndex(w http.ResponseWriter, r *http.Request) {
	var projects []Project
	db.Find(&projects)

	w.Header().Set("Content-Type", "application/json;charset=UTF8")
	json.NewEncoder(w).Encode(projects)
}

// ProjectCreate represents the create handler on POST /projects
func ProjectCreate(w http.ResponseWriter, r *http.Request) {
	project := createProjectFromJSON(r.Body)

	db.NewRecord(project)
	db.Create(&project)

	w.Header().Set("Content-Type", "application/json;charset=UTF8")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(project)
}

func createProjectFromJSON(body io.Reader) Project {
	var project = Project{}
	decoder := json.NewDecoder(body)

	err := decoder.Decode(&project)

	if err != nil {
		log.Println(err.Error())
		panic("could not parse JSON")
	}

	return project
}
