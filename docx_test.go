package docx

import (
	"strings"
	"testing"
)

const testFile = "./TestDocument.docx"
const testFileResult = "./TestDocumentResult.docx"

func loadFile(file string) *Docx {
	r, err := ReadDocxFile(file)
	if err != nil {
		panic(err)
	}

	return r.Editable()
}



func TestReplaceHeader(t *testing.T) {
	d := loadFile(testFile)
	d.ReplaceHeader("This is a header.", "newHeader")
	d.WriteToFile(testFileResult)

	d = loadFile(testFileResult)

	headers := d.headers
	found := false
	for _, v := range headers {
		if strings.Contains(v, "This is a header.") {
			t.Error("Missing 'This is a header.', got ", d.content)
		}

		if strings.Contains(v, "newHeader") {
			found = true
		}
	}
	if !found {
		t.Error("Expected 'newHeader', got ", d.headers)
	}
}

func TestReplaceFooter(t *testing.T) {
	d := loadFile(testFile)
	d.ReplaceFooter("This is a footer.", "newFooter")
	d.WriteToFile(testFileResult)

	d = loadFile(testFileResult)

	footers := d.footers
	found := false
	for _, v := range footers {
		if strings.Contains(v, "This is a footer.") {
			t.Error("Missing 'This is a footer.', got ", d.content)
		}

		if strings.Contains(v, "newFooter") {
			found = true
		}
	}
	if !found {
		t.Error("Expected 'newFooter', got ", d.headers)
	}
}

