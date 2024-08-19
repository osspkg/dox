package docx

import (
	"fmt"

	"go.osspkg.com/dox/docx/node"
	"go.osspkg.com/dox/docx/types"
)

func (v *Docx) Paragraph(text string, opts ...types.Opt) {
	v.nodes = append(v.nodes, &node.Paragraph{
		Text: text,
		Style: &node.Style{
			ID:   "Paragraph",
			Name: "Paragraph",
		},
		Opts: opts,
	})
}

func (v *Docx) Head(text string, level uint8, opts ...types.Opt) {
	v.nodes = append(v.nodes, &node.Head{
		Text: text,
		Style: &node.Style{
			ID:   fmt.Sprintf("Heading %d", level),
			Name: "Paragraph",
		},
		Opts: opts,
	})
}
