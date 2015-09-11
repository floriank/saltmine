package api

import (
	"encoding/json"
	. "github.com/floriank/saltmine/datastore"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (s *SaltmineAPI) ProjectsList() func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		projects, _ := s.projects.List()
		json.NewEncoder(w).Encode(projects)
	}
}

func (s *SaltmineAPI) ProjectGet() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		project, err := s.projects.Get(id)

		if err != nil {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}

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
		project := Project{}

		decoderErr := json.NewDecoder(r.Body).Decode(&project)
		if decoderErr != nil {
			// golang does not support "Unprocessable entity"
			http.Error(w, "Could not read from input", 422)
			return
		}

		createdProject, err := s.projects.Create(&project)

		if err != nil {
			http.Error(w, "Could not create project", 422)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(&createdProject)
	}
}
