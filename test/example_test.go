package example

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEmpty(t *testing.T) {
	var m Bits

	Convey("New Bits should be empty.", t, func() {
		So(m, ShouldEqual, 0)
	})
}

func TestSetSingle(t *testing.T) {
	var m Bits

	m.Set(Bit01)
	Convey("Bits should be 1.", t, func() {
		So(m, ShouldEqual, 1)
	})

	Convey("Bits should have SystemAdmin flag set.", t, func() {
		So(m.IsSet(Bit01), ShouldEqual, true)
	})
}

func TestSetMultiple(t *testing.T) {
	var m Bits

	m.Set(Bit01 | Bit32)
	Convey("Bits should have Bit01 & Bit32 flag set.", t, func() {
		So(m.IsSet(Bit01|Bit32), ShouldEqual, true)
	})

	Convey("Bits should have Bit01 flag set.", t, func() {
		So(m.IsSet(Bit01), ShouldEqual, true)
	})

	Convey("Bits should not have Bit31 set.", t, func() {
		So(m.IsSet(Bit31), ShouldEqual, false)
	})
}

func TestUnsetSingle(t *testing.T) {
	var m Bits

	m.Set(Bit01)
	Convey("Bits should have Bit01 flag set.", t, func() {
		So(m.IsSet(Bit01), ShouldEqual, true)
	})

	m.Unset(Bit01)
	Convey("Bits should not have Bit01 flag set.", t, func() {
		So(m.IsSet(Bit01), ShouldEqual, false)
	})
}

func TestUnsetMultple(t *testing.T) {
	var m Bits

	m.Set(Bit01 | Bit32)
	Convey("Bits should have Bit01 & Bit32 flag set.", t, func() {
		So(m.IsSet(Bit01|Bit32), ShouldEqual, true)
	})

	m.Unset(Bit32)

	Convey("Bits should have Bit01 flag set.", t, func() {
		So(m.IsSet(Bit01), ShouldEqual, true)
	})

	Convey("Bits should not have Bit32 flag set.", t, func() {
		So(m.IsSet(Bit32), ShouldEqual, false)
	})
}

func TestStringer(t *testing.T) {
	Convey("Bits stringer names are generated.", t, func() {
		So(Bit01.String(), ShouldEqual, "Bit01")
	})
}

func TestJSONEnumsMarshal(t *testing.T) {
	j, _ := json.Marshal(Bit01)
	Convey("Bits jsonenum marshal works.", t, func() {
		So(j, ShouldResemble, []byte("\"Bit01\""))
	})
}

func TestJSONEnumsUnmarshal(t *testing.T) {
	var m Bits
	json.Unmarshal([]byte("\"Bit01\""), &m)
	Convey("Bits jsonenum unmarshal works.", t, func() {
		So(m, ShouldEqual, Bit01)
	})
}

func TestUnstring(t *testing.T) {
	f, _ := Unstring("Bit01")
	Convey("Bits unstring works.", t, func() {
		So(f, ShouldEqual, Bit01)
	})
}

func TestMutex(t *testing.T) {
	var m Bits
	m, err := m.Set(MutuallyExclusive)
	Convey("Bits mutex works.", t, func() {
		So(err, ShouldEqual, ErrMutex)
	})
}

func TestToggle(t *testing.T) {
	var m Bits
	m.Set(Bit01)
	m = m.Punch(Bit02 | Bit03)
	Convey("Bits toggle works.", t, func() {
		So(m, ShouldEqual, Bit03)
	})
}
