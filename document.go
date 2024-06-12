package godocx

import (
	_ "embed"
	"os"
	"path/filepath"

	"github.com/gomutex/godocx/ooxml"
	"github.com/gomutex/godocx/opc"
)

//go:embed templates/default.docx
var defaultDocx []byte

// NewDocument creates a new document from the default template.
func NewDocument() (*ooxml.RootDoc, error) {
	return opc.Unpack(&defaultDocx)
}

// OpenDocument opens a document from the given file name.
func OpenDocument(fileName string) (*ooxml.RootDoc, error) {
	docxContent, err := os.ReadFile(filepath.Clean(fileName))
	if err != nil {
		return nil, err
	}
	return opc.Unpack(&docxContent)
}
