package flags

import "github.com/clipperhouse/typewriter"

var flagstest = &typewriter.Template{
	Name: "flags_test",
	Text: `
package example

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEmpty(t *testing.T) {
	var m {{.Type}}

	Convey("New {{.Type}} should be empty.", t, func() {
		So(m, ShouldEqual, 0)
	})
}

func TestSetSingle(t *testing.T) {
	var m {{.Type}}

	m.Set({{.Type}}(0x00000001))
	Convey("{{.Type}} should be 1.", t, func() {
		So(m, ShouldEqual, 1)
	})

	Convey("{{.Type}} should have SystemAdmin flag set.", t, func() {
		So(m.IsSet({{.Type}}(0x00000001)), ShouldEqual, true)
	})
}

func TestSetMultiple(t *testing.T) {
	var m {{.Type}}

	m.Set({{.Type}}(0x00000001) | {{.Type}}(0x80000000))
	Convey("{{.Type}} should have {{.Type}}(0x00000001) & {{.Type}}(0x80000000) flag set.", t, func() {
		So(m.IsSet({{.Type}}(0x00000001)|{{.Type}}(0x80000000)), ShouldEqual, true)
	})

	Convey("{{.Type}} should have {{.Type}}(0x00000001) flag set.", t, func() {
		So(m.IsSet({{.Type}}(0x00000001)), ShouldEqual, true)
	})

	Convey("{{.Type}} should not have Bit31 set.", t, func() {
		So(m.IsSet(Bit31), ShouldEqual, false)
	})
}

func TestUnsetSingle(t *testing.T) {
	var m {{.Type}}

	m.Set({{.Type}}(0x00000001))
	Convey("{{.Type}} should have {{.Type}}(0x00000001) flag set.", t, func() {
		So(m.IsSet({{.Type}}(0x00000001)), ShouldEqual, true)
	})

	m.Unset({{.Type}}(0x00000001))
	Convey("{{.Type}} should not have {{.Type}}(0x00000001) flag set.", t, func() {
		So(m.IsSet({{.Type}}(0x00000001)), ShouldEqual, false)
	})
}

func TestUnsetMultple(t *testing.T) {
	var m {{.Type}}

	m.Set({{.Type}}(0x00000001) | {{.Type}}(0x80000000))
	Convey("{{.Type}} should have {{.Type}}(0x00000001) & {{.Type}}(0x80000000) flag set.", t, func() {
		So(m.IsSet({{.Type}}(0x00000001)|{{.Type}}(0x80000000)), ShouldEqual, true)
	})

	m.Unset({{.Type}}(0x80000000))

	Convey("{{.Type}} should have {{.Type}}(0x00000001) flag set.", t, func() {
		So(m.IsSet({{.Type}}(0x00000001)), ShouldEqual, true)
	})

	Convey("{{.Type}} should not have {{.Type}}(0x80000000) flag set.", t, func() {
		So(m.IsSet({{.Type}}(0x80000000)), ShouldEqual, false)
	})
}

func TestStringer(t *testing.T) {
	Convey("{{.Type}} stringer names are generated.", t, func() {
		So({{.Type}}(0x00000001).String(), ShouldEqual, "{{.Type}}(0x00000001)")
	})
}

func TestJSONEnumsMarshal(t *testing.T) {
	j, _ := json.Marshal({{.Type}}(0x00000001))
	Convey("{{.Type}} jsonenum marshal works.", t, func() {
		So(j, ShouldResemble, []byte("\"{{.Type}}(0x00000001)\""))
	})
}

func TestJSONEnumsUnmarshal(t *testing.T) {
	var m {{.Type}}
	json.Unmarshal([]byte("\"{{.Type}}(0x00000001)\""), &m)
	Convey("{{.Type}} jsonenum unmarshal works.", t, func() {
		So(m, ShouldEqual, {{.Type}}(0x00000001))
	})
}

func TestUnstring(t *testing.T) {
	f, _ := Unstring("{{.Type}}(0x00000001)")
	Convey("{{.Type}} unstring works.", t, func() {
		So(f, ShouldEqual, {{.Type}}(0x00000001))
	})
}

func TestMutex(t *testing.T) {
	var m {{.Type}}
	m, err := m.Set(MutuallyExclusive)
	Convey("{{.Type}} mutex works.", t, func() {
		So(err, ShouldEqual, ErrMutex)
	})
}
`,
}
