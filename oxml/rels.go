package oxml

import "encoding/xml"

type Relationship struct {
	XMLName    xml.Name `xml:"Relationship"`
	ID         string   `xml:"Id,attr"`
	Type       string   `xml:"Type,attr"`
	Target     string   `xml:"Target,attr"`
	TargetMode string   `xml:"TargetMode,attr,omitempty"`
}

type Relationships struct {
	RelativePath  string          `xml:"-"`
	XMLName       xml.Name        `xml:"Relationships"`
	Xmlns         string          `xml:"xmlns,attr"`
	Relationships []*Relationship `xml:"Relationship"`
}
