package stypes

import (
	"encoding/xml"
	"testing"
)

func TestTblStyleOverrideTypeFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected TblStyleOverrideType
	}{
		{"wholeTable", TblStyleOverrideWholeTable},
		{"firstRow", TblStyleOverrideFirstRow},
		{"lastRow", TblStyleOverrideLastRow},
		{"firstCol", TblStyleOverrideFirstCol},
		{"lastCol", TblStyleOverrideLastCol},
		{"band1Vert", TblStyleOverrideBand1Vert},
		{"band2Vert", TblStyleOverrideBand2Vert},
		{"band1Horz", TblStyleOverrideBand1Horz},
		{"band2Horz", TblStyleOverrideBand2Horz},
		{"neCell", TblStyleOverrideNeCell},
		{"nwCell", TblStyleOverrideNwCell},
		{"seCell", TblStyleOverrideSeCell},
		{"swCell", TblStyleOverrideSwCell},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := TblStyleOverrideTypeFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestTblStyleOverrideTypeFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := TblStyleOverrideTypeFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "invalid TblStyleOverrideType value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestTblStyleOverrideType_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected TblStyleOverrideType
	}{
		{`<element type="wholeTable"></element>`, TblStyleOverrideWholeTable},
		{`<element type="firstRow"></element>`, TblStyleOverrideFirstRow},
		{`<element type="lastRow"></element>`, TblStyleOverrideLastRow},
		{`<element type="firstCol"></element>`, TblStyleOverrideFirstCol},
		{`<element type="lastCol"></element>`, TblStyleOverrideLastCol},
		{`<element type="band1Vert"></element>`, TblStyleOverrideBand1Vert},
		{`<element type="band2Vert"></element>`, TblStyleOverrideBand2Vert},
		{`<element type="band1Horz"></element>`, TblStyleOverrideBand1Horz},
		{`<element type="band2Horz"></element>`, TblStyleOverrideBand2Horz},
		{`<element type="neCell"></element>`, TblStyleOverrideNeCell},
		{`<element type="nwCell"></element>`, TblStyleOverrideNwCell},
		{`<element type="seCell"></element>`, TblStyleOverrideSeCell},
		{`<element type="swCell"></element>`, TblStyleOverrideSwCell},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName xml.Name             `xml:"element"`
				Type    TblStyleOverrideType `xml:"type,attr"`
			}

			var elem Element

			err := xml.Unmarshal([]byte(tt.inputXML), &elem)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if elem.Type != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, elem.Type)
			}
		})
	}
}

func TestTblStyleOverrideType_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element type="invalidValue"></element>`

	type Element struct {
		XMLName xml.Name             `xml:"element"`
		Type    TblStyleOverrideType `xml:"type,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "invalid TblStyleOverrideType value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
