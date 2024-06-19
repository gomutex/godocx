package stypes

import (
	"encoding/xml"
	"strconv"
	"testing"
)

func TestTextScaleFromUint16_ValidValues(t *testing.T) {
	tests := []struct {
		input    uint16
		expected TextScale
	}{
		{50, 50},
		{600, 600},
	}

	for _, tt := range tests {
		t.Run(strconv.Itoa(int(tt.input)), func(t *testing.T) {
			result, err := TextScaleFromUint16(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %d but got %d", tt.expected, result)
			}
		})
	}
}

func TestTextScaleFromUint16_InvalidValue(t *testing.T) {
	input := uint16(700)

	_, err := TextScaleFromUint16(input)
	if err == nil {
		t.Fatalf("Expected error for invalid value %d, but got none", input)
	}

	expectedError := "Invalid Text Scale"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestTextScaleFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected TextScale
	}{
		{"50", 50},
		{"600", 600},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := TextScaleFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %d but got %d", tt.expected, result)
			}
		})
	}
}

func TestTextScaleFromStr_InvalidValue(t *testing.T) {
	input := "700"

	_, err := TextScaleFromStr(input)
	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none", input)
	}

	expectedError := "Invalid Text Scale"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestTextScale_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected TextScale
	}{
		{`<element scale="50"></element>`, 50},
		{`<element scale="600"></element>`, 600},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName xml.Name  `xml:"element"`
				Scale   TextScale `xml:"scale,attr"`
			}

			var elem Element

			err := xml.Unmarshal([]byte(tt.inputXML), &elem)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if elem.Scale != tt.expected {
				t.Errorf("Expected %d but got %d", tt.expected, elem.Scale)
			}
		})
	}
}

func TestTextScale_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element scale="700"></element>`

	type Element struct {
		XMLName xml.Name  `xml:"element"`
		Scale   TextScale `xml:"scale,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "Invalid Text Scale"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestTextScale_ToStr(t *testing.T) {
	tests := []struct {
		input    TextScale
		expected string
	}{
		{50, "50"},
		{600, "600"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := tt.input.ToStr()

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}
