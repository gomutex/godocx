package dmlst

import (
	"encoding/xml"
	"testing"
)

func TestTileFlipModeFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected TileFlipMode
	}{
		{"none", TileFlipModeNone},
		{"x", TileFlipModeHorizontal},
		{"y", TileFlipModeVertical},
		{"xy", TileFlipModeBoth},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := TileFlipModeFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestTileFlipModeFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := TileFlipModeFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "Invalid TileFlipMode value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestTileFlipMode_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected TileFlipMode
	}{
		{`<element val="none"></element>`, TileFlipModeNone},
		{`<element val="x"></element>`, TileFlipModeHorizontal},
		{`<element val="y"></element>`, TileFlipModeVertical},
		{`<element val="xy"></element>`, TileFlipModeBoth},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName xml.Name     `xml:"element"`
				Val     TileFlipMode `xml:"val,attr"`
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

func TestTileFlipMode_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element val="invalidValue"></element>`

	type Element struct {
		XMLName xml.Name     `xml:"element"`
		Val     TileFlipMode `xml:"val,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "Invalid TileFlipMode value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
