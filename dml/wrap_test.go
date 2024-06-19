package dml

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"

	"github.com/gomutex/godocx/dml/dmlct"
	"github.com/gomutex/godocx/dml/dmlst"
	"github.com/gomutex/godocx/internal"
)

func TestWrapSquare_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    WrapSquare
		expected string
	}{
		{
			name: "With all attributes and EffectExtent",
			input: WrapSquare{
				WrapText: dmlst.WrapTextLeft,
				DistT:    internal.ToPtr(uint(10)),
				DistB:    internal.ToPtr(uint(15)),
				DistL:    internal.ToPtr(uint(5)),
				DistR:    internal.ToPtr(uint(8)),
				EffectExtent: &EffectExtent{
					LeftEdge:   1,
					TopEdge:    2,
					RightEdge:  3,
					BottomEdge: 4,
				},
			},
			expected: `<wp:wrapSquare wrapText="left" distT="10" distB="15" distL="5" distR="8"><wp:effectExtent l="1" t="2" r="3" b="4"></wp:effectExtent></wp:wrapSquare>`,
		},
		{
			name: "Only WrapText attribute",
			input: WrapSquare{
				WrapText: dmlst.WrapTextLeft,
			},
			expected: `<wp:wrapSquare wrapText="left"></wp:wrapSquare>`,
		},
		{
			name: "With DistT and DistR attributes",
			input: WrapSquare{
				WrapText: dmlst.WrapTextLeft,
				DistT:    internal.ToPtr(uint(10)),
				DistR:    internal.ToPtr(uint(8)),
			},
			expected: `<wp:wrapSquare wrapText="left" distT="10" distR="8"></wp:wrapSquare>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "wrapSquare"}}

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

func TestWrapSquare_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected WrapSquare
	}{
		{
			name:     "With all attributes and EffectExtent",
			inputXML: `<wp:wrapSquare wrapText="left" distT="10" distB="15" distL="5" distR="8"><wp:effectExtent l="1" t="2" r="3" b="4"></wp:effectExtent></wp:wrapSquare>`,
			expected: WrapSquare{
				WrapText: dmlst.WrapTextLeft,
				DistT:    internal.ToPtr(uint(10)),
				DistB:    internal.ToPtr(uint(15)),
				DistL:    internal.ToPtr(uint(5)),
				DistR:    internal.ToPtr(uint(8)),
				EffectExtent: &EffectExtent{
					LeftEdge:   1,
					TopEdge:    2,
					RightEdge:  3,
					BottomEdge: 4,
				},
			},
		},
		{
			name:     "Only WrapText attribute",
			inputXML: `<wp:wrapSquare wrapText="left"></wp:wrapSquare>`,
			expected: WrapSquare{
				WrapText: dmlst.WrapTextLeft,
			},
		},
		{
			name:     "With DistT and DistR attributes",
			inputXML: `<wp:wrapSquare wrapText="left" distT="10" distR="8"></wp:wrapSquare>`,
			expected: WrapSquare{
				WrapText: dmlst.WrapTextLeft,
				DistT:    internal.ToPtr(uint(10)),
				DistR:    internal.ToPtr(uint(8)),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result WrapSquare

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if tt.expected.WrapText != result.WrapText {
				t.Errorf("WrapText expected %v, but got %v", tt.expected.WrapText, result.WrapText)
			}

			if err := internal.ComparePtr("DistT", tt.expected.DistT, result.DistT); err != nil {
				t.Errorf("DistT comparison error: %v", err)
			}
			if err := internal.ComparePtr("DistB", tt.expected.DistB, result.DistB); err != nil {
				t.Errorf("DistB comparison error: %v", err)
			}
			if err := internal.ComparePtr("DistL", tt.expected.DistL, result.DistL); err != nil {
				t.Errorf("DistL comparison error: %v", err)
			}
			if err := internal.ComparePtr("DistR", tt.expected.DistR, result.DistR); err != nil {
				t.Errorf("DistR comparison error: %v", err)
			}
			if err := internal.ComparePtr("EffectExtent", tt.expected.EffectExtent, result.EffectExtent); err != nil {
				t.Errorf("EffectExtent comparison error: %v", err)
			}
		})
	}
}

// Tests for WrapSquare ends

// Tests for WrapPolygon starts
func TestWrapPolygon_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    WrapPolygon
		expected string
	}{
		{
			name: "With all attributes and edited true",
			input: WrapPolygon{
				Start:  dmlct.NewPoint2D(1, 2),
				LineTo: []dmlct.Point2D{dmlct.NewPoint2D(3, 4), dmlct.NewPoint2D(5, 6)},
				Edited: internal.ToPtr(true),
			},
			expected: `<wp:wrapPolygon xmlns="wp" edited="true"><wp:start x="1" y="2"></wp:start><wp:lineTo x="3" y="4"></wp:lineTo><wp:lineTo x="5" y="6"></wp:lineTo></wp:wrapPolygon>`,
		},
		{
			name: "With no edited attribute",
			input: WrapPolygon{
				Start:  dmlct.NewPoint2D(1, 2),
				LineTo: []dmlct.Point2D{dmlct.NewPoint2D(3, 4), dmlct.NewPoint2D(5, 6), dmlct.NewPoint2D(7, 8)},
				Edited: nil,
			},
			expected: `<wp:wrapPolygon xmlns="wp"><wp:start x="1" y="2"></wp:start><wp:lineTo x="3" y="4"></wp:lineTo><wp:lineTo x="5" y="6"></wp:lineTo><wp:lineTo x="7" y="8"></wp:lineTo></wp:wrapPolygon>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Space: "wp", Local: "wrapPolygon"}}

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

