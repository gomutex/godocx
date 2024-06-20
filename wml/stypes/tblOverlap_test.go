package stypes

import (
	"encoding/xml"
	"testing"
)

func TestTblOverlapFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected TblOverlap
	}{
		{"never", TblOverlapNever},
		{"overlap", TblOverlapOverlap},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := TblOverlapFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestTblOverlapFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := TblOverlapFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "Invalid TblOverlap value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestTblOverlap_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected TblOverlap
	}{
		{`<element overlap="never"></element>`, TblOverlapNever},
		{`<element overlap="overlap"></element>`, TblOverlapOverlap},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName xml.Name   `xml:"element"`
				Overlap TblOverlap `xml:"overlap,attr"`
			}

			var elem Element

			err := xml.Unmarshal([]byte(tt.inputXML), &elem)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if elem.Overlap != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, elem.Overlap)
			}
		})
	}
}

func TestTblOverlap_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element overlap="invalidValue"></element>`

	type Element struct {
		XMLName xml.Name   `xml:"element"`
		Overlap TblOverlap `xml:"overlap,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "Invalid TblOverlap value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
