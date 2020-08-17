package yamlloader

import (
	"bytes"
	"testing"

	. "github.com/sehejss/onion"
	. "github.com/smartystreets/goconvey/convey"
)

func TestYamlLoader(t *testing.T) {
	Convey("Load a yaml structure into a json", t, func() {

		buf := bytes.NewBufferString(`---
  str: "string_data"
  bool: true
  integer: 10
  nested:
    key1: "string"
    key2: 100
`)
		bufInvalid := bytes.NewBufferString(`---
str: - inv
  lid
 s
ALALA`)

		Convey("Check if the file is loaded correctly ", func() {
			fl, err := NewStreamLayer(buf, "yml", nil)
			So(err, ShouldBeNil)
			o := New(fl)
			So(o.GetStringDefault("str", ""), ShouldEqual, "string_data")
			So(o.GetStringDefault("nested.key1", ""), ShouldEqual, "string")
			So(o.GetIntDefault("nested.key2", 0), ShouldEqual, 100)
			So(o.GetBoolDefault("bool", false), ShouldBeTrue)
		})

		Convey("Check for the invalid file content", func() {
			_, err := NewStreamLayer(bufInvalid, "yaml", nil)
			So(err, ShouldNotBeNil)
		})
	})
}
