package ctypes

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"
)

func TestGridChange_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    GridChange
		expected string
	}{
		{
			name:     "With ID",
			input:    GridChange{ID: 1},
			expected: `<w:tblGridChange w:id="1"></w:tblGridChange>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:tblGridChange"}}

			err := tt.input.MarshalXML(encoder, start)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			encoder.Flush()

			if result.String() != tt.expected {
				t.Errorf("Expected XML:\n%s\nGot:\n%s", tt.expected, result.String())
			}
		})
	}
}

func TestGridChange_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name       string
		inputXML   string
		expected   GridChange
		expectFail bool // Whether unmarshalling is expected to fail
	}{
		{
			name:     "With ID",
			inputXML: `<w:tblGridChange w:id="1"></w:tblGridChange>`,
			expected: GridChange{ID: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decoder := xml.NewDecoder(strings.NewReader(tt.inputXML))
			var result GridChange

			err := decoder.Decode(&result)

			if tt.expectFail {
				if err == nil {
					t.Error("Expected unmarshaling to fail but it did not")
				}
				return
			}

			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Unmarshaled GridChange struct does not match expected:\nExpected: %+v\nActual:   %+v", tt.expected, result)
			}
		})
	}
}
