package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBogus(t *testing.T) {

	Convey("True should be true", t, func() {
		So(true, ShouldBeTrue)
	})

}
