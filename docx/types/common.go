package types

import (
	"io"

	"go.osspkg.com/dox/xml"
)

type Opt interface {
	Tag() string
	Exec(w io.Writer) (err error)
}

type Node interface {
	XML(x *xml.XML)
	Styles() Styles
}

type Styles interface {
	StyleId() string
	XML(x *xml.XML)
}