func TestWrapPolygon_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected WrapPolygon
	}{
		{
			name:     "With all attributes and edited true",
			inputXML: `<wp:wrapPolygon xmlns:wp="wp" edited="true"><wp:start x="1" y="2"></wp:start><wp:lineTo x="3" y="4"></wp:lineTo><wp:lineTo x="5" y="6"></wp:lineTo></wp:wrapPolygon>`,
			expected: WrapPolygon{
				Start:  dmlct.Point2D{XAxis: 1, YAxis: 2},
				LineTo: []dmlct.Point2D{{XAxis: 3, YAxis: 4}, {XAxis: 5, YAxis: 6}},
				Edited: internal.ToPtr(true),
			},
		},
		{
			name:     "With no edited attribute",
			inputXML: `<wp:wrapPolygon xmlns:wp="wp"><wp:start x="1" y="2"></wp:start><wp:lineTo x="3" y="4"></wp:lineTo><wp:lineTo x="5" y="6"></wp:lineTo><wp:lineTo x="7" y="8"></wp:lineTo></wp:wrapPolygon>`,
			expected: WrapPolygon{
				Start:  dmlct.Point2D{XAxis: 1, YAxis: 2},
				LineTo: []dmlct.Point2D{{XAxis: 3, YAxis: 4}, {XAxis: 5, YAxis: 6}, {XAxis: 7, YAxis: 8}},
				Edited: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result WrapPolygon

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if len(tt.expected.LineTo) != len(result.LineTo) {
				t.Errorf("LineTo length expected %d, but got %d", len(tt.expected.LineTo), len(result.LineTo))
			} else {
				for i := range tt.expected.LineTo {
					if tt.expected.LineTo[i] != result.LineTo[i] {
						t.Errorf("LineTo[%d] expected %+v, but got %+v", i, tt.expected.LineTo[i], result.LineTo[i])
					}
				}
			}

			if tt.expected.Start != result.Start {
				t.Errorf("Start expected %+v, but got %+v", tt.expected.Start, result.Start)
			}

			if err := internal.ComparePtr("Edited", tt.expected.Edited, result.Edited); err != nil {
				t.Errorf("Edited comparison error: %v", err)
			}
		})
	}
}

