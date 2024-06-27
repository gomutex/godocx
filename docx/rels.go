package docx

import (
	"encoding/xml"
)

// Relationship represents a relationship between elements in an Office Open XML (OOXML) document.
// It includes essential information such as ID, type, target, and target mode.
type Relationship struct {
	XMLName    xml.Name `xml:"Relationship"`
	ID         string   `xml:"Id,attr"`
	Type       string   `xml:"Type,attr"`
	Target     string   `xml:"Target,attr"`
	TargetMode string   `xml:"TargetMode,attr,omitempty"`
}

// Relationships represents a collection of relationships in an OOXML document.
// It includes the relative path, XML namespace, and a slice of Relationship instances.
type Relationships struct {
	RelativePath  string          `xml:"-"`
	XMLName       xml.Name        `xml:"Relationships"`
	Xmlns         string          `xml:"xmlns,attr"`
	Relationships []*Relationship `xml:"Relationship"`
}

func (r Relationship) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "Relationship"
	start.Attr = []xml.Attr{}

	if r.ID != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "Id"}, Value: r.ID})
	}

	if r.Type != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "Type"}, Value: r.Type})
	}

	if r.Target != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "Target"}, Value: r.Target})
	}

	if r.TargetMode != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "TargetMode"}, Value: r.TargetMode})
	}

	return e.EncodeElement("", start)
}
