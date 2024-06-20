package stypes

import (
	"encoding/xml"
	"testing"
)

func TestXAlignFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected XAlign
	}{
		{"left", XAlignLeft},
		{"center", XAlignCenter},
		{"right", XAlignRight},
		{"inside", XAlignInside},
		{"outside", XAlignOutside},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := XAlignFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestXAlignFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := XAlignFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "Invalid XAlign value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestXAlign_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected XAlign
	}{
		{`<element val="left"></element>`, XAlignLeft},
		{`<element val="center"></element>`, XAlignCenter},
		{`<element val="right"></element>`, XAlignRight},
		{`<element val="inside"></element>`, XAlignInside},
		{`<element val="outside"></element>`, XAlignOutside},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName xml.Name `xml:"element"`
				Val     XAlign   `xml:"val,attr"`
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

func TestXAlign_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element val="invalidValue"></element>`

	type Element struct {
		XMLName xml.Name `xml:"element"`
		Val     XAlign   `xml:"val,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "Invalid XAlign value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestYAlignFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected YAlign
	}{
		{"inline", YAlignInline},
		{"top", YAlignTop},
		{"center", YAlignCenter},
		{"bottom", YAlignBottom},
		{"inside", YAlignInside},
		{"outside", YAlignOutside},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := YAlignFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestYAlignFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := YAlignFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "Invalid YAlign value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestYAlign_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected YAlign
	}{
		{`<element val="inline"></element>`, YAlignInline},
		{`<element val="top"></element>`, YAlignTop},
		{`<element val="center"></element>`, YAlignCenter},
		{`<element val="bottom"></element>`, YAlignBottom},
		{`<element val="inside"></element>`, YAlignInside},
		{`<element val="outside"></element>`, YAlignOutside},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName xml.Name `xml:"element"`
				Val     YAlign   `xml:"val,attr"`
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

func TestYAlign_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element val="invalidValue"></element>`

	type Element struct {
		XMLName xml.Name `xml:"element"`
		Val     YAlign   `xml:"val,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "Invalid YAlign value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
