package stypes

import (
	"encoding/xml"
	"testing"
)

func TestThemeFontFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected ThemeFont
	}{
		{"majorEastAsia", ThemeFontMajorEastAsia},
		{"majorBidi", ThemeFontMajorBidi},
		{"majorAscii", ThemeFontMajorAscii},
		{"majorHAnsi", ThemeFontMajorHAnsi},
		{"minorEastAsia", ThemeFontMinorEastAsia},
		{"minorBidi", ThemeFontMinorBidi},
		{"minorAscii", ThemeFontMinorAscii},
		{"minorHAnsi", ThemeFontMinorHAnsi},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ThemeFontFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestThemeFontFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := ThemeFontFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "invalid ThemeFont value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestThemeFont_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected ThemeFont
	}{
		{`<element theme="majorEastAsia"></element>`, ThemeFontMajorEastAsia},
		{`<element theme="majorBidi"></element>`, ThemeFontMajorBidi},
		{`<element theme="majorAscii"></element>`, ThemeFontMajorAscii},
		{`<element theme="majorHAnsi"></element>`, ThemeFontMajorHAnsi},
		{`<element theme="minorEastAsia"></element>`, ThemeFontMinorEastAsia},
		{`<element theme="minorBidi"></element>`, ThemeFontMinorBidi},
		{`<element theme="minorAscii"></element>`, ThemeFontMinorAscii},
		{`<element theme="minorHAnsi"></element>`, ThemeFontMinorHAnsi},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName xml.Name  `xml:"element"`
				Theme   ThemeFont `xml:"theme,attr"`
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

func TestThemeFont_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element theme="invalidValue"></element>`

	type Element struct {
		XMLName xml.Name  `xml:"element"`
		Theme   ThemeFont `xml:"theme,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "invalid ThemeFont value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
