package ctypes_test

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/gomutex/godocx/wml/ctypes"
	"github.com/gomutex/godocx/wml/stypes"
)

func TestTableStyleProp_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		prop     *ctypes.TableStyleProp
		expected string
	}{
		{
			name: "all properties set",
			prop: &ctypes.TableStyleProp{
				ParaProp:  &ctypes.ParagraphProp{},
				RunProp:   &ctypes.RunProperty{},
				TableProp: &ctypes.TableProp{},
				RowProp:   &ctypes.RowProperty{},
				CellProp:  &ctypes.CellProperty{},
				Type:      stypes.TblStyleOverrideType("testType"),
			},
			expected: `<w:tblStylePr w:type="testType"><w:pPr></w:pPr><w:rPr></w:rPr><w:tblPr></w:tblPr><w:trPr></w:trPr><w:tcPr></w:tcPr></w:tblStylePr>`,
		},
		{
			name: "some properties nil",
			prop: &ctypes.TableStyleProp{
				ParaProp:  nil,
				RunProp:   &ctypes.RunProperty{},
				TableProp: nil,
				RowProp:   &ctypes.RowProperty{},
				CellProp:  nil,
				Type:      stypes.TblStyleOverrideType("testType"),
			},
			expected: `<w:tblStylePr w:type="testType"><w:rPr></w:rPr><w:trPr></w:trPr></w:tblStylePr>`,
		},
		{
			name: "all properties nil",
			prop: &ctypes.TableStyleProp{
				ParaProp:  nil,
				RunProp:   nil,
				TableProp: nil,
				RowProp:   nil,
				CellProp:  nil,
				Type:      stypes.TblStyleOverrideType("testType"),
			},
			expected: `<w:tblStylePr w:type="testType"></w:tblStylePr>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(tt.prop)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got := string(output); got != tt.expected {
				t.Errorf("expected %s, but got %s", tt.expected, got)
			}
		})
	}
}

func TestTableStyleProp_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		xmlInput string
		expected *ctypes.TableStyleProp
	}{
		{
			name: "all properties set",
			xmlInput: `<w:tblStylePr w:type="neCell">
							<w:pPr></w:pPr>
							<w:rPr></w:rPr>
							<w:tblPr></w:tblPr>
							<w:trPr></w:trPr>
							<w:tcPr></w:tcPr>
						</w:tblStylePr>`,
			expected: &ctypes.TableStyleProp{
				ParaProp:  &ctypes.ParagraphProp{},
				RunProp:   &ctypes.RunProperty{},
				TableProp: &ctypes.TableProp{},
				RowProp:   &ctypes.RowProperty{},
				CellProp:  &ctypes.CellProperty{},
				Type:      stypes.TblStyleOverrideNeCell,
			},
		},
		{
			name: "some properties nil",
			xmlInput: `<w:tblStylePr w:type="neCell">
							<w:rPr></w:rPr>
							<w:trPr></w:trPr>
						</w:tblStylePr>`,
			expected: &ctypes.TableStyleProp{
				RunProp: &ctypes.RunProperty{},
				RowProp: &ctypes.RowProperty{},
				Type:    stypes.TblStyleOverrideNeCell,
			},
		},
		{
			name:     "all properties nil",
			xmlInput: `<w:tblStylePr w:type="neCell"></w:tblStylePr>`,
			expected: &ctypes.TableStyleProp{
				Type: stypes.TblStyleOverrideNeCell,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var prop ctypes.TableStyleProp
			err := xml.Unmarshal([]byte(tt.xmlInput), &prop)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(&prop, tt.expected) {
				t.Errorf("expected %+v, but got %+v", tt.expected, &prop)
			}
		})
	}
}