func TestWrapTight_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    WrapTight
		expected string
	}{
		{
			name: "With all attributes and DistL, DistR",
			input: WrapTight{
				WrapPolygon: WrapPolygon{
					Start:  dmlct.Point2D{XAxis: 1, YAxis: 2},
					LineTo: []dmlct.Point2D{{XAxis: 3, YAxis: 4}, {XAxis: 5, YAxis: 6}},
					Edited: internal.ToPtr(true),
				},
				WrapText: dmlst.WrapTextRight,
				DistL:    internal.ToPtr(uint(10)),
				DistR:    internal.ToPtr(uint(5)),
			},
			expected: `<wp:wrapTight wrapText="right" distL="10" distR="5"><wp:wrapPolygon edited="true"><wp:start x="1" y="2"></wp:start><wp:lineTo x="3" y="4"></wp:lineTo><wp:lineTo x="5" y="6"></wp:lineTo></wp:wrapPolygon></wp:wrapTight>`,
		},
		{
			name: "With no DistL and DistR",
			input: WrapTight{
				WrapPolygon: WrapPolygon{
					Start:  dmlct.Point2D{XAxis: 1, YAxis: 2},
					LineTo: []dmlct.Point2D{{XAxis: 3, YAxis: 4}, {XAxis: 5, YAxis: 6}},
					Edited: internal.ToPtr(false),
				},
				WrapText: dmlst.WrapTextRight,
				DistL:    nil,
				DistR:    nil,
			},
			expected: `<wp:wrapTight wrapText="right"><wp:wrapPolygon edited="false"><wp:start x="1" y="2"></wp:start><wp:lineTo x="3" y="4"></wp:lineTo><wp:lineTo x="5" y="6"></wp:lineTo></wp:wrapPolygon></wp:wrapTight>`,
		},
		{
			name: "With minimal attributes",
			input: WrapTight{
				WrapPolygon: WrapPolygon{
					Start:  dmlct.Point2D{XAxis: 1, YAxis: 2},
					LineTo: []dmlct.Point2D{{XAxis: 3, YAxis: 4}},
					Edited: nil,
				},
				WrapText: dmlst.WrapTextRight,
				DistL:    nil,
				DistR:    nil,
			},
			expected: `<wp:wrapTight wrapText="right"><wp:wrapPolygon><wp:start x="1" y="2"></wp:start><wp:lineTo x="3" y="4"></wp:lineTo></wp:wrapPolygon></wp:wrapTight>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{}

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

func compareWrapTight(a, b WrapTight) bool {
	if a.WrapText != b.WrapText {
		return false
	}

	if a.DistL == nil || b.DistL == nil {
		if a.DistL != b.DistL {
			return false
		}
	} else if *a.DistL != *b.DistL {
		return false
	}

	if a.DistR == nil || b.DistR == nil {
		if a.DistR != b.DistR {
			return false
		}
	} else if *a.DistR != *b.DistR {
		return false
	}

	if !reflect.DeepEqual(a.WrapPolygon, b.WrapPolygon) {
		return false
	}

	return true
}

func TestWrapTight_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name        string
		inputXML    string
		expected    WrapTight
		expectedErr bool
	}{
		{
			name: "With all attributes and DistL, DistR",
			inputXML: `<wp:wrapTight xmlns:wp="wp" wrapText="right" distL="10" distR="5">
						  <wp:wrapPolygon edited="true">
						    <wp:start x="1" y="2"></wp:start>
						    <wp:lineTo x="3" y="4"></wp:lineTo>
						    <wp:lineTo x="5" y="6"></wp:lineTo>
						  </wp:wrapPolygon>
						</wp:wrapTight>`,
			expected: WrapTight{
				WrapPolygon: WrapPolygon{
					Start:  dmlct.Point2D{XAxis: 1, YAxis: 2},
					LineTo: []dmlct.Point2D{{XAxis: 3, YAxis: 4}, {XAxis: 5, YAxis: 6}},
					Edited: internal.ToPtr(true),
				},
				WrapText: dmlst.WrapTextRight,
				DistL:    internal.ToPtr(uint(10)),
				DistR:    internal.ToPtr(uint(5)),
			},
			expectedErr: false,
		},
		{
			name: "With minimal attributes",
			inputXML: `<wp:wrapTight xmlns:wp="wp" wrapText="right">
						  <wp:wrapPolygon>
						    <wp:start x="1" y="2"></wp:start>
						    <wp:lineTo x="3" y="4"></wp:lineTo>
						  </wp:wrapPolygon>
						</wp:wrapTight>`,
			expected: WrapTight{
				WrapPolygon: WrapPolygon{
					Start:  dmlct.Point2D{XAxis: 1, YAxis: 2},
					LineTo: []dmlct.Point2D{{XAxis: 3, YAxis: 4}},
					Edited: nil,
				},
				WrapText: dmlst.WrapTextRight,
				DistL:    nil,
				DistR:    nil,
			},
			expectedErr: false,
		},
		{
			name: "With no WrapPolygon",
			inputXML: `<wp:wrapTight xmlns:wp="wp" wrapText="right">
						</wp:wrapTight>`,
			expected: WrapTight{
				WrapPolygon: WrapPolygon{},
				WrapText:    dmlst.WrapTextRight,
				DistL:       nil,
				DistR:       nil,
			},
			expectedErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result WrapTight

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if (err != nil) != tt.expectedErr {
				t.Fatalf("Error unexpected during unmarshaling XML: %v", err)
			}

			if !compareWrapTight(result, tt.expected) {
				t.Errorf("Expected result:\n%+v\nGot:\n%+v", tt.expected, result)
			}
		})
	}
}

func TestWrapThrough_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    WrapThrough
		expected string
	}{
		{
			name: "With all attributes and DistL, DistR",
			input: WrapThrough{
				WrapPolygon: WrapPolygon{
					Start:  dmlct.Point2D{XAxis: 1, YAxis: 2},
					LineTo: []dmlct.Point2D{{XAxis: 3, YAxis: 4}, {XAxis: 5, YAxis: 6}},
					Edited: internal.ToPtr(true),
				},
				WrapText: dmlst.WrapTextRight,
				DistL:    internal.ToPtr(uint(10)),
				DistR:    internal.ToPtr(uint(5)),
			},
			expected: `<wp:wrapThrough wrapText="right" distL="10" distR="5"><wp:wrapPolygon edited="true"><wp:start x="1" y="2"></wp:start><wp:lineTo x="3" y="4"></wp:lineTo><wp:lineTo x="5" y="6"></wp:lineTo></wp:wrapPolygon></wp:wrapThrough>`,
		},
		{
			name: "With no DistL and DistR",
			input: WrapThrough{
				WrapPolygon: WrapPolygon{
					Start:  dmlct.Point2D{XAxis: 1, YAxis: 2},
					LineTo: []dmlct.Point2D{{XAxis: 3, YAxis: 4}, {XAxis: 5, YAxis: 6}},
					Edited: internal.ToPtr(false),
				},
				WrapText: dmlst.WrapTextRight,
				DistL:    nil,
				DistR:    nil,
			},
			expected: `<wp:wrapThrough wrapText="right"><wp:wrapPolygon edited="false"><wp:start x="1" y="2"></wp:start><wp:lineTo x="3" y="4"></wp:lineTo><wp:lineTo x="5" y="6"></wp:lineTo></wp:wrapPolygon></wp:wrapThrough>`,
		},
		{
			name: "With minimal attributes",
			input: WrapThrough{
				WrapPolygon: WrapPolygon{
					Start:  dmlct.Point2D{XAxis: 1, YAxis: 2},
					LineTo: []dmlct.Point2D{{XAxis: 3, YAxis: 4}},
					Edited: nil,
				},
				WrapText: dmlst.WrapTextRight,
				DistL:    nil,
				DistR:    nil,
			},
			expected: `<wp:wrapThrough wrapText="right"><wp:wrapPolygon><wp:start x="1" y="2"></wp:start><wp:lineTo x="3" y="4"></wp:lineTo></wp:wrapPolygon></wp:wrapThrough>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{}

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
func compareWrapThrough(a, b WrapThrough) bool {
	if a.WrapText != b.WrapText {
		return false
	}

	if a.DistL == nil || b.DistL == nil {
		if a.DistL != b.DistL {
			return false
		}
	} else if *a.DistL != *b.DistL {
		return false
	}

	if a.DistR == nil || b.DistR == nil {
		if a.DistR != b.DistR {
			return false
		}
	} else if *a.DistR != *b.DistR {
		return false
	}

	if !reflect.DeepEqual(a.WrapPolygon, b.WrapPolygon) {
		return false
	}

	return true
}

func TestWrapThrough_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name        string
		inputXML    string
		expected    WrapThrough
		expectedErr bool
	}{
		{
			name: "With all attributes and DistL, DistR",
			inputXML: `<wp:wrapThrough xmlns:wp="wp" wrapText="right" distL="10" distR="5">
						  <wp:wrapPolygon edited="true">
						    <wp:start x="1" y="2"></wp:start>
						    <wp:lineTo x="3" y="4"></wp:lineTo>
						    <wp:lineTo x="5" y="6"></wp:lineTo>
						  </wp:wrapPolygon>
						</wp:wrapThrough>`,
			expected: WrapThrough{
				WrapPolygon: WrapPolygon{
					Start:  dmlct.Point2D{XAxis: 1, YAxis: 2},
					LineTo: []dmlct.Point2D{{XAxis: 3, YAxis: 4}, {XAxis: 5, YAxis: 6}},
					Edited: internal.ToPtr(true),
				},
				WrapText: dmlst.WrapTextRight,
				DistL:    internal.ToPtr(uint(10)),
				DistR:    internal.ToPtr(uint(5)),
			},
			expectedErr: false,
		},
		{
			name: "With minimal attributes",
			inputXML: `<wp:wrapThrough xmlns:wp="wp" wrapText="right">
						  <wp:wrapPolygon>
						    <wp:start x="1" y="2"></wp:start>
						    <wp:lineTo x="3" y="4"></wp:lineTo>
						  </wp:wrapPolygon>
						</wp:wrapThrough>`,
			expected: WrapThrough{
				WrapPolygon: WrapPolygon{
					Start:  dmlct.Point2D{XAxis: 1, YAxis: 2},
					LineTo: []dmlct.Point2D{{XAxis: 3, YAxis: 4}},
					Edited: nil,
				},
				WrapText: dmlst.WrapTextRight,
				DistL:    nil,
				DistR:    nil,
			},
			expectedErr: false,
		},
		{
			name: "With no WrapPolygon",
			inputXML: `<wp:wrapThrough xmlns:wp="wp" wrapText="right">
						</wp:wrapThrough>`,
			expected: WrapThrough{
				WrapPolygon: WrapPolygon{},
				WrapText:    dmlst.WrapTextRight,
				DistL:       nil,
				DistR:       nil,
			},
			expectedErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result WrapThrough

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if (err != nil) != tt.expectedErr {
				t.Fatalf("Error unexpected during unmarshaling XML: %v", err)
			}

			if !compareWrapThrough(result, tt.expected) {
				t.Errorf("Expected result:\n%+v\nGot:\n%+v", tt.expected, result)
			}
		})
	}
}

