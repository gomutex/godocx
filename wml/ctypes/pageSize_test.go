package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/wml/stypes"
)

func TestPageSize_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    PageSize
		expected string
	}{
		{
			name: "All attributes",
			input: PageSize{
				Width:  uint64Ptr(12240),
				Height: uint64Ptr(15840),
				Orient: stypes.PageOrientLandscape,
				Code:   intPtr(1),
			},
			expected: `<w:pgSz w:w="12240" w:h="15840" w:orient="landscape" w:code="1"></w:pgSz>`,
		},
		{
			name: "Some attributes",
			input: PageSize{
				Width:  uint64Ptr(12240),
				Height: uint64Ptr(15840),
			},
			expected: `<w:pgSz w:w="12240" w:h="15840"></w:pgSz>`,
		},
		{
			name:     "No attributes",
			input:    PageSize{},
			expected: `<w:pgSz></w:pgSz>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:pgSz"}}

			err := tt.input.MarshalXML(encoder, start)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}
			encoder.Flush()

			if result.String() != tt.expected {
				t.Errorf("Expected XML:\n%s\n\nGot:\n%s", tt.expected, result.String())
			}
		})
	}
}

func TestPageSize_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected PageSize
	}{
		{
			name:     "All attributes",
			inputXML: `<w:pgSz w:w="12240" w:h="15840" w:orient="landscape" w:code="1"></w:pgSz>`,
			expected: PageSize{
				Width:  uint64Ptr(12240),
				Height: uint64Ptr(15840),
				Orient: stypes.PageOrientLandscape,
				Code:   intPtr(1),
			},
		},
		{
			name:     "Some attributes",
			inputXML: `<w:pgSz w:w="12240" w:h="15840"></w:pgSz>`,
			expected: PageSize{
				Width:  uint64Ptr(12240),
				Height: uint64Ptr(15840),
			},
		},
		{
			name:     "No attributes",
			inputXML: `<w:pgSz></w:pgSz>`,
			expected: PageSize{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result PageSize

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error during unmarshaling: %v", err)
			}

			comparePageSizes(t, result, tt.expected)
		})
	}
}

func uint64Ptr(i uint64) *uint64 {
	return &i
}

func comparePageSizes(t *testing.T, got, want PageSize) {
	if !compareUint64Ptr(got.Width, want.Width) {
		t.Errorf("Width = %v, want %v", got.Width, want.Width)
	}
	if !compareUint64Ptr(got.Height, want.Height) {
		t.Errorf("Height = %v, want %v", got.Height, want.Height)
	}
	if got.Orient != want.Orient {
		t.Errorf("Orient = %v, want %v", got.Orient, want.Orient)
	}
	if !compareIntPtr(got.Code, want.Code) {
		t.Errorf("Code = %v, want %v", got.Code, want.Code)
	}
}

func compareUint64Ptr(got, want *uint64) bool {
	if got == nil && want == nil {
		return true
	}
	if got == nil || want == nil {
		return false
	}
	return *got == *want
}
