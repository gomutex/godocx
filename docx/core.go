package docx

import (
	"encoding/xml"

	"github.com/gomutex/godocx/common/constants"
)

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
func marshal(data any) ([]byte, error) {
	body, err := xml.Marshal(data)
	if err != nil {
		return nil, err
	}

	out := append(constants.XMLHeader, body...)
	return out, nil
}
