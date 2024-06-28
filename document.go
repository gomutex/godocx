package godocx

import (
	_ "embed"
	"os"
	"path/filepath"

	"github.com/gomutex/godocx/docx"
	"github.com/gomutex/godocx/packager"
)

//go:embed templates/default.docx
var defaultDocx []byte

// NewDocument creates a new document from the default template.
func NewDocument() (*docx.RootDoc, error) {
	return packager.Unpack(&defaultDocx)
}

// OpenDocument opens a document from the given file name.
func OpenDocument(fileName string) (*docx.RootDoc, error) {
	docxContent, err := os.ReadFile(filepath.Clean(fileName))
	if err != nil {
		return nil, err
	}
	return packager.Unpack(&docxContent)
}
