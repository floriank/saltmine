package api

import (
	"encoding/json"
	. "github.com/floriank/saltmine/datastore"
	"net/http"
)

func (s *SaltmineAPI) ProjectsList() func(http.ResponseWriter, *http.Request) {
	projects, _ := s.projects.List()

	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(projects)
	}
}

func (s *SaltmineAPI) ProjectGet() func(http.ResponseWriter, *http.Request) {
	project, _ := s.projects.Get(1)
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(project)
	}
}

func (s *SaltmineAPI) ProjectUpdate() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(&Project{})
	}
}

func (s *SaltmineAPI) ProjectDelete() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(&Project{})
	}
}

func (s *SaltmineAPI) ProjectCreate() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(&Project{})
	}
}
