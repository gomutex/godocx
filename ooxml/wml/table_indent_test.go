package wml

import (
	"encoding/xml"
	"testing"
)

func TestTableIndent(t *testing.T) {
	tests := []struct {
		name      string
		width     int
		widthType string
		expected  string
	}{
		{
			name:      "WithWidthType",
			width:     0,
			widthType: "dxa",
			expected:  `<w:tblInd w:w="0" w:type="dxa"></w:tblInd>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widthElement := &TableIndent{
				Width:     tt.width,
				WidthType: WidthType(tt.widthType),
			}

			output, err := xml.Marshal(widthElement)
			if err != nil {
				t.Errorf("Unexpected error during XML marshalling: %v", err)
				return
			}

			actual := string(output)
			if actual != tt.expected {
				t.Errorf("Expected:\n%s\nActual:\n%s", tt.expected, actual)
			}

			var unmarshaledWidthElement TableIndent
			err = xml.Unmarshal(output, &unmarshaledWidthElement)
			if tt.expected != "" {
				if err != nil {
					t.Errorf("Unexpected error during XML unmarshalling: %v", err)
					return
				}
				if unmarshaledWidthElement.Width != tt.width || unmarshaledWidthElement.WidthType != WidthType(tt.widthType) {
					t.Errorf("Unmarshalled values do not match expected values")
				}
			} else {
				if err == nil {
					t.Errorf("Expected error during XML unmarshalling")
				}
			}

		})
	}
}
