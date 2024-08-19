package style

import (
	"fmt"
	"io"

	"go.osspkg.com/dox/docx/types"
)

type align struct {
	val string
}

func (*align) Tag() string {
	return "align"
}

func (v *align) Exec(w io.Writer) (err error) {
	_, err = fmt.Fprintf(w, `<w:jc w:val="%s" />`, v.val)
	return
}

func AlignLeft() types.Opt {
	return &align{val: "left"}
}

func AlignCenter() types.Opt {
	return &align{val: "center"}
}

func AlignRight() types.Opt {
	return &align{val: "right"}
}

func AlignJustify() types.Opt {
	return &align{val: "both"}
}
