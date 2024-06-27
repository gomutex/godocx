// Package oxml provides utilities for working with Office Open XML (OOXML) documents,
// including functions related to encoding and decoding XML elements.
package docx

import "encoding/xml"

// docxDcTerms represents an XML element with text content and an xsi:type attribute.
// Core Document Properties
type docxDcTerms struct {
	Text string `xml:",chardata"`
	Type string `xml:"xsi:type,attr"`
}

// Core Document Properties
type decodeDcTerms struct {
	Text string `xml:",chardata"`
	Type string `xml:"xsi:type,attr"`
}

// The function adds an XML header to the encoded data.
func marshal(data any) (out string, err error) {
	body, err := xml.Marshal(data)
	if err != nil {
		return
	}

	out = xml.Header + string(body)
	return
}
