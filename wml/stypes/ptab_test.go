package stypes

import (
	"encoding/xml"
	"testing"
)

func TestPTabLeaderFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected PTabLeader
	}{
		{"none", PTabLeaderNone},
		{"dot", PTabLeaderDot},
		{"hyphen", PTabLeaderHyphen},
		{"underscore", PTabLeaderUnderscore},
		{"middleDot", PTabLeaderMiddleDot},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := PTabLeaderFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestPTabLeaderFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := PTabLeaderFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "invalid PTabLeader value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestPTabLeader_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected PTabLeader
	}{
		{`<element tabLeader="none"></element>`, PTabLeaderNone},
		{`<element tabLeader="dot"></element>`, PTabLeaderDot},
		{`<element tabLeader="hyphen"></element>`, PTabLeaderHyphen},
		{`<element tabLeader="underscore"></element>`, PTabLeaderUnderscore},
		{`<element tabLeader="middleDot"></element>`, PTabLeaderMiddleDot},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName   xml.Name   `xml:"element"`
				TabLeader PTabLeader `xml:"tabLeader,attr"`
			}

			var elem Element

			err := xml.Unmarshal([]byte(tt.inputXML), &elem)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if elem.TabLeader != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, elem.TabLeader)
			}
		})
	}
}

func TestPTabLeader_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element tabLeader="invalidValue"></element>`

	type Element struct {
		XMLName   xml.Name   `xml:"element"`
		TabLeader PTabLeader `xml:"tabLeader,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "invalid PTabLeader value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestPTabRelativeToFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected PTabRelativeTo
	}{
		{"margin", PTabRelativeToMargin},
		{"indent", PTabRelativeToIndent},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := PTabRelativeToFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestPTabRelativeToFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := PTabRelativeToFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "invalid PTabRelativeTo value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestPTabRelativeTo_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected PTabRelativeTo
	}{
		{`<element tabRelativeTo="margin"></element>`, PTabRelativeToMargin},
		{`<element tabRelativeTo="indent"></element>`, PTabRelativeToIndent},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName       xml.Name       `xml:"element"`
				TabRelativeTo PTabRelativeTo `xml:"tabRelativeTo,attr"`
			}

			var elem Element

			err := xml.Unmarshal([]byte(tt.inputXML), &elem)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if elem.TabRelativeTo != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, elem.TabRelativeTo)
			}
		})
	}
}

func TestPTabRelativeTo_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element tabRelativeTo="invalidValue"></element>`

	type Element struct {
		XMLName       xml.Name       `xml:"element"`
		TabRelativeTo PTabRelativeTo `xml:"tabRelativeTo,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "invalid PTabRelativeTo value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestPTabAlignmentFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected PTabAlignment
	}{
		{"left", PTabAlignmentLeft},
		{"center", PTabAlignmentCenter},
		{"right", PTabAlignmentRight},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := PTabAlignmentFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestPTabAlignmentFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := PTabAlignmentFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "invalid PTabAlignment value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestPTabAlignment_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected PTabAlignment
	}{
		{`<element tabAlignment="left"></element>`, PTabAlignmentLeft},
		{`<element tabAlignment="center"></element>`, PTabAlignmentCenter},
		{`<element tabAlignment="right"></element>`, PTabAlignmentRight},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName      xml.Name      `xml:"element"`
				TabAlignment PTabAlignment `xml:"tabAlignment,attr"`
			}

			var elem Element

			err := xml.Unmarshal([]byte(tt.inputXML), &elem)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if elem.TabAlignment != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, elem.TabAlignment)
			}
		})
	}
}

func TestPTabAlignment_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element tabAlignment="invalidValue"></element>`

	type Element struct {
		XMLName      xml.Name      `xml:"element"`
		TabAlignment PTabAlignment `xml:"tabAlignment,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "invalid PTabAlignment value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
