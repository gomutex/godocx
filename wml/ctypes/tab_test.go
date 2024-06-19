package ctypes

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"

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
				Val:        CTabStopPtr(stypes.CustTabStopCenter),
				Position:   IntPtr(720),
				LeaderChar: LeadCharPtr(stypes.CustLeadCharDot),
			},
			expected: `<w:tab val="center" pos="720" leader="dot"></w:tab>`,
		},
		{
			name: "Without optional attributes",
			input: Tab{
				Val:      CTabStopPtr(stypes.CustTabStopRight),
				Position: IntPtr(1440),
			},
			expected: `<w:tab val="right" pos="1440"></w:tab>`,
		},
		{
			name:     "Empty struct",
			input:    Tab{},
			expected: `<w:tab></w:tab>`,
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
			inputXML: `<w:tab val="center" pos="720" leader="dot"></w:tab>`,
			expected: Tab{
				Val:        CTabStopPtr(stypes.CustTabStopCenter),
				Position:   IntPtr(720),
				LeaderChar: LeadCharPtr(stypes.CustLeadCharDot),
			},
		},
		{
			name:     "Without optional attributes",
			inputXML: `<w:tab val="right" pos="1440"></w:tab>`,
			expected: Tab{
				Val:      CTabStopPtr(stypes.CustTabStopRight),
				Position: IntPtr(1440),
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
			if tt.expected.Val != nil {
				if result.Val == nil || *result.Val != *tt.expected.Val {
					t.Errorf("Expected Val %s but got %v", *tt.expected.Val, *result.Val)
				}
			} else if result.Val != nil {
				t.Errorf("Expected Val nil but got %v", *result.Val)
			}

			// Compare Position
			if tt.expected.Position != nil {
				if result.Position == nil || *result.Position != *tt.expected.Position {
					t.Errorf("Expected Position %d but got %v", *tt.expected.Position, *result.Position)
				}
			} else if result.Position != nil {
				t.Errorf("Expected Position nil but got %v", *result.Position)
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
					{Val: CTabStopPtr(stypes.CustTabStopCenter), Position: IntPtr(100), LeaderChar: LeadCharPtr(stypes.CustLeadCharDot)},
					{Val: CTabStopPtr(stypes.CustTabStopLeft), Position: IntPtr(200), LeaderChar: LeadCharPtr(stypes.CustLeadCharHyphen)},
				},
			},
			expected: `<w:tabs><w:tab val="center" pos="100" leader="dot"></w:tab><w:tab val="left" pos="200" leader="hyphen"></w:tab></w:tabs>`,
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
					{Val: CTabStopPtr(stypes.CustTabStopCenter), Position: IntPtr(100), LeaderChar: LeadCharPtr(stypes.CustLeadCharDot)},
					{Val: CTabStopPtr(stypes.CustTabStopLeft), Position: IntPtr(200), LeaderChar: LeadCharPtr(stypes.CustLeadCharHyphen)},
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

func IntPtr(v int) *int {
	return &v
}

func LeadCharPtr(lc stypes.CustLeadChar) *stypes.CustLeadChar {
	return &lc
}

func CTabStopPtr(v stypes.CustTabStop) *stypes.CustTabStop {
	return &v
}