func TestWrapTopBtm_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    WrapTopBtm
		expected string
	}{
		{
			name: "With all attributes and EffectExtent",
			input: WrapTopBtm{
				DistT: internal.ToPtr(uint(10)),
				DistB: internal.ToPtr(uint(15)),
				EffectExtent: &EffectExtent{
					LeftEdge:   1,
					TopEdge:    2,
					RightEdge:  3,
					BottomEdge: 4,
				},
			},
			expected: `<wp:wrapTopAndBottom distT="10" distB="15"><wp:effectExtent l="1" t="2" r="3" b="4"></wp:effectExtent></wp:wrapTopAndBottom>`,
		},
		{
			name: "With DistT and EffectExtent",
			input: WrapTopBtm{
				DistT: internal.ToPtr(uint(10)),
				EffectExtent: &EffectExtent{
					LeftEdge:   1,
					TopEdge:    2,
					RightEdge:  3,
					BottomEdge: 4,
				},
			},
			expected: `<wp:wrapTopAndBottom distT="10"><wp:effectExtent l="1" t="2" r="3" b="4"></wp:effectExtent></wp:wrapTopAndBottom>`,
		},
		{
			name: "With DistB only",
			input: WrapTopBtm{
				DistB: internal.ToPtr(uint(15)),
			},
			expected: `<wp:wrapTopAndBottom distB="15"></wp:wrapTopAndBottom>`,
		},
		{
			name:     "With no attributes",
			input:    WrapTopBtm{},
			expected: `<wp:wrapTopAndBottom></wp:wrapTopAndBottom>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{}

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

// Custom comparison function for WrapTopBtm
func compareWrapTopBtm(a, b WrapTopBtm) bool {
	if a.DistT == nil || b.DistT == nil {
		if a.DistT != b.DistT {
			return false
		}
	} else if *a.DistT != *b.DistT {
		return false
	}

	if a.DistB == nil || b.DistB == nil {
		if a.DistB != b.DistB {
			return false
		}
	} else if *a.DistB != *b.DistB {
		return false
	}

	if a.EffectExtent == nil || b.EffectExtent == nil {
		if a.EffectExtent != b.EffectExtent {
			return false
		}
	} else if !reflect.DeepEqual(*a.EffectExtent, *b.EffectExtent) {
		return false
	}

	return true
}

func TestWrapTopBtm_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		xmlInput string
		expected WrapTopBtm
	}{
		{
			name:     "With all attributes and EffectExtent",
			xmlInput: `<wp:wrapTopAndBottom distT="10" distB="15"><wp:effectExtent l="1" t="2" r="3" b="4"></wp:effectExtent></wp:wrapTopAndBottom>`,
			expected: WrapTopBtm{
				DistT: internal.ToPtr(uint(10)),
				DistB: internal.ToPtr(uint(15)),
				EffectExtent: &EffectExtent{
					LeftEdge:   1,
					TopEdge:    2,
					RightEdge:  3,
					BottomEdge: 4,
				},
			},
		},
		{
			name:     "With DistT and EffectExtent",
			xmlInput: `<wp:wrapTopAndBottom distT="10"><wp:effectExtent l="1" t="2" r="3" b="4"></wp:effectExtent></wp:wrapTopAndBottom>`,
			expected: WrapTopBtm{
				DistT: internal.ToPtr(uint(10)),
				EffectExtent: &EffectExtent{
					LeftEdge:   1,
					TopEdge:    2,
					RightEdge:  3,
					BottomEdge: 4,
				},
			},
		},
		{
			name:     "With DistB only",
			xmlInput: `<wp:wrapTopAndBottom distB="15"></wp:wrapTopAndBottom>`,
			expected: WrapTopBtm{
				DistB: internal.ToPtr(uint(15)),
			},
		},
		{
			name:     "With no attributes",
			xmlInput: `<wp:wrapTopAndBottom></wp:wrapTopAndBottom>`,
			expected: WrapTopBtm{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var w WrapTopBtm
			err := xml.Unmarshal([]byte(tt.xmlInput), &w)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if !compareWrapTopBtm(w, tt.expected) {
				t.Errorf("Expected:\n%+v\nGot:\n%+v", tt.expected, w)
			}
		})
	}
}
