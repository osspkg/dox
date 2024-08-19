package docx_test

import (
	"testing"

	"go.osspkg.com/dox/docx"
	"go.osspkg.com/dox/docx/style"
	"golang.org/x/text/language"
)

func TestUnit_New(t *testing.T) {
	doc := docx.New(language.Russian)
	doc.Paragraph("aaaa", style.AlignRight())
	if err := doc.Save("/tmp/test.docx"); err != nil {
		t.Error(err)
	}
}
