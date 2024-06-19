package stypes

import (
	"encoding/xml"
	"testing"
)

func TestCustTabStopFromStr(t *testing.T) {
	tests := []struct {
		input    string
		expected CustTabStop
	}{
		{"clear", CustTabStopClear},
		{"left", CustTabStopLeft},
		{"center", CustTabStopCenter},
		{"right", CustTabStopRight},
		{"decimal", CustTabStopDecimal},
		{"bar", CustTabStopBar},
		{"num", CustTabStopNum},
		{"invalid", CustTabStopInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := CustTabStopFromStr(tt.input)
			if tt.expected == CustTabStopInvalid && err == nil {
				t.Fatalf("Expected error for input %s but got none", tt.input)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestCustTabStop_UnmarshalXMLAttr(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected CustTabStop
	}{
		{
			name:     "Valid attribute clear",
			inputXML: `<element tab="clear"></element>`,
			expected: CustTabStopClear,
		},
		{
			name:     "Valid attribute left",
			inputXML: `<element tab="left"></element>`,
			expected: CustTabStopLeft,
		},
		{
			name:     "Valid attribute center",
			inputXML: `<element tab="center"></element>`,
			expected: CustTabStopCenter,
		},
		{
			name:     "Invalid attribute",
			inputXML: `<element tab="invalid"></element>`,
			expected: CustTabStopInvalid,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			type Element struct {
				XMLName xml.Name    `xml:"element"`
				Tab     CustTabStop `xml:"tab,attr"`
			}

			var elem Element

			err := xml.Unmarshal([]byte(tt.inputXML), &elem)
			if tt.expected == CustTabStopInvalid && err == nil {
				t.Fatalf("Expected error for input XML %s but got none", tt.inputXML)
			}

			if elem.Tab != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, elem.Tab)
			}
		})
	}
}
