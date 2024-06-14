package godocx

import (
	_ "embed"
	"os"
	"path/filepath"

	"github.com/gomutex/godocx/doc"
	"github.com/gomutex/godocx/packager"
)

//go:embed templates/default.docx
var defaultDocx []byte

// NewDocument creates a new document from the default template.
func NewDocument() (*doc.RootDoc, error) {
	return packager.Unpack(&defaultDocx)
}

// OpenDocument opens a document from the given file name.
func OpenDocument(fileName string) (*doc.RootDoc, error) {
	docxContent, err := os.ReadFile(filepath.Clean(fileName))
	if err != nil {
		return nil, err
	}
	return packager.Unpack(&docxContent)
}
