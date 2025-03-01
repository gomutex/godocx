package dmlct

import (
	"encoding/xml"
	"strconv"
)

// Non-Visual Drawing Properties
type CNvPr struct {
	ID   uint   `xml:"id,attr,omitempty"`
	Name string `xml:"name,attr,omitempty"`

	//Alternative Text for Object - Default value is "".
	Description string `xml:"descr,attr,omitempty"`

	// Hidden - Default value is "false".
	Hidden *bool `xml:"hidden,attr,omitempty"`

	//TODO: implement child elements
	// Sequence [1..1]
	// a:hlinkClick [0..1]    Drawing Element On Click Hyperlink
	// a:hlinkHover [0..1]    Hyperlink for Hover
	// a:extLst [0..1]    Extension List
}

func NewNonVisProp(id uint, name string) *CNvPr {
	return &CNvPr{
		ID:   id,
		Name: name,
	}
}

func (c CNvPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// ! NOTE: Disabling the empty name check for the Picture
	//  since popular docx tools allow them
	// if c.Name == "" {
	// 	return fmt.Errorf("invalid Name for Non-Visual Drawing Properties when marshaling")
	// }

	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "id"}, Value: strconv.FormatUint(uint64(c.ID), 10)},
		{Name: xml.Name{Local: "name"}, Value: c.Name},
	}

	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "descr"}, Value: c.Description})

	if c.Hidden != nil {
		if *c.Hidden {
			start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hidden"}, Value: "true"})
		} else {
			start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hidden"}, Value: "false"})
		}
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
