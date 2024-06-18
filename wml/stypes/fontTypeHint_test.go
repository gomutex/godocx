package stypes

import (
	"encoding/xml"
	"testing"
)

func TestFontTypeHintFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected FontTypeHint
	}{
		{"default", FontTypeHintDefault},
		{"eastAsia", FontTypeHintEastAsia},
		{"cs", FontTypeHintCS},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := FontTypeHintFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestFontTypeHintFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := FontTypeHintFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "invalid FontTypeHint value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestFontTypeHint_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected FontTypeHint
	}{
		{`<element theme="default"></element>`, FontTypeHintDefault},
		{`<element theme="eastAsia"></element>`, FontTypeHintEastAsia},
		{`<element theme="cs"></element>`, FontTypeHintCS},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName xml.Name     `xml:"element"`
				Theme   FontTypeHint `xml:"theme,attr"`
			}

			var elem Element

			err := xml.Unmarshal([]byte(tt.inputXML), &elem)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if elem.Theme != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, elem.Theme)
			}
		})
	}
}

func TestFontTypeHint_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element theme="invalidValue"></element>`

	type Element struct {
		XMLName xml.Name     `xml:"element"`
		Theme   FontTypeHint `xml:"theme,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "invalid FontTypeHint value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
