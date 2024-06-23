package shapes

import (
	"encoding/xml"
	"strconv"

	"github.com/gomutex/godocx/dml/dmlst"
)

// Tile
type Tile struct {
	Tx   *int64               `xml:"tx,attr,omitempty"`   // Horizontal Offset
	Ty   *int64               `xml:"ty,attr,omitempty"`   // Vertical Offset
	Sx   *int                 `xml:"sx,attr,omitempty"`   // Horizontal Ratio
	Sy   *int                 `xml:"sy,attr,omitempty"`   // Vertical Ratio
	Flip *dmlst.TileFlipMode  `xml:"flip,attr,omitempty"` // Tile Flipping
	Algn *dmlst.RectAlignment `xml:"algn,attr,omitempty"` // Alignment
}

// MarshalXML marshals the Tile struct into XML.
func (t Tile) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "a:tile"
	start.Attr = []xml.Attr{}

	if t.Tx != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "tx"}, Value: strconv.FormatInt(*t.Tx, 10)})
	}
	if t.Ty != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ty"}, Value: strconv.FormatInt(*t.Ty, 10)})
	}
	if t.Sx != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sx"}, Value: strconv.Itoa(*t.Sx)})
	}
	if t.Sy != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sy"}, Value: strconv.Itoa(*t.Sy)})
	}
	if t.Flip != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "flip"}, Value: string(*t.Flip)})
	}
	if t.Algn != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "algn"}, Value: string(*t.Algn)})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
