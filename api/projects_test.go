package api

import (
	"bytes"
	"encoding/json"
	. "github.com/floriank/saltmine/datastore"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mockApi *SaltmineAPI
	api     *SaltmineAPI
	db      gorm.DB
)

func init() {
	mockApi = &SaltmineAPI{
		projects: &MockProjectsStore{},
		tickets:  &MockTicketsStore{},
		version:  "0.0.0",
	}
	db, _ = gorm.Open("sqlite3", "./saltmine_test.db")
	api = NewSaltmineAPI(&db, "0.0.0")
}

func reset(db *gorm.DB) {
	db.DropTableIfExists(&Project{})
	db.CreateTable(&Project{})
	db.DropTableIfExists(&Ticket{})
	db.CreateTable(&Ticket{})
}

func TestProjectAPIGet(t *testing.T) {
	reset(&db)
	router := api.GetRouter()

	project := Project{
		Identifier:  "foo",
		Title:       "test",
		Description: "lorem",
	}

	db.NewRecord(project)
	db.Create(&project)

	Convey("that getting a project from the api should work", t, func() {
		Convey("requesting a non existing project", func() {
			request, _ := http.NewRequest("GET", "/projects/123", nil)
			response := httptest.NewRecorder()

			router.ServeHTTP(response, request)

			So(response.Code, ShouldEqual, http.StatusNotFound)
		})

		Convey("requesting an existing project", func() {
			request, _ := http.NewRequest("GET", "/projects/1", nil)
			response := httptest.NewRecorder()

			router.ServeHTTP(response, request)

			So(response.Code, ShouldEqual, http.StatusOK)

			foundProject := &Project{}
			json.NewDecoder(response.Body).Decode(foundProject)

			So(foundProject.ID, ShouldEqual, 1)
			So(foundProject.Identifier, ShouldEqual, "foo")
		})
	})
}

func TestProjectAPICreate(t *testing.T) {
	reset(&db)
	router := api.GetRouter()

	Convey("that creating a project via the API shoud work", t, func() {
		Convey("using correct JSON", func() {
			content := `
				{
					"identifier": "foo",
					"title": "Foo & bar",
					"description": "bar foo baz"
				}
			`
			request, _ := http.NewRequest("POST", "/projects/", bytes.NewBufferString(content))
			response := httptest.NewRecorder()

			router.ServeHTTP(response, request)

			So(response.Code, ShouldEqual, http.StatusCreated)

			project := Project{}
			json.NewDecoder(response.Body).Decode(&project)

			So(project.ID, ShouldBeGreaterThan, 0)
			So(project.Identifier, ShouldEqual, "foo")
			So(project.Description, ShouldEqual, "bar foo baz")
			So(project.Title, ShouldEqual, "Foo & bar")
		})

		Convey("using broken JSON", func() {
			content := `
				{
					"identifier": "foo",
					"title": "Foo & bar",
					"de
				}
			`
			request, _ := http.NewRequest("POST", "/projects/", bytes.NewBufferString(content))
			response := httptest.NewRecorder()

			router.ServeHTTP(response, request)

			So(response.Code, ShouldEqual, 422)
			projects, _ := api.projects.List()
			So(len(projects), ShouldEqual, 0)
		})

		Reset(func() {
			reset(&db)
		})
	})
}

func TestProjectAPIStatusResponse(t *testing.T) {
	router := mockApi.GetRouter()

	Convey("when talking to the API", t, func() {

		Convey("using GET for a single project", func() {
			request, err := http.NewRequest("GET", "/projects/1", nil)
			response := httptest.NewRecorder()

			if err != nil {
				t.Fatal("Could not create request!")
			}

			router.ServeHTTP(response, request)

			So(response.Code, ShouldEqual, http.StatusOK)
		})

		Convey("using GET for getting all projects", func() {
			request, err := http.NewRequest("GET", "/projects/", nil)
			response := httptest.NewRecorder()

			if err != nil {
				t.Fatal("Could not create request!")
			}

			router.ServeHTTP(response, request)

			So(response.Code, ShouldEqual, http.StatusOK)
		})

		Convey("using POST to create projects", func() {
			content := `{
				"foo": "bar"
			}`
			request, err := http.NewRequest("POST", "/projects/", bytes.NewBufferString(content))
			response := httptest.NewRecorder()

			if err != nil {
				t.Fatal("Could not create request!")
			}

			router.ServeHTTP(response, request)
			So(response.Code, ShouldEqual, http.StatusCreated)
		})

		Convey("using PATCH or POST to update projects", func() {
			var content string = "bar"

			Convey("w/ PATCH", func() {
				request, err := http.NewRequest("PATCH", "/projects/12", bytes.NewBufferString(content))
				response := httptest.NewRecorder()

				if err != nil {
					t.Fatal("Could not create request!")
				}

				router.ServeHTTP(response, request)
				So(response.Code, ShouldEqual, http.StatusOK)
			})

			Convey("w/ PUT", func() {
				request, err := http.NewRequest("PUT", "/projects/12", bytes.NewBufferString(content))
				response := httptest.NewRecorder()

				if err != nil {
					t.Fatal("Could not create request!")
				}

				router.ServeHTTP(response, request)
				So(response.Code, ShouldEqual, http.StatusOK)
			})
		})
	})
}
