package flags

import "github.com/clipperhouse/typewriter"

var flags = &typewriter.Template{
	Name: "flags",
	Text: `
	// {{.SliceName}} is a slice of {{.Type}}.
	type {{.SliceName}} []{{.Type}}

	// ErrMutex says that a result had more than one mutually exclusive bit set.
	var ErrMutex = errors.New("attempt to set mutually exclusive {{.Type}}")

	// Punch turns a particular bit or set of {{.Type}} on.
	// If a bit that is part of a mutually exclusive set is
	// turned on, all of the other bits in that set will be
	// turned off. If multiple bits in the mutually exclusive
	// set are turned on, only the highest bit will remain on.
	func (f *{{.Type}}) Punch(v {{.Type}}) {{.Type}} {
		var x {{.Type}}

		s := v.Unpack()
		for _, b := range s {
			x = *f | b
			y := x & MutuallyExclusive
			if !y.IsEmpty() && !(y & (y - 1)).IsEmpty() {
				z := x &^ MutuallyExclusive
				x = z | b
			}
		}

		*f = x
		return x
	}


	// Set turns a particular bit or set of {{.Type}} on, and
	// returns ErrMutex if an attempt is made to turn mutually exclusive
	// bits on at the same time.
	func (f *{{.Type}}) Set(v {{.Type}}) ({{.Type}}, error) {
		x := *f | v

		y := x & MutuallyExclusive
		if !y.IsEmpty() && !(y & (y - 1)).IsEmpty() {
			return *f, ErrMutex
		}

		*f = x
		return x, nil
	}

	// Unset turns a particular bit or set of {{.Type}} off.
	func (f *{{.Type}}) Unset(v {{.Type}}) ({{.Type}}, error) {
		x := *f &^ v

		*f = x
		return x, nil
	}

	// IsEmpty is true if no {{.Type}} are on.
	func (f {{.Type}}) IsEmpty() bool {
		return f == 0
	}

	// IsSet determines if a Flags has a particular bit or set of {{.Type}} on.
	func (f {{.Type}}) IsSet(v {{.Type}}) bool {
		return f&v != 0
	}

	// Unpack creates a {{.SliceName}} from a single {{.Type}}, with each item in the list
	// having only a single bit on.
	func (f {{.Type}}) Unpack() {{.SliceName}} {
		var v {{.Type}}
		fl := {{.SliceName}}{}

		for v = 1; v <= 1<<31; v <<= 1 {
			if f.IsSet(v) {
				fl = append(fl, v)
			}
			if v == 1<<31 {
				break
			}
		}

		return fl
	}

	// Pack packs a {{.SliceName}} into a single {{.Type}} value.
	func (fl {{.SliceName}}) Pack() ({{.Type}}, error) {
		var f {{.Type}}
		var err error

		for _, v := range fl {
			f, err = f.Set(v)
			if err != nil {
				return 0, err
			}
		}

		return f, nil
	}

	// Unstring turns the string representation of a {{.Type}} into a {{.Type}} value.
	// This is meant to be used with constants that have had String() and
	// JSON Marshal/Unmarshal routines generated with stringer and jsonenums.
	func Unstring(s string) ({{.Type}}, error) {
		var f {{.Type}}
		err := json.Unmarshal([]byte("\""+s+"\""), &f)
		if err != nil {
			return 0, err
		}
		return f, nil
	}

	// PackStrings takes a slice of string representations of {{.Type}} and packs
	// them into a single Flags value.
	func PackStrings(fs []string) ({{.Type}}, error) {
		var f, v {{.Type}}
		var err error

		for _, s := range fs {
			v, err = Unstring(s)
			if err != nil {
				return 0, err
			}

			f, err = f.Set(v)
			if err != nil {
				return 0, err
			}
		}

		return f, nil
	}
	`,
}
