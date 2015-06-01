package flags

import (
	"io"

	"github.com/clipperhouse/typewriter"
)

type model struct {
	Type          typewriter.Type
	SliceName     string
	TypeParameter typewriter.Type
	typewriter.TagValue
}

var templates = typewriter.TemplateSlice{
	flags,
}

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

	return nil
}
