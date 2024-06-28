package ctypes

import (
	"bytes"
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/internal"
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

func TestEmpty_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    Empty
		expected string
	}{
		{
			name:     "Empty element",
			input:    Empty{},
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

func TestEmpty_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
	}{
		{
			name:     "Empty element",
			inputXML: `<w:tab></w:tab>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result Empty

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}
		})
	}
}

// !--- Tests of Empty ends here ---!

func TestMarkup_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    Markup
		expected string
	}{
		{
			name:     "With ID",
			input:    Markup{ID: 42},
			expected: `<Markup w:id="42"></Markup>`,
		},
		{
			name:     "Zero ID",
			input:    Markup{ID: 0},
			expected: `<Markup w:id="0"></Markup>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "Markup"}}

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

func TestMarkup_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected Markup
	}{
		{
			name:     "With ID",
			inputXML: `<Markup w:id="42"></Markup>`,
			expected: Markup{ID: 42},
		},
		{
			name:     "Zero ID",
			inputXML: `<Markup w:id="0"></Markup>`,
			expected: Markup{ID: 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result Markup

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if result.ID != tt.expected.ID {
				t.Errorf("Expected ID %d but got %d", tt.expected.ID, result.ID)
			}
		})
	}
}

// !--- Tests of Markup ends here ---!

// Test function for MarshalXML method
func TestMarshalXML(t *testing.T) {
	tests := []struct {
		name        string
		instance    GenOptStrVal[string]
		expectedXML string
	}{
		{
			name:        "WithValue",
			instance:    GenOptStrVal[string]{Val: internal.ToPtr("test")},
			expectedXML: `<element w:val="test"></element>`,
		},
		{
			name:        "WithNilValue",
			instance:    GenOptStrVal[string]{Val: nil},
			expectedXML: `<element></element>`,
		},
		{
			name:        "EmptyValue",
			instance:    GenOptStrVal[string]{Val: internal.ToPtr("")},
			expectedXML: `<element w:val=""></element>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			encoder := xml.NewEncoder(&buf)

			start := xml.StartElement{Name: xml.Name{Local: "element"}}
			err := tt.instance.MarshalXML(encoder, start)
			if err != nil {
				t.Errorf("MarshalXML error: %v", err)
				return
			}

			// Convert expected and actual values to pointers for comparison
			expectedPtr := internal.ToPtr(tt.expectedXML)
			actualPtr := internal.ToPtr(buf.String())

			// Compare using ComparePtr function
			err = internal.ComparePtr("XML Output", expectedPtr, actualPtr)
			if err != nil {
				t.Errorf("Comparison error: %v", err)
			}
		})
	}
}

func TestUnmarshalXML(t *testing.T) {
	tests := []struct {
		name           string
		xmlString      string
		expectedStruct GenOptStrVal[string]
	}{
		{
			name:      "WithValAttribute",
			xmlString: `<element w:val="test"></element>`,
			expectedStruct: GenOptStrVal[string]{
				Val: internal.ToPtr("test"),
			},
		},
		{
			name:      "WithoutValAttribute",
			xmlString: `<element></element>`,
			expectedStruct: GenOptStrVal[string]{
				Val: nil,
			},
		},
		{
			name:      "EmptyValAttribute",
			xmlString: `<element w:val=""></element>`,
			expectedStruct: GenOptStrVal[string]{
				Val: internal.ToPtr(""),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new XML decoder for the test XML string
			decoder := xml.NewDecoder(strings.NewReader(tt.xmlString))

			// Create an empty instance of GenOptStrVal to store the decoded result
			var result GenOptStrVal[string]

			// Decode XML into result
			err := decoder.Decode(&result)
			if err != nil {
				t.Errorf("Decode error: %v", err)
				return
			}

			// Compare the decoded result with the expectedStruct using ComparePtr
			err = internal.ComparePtr("Val", tt.expectedStruct.Val, result.Val)
			if err != nil {
				t.Errorf("Comparison error: %v", err)
			}
		})
	}
}

// !--- Tests of GenOptStrVal[string] ends here ---!
