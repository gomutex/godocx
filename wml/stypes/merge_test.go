package stypes

import (
	"encoding/xml"
	"testing"
)

func TestMergeCellFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected MergeCell
	}{
		{"continue", MergeCellContinue},
		{"restart", MergeCellRestart},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := MergeCellFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestMergeCellFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := MergeCellFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "invalid MergeCell value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestMergeCell_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected MergeCell
	}{
		{`<element val="continue"></element>`, MergeCellContinue},
		{`<element val="restart"></element>`, MergeCellRestart},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName xml.Name  `xml:"element"`
				Val     MergeCell `xml:"val,attr"`
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

func TestMergeCell_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element val="invalidValue"></element>`

	type Element struct {
		XMLName xml.Name  `xml:"element"`
		Val     MergeCell `xml:"val,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "invalid MergeCell value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
