package stypes

import (
	"encoding/xml"
	"testing"
)

func TestVerticalJc_MarshalXMLAttr(t *testing.T) {
	tests := []struct {
		name     string
		input    VerticalJc
		expected xml.Attr
	}{
		{
			name:     "Marshal 'top'",
			input:    VerticalJcTop,
			expected: xml.Attr{Name: xml.Name{Local: "alignment"}, Value: "top"},
		},
		{
			name:     "Marshal 'center'",
			input:    VerticalJcCenter,
			expected: xml.Attr{Name: xml.Name{Local: "alignment"}, Value: "center"},
		},
		{
			name:     "Marshal 'both'",
			input:    VerticalJcBoth,
			expected: xml.Attr{Name: xml.Name{Local: "alignment"}, Value: "both"},
		},
		{
			name:     "Marshal 'bottom'",
			input:    VerticalJcBottom,
			expected: xml.Attr{Name: xml.Name{Local: "alignment"}, Value: "bottom"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attr, err := tt.input.MarshalXMLAttr(xml.Name{Local: "alignment"})
			if err != nil {
				t.Fatalf("Error marshaling XML attribute: %v", err)
			}

			if attr.Name.Local != "alignment" {
				t.Errorf("Expected attribute name 'alignment', got '%s'", attr.Name.Local)
			}

			if attr.Value != tt.expected.Value {
				t.Errorf("Expected attribute value '%s', got '%s'", tt.expected.Value, attr.Value)
			}
		})
	}
}

func TestVerticalJc_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		name     string
		input    xml.Attr
		expected VerticalJc
	}{
		{
			name:     "Unmarshal attribute 'top'",
			input:    xml.Attr{Name: xml.Name{Local: "alignment"}, Value: "top"},
			expected: VerticalJcTop,
		},
		{
			name:     "Unmarshal attribute 'center'",
			input:    xml.Attr{Name: xml.Name{Local: "alignment"}, Value: "center"},
			expected: VerticalJcCenter,
		},
		{
			name:     "Unmarshal attribute 'both'",
			input:    xml.Attr{Name: xml.Name{Local: "alignment"}, Value: "both"},
			expected: VerticalJcBoth,
		},
		{
			name:     "Unmarshal attribute 'bottom'",
			input:    xml.Attr{Name: xml.Name{Local: "alignment"}, Value: "bottom"},
			expected: VerticalJcBottom,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var v VerticalJc
			err := v.UnmarshalXMLAttr(tt.input)
			if err != nil {
				t.Fatalf("Error unmarshaling XML attribute: %v", err)
			}

			if v != tt.expected {
				t.Errorf("Expected VerticalJc '%s', got '%s'", tt.expected, v)
			}
		})
	}
}

func TestVerticalJc_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	input := xml.Attr{Name: xml.Name{Local: "alignment"}, Value: "invalid"}

	var v VerticalJc
	err := v.UnmarshalXMLAttr(input)
	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "unexpected value for VerticalJc: invalid"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
