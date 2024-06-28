package ctypes

import (
	"bytes"
	"encoding/xml"
	"strings"
	"testing"
)

func TestCTString_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    CTString
		expected string
	}{
		{
			name:     "With value",
			input:    CTString{Val: "example"},
			expected: `<w:rStyle w:val="example"></w:rStyle>`,
		},
		{
			name:     "Empty value",
			input:    CTString{Val: ""},
			expected: `<w:rStyle w:val=""></w:rStyle>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:rStyle"}}

			err := tt.input.MarshalXML(encoder, start)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			// Finalize encoding
			encoder.Flush()

			if result.String() != tt.expected {
				t.Errorf("Expected XML:\n%s\nGot:\n%s", tt.expected, result.String())
			}
		})
	}
}

func TestCTString_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected CTString
	}{
		{
			name:     "With value",
			inputXML: `<w:rStyle w:val="example"></w:rStyle>`,
			expected: CTString{Val: "example"},
		},
		{
			name:     "Empty value",
			inputXML: `<w:rStyle w:val=""></w:rStyle>`,
			expected: CTString{Val: ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result CTString

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if result.Val != tt.expected.Val {
				t.Errorf("Expected Val %s but got %s", tt.expected.Val, result.Val)
			}
		})
	}
}

// !--- Tests of CTString ends here ---!

func TestDecimalNum_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    DecimalNum
		expected string
	}{
		{
			name:     "With value",
			input:    DecimalNum{Val: 10},
			expected: `<w:outlineLvl w:val="10"></w:outlineLvl>`,
		},
		{
			name:     "Empty value",
			input:    DecimalNum{Val: -1},
			expected: `<w:outlineLvl w:val="-1"></w:outlineLvl>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:outlineLvl"}}

			err := tt.input.MarshalXML(encoder, start)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			// Finalize encoding
			encoder.Flush()

			if result.String() != tt.expected {
				t.Errorf("Expected XML:\n%s\nGot:\n%s", tt.expected, result.String())
			}
		})
	}
}

func TestDecimalNum_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected DecimalNum
	}{
		{
			name:     "With value",
			inputXML: `<w:outlineLvl w:val="00122"></w:outlineLvl>`,
			expected: DecimalNum{Val: 122},
		},
		{
			name:     "Empty value",
			inputXML: `<w:outlineLvl w:val="+3"></w:outlineLvl>`,
			expected: DecimalNum{Val: 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result DecimalNum

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if result.Val != tt.expected.Val {
				t.Errorf("Expected Val %d but got %d", tt.expected.Val, result.Val)
			}
		})
	}
}

// !--- Tests of DecimalNum ends here ---!

// !--- Tests of Uint64 starts here ---!

func TestUint64Elem_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    Uint64Elem
		expected string
	}{
		{
			name:     "With value",
			input:    Uint64Elem{Val: 10},
			expected: `<w:kern w:val="10"></w:kern>`,
		},
		{
			name:     "Empty value",
			input:    Uint64Elem{Val: 18446744073709551615},
			expected: `<w:kern w:val="18446744073709551615"></w:kern>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:kern"}}

			err := tt.input.MarshalXML(encoder, start)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			// Finalize encoding
			encoder.Flush()

			if result.String() != tt.expected {
				t.Errorf("Expected XML:\n%s\nGot:\n%s", tt.expected, result.String())
			}
		})
	}
}

func TestUint64Elem_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected Uint64Elem
	}{
		{
			name:     "With value",
			inputXML: `<w:kern w:val="00122"></w:kern>`,
			expected: Uint64Elem{Val: 122},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result Uint64Elem

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if result.Val != tt.expected.Val {
				t.Errorf("Expected Val %d but got %d", tt.expected.Val, result.Val)
			}
		})
	}
}

// !--- Tests of Uint64 ends here ---!

func TestGenSingleStrVal_MarshalXML(t *testing.T) {
	tests := []struct {
		name string
		val  string
		want string
	}{
		{"Test1", "Hello", `<GenSingleStrVal w:val="Hello"></GenSingleStrVal>`},
		{"Test2", "World", `<GenSingleStrVal w:val="World"></GenSingleStrVal>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := NewGenSingleStrVal(tt.val)

			var buf bytes.Buffer
			enc := xml.NewEncoder(&buf)
			start := xml.StartElement{Name: xml.Name{Local: "GenSingleStrVal"}}
			if err := gen.MarshalXML(enc, start); err != nil {
				t.Errorf("MarshalXML() error = %v", err)
				return
			}
			if got := buf.String(); got != tt.want {
				t.Errorf("MarshalXML() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenSingleStrVal_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name string
		xml  string
		want string
	}{
		{"Test1", `<GenSingleStrVal w:val="Hello"></GenSingleStrVal>`, "Hello"},
		{"Test2", `<GenSingleStrVal w:val="World"></GenSingleStrVal>`, "World"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gen GenSingleStrVal[string]

			err := xml.Unmarshal([]byte(tt.xml), &gen)
			if err != nil {
				t.Errorf("UnmarshalXML() error = %v", err)
				return
			}

			if gen.Val != tt.want {
				t.Errorf("UnmarshalXML() = %v, want %v", gen.Val, tt.want)
			}
		})
	}
}

// !--- Tests of GenSingleStrVal ends here ---!
