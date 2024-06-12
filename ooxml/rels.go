package ooxml

import "encoding/xml"

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
