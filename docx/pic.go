package docx

import (
	"github.com/gomutex/godocx/common/units"
	"github.com/gomutex/godocx/dml"
)

type PicMeta struct {
	Para   *Paragraph
	Inline *dml.Inline
}

// AddPicture adds a new image to the document.
//
// Example usage:
//
//	// Add a picture to the document
//	_, err = document.AddPicture("gopher.png", units.Inch(2.9), units.Inch(2.9))
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// Parameters:
//   - path: The path of the image file to be added.
//   - width: The width of the image in inches.
//   - height: The height of the image in inches.
//
// Returns:
//   - *PicMeta: Metadata about the added picture, including the Paragraph instance and Inline element.
//   - error: An error, if any occurred during the process.
func (rd *RootDoc) AddPicture(path string, width units.Inch, height units.Inch) (*PicMeta, error) {

	p := newParagraph(rd)

	bodyElem := DocumentChild{
		Para: p,
	}
	rd.Document.Body.Children = append(rd.Document.Body.Children, bodyElem)

	return p.AddPicture(path, width, height)
}
