package api

import (
	"encoding/json"
	. "github.com/floriank/saltmine/datastore"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProjectGet(t *testing.T) {
	router := api.GetRouter()
	Convey("that getting a project through the api is possible", t, func() {
		request, err := http.NewRequest("GET", "/projects/1", nil)
		response := httptest.NewRecorder()

		if err != nil {
			t.Fatal("Could not create request!")
		}

		router.ServeHTTP(response, request)

		So(response.Code, ShouldEqual, http.StatusOK)

		project := &Project{}

		json.NewDecoder(response.Body).Decode(project)

		So(project.ID, ShouldEqual, 12)
	})
}

func TestProjectList(t *testing.T) {
	router := api.GetRouter()
	Convey("that getting a list of projects through the api is possible", t, func() {
		request, err := http.NewRequest("GET", "/projects/", nil)
		response := httptest.NewRecorder()

		if err != nil {
			t.Fatal("Could not create request!")
		}

		router.ServeHTTP(response, request)

		So(response.Code, ShouldEqual, http.StatusOK)

		var projects []*Project

		json.NewDecoder(response.Body).Decode(&projects)

		So(projects[0].ID, ShouldEqual, 1)
		So(projects[1].ID, ShouldEqual, 2)
		So(projects[0].Identifier, ShouldEqual, "foo")
		So(projects[1].Identifier, ShouldEqual, "bar")
	})
}
