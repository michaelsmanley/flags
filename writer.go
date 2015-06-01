package flags

import (
	"io"

	"github.com/clipperhouse/typewriter"
)

func init() {
	err := typewriter.Register(NewWriter())
	if err != nil {
		panic(err)
	}
}

func NewWriter() *Writer {
	return &Writer{}
}

func SliceName(typ typewriter.Type) string {
	return typ.Name + "Slice"
}

type Writer struct{}

func (fw *Writer) Name() string {
	return "flags"
}

func (fw *Writer) Imports(typ typewriter.Type) (result []typewriter.ImportSpec) {
	return []typewriter.ImportSpec{
		{Path: "encoding/json"},
		{Path: "errors"},
	}
}

func (fw *Writer) Write(w io.Writer, typ typewriter.Type) error {
	tag, found := typ.FindTag(fw)

	if !found {
		return nil
	}

	// start with the slice template
	tmpl, err := templates.ByTag(typ, tag)

	if err != nil {
		return err
	}

	m := model{
		Type:      typ,
		SliceName: SliceName(typ),
	}

	if err := tmpl.Execute(w, m); err != nil {
		return err
	}

	for _, v := range tag.Values {
		var tp typewriter.Type

		if len(v.TypeParameters) > 0 {
			tp = v.TypeParameters[0]
		}

		m := model{
			Type:          typ,
			SliceName:     SliceName(typ),
			TypeParameter: tp,
			TagValue:      v,
		}

		tmpl, err := templates.ByTagValue(typ, v)

		if err != nil {
			return err
		}

		if err := tmpl.Execute(w, m); err != nil {
			return err
		}
	}

	return nil
}
