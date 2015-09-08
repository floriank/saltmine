package api

import (
	. "github.com/floriank/saltmine/datastore"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"io"
	"net/http"
)

type SaltmineAPI struct {
	projects ProjectStorer
	tickets  TicketStorer
	version  string
}

func NewSaltmineAPI(db *gorm.DB, version string) *SaltmineAPI {
	return &SaltmineAPI{
		projects: NewProjectStore(db),
		tickets:  NewTicketStore(db),
		version:  version,
	}
}

func (s *SaltmineAPI) GetRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/version", s.VersionGet()).Methods("GET")

	p := r.PathPrefix("/projects").Subrouter()
	// tickets := r.PathPrefix("/tickets").Subrouter()

	// projects.HandleFunc("/", s.ProjectsList()).Methods("GET")
	p.HandleFunc("/", s.ProjectsList()).Methods("GET")
	p.HandleFunc("/", s.ProjectCreate()).Methods("POST")
	p.HandleFunc("/{id}", s.ProjectUpdate()).Methods("PUT", "PATCH")
	p.HandleFunc("/{id}", s.ProjectGet()).Methods("GET")
	// 	Methods("GET").Path("/").HandlerFunc(s.ProjectsList()).
	// 	Methods("GET").Path("/{id}/").HandlerFunc(s.ProjectGet()).
	// 	Methods("POST").Path("/").HandlerFunc(s.ProjectUpdate()).
	// 	Methods("DELETE").Path("/{id}").HandlerFunc(s.ProjectDelete()).
	// 	Methods("PATCH").Path("/{id}").HandlerFunc(s.ProjectUpdate())

	// tickets.
	// 	Methods("GET").Path("/").HandlerFunc(s.TicketsList()).
	// 	Methods("GET").Path("/{id}").HandlerFunc(s.TicketGet()).
	// 	Methods("POST").Path("/").HandlerFunc(s.TicketCreate()).
	// 	Methods("DELETE").Path("/{id}").HandlerFunc(s.TicketDelete()).
	// 	Methods("PATCH").Path("/{id}").HandlerFunc(s.TicketUpdate())

	return r
}

func (s *SaltmineAPI) VersionGet() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, s.version)
	}
}
