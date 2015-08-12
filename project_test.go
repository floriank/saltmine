package main

import (
	"bytes"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	db = Connect("./saltmine_test.db")
	db.DropTable(&Project{})
	db.CreateTable(&Project{})
}

func TestProjectCreate(t *testing.T) {
	Convey("That project create should work", t, func() {
		const body string = `
    {
      "identifier": "foo",
      "title": "FooBar Demo",
      "description": "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet eos alias ratione, iusto obcaecati nesciunt et minus dignissimos illum. Animi pariatur eum veritatis, nam voluptates incidunt, dignissimos eos quaerat ipsum."
    }
    `
		request, err := http.NewRequest("POST", "/projects", bytes.NewBufferString(body))

		if err != nil {
			t.Errorf("could not POST to endpoint: %v", err.Error())
		}

		response := httptest.NewRecorder()

		ProjectCreate(response, request)

		So(response.Body.String(), ShouldEqual, "ok")
	})
}
