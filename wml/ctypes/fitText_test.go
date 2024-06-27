package ctypes

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"
)

func TestFitText_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    FitText
		expected string
	}{
		{
			name:     "With ID",
			input:    FitText{Val: 123, ID: IntPtr(456)},
			expected: `<w:fitText w:val="123" w:id="456"></w:fitText>`,
		},
		{
			name:     "Without ID",
			input:    FitText{Val: 789},
			expected: `<w:fitText w:val="789"></w:fitText>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:fitText"}}

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

func TestFitText_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected FitText
	}{
		{
			name:     "With ID",
			inputXML: `<w:fitText w:val="123" w:id="456"></w:fitText>`,
			expected: FitText{Val: 123, ID: IntPtr(456)},
		},
		{
			name:     "Without ID",
			inputXML: `<w:fitText w:val="789"></w:fitText>`,
			expected: FitText{Val: 789},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result FitText

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %+v but got %+v", tt.expected, result)
			}
		})
	}
}

func IntPtr(i int) *int {
	return &i
}
