package node

import (
	"io"

	"go.osspkg.com/dox/docx/types"
	"go.osspkg.com/dox/internal/do"
)

type Paragraph struct {
	Text  string
	Style *Style
	Opts  []types.Opt
}

func (v *Paragraph) Styles() types.Styles {
	return v.Style
}

func (v *Paragraph) XML(w io.Writer) error {
	return do.Multi(w,
		do.Write(`<w:p>`),
		do.Write(`<w:pPr>`),
		do.Write(`<w:pStyle w:val="%s" />`, v.Style.ID),
		do.Write(`<w:bidi w:val="0" />`),
		&do.Pack[types.Opt]{Data: v.Opts},
		do.Write(`<w:rPr></w:rPr>`),
		do.Write(`</w:pPr>`),
		do.Write(`<w:r>`),
		do.Write(`<w:rPr></w:rPr>`),
		do.Write(`<w:t>%s</w:t>`, v.Text),
		do.Write(`</w:r>`),
		do.Write(`</w:p>`),
	)
}
