package stypes

import (
	"encoding/xml"
	"testing"
)

func TestSectionMarkFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected SectionMark
	}{
		{"nextPage", SectionMarkNextPage},
		{"nextColumn", SectionMarkNextColumn},
		{"continuous", SectionMarkNextContinuous},
		{"evenPage", SectionMarkEvenPage},
		{"oddPage", SectionMarkOddPage},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := SectionMarkFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestSectionMarkFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := SectionMarkFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "Invalid Section Mark"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestSectionMark_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected SectionMark
	}{
		{`<element mark="nextPage"></element>`, SectionMarkNextPage},
		{`<element mark="nextColumn"></element>`, SectionMarkNextColumn},
		{`<element mark="continuous"></element>`, SectionMarkNextContinuous},
		{`<element mark="evenPage"></element>`, SectionMarkEvenPage},
		{`<element mark="oddPage"></element>`, SectionMarkOddPage},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName xml.Name    `xml:"element"`
				Mark    SectionMark `xml:"mark,attr"`
			}

			var elem Element

			err := xml.Unmarshal([]byte(tt.inputXML), &elem)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if elem.Mark != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, elem.Mark)
			}
		})
	}
}

func TestSectionMark_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element mark="invalidValue"></element>`

	type Element struct {
		XMLName xml.Name    `xml:"element"`
		Mark    SectionMark `xml:"mark,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "Invalid Section Mark"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
