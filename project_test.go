package main

import (
	"bytes"
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	db = Connect("./saltmine_test.db")
	db.DropTableIfExists(&Project{})
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

		So(err, ShouldBeNil)

		response := httptest.NewRecorder()

		ProjectCreate(response, request)

		So(response.Body.String(), ShouldEqual, "ok")
		So(response.Code, ShouldEqual, http.StatusCreated)
	})
}

func TestProjectIndex(t *testing.T) {
	db.Exec("DELETE FROM projects;")

	project := Project{
		Identifier:  "foo",
		Title:       "Bar",
		Description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Repudiandae minus fugiat quisquam modi pariatur deserunt unde nam, sed, beatae nemo tempora odit maxime doloribus possimus ratione. Aperiam tempore minus nulla!",
	}

	db.NewRecord(project)
	db.Create(&project)

	Convey("That the project index returns a list of projects", t, func() {
		request, err := http.NewRequest("GET", "/projects", nil)

		So(err, ShouldBeNil)

		response := httptest.NewRecorder()

		ProjectIndex(response, request)
		So(response.Code, ShouldEqual, http.StatusOK)

		var data []Project
		json.NewDecoder(response.Body).Decode(&data)

		So(err, ShouldBeNil)
		So(len(data), ShouldEqual, 1)
		So(data[0].Identifier, ShouldEqual, "foo")
	})
}
