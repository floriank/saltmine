package datastore

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var project Project

func init() {
	project = Project{
		Identifier:  "foo",
		Title:       "Bar",
		Description: "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Repudiandae minus fugiat quisquam modi pariatur deserunt unde nam, sed, beatae <ne></ne>mo tempora odit maxime doloribus possimus ratione. Aperiam tempore minus nulla!",
	}

	db.NewRecord(project)
	db.Create(&project)
}

func TestNewTicketStore(t *testing.T) {
	reset(&db)
	Convey("creating a new store", t, func() {
		store := NewTicketStore(&db)
		So(store, ShouldNotBeNil)
	})
}

func TestTicketStoreGet(t *testing.T) {
	reset(&db)
	store := NewTicketStore(&db)
	ticket := &Ticket{
		Subject:     "Foobar",
		Description: "formatted text",
		Project:     project,
	}

	db.NewRecord(ticket)
	db.Create(&ticket)

	Convey("getting a single ticket by id", t, func() {
		Convey("with no ticket present should give an error", func() {
			_, err := store.Get(42)
			So(err, ShouldNotBeNil)
		})

		Convey("with the ticket present", func() {
			ticket, err := store.Get(1)
			So(ticket, ShouldNotBeNil)
			So(err, ShouldBeNil)
			So(ticket.Subject, ShouldEqual, "Foobar")
		})
	})
}

func TestTicketStoreList(t *testing.T) {
	reset(&db)
	store := NewTicketStore(&db)
	ticket := &Ticket{
		Subject:     "Foobar",
		Description: "formatted text",
		Project:     project,
	}
	ticket1 := &Ticket{
		Subject:     "Barfoo",
		Description: "formatted text, for the second time",
		Project:     project,
	}

	db.NewRecord(ticket)
	db.Create(&ticket)
	db.NewRecord(ticket1)
	db.Create(&ticket1)

	Convey("getting tickets with the list method", t, func() {
		tickets, _ := store.List()

		So(len(tickets), ShouldEqual, 2)
		So(tickets[1].Subject, ShouldEqual, "Barfoo")
		So(tickets[0].Description, ShouldEqual, "formatted text")
	})
}

func TestTicketDelete(t *testing.T) {
	reset(&db)
	store := NewTicketStore(&db)
	ticket := &Ticket{
		Subject:     "Foobar",
		Description: "formatted text",
		Project:     project,
	}

	db.NewRecord(ticket)
	db.Create(&ticket)

	Convey("deleting a ticket", t, func() {
		ticket, _ := store.Delete(ticket)
		var tickets []*Ticket
		db.Find(&tickets)

		So(len(tickets), ShouldEqual, 0)
		So(ticket.Subject, ShouldEqual, "Foobar")
	})
}

func TestTicketStoreUpdate(t *testing.T) {
	reset(&db)
	store := NewTicketStore(&db)
	ticket := Ticket{
		Subject:     "Bar",
		Description: "Lorem",
		Project:     project,
	}

	db.NewRecord(ticket)
	db.Create(&ticket)

	Convey("that update will update a single ticket", t, func() {
		ticket.Subject = "Foo"
		ticket.Description = "ipsum"
		store.Update(&ticket)

		var tickets []*Ticket
		db.Find(&tickets)
		So(len(tickets), ShouldEqual, 1)

		So(tickets[0].Subject, ShouldEqual, "Foo")
		So(tickets[0].Description, ShouldEqual, "ipsum")
	})
}
