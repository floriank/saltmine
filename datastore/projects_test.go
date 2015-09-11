package datastore

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewProjectStore(t *testing.T) {
	reset(&db)
	Convey("creating a new store", t, func() {
		store := NewProjectStore(&db)
		So(store, ShouldNotBeNil)
	})
}

func TestProjectStoreGet(t *testing.T) {
	reset(&db)
	store := NewProjectStore(&db)
	project := Project{
		Identifier:  "foo",
		Title:       "Bar",
		Description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Repudiandae minus fugiat quisquam modi pariatur deserunt unde nam, sed, beatae <ne></ne>mo tempora odit maxime doloribus possimus ratione. Aperiam tempore minus nulla!",
	}

	db.NewRecord(project)
	db.Create(&project)

	Convey("getting a single project by id", t, func() {
		Convey("with no project present should give an error", func() {
			_, err := store.Get(42)
			So(err, ShouldNotBeNil)
		})

		Convey("with the project present", func() {
			foundProject, err := store.Get(1)
			So(foundProject, ShouldNotBeNil)
			So(err, ShouldBeNil)
		})
	})
}

func TestProjectStoreList(t *testing.T) {
	reset(&db)
	store := NewProjectStore(&db)
	project := Project{
		Identifier:  "foo",
		Title:       "Bar",
		Description: "Lorem",
	}

	project2 := Project{
		Identifier:  "bar",
		Title:       "Baz",
		Description: "ipsum",
	}

	db.NewRecord(&project)
	db.Create(&project)

	db.NewRecord(&project2)
	db.Create(&project2)

	Convey("Using List will retrieve all the projects from the database", t, func() {
		projects, err := store.List()
		So(err, ShouldBeNil)
		So(len(projects), ShouldEqual, 2)
		So(projects[0].Identifier, ShouldEqual, "foo")
		So(projects[1].Identifier, ShouldEqual, "bar")
	})
}

func TestProjectCreate(t *testing.T) {
	reset(&db)
	store := NewProjectStore(&db)
	project := Project{
		Identifier:  "foo",
		Title:       "Bar",
		Description: "Lorem",
	}

	Convey("that create will create a new project in the db", t, func() {
		store.Create(&project)

		var projects []*Project
		db.Find(&projects)
		So(len(projects), ShouldEqual, 1)
		So(projects[0].Title, ShouldEqual, "Bar")
	})
}

func TestProjectStoreDelete(t *testing.T) {
	reset(&db)
	store := NewProjectStore(&db)
	project := Project{
		Identifier:  "foo",
		Title:       "Bar",
		Description: "Lorem",
	}

	db.NewRecord(&project)
	db.Create(&project)

	Convey("that delete will remove a project from the db", t, func() {
		store.Delete(&project)

		var projects []*Project
		db.Find(&projects)
		So(len(projects), ShouldEqual, 0)
	})
}

func TestProjectStoreUpdate(t *testing.T) {
	reset(&db)
	store := NewProjectStore(&db)
	project := Project{
		Identifier:  "foo",
		Title:       "Bar",
		Description: "Lorem",
	}

	db.NewRecord(project)
	db.Create(&project)

	Convey("that delete will remove a project from the db", t, func() {
		project.Identifier = "newIdentifier"
		store.Update(&project)

		var projects []*Project
		db.Find(&projects)
		So(len(projects), ShouldEqual, 1)

		So(projects[0].Identifier, ShouldEqual, "newIdentifier")
	})
}
