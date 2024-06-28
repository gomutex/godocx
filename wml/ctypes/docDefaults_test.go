package ctypes

import (
	"bytes"
	"encoding/xml"
	"reflect"
	"testing"
)

func TestDocDefault_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		doc      DocDefault
		expected string
	}{
		{
			name: "Marshal DocDefault with RunPropDefault and ParaPropDefault",
			doc: DocDefault{
				RunProp: &RunPropDefault{
					RunProp: &RunProperty{},
				},
				ParaProp: &ParaPropDefault{
					ParaProp: &ParagraphProp{},
				},
			},
			expected: `<w:docDefaults><w:rPrDefault><w:rPr></w:rPr></w:rPrDefault><w:pPrDefault><w:pPr></w:pPr></w:pPrDefault></w:docDefaults>`,
		},
		{
			name:     "Marshal DocDefault with nil RunPropDefault and ParaPropDefault",
			doc:      DocDefault{},
			expected: `<w:docDefaults></w:docDefaults>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			encoder := xml.NewEncoder(&buf)

			start := xml.StartElement{Name: xml.Name{Local: "w:docDefaults"}}
			if err := tt.doc.MarshalXML(encoder, start); err != nil {
				t.Fatalf("Failed to marshal DocDefault: %v", err)
			}

			if err := encoder.Flush(); err != nil {
				t.Fatalf("Failed to flush encoder: %v", err)
			}

			got := buf.String()
			if got != tt.expected {
				t.Errorf("MarshalXML output:\n%s\nExpected:\n%s", got, tt.expected)
			}
		})
	}
}

func TestRunPropDefault_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		runProp  RunPropDefault
		expected string
	}{
		{
			name: "Marshal RunPropDefault with RunProperty",
			runProp: RunPropDefault{
				RunProp: &RunProperty{},
			},
			expected: `<w:rPrDefault><w:rPr></w:rPr></w:rPrDefault>`,
		},
		{
			name:     "Marshal RunPropDefault with nil RunProperty",
			runProp:  RunPropDefault{},
			expected: `<w:rPrDefault></w:rPrDefault>`,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			encoder := xml.NewEncoder(&buf)

			start := xml.StartElement{Name: xml.Name{Local: "w:rPrDefault"}}
			if err := tt.runProp.MarshalXML(encoder, start); err != nil {
				t.Fatalf("Failed to marshal RunPropDefault: %v", err)
			}

			if err := encoder.Flush(); err != nil {
				t.Fatalf("Failed to flush encoder: %v", err)
			}

			got := buf.String()
			if got != tt.expected {
				t.Errorf("MarshalXML output:\n%s\nExpected:\n%s", got, tt.expected)
			}
		})
	}
}

func TestParaPropDefault_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		paraProp ParaPropDefault
		expected string
	}{
		{
			name: "Marshal ParaPropDefault with ParagraphProp",
			paraProp: ParaPropDefault{
				ParaProp: &ParagraphProp{},
			},
			expected: `<w:pPrDefault><w:pPr></w:pPr></w:pPrDefault>`,
		},
		{
			name:     "Marshal ParaPropDefault with nil ParagraphProp",
			paraProp: ParaPropDefault{},
			expected: `<w:pPrDefault></w:pPrDefault>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			encoder := xml.NewEncoder(&buf)

			start := xml.StartElement{Name: xml.Name{Local: "w:pPrDefault"}}
			if err := tt.paraProp.MarshalXML(encoder, start); err != nil {
				t.Fatalf("Failed to marshal ParaPropDefault: %v", err)
			}

			if err := encoder.Flush(); err != nil {
				t.Fatalf("Failed to flush encoder: %v", err)
			}

			got := buf.String()
			if got != tt.expected {
				t.Errorf("MarshalXML output:\n%s\nExpected:\n%s", got, tt.expected)
			}
		})
	}
}
func TestDocDefault_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		xml      string
		expected DocDefault
	}{
		{
			name: "Unmarshal DocDefault with RunPropDefault and ParaPropDefault",
			xml:  `<w:docDefaults><w:rPrDefault><w:rPr></w:rPr></w:rPrDefault><w:pPrDefault><w:pPr></w:pPr></w:pPrDefault></w:docDefaults>`,
			expected: DocDefault{
				RunProp: &RunPropDefault{
					RunProp: &RunProperty{},
				},
				ParaProp: &ParaPropDefault{
					ParaProp: &ParagraphProp{},
				},
			},
		},
		{
			name:     "Unmarshal DocDefault with empty elements",
			xml:      `<w:docDefaults></w:docDefaults>`,
			expected: DocDefault{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var doc DocDefault
			if err := xml.Unmarshal([]byte(tt.xml), &doc); err != nil {
				t.Fatalf("Failed to unmarshal DocDefault XML: %v", err)
			}

			if !reflect.DeepEqual(doc, tt.expected) {
				t.Errorf("UnmarshalXML result:\n%+v\nExpected:\n%+v", doc, tt.expected)
			}
		})
	}
}

func TestRunPropDefault_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		xml      string
		expected RunPropDefault
	}{
		{
			name: "Unmarshal RunPropDefault with RunProperty",
			xml:  `<w:rPrDefault><w:rPr></w:rPr></w:rPrDefault>`,
			expected: RunPropDefault{
				RunProp: &RunProperty{},
			},
		},
		{
			name:     "Unmarshal RunPropDefault with empty element",
			xml:      `<w:rPrDefault></w:rPrDefault>`,
			expected: RunPropDefault{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var runProp RunPropDefault
			if err := xml.Unmarshal([]byte(tt.xml), &runProp); err != nil {
				t.Fatalf("Failed to unmarshal RunPropDefault XML: %v", err)
			}

			if !reflect.DeepEqual(runProp, tt.expected) {
				t.Errorf("UnmarshalXML result:\n%+v\nExpected:\n%+v", runProp, tt.expected)
			}
		})
	}
}

func TestParaPropDefault_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		xml      string
		expected ParaPropDefault
	}{
		{
			name: "Unmarshal ParaPropDefault with ParagraphProp",
			xml:  `<w:pPrDefault><w:pPr></w:pPr></w:pPrDefault>`,
			expected: ParaPropDefault{
				ParaProp: &ParagraphProp{},
			},
		},
		{
			name:     "Unmarshal ParaPropDefault with empty element",
			xml:      `<w:pPrDefault></w:pPrDefault>`,
			expected: ParaPropDefault{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var paraProp ParaPropDefault
			if err := xml.Unmarshal([]byte(tt.xml), &paraProp); err != nil {
				t.Fatalf("Failed to unmarshal ParaPropDefault XML: %v", err)
			}

			if !reflect.DeepEqual(paraProp, tt.expected) {
				t.Errorf("UnmarshalXML result:\n%+v\nExpected:\n%+v", paraProp, tt.expected)
			}
		})
	}
}
