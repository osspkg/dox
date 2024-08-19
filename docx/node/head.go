package node

import (
	"io"

	"go.osspkg.com/dox/docx/types"
	"go.osspkg.com/dox/internal/do"
)

type Head struct {
	Text  string
	Level uint8
	Style *Style
	Opts  []types.Opt
}

func (v *Head) Styles() types.Styles {
	return v.Style
}

func (v *Head) XML(w io.Writer) error {
	return do.Multi(w,
		do.Write(`<w:p>`),
		do.Write(`<w:pPr>`),
		do.Write(`<w:pStyle w:val="%s" />`, v.Style.ID),
		do.Write(`<w:bidi w:val="0" />`),
		&do.Pack[types.Opt]{Data: v.Opts},
		do.Write(`<w:pPr>`),
		do.Write(`<w:numPr><w:ilvl w:val="%d" /><w:numId w:val="%d" /></w:numPr>`, v.Level, v.Level+1),
		do.Write(`<w:spacing w:before="%d" w:after="120" />`, 240-v.Level*20),
		do.Write(`<w:outlineLvl w:val="%d" />`, v.Level),
		do.Write(`</w:pPr>`),
		do.Write(`<w:r>`),
		do.Write(`<w:rPr></w:rPr>`),
		do.Write(`<w:t>%s</w:t>`, v.Text),
		do.Write(`</w:r>`),
		do.Write(`</w:p>`),
	)
}
