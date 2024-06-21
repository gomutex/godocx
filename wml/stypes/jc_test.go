package stypes

import (
	"encoding/xml"
	"testing"
)

func TestJustificationFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected Justification
	}{
		{"left", JustificationLeft},
		{"center", JustificationCenter},
		{"right", JustificationRight},
		{"both", JustificationBoth},
		{"mediumKashida", JustificationMediumKashida},
		{"distribute", JustificationDistribute},
		{"numTab", JustificationNumTab},
		{"highKashida", JustificationHighKashida},
		{"lowKashida", JustificationLowKashida},
		{"thaiDistribute", JustificationThaiDistribute},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := JustificationFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestJustificationFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := JustificationFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "invalid Justification value"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestJustification_UnmarshalXMLAttr(t *testing.T) {
	tests := []struct {
		name         string
		xmlAttrValue string
		expected     Justification
		expectError  bool
	}{
		{
			name:         "Valid attribute 'left'",
			xmlAttrValue: "left",
			expected:     JustificationLeft,
			expectError:  false,
		},
		{
			name:         "Valid attribute 'center'",
			xmlAttrValue: "center",
			expected:     JustificationCenter,
			expectError:  false,
		},
		{
			name:         "Valid attribute 'right'",
			xmlAttrValue: "right",
			expected:     JustificationRight,
			expectError:  false,
		},
		{
			name:         "Valid attribute 'both'",
			xmlAttrValue: "both",
			expected:     JustificationBoth,
			expectError:  false,
		},
		{
			name:         "Valid attribute 'mediumKashida'",
			xmlAttrValue: "mediumKashida",
			expected:     JustificationMediumKashida,
			expectError:  false,
		},
		{
			name:         "Valid attribute 'distribute'",
			xmlAttrValue: "distribute",
			expected:     JustificationDistribute,
			expectError:  false,
		},
		{
			name:         "Valid attribute 'numTab'",
			xmlAttrValue: "numTab",
			expected:     JustificationNumTab,
			expectError:  false,
		},
		{
			name:         "Valid attribute 'highKashida'",
			xmlAttrValue: "highKashida",
			expected:     JustificationHighKashida,
			expectError:  false,
		},
		{
			name:         "Valid attribute 'lowKashida'",
			xmlAttrValue: "lowKashida",
			expected:     JustificationLowKashida,
			expectError:  false,
		},
		{
			name:         "Valid attribute 'thaiDistribute'",
			xmlAttrValue: "thaiDistribute",
			expected:     JustificationThaiDistribute,
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
			var j Justification
			attr := xml.Attr{Name: xml.Name{Local: "align"}, Value: tt.xmlAttrValue}

			err := j.UnmarshalXMLAttr(attr)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}

				if j != tt.expected {
					t.Errorf("Expected Justification %s but got %s", tt.expected, j)
				}
			}
		})
	}
}
