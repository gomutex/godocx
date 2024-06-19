package dmlst

import (
	"encoding/xml"
	"testing"
)

func TestRelFromHFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected RelFromH
	}{
		{"character", RelFromHCharacter},
		{"column", RelFromHColumn},
		{"insideMargin", RelFromHInsideMargin},
		{"leftMargin", RelFromHLeftMargin},
		{"margin", RelFromHMargin},
		{"outsizeMargin", RelFromHOutsizeMargin},
		{"page", RelFromHPage},
		{"rightMargin", RelFromHRightMargin},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := RelFromHFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestRelFromHFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := RelFromHFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "Invalid RelFromH value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestRelFromH_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected RelFromH
	}{
		{`<element val="character"></element>`, RelFromHCharacter},
		{`<element val="column"></element>`, RelFromHColumn},
		{`<element val="insideMargin"></element>`, RelFromHInsideMargin},
		{`<element val="leftMargin"></element>`, RelFromHLeftMargin},
		{`<element val="margin"></element>`, RelFromHMargin},
		{`<element val="outsizeMargin"></element>`, RelFromHOutsizeMargin},
		{`<element val="page"></element>`, RelFromHPage},
		{`<element val="rightMargin"></element>`, RelFromHRightMargin},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName xml.Name `xml:"element"`
				Val     RelFromH `xml:"val,attr"`
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

func TestRelFromH_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element val="invalidValue"></element>`

	type Element struct {
		XMLName xml.Name `xml:"element"`
		Val     RelFromH `xml:"val,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "Invalid RelFromH value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestRelFromVFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected RelFromV
	}{
		{"bottomMargin", RelFromVBottomMargin},
		{"insideMargin", RelFromVInsideMargin},
		{"line", RelFromVLine},
		{"margin", RelFromVMargin},
		{"outsizeMargin", RelFromVOutsizeMargin},
		{"page", RelFromVPage},
		{"paragraph", RelFromVParagraph},
		{"topMargin", RelFromVTopMargin},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := RelFromVFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestRelFromVFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := RelFromVFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "Invalid RelFromV value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestRelFromV_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected RelFromV
	}{
		{`<element val="bottomMargin"></element>`, RelFromVBottomMargin},
		{`<element val="insideMargin"></element>`, RelFromVInsideMargin},
		{`<element val="line"></element>`, RelFromVLine},
		{`<element val="margin"></element>`, RelFromVMargin},
		{`<element val="outsizeMargin"></element>`, RelFromVOutsizeMargin},
		{`<element val="page"></element>`, RelFromVPage},
		{`<element val="paragraph"></element>`, RelFromVParagraph},
		{`<element val="topMargin"></element>`, RelFromVTopMargin},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName xml.Name `xml:"element"`
				Val     RelFromV `xml:"val,attr"`
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

func TestRelFromV_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element val="invalidValue"></element>`

	type Element struct {
		XMLName xml.Name `xml:"element"`
		Val     RelFromV `xml:"val,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "Invalid RelFromV value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
