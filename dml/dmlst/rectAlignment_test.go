package dmlst

import (
	"encoding/xml"
	"testing"
)

func TestRectAlignmentFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected RectAlignment
	}{
		{"tl", RectAlignmentTopLeft},
		{"t", RectAlignmentTop},
		{"tr", RectAlignmentTopRight},
		{"l", RectAlignmentLeft},
		{"ctr", RectAlignmentCenter},
		{"r", RectAlignmentRight},
		{"bl", RectAlignmentBottomLeft},
		{"b", RectAlignmentBottom},
		{"br", RectAlignmentBottomRight},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := RectAlignmentFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestRectAlignmentFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := RectAlignmentFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "Invalid RectAlignment value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestRectAlignment_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected RectAlignment
	}{
		{`<element val="tl"></element>`, RectAlignmentTopLeft},
		{`<element val="t"></element>`, RectAlignmentTop},
		{`<element val="tr"></element>`, RectAlignmentTopRight},
		{`<element val="l"></element>`, RectAlignmentLeft},
		{`<element val="ctr"></element>`, RectAlignmentCenter},
		{`<element val="r"></element>`, RectAlignmentRight},
		{`<element val="bl"></element>`, RectAlignmentBottomLeft},
		{`<element val="b"></element>`, RectAlignmentBottom},
		{`<element val="br"></element>`, RectAlignmentBottomRight},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName xml.Name      `xml:"element"`
				Val     RectAlignment `xml:"val,attr"`
			}

			var elem Element

			err := xml.Unmarshal([]byte(tt.inputXML), &elem)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if elem.Val != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, elem.Val)
			}
		})
	}
}

func TestRectAlignment_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element val="invalidValue"></element>`

	type Element struct {
		XMLName xml.Name      `xml:"element"`
		Val     RectAlignment `xml:"val,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "Invalid RectAlignment value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
