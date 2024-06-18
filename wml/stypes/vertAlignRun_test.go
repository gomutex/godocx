package stypes

import (
	"encoding/xml"
	"testing"
)

func TestVerticalAlignRunFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected VerticalAlignRun
	}{
		{"baseline", VerticalAlignRunBaseline},
		{"superscript", VerticalAlignRunSuperscript},
		{"subscript", VerticalAlignRunSubscript},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := VerticalAlignRunFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestVerticalAlignRunFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := VerticalAlignRunFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "Invalid VerticalAlignRun Type"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestVerticalAlignRun_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected VerticalAlignRun
	}{
		{`<element verticalAlign="baseline"></element>`, VerticalAlignRunBaseline},
		{`<element verticalAlign="superscript"></element>`, VerticalAlignRunSuperscript},
		{`<element verticalAlign="subscript"></element>`, VerticalAlignRunSubscript},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName       xml.Name         `xml:"element"`
				VerticalAlign VerticalAlignRun `xml:"verticalAlign,attr"`
			}

			var elem Element

			err := xml.Unmarshal([]byte(tt.inputXML), &elem)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if elem.VerticalAlign != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, elem.VerticalAlign)
			}
		})
	}
}

func TestVerticalAlignRun_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element verticalAlign="invalidValue"></element>`

	type Element struct {
		XMLName       xml.Name         `xml:"element"`
		VerticalAlign VerticalAlignRun `xml:"verticalAlign,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "Invalid VerticalAlignRun Type"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
