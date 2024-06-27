package shapes

import (
	"encoding/xml"

	"github.com/gomutex/godocx/dml/dmlct"
)

type Stretch struct {
	FillRect *dmlct.RelativeRect `xml:"fillRect,omitempty"`
}

func (s Stretch) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "a:stretch"

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if s.FillRect != nil {
		if err := s.FillRect.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "a:fillRect"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
