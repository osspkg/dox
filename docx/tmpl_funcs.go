package docx

import (
	"text/template"
	"time"
)

var tmplFuncs = template.FuncMap{
	"date": func(t time.Time) string {
		return t.Format("2006-01-02T15:04:05Z")
	},
}
