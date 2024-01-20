package elements

import (
	"encoding/xml"
)

type Drawing struct {
	Data []*DrawingData
}

type DrawingData struct {
	Pic *Pic
	// TextBox *TextBox
}

func (d *Drawing) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {

	start.Name.Local = "w:drawing"

	err = e.EncodeToken(start)
	if err != nil {
		return err
	}

	// for _, data := range d.Data {

	// }

	return e.EncodeToken(xml.EndElement{Name: start.Name})

}
