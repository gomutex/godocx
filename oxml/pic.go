package oxml

import (
	"fmt"
	"path/filepath"

	"github.com/gomutex/godocx/constants"
	"github.com/gomutex/godocx/dml"
	"github.com/gomutex/godocx/internal"
	"github.com/gomutex/godocx/oxml/elements"
	"github.com/gomutex/godocx/shared/units"
)

type PicMeta struct {
	Para   *elements.Paragraph
	Inline *dml.Inline
}

// AddPicture adds a new image to the document.
//
// Parameters:
//   - path: The path of the image file to be added.
//   - width: The width of the image in inches.
//   - height: The height of the image in inches.
//
// Returns:
//   - p: The created Paragraph instance containing the image.
//   - err: An error, if any occurred during the process.
func (rd *RootDoc) AddPicture(path string, width units.Inch, height units.Inch) (*PicMeta, error) {

	imgBytes, err := internal.FileToByte(path)
	if err != nil {
		return nil, err
	}

	imgExt := filepath.Ext(path)
	rd.ImageCount += 1
	fileName := fmt.Sprintf("image%d%s", rd.ImageCount, imgExt)
	fileIdxPath := fmt.Sprintf("%s%s", constants.MediaPath, fileName)

	rd.FileMap.Store(fileIdxPath, imgBytes)

	relName := fmt.Sprintf("media/%s", fileName)

	rID := rd.Document.addRelation(constants.SourceRelationshipImage, relName)

	p := elements.NewParagraph()

	bodyElem := DocumentChild{
		Para: p,
	}
	rd.Document.Body.Children = append(rd.Document.Body.Children, bodyElem)

	inline := p.AddDrawing(rID, width, height)

	return &PicMeta{
		Para:   p,
		Inline: inline,
	}, nil
}
