package docx

import (
	"fmt"
	"io"
	"time"

	"go.osspkg.com/dox/docx/types"
	"go.osspkg.com/dox/internal/archive"
	"golang.org/x/text/language"
)

type Docx struct {
	option tmplOption
	nodes  []types.Node
}

func New(lang language.Tag) *Docx {
	return &Docx{
		option: tmplOption{
			DateTime: time.Now(),
			Lang:     lang.String(),
		},
		nodes: make([]types.Node, 0, 1024),
	}
}

func (v *Docx) Save(filepath string) error {
	z := archive.New()
	for _, file := range allStaticFiles {
		err := z.File(file.Filepath, func(w io.Writer) error {
			return file.Build(w, v.option)
		})
		if err != nil {
			return err
		}
	}
	styles := make(map[string]types.Styles, len(v.nodes))
	err := z.File(tmpDynamicWordDocument.Filepath, func(w io.Writer) (err error) {
		if _, err = fmt.Fprintf(w, tmpDynamicWordDocument.AfterBody); err != nil {
			return err
		}
		for _, node := range v.nodes {
			if sval := node.Styles(); sval != nil {
				styles[sval.StyleId()] = sval
			}
			if err = node.XML(w); err != nil {
				return err
			}
		}
		if _, err = fmt.Fprintf(w, tmpDynamicWordDocument.BeforeBody); err != nil {
			return err
		}
		return
	})
	if err != nil {
		return err
	}
	err = z.File(tmpDynamicWordStyles.Filepath, func(w io.Writer) (err error) {
		if _, err = fmt.Fprintf(w, tmpDynamicWordStyles.AfterBody); err != nil {
			return err
		}
		for _, node := range styles {
			if err = node.XML(w); err != nil {
				return err
			}
		}
		if _, err = fmt.Fprintf(w, tmpDynamicWordStyles.BeforeBody); err != nil {
			return err
		}
		return
	})
	if err != nil {
		return err
	}
	return z.Save(filepath)
}
