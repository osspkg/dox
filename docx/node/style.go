package node

import (
	"go.osspkg.com/dox/xml"
)

type Style struct {
	ID   string
	Name string
}

func (v *Style) StyleId() string {
	return v.ID
}

func (v *Style) XML(x *xml.XML) {
	g := x.Group("w", "styles/style", "type", "paragraph", "styleId", v.ID)

	g.Tag("name", "", "val", v.Name)
	g.Tag("basedOn", "", "val", "Normal")
	g.Tag("qFormat", "")
	g.Tag("pPr/numPr/ilvl", "", "val", 0)
	g.Tag("pPr/numPr/numId", "", "val", 1)
	g.Tag("pPr/spacing", "", "before", 240, "after", 120)
	g.Tag("pPr/outlineLvl", "", "val", 0)
	g.Tag("rPr/b", "")
	g.Tag("rPr/bCs", "")
	g.Tag("rPr/sz", "", "val", 36)
	g.Tag("rPr/szCs", "", "val", 36)
}
