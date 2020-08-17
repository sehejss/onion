package tomlloader

import (
	"bytes"
	"testing"

	. "github.com/sehejss/onion"
	. "github.com/smartystreets/goconvey/convey"
)

const testNormal = `str = "string_data"
bool = true
integer = 10
[nested]
  key1 = "string"
  key2 = 100`

const testWithDottedKeys = `name = "Orange"
physical.color = "orange"
physical.weight = 3`

func TestTomlLoader(t *testing.T) {
	Convey("Load a toml structure into a layer", t, func() {
		buf := bytes.NewBufferString(testNormal)
		Convey("Check if the file is loaded correctly ", func() {
			layer, err := NewStreamLayer(buf, "toml", nil)
			So(err, ShouldBeNil)

			o := New(layer)
			So(o.GetStringDefault("str", ""), ShouldEqual, "string_data")
			So(o.GetStringDefault("nested.key1", ""), ShouldEqual, "string")
			So(o.GetIntDefault("nested.key2", 0), ShouldEqual, 100)
			So(o.GetBoolDefault("bool", false), ShouldBeTrue)

		})

		bufInvalid := bytes.NewBufferString(`invalid toml file`)
		Convey("Check for the invalid file content", func() {
			_, err := NewStreamLayer(bufInvalid, "toml", nil)
			So(err, ShouldNotBeNil)
		})

		bufferWithDottedKeys := bytes.NewBufferString(testWithDottedKeys)
		Convey("Check if the file is loaded correctly, even with dots ", func() {
			layer, err := NewStreamLayer(bufferWithDottedKeys, "toml", nil)
			So(err, ShouldBeNil)

			o := New(layer)
			So(o.GetStringDefault("name", ""), ShouldEqual, "Orange")
			So(o.GetStringDefault("physical.color", ""), ShouldEqual, "orange")
			So(o.GetIntDefault("physical.weight", 0), ShouldEqual, 3)

			mergedLayers := NewMapLayer(o.LayersData()...).Load()
			physicalMap, isAMap := mergedLayers["physical"].(map[string]interface{})
			So(isAMap, ShouldBeTrue)
			So(physicalMap["color"], ShouldEqual, "orange")
			So(physicalMap["weight"], ShouldEqual, 3)
		})
	})
}
