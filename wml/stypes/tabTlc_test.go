package stypes

import (
	"encoding/xml"
	"testing"
)

func TestCustLeadCharFromStr(t *testing.T) {
	tests := []struct {
		input    string
		expected CustLeadChar
	}{
		{"none", CustLeadCharNone},
		{"dot", CustLeadCharDot},
		{"hyphen", CustLeadCharHyphen},
		{"underscore", CustLeadCharUnderScore},
		{"heavy", CustLeadCharHeavy},
		{"middleDot", CustLeadCharMiddleDot},
		{"invalid", CustLeadCharInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := CustLeadCharFromStr(tt.input)
			if tt.expected == CustLeadCharInvalid && err == nil {
				t.Fatalf("Expected error for input %s but got none", tt.input)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestCustLeadChar_UnmarshalXMLAttr(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected CustLeadChar
	}{
		{
			name:     "Valid attribute none",
			inputXML: `<element leader="none"></element>`,
			expected: CustLeadCharNone,
		},
		{
			name:     "Valid attribute dot",
			inputXML: `<element leader="dot"></element>`,
			expected: CustLeadCharDot,
		},
		{
			name:     "Valid attribute hyphen",
			inputXML: `<element leader="hyphen"></element>`,
			expected: CustLeadCharHyphen,
		},
		{
			name:     "Invalid attribute",
			inputXML: `<element leader="invalid"></element>`,
			expected: CustLeadCharInvalid,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			type Element struct {
				XMLName xml.Name     `xml:"element"`
				Leader  CustLeadChar `xml:"leader,attr"`
			}

			var elem Element

			err := xml.Unmarshal([]byte(tt.inputXML), &elem)
			if tt.expected == CustLeadCharInvalid && err == nil {
				t.Fatalf("Expected error for input XML %s but got none", tt.inputXML)
			}

			if elem.Leader != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, elem.Leader)
			}
		})
	}
}
