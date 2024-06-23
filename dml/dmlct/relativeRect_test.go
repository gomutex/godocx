package dmlct

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"

	"github.com/gomutex/godocx/internal"
)

func TestRelativeRect_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    RelativeRect
		expected string
	}{
		{
			name: "All fields set",
			input: RelativeRect{
				Left:   internal.ToPtr(10),
				Right:  internal.ToPtr(20),
				Top:    internal.ToPtr(30),
				Bottom: internal.ToPtr(40),
			},
			expected: `<RelativeRect t="30" l="10" b="40" r="20"></RelativeRect>`,
		},
		{
			name: "Some fields set",
			input: RelativeRect{
				Top:    internal.ToPtr(30),
				Bottom: internal.ToPtr(40),
			},
			expected: `<RelativeRect t="30" b="40"></RelativeRect>`,
		},
		{
			name:     "No fields set",
			input:    RelativeRect{},
			expected: `<RelativeRect></RelativeRect>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)

			start := xml.StartElement{Name: xml.Name{Local: "RelativeRect"}}
			if err := tt.input.MarshalXML(encoder, start); err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			encoder.Flush()

			got := strings.TrimSpace(result.String())
			if got != tt.expected {
				t.Errorf("Expected XML:\n%s\nGot:\n%s", tt.expected, got)
			}
		})
	}
}

func TestRelativeRect_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected RelativeRect
	}{
		{
			name:     "All fields set",
			inputXML: `<RelativeRect l="10" r="20" t="30" b="40"></RelativeRect>`,
			expected: RelativeRect{
				Left:   internal.ToPtr(10),
				Right:  internal.ToPtr(20),
				Top:    internal.ToPtr(30),
				Bottom: internal.ToPtr(40),
			},
		},
		{
			name:     "Some fields set",
			inputXML: `<RelativeRect t="30" b="40"></RelativeRect>`,
			expected: RelativeRect{
				Top:    internal.ToPtr(30),
				Bottom: internal.ToPtr(40),
			},
		},
		{
			name:     "No fields set",
			inputXML: `<RelativeRect></RelativeRect>`,
			expected: RelativeRect{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result RelativeRect

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			// Compare unmarshaled result with expected
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Unmarshaled RelativeRect does not match expected:\nExpected: %+v\nGot: %+v", tt.expected, result)
			}
		})
	}
}
