package stypes

import (
	"encoding/xml"
	"testing"
)

func TestTableWidthFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected TableWidth
	}{
		{"dxa", TableWidthDxa},
		{"auto", TableWidthAuto},
		{"pct", TableWidthPct},
		{"nil", TableWidthNil},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := TableWidthFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestTableWidthFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := TableWidthFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "Invalid TableWidth value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestTableWidth_UnmarshalXMLAttr(t *testing.T) {
	tests := []struct {
		name         string
		xmlAttrValue string
		expected     TableWidth
		expectError  bool
	}{
		{
			name:         "Valid attribute 'dxa'",
			xmlAttrValue: "dxa",
			expected:     TableWidthDxa,
			expectError:  false,
		},
		{
			name:         "Valid attribute 'auto'",
			xmlAttrValue: "auto",
			expected:     TableWidthAuto,
			expectError:  false,
		},
		{
			name:         "Valid attribute 'pct'",
			xmlAttrValue: "pct",
			expected:     TableWidthPct,
			expectError:  false,
		},
		{
			name:         "Valid attribute 'nil'",
			xmlAttrValue: "nil",
			expected:     TableWidthNil,
			expectError:  false,
		},
		{
			name:         "Invalid attribute value",
			xmlAttrValue: "invalid",
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tw TableWidth
			attr := xml.Attr{Name: xml.Name{Local: "width"}, Value: tt.xmlAttrValue}

			err := tw.UnmarshalXMLAttr(attr)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}

				if tw != tt.expected {
					t.Errorf("Expected TableWidth %s but got %s", tt.expected, tw)
				}
			}
		})
	}
}
