package geom

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestMarshalAdjustValues(t *testing.T) {
	tests := []struct {
		adjustValues *AdjustValues
		expectedXML  string
	}{
		{
			adjustValues: &AdjustValues{
				ShapeGuides: []ShapeGuide{
					{Name: "guide1", Formula: "formula1"},
					{Name: "guide2", Formula: "formula2"},
				},
			},
			expectedXML: `<a:avLst><a:gd name="guide1" fmla="formula1"></a:gd><a:gd name="guide2" fmla="formula2"></a:gd></a:avLst>`,
		},
		{
			adjustValues: &AdjustValues{
				ShapeGuides: []ShapeGuide{},
			},
			expectedXML: `<a:avLst></a:avLst>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.expectedXML, func(t *testing.T) {
			generatedXML, err := xml.Marshal(tt.adjustValues)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			if strings.TrimSpace(string(generatedXML)) != tt.expectedXML {
				t.Errorf("Expected XML:\n%s\nBut got:\n%s", tt.expectedXML, generatedXML)
			}
		})
	}
}

func TestUnmarshalAdjustValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected AdjustValues
	}{
		{
			inputXML: `<a:avLst><a:gd name="guide1" fmla="formula1"></a:gd><a:gd name="guide2" fmla="formula2"></a:gd></a:avLst>`,
			expected: AdjustValues{
				ShapeGuides: []ShapeGuide{
					{Name: "guide1", Formula: "formula1"},
					{Name: "guide2", Formula: "formula2"},
				},
			},
		},
		{
			inputXML: `<a:avLst></a:avLst>`,
			expected: AdjustValues{
				ShapeGuides: []ShapeGuide{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			var adjustValues AdjustValues

			err := xml.Unmarshal([]byte(tt.inputXML), &adjustValues)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if len(adjustValues.ShapeGuides) != len(tt.expected.ShapeGuides) {
				t.Errorf("Expected %d ShapeGuides, but got %d", len(tt.expected.ShapeGuides), len(adjustValues.ShapeGuides))
			} else {
				for i := range tt.expected.ShapeGuides {
					if adjustValues.ShapeGuides[i].Name != tt.expected.ShapeGuides[i].Name {
						t.Errorf("Expected ShapeGuide Name %s, but got %s", tt.expected.ShapeGuides[i].Name, adjustValues.ShapeGuides[i].Name)
					}
					if adjustValues.ShapeGuides[i].Formula != tt.expected.ShapeGuides[i].Formula {
						t.Errorf("Expected ShapeGuide Formula %s, but got %s", tt.expected.ShapeGuides[i].Formula, adjustValues.ShapeGuides[i].Formula)
					}
				}
			}
		})
	}
}
