package docx

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gomutex/godocx/common/constants"
	"github.com/gomutex/godocx/common/units"
	"github.com/gomutex/godocx/dml"
	"github.com/gomutex/godocx/internal"
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

	imgBytes, err := internal.FileToByte(path)
	if err != nil {
		return nil, err
	}

	imgExt := filepath.Ext(path)
	rd.ImageCount += 1
	fileName := fmt.Sprintf("image%d%s", rd.ImageCount, imgExt)
	fileIdxPath := fmt.Sprintf("%s%s", constants.MediaPath, fileName)

	imgExtStripDot := strings.TrimPrefix(imgExt, ".")
	imgMIME, err := MIMEFromExt(imgExtStripDot)
	if err != nil {
		return nil, err
	}

	err = rd.ContentType.AddExtension(imgExtStripDot, imgMIME)
	if err != nil {
		return nil, err
	}

	overridePart := fmt.Sprintf("/%s%s", constants.MediaPath, fileName)
	err = rd.ContentType.AddOverride(overridePart, imgMIME)
	if err != nil {
		return nil, err
	}

	rd.FileMap.Store(fileIdxPath, imgBytes)

	relName := fmt.Sprintf("media/%s", fileName)

	rID := rd.Document.addRelation(constants.SourceRelationshipImage, relName)

	p := newParagraph(rd)

	bodyElem := DocumentChild{
		Para: p,
	}
	rd.Document.Body.Children = append(rd.Document.Body.Children, bodyElem)

	inline := p.addDrawing(rID, rd.ImageCount, width, height)

	return &PicMeta{
		Para:   p,
		Inline: inline,
	}, nil
}
