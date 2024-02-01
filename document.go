package godocx

import (
	_ "embed"
	"os"
	"path/filepath"

	"github.com/gomutex/godocx/opc"
	"github.com/gomutex/godocx/oxml"
)

//go:embed templates/default.docx
var defaultDocx []byte

// NewDocument creates a new document from the default template.
func NewDocument() (*oxml.RootDoc, error) {
	return opc.Unpack(&defaultDocx)
}

// OpenDocument opens a document from the given file name.
func OpenDocument(fileName string) (*oxml.RootDoc, error) {
	docxContent, err := os.ReadFile(filepath.Clean(fileName))
	if err != nil {
		return nil, err
	}
	return opc.Unpack(&docxContent)
}
