package stypes

import (
	"encoding/xml"
	"testing"
)

func TestCombineBracketsFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected CombineBrackets
	}{
		{"none", CombineBracketsNone},
		{"round", CombineBracketsRound},
		{"square", CombineBracketsSquare},
		{"angle", CombineBracketsAngle},
		{"curly", CombineBracketsCurly},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := CombineBracketsFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestCombineBracketsFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := CombineBracketsFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "invalid CombineBrackets value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestCombineBrackets_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected CombineBrackets
	}{
		{`<element val="none"></element>`, CombineBracketsNone},
		{`<element val="round"></element>`, CombineBracketsRound},
		{`<element val="square"></element>`, CombineBracketsSquare},
		{`<element val="angle"></element>`, CombineBracketsAngle},
		{`<element val="curly"></element>`, CombineBracketsCurly},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName xml.Name        `xml:"element"`
				Val     CombineBrackets `xml:"val,attr"`
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

func TestCombineBrackets_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element val="invalidValue"></element>`

	type Element struct {
		XMLName xml.Name        `xml:"element"`
		Val     CombineBrackets `xml:"val,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "invalid CombineBrackets value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
