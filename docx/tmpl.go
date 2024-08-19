package docx

import (
	"fmt"
	"io"
	"text/template"
)

type staticFile struct {
	Filepath string
	Body     string
}

func (v staticFile) Build(w io.Writer, data interface{}) error {
	t, err := template.New(v.Filepath).Funcs(tmplFuncs).Parse(v.Body)
	if err != nil {
		return fmt.Errorf("build %s: %w", v.Filepath, err)
	}
	return t.Execute(w, data)
}

var allStaticFiles = []staticFile{
	tmplContentTypes,
	tmplDocPropsApp,
	tmplDocPropsCore,
	tmplRels,
	tmplWordDocumentRels,
	tmplWordFontTable,
}

type dynamicFile struct {
	Filepath              string
	AfterBody, BeforeBody string
}
