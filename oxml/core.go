package oxml

import "encoding/xml"

type docxDcTerms struct {
	Text string `xml:",chardata"`
	Type string `xml:"xsi:type,attr"`
}

type decodeDcTerms struct {
	Text string `xml:",chardata"`
	Type string `xml:"xsi:type,attr"`
}

func marshal(data any) (out string, err error) {
	body, err := xml.Marshal(data)
	if err != nil {
		return
	}

	out = xml.Header + string(body)
	return
}
