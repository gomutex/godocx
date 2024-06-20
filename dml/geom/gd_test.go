package geom

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestShapeGuide_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    ShapeGuide
		expected string
	}{
		{
			name: "With name and formula",
			input: ShapeGuide{
				Name:    "height",
				Formula: "val 100",
			},
			expected: `<a:gd name="height" fmla="val 100"></a:gd>`,
		},
		{
			name: "Empty name and formula",
			input: ShapeGuide{
				Name:    "",
				Formula: "",
			},
			expected: `<a:gd name="" fmla=""></a:gd>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "a:gd"}}

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

func TestShapeGuide_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected ShapeGuide
	}{
		{
			name:     "With name and formula",
			inputXML: `<a:gd name="height" fmla="val 100"></a:gd>`,
			expected: ShapeGuide{
				Name:    "height",
				Formula: "val 100",
			},
		},
		{
			name:     "Empty name and formula",
			inputXML: `<a:gd name="" fmla=""></a:gd>`,
			expected: ShapeGuide{
				Name:    "",
				Formula: "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result ShapeGuide

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if result.Name != tt.expected.Name {
				t.Errorf("Expected Name %s but got %s", tt.expected.Name, result.Name)
			}
			if result.Formula != tt.expected.Formula {
				t.Errorf("Expected Formula %s but got %s", tt.expected.Formula, result.Formula)
			}
		})
	}
}
