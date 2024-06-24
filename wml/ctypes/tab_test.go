package ctypes

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"

	"github.com/gomutex/godocx/internal"
	"github.com/gomutex/godocx/wml/stypes"
)

func TestTab_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    Tab
		expected string
	}{
		{
			name: "With all attributes",
			input: Tab{
				Val:        stypes.CustTabStopCenter,
				Position:   720,
				LeaderChar: internal.ToPtr(stypes.CustLeadCharDot),
			},
			expected: `<w:tab w:val="center" w:pos="720" w:leader="dot"></w:tab>`,
		},
		{
			name: "Without optional attributes",
			input: Tab{
				Val:      stypes.CustTabStopRight,
				Position: 1440,
			},
			expected: `<w:tab w:val="right" w:pos="1440"></w:tab>`,
		},
		{
			name:     "Empty struct",
			input:    Tab{},
			expected: `<w:tab w:val="" w:pos="0"></w:tab>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:tab"}}

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
func TestTab_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected Tab
	}{
		{
			name:     "With all attributes",
			inputXML: `<w:tab w:val="center" w:pos="720" w:leader="dot"></w:tab>`,
			expected: Tab{
				Val:        stypes.CustTabStopCenter,
				Position:   720,
				LeaderChar: internal.ToPtr(stypes.CustLeadCharDot),
			},
		},
		{
			name:     "Without optional attributes",
			inputXML: `<w:tab w:val="right" w:pos="1440"></w:tab>`,
			expected: Tab{
				Val:      stypes.CustTabStopRight,
				Position: 1440,
			},
		},
		{
			name:     "Empty struct",
			inputXML: `<w:tab></w:tab>`,
			expected: Tab{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result Tab

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			// Compare Val
			if result.Val != tt.expected.Val {
				t.Errorf("Expected Val %s but got %v", tt.expected.Val, result.Val)
			}

			// Compare Position
			if result.Position != tt.expected.Position {
				t.Errorf("Expected Position %d but got %v", tt.expected.Position, result.Position)
			}

			// Compare LeaderChar
			if tt.expected.LeaderChar != nil {
				if result.LeaderChar == nil || *result.LeaderChar != *tt.expected.LeaderChar {
					t.Errorf("Expected LeaderChar %s but got %v", *tt.expected.LeaderChar, *result.LeaderChar)
				}
			} else if result.LeaderChar != nil {
				t.Errorf("Expected LeaderChar nil but got %v", *result.LeaderChar)
			}
		})
	}
}
func TestTabs_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		tabs     Tabs
		expected string
	}{
		{
			name:     "Empty Tabs",
			tabs:     Tabs{},
			expected: ``,
		},
		{
			name: "Tabs with Multiple Tab elements",
			tabs: Tabs{
				Tab: []Tab{
					{Val: stypes.CustTabStopCenter, Position: 100, LeaderChar: internal.ToPtr(stypes.CustLeadCharDot)},
					{Val: stypes.CustTabStopLeft, Position: 200, LeaderChar: internal.ToPtr(stypes.CustLeadCharHyphen)},
				},
			},
			expected: `<w:tabs><w:tab w:val="center" w:pos="100" w:leader="dot"></w:tab><w:tab w:val="left" w:pos="200" w:leader="hyphen"></w:tab></w:tabs>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			xmlBytes, err := xml.Marshal(tt.tabs)
			if err != nil {
				t.Fatalf("Unexpected error during Marshal: %v", err)
			}

			actual := strings.TrimSpace(string(xmlBytes))
			if actual != tt.expected {
				t.Fatalf("Unexpected XML output: expected '%s' but got '%s'", tt.expected, actual)
			}
		})
	}
}
func TestTabs_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		xmlInput string
		expected Tabs
	}{
		{
			name:     "Empty Tabs",
			xmlInput: `<w:tabs></w:tabs>`,
			expected: Tabs{},
		},
		{
			name: "Tabs with Multiple Tab elements",
			xmlInput: `<tabs>
                <w:tab val="center" pos="100" leader="dot"/>
                <w:tab val="left" pos="200" leader="hyphen"/>
            </tabs>`,
			expected: Tabs{
				Tab: []Tab{
					{Val: stypes.CustTabStopCenter, Position: 100, LeaderChar: internal.ToPtr(stypes.CustLeadCharDot)},
					{Val: stypes.CustTabStopLeft, Position: 200, LeaderChar: internal.ToPtr(stypes.CustLeadCharHyphen)},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tabs Tabs
			err := xml.Unmarshal([]byte(tt.xmlInput), &tabs)
			if err != nil {
				t.Fatalf("Unexpected error during Unmarshal: %v", err)
			}

			// Compare individual fields of Tabs struct
			if !reflect.DeepEqual(tabs, tt.expected) {
				t.Errorf("Unmarshaled Tabs struct does not match expected:\nExpected: %+v\nActual:   %+v", tt.expected, tabs)
			}
		})
	}
}
