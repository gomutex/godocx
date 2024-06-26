package docxrun

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"

	"github.com/gomutex/godocx/wml/ctypes"
)

func optBoolElemPtr(value ctypes.OnOff) *ctypes.OnOff {
	return &value
}

func singleUint64ValPtr(value ctypes.Uint64Elem) *ctypes.Uint64Elem {
	return &value
}

func singleIntValPtr(value ctypes.DecimalNum) *ctypes.DecimalNum {
	return &value
}

func singleStrValPtr(value ctypes.CTString) *ctypes.CTString {
	return &value
}

func TestRunProperty_MarshalXML(t *testing.T) {
	// trueOptBool := types.NewOptBool(true)
	tests := []struct {
		name     string
		prop     RunProperty
		expected string
	}{
		{
			name: "All attributes set",
			prop: RunProperty{
				Bold:         optBoolElemPtr(ctypes.OnOff{}),
				BoldCS:       optBoolElemPtr(ctypes.OnOff{}),
				Italic:       optBoolElemPtr(ctypes.OnOff{}),
				ItalicCS:     optBoolElemPtr(ctypes.OnOff{}),
				Strike:       optBoolElemPtr(ctypes.OnOff{}),
				DoubleStrike: optBoolElemPtr(ctypes.OnOff{}),
				Outline:      optBoolElemPtr(ctypes.OnOff{}),
				Shadow:       optBoolElemPtr(ctypes.OnOff{}),
				Caps:         optBoolElemPtr(ctypes.OnOff{}),
				SmallCaps:    optBoolElemPtr(ctypes.OnOff{}),
				Emboss:       optBoolElemPtr(ctypes.OnOff{}),
				Imprint:      optBoolElemPtr(ctypes.OnOff{}),
				NoGrammar:    optBoolElemPtr(ctypes.OnOff{}),
				SnapToGrid:   optBoolElemPtr(ctypes.OnOff{}),
				Vanish:       optBoolElemPtr(ctypes.OnOff{}),
				WebHidden:    optBoolElemPtr(ctypes.OnOff{}),
				RightToLeft:  optBoolElemPtr(ctypes.OnOff{}),
				CSFormat:     optBoolElemPtr(ctypes.OnOff{}),
				SpecVanish:   optBoolElemPtr(ctypes.OnOff{}),
				OMath:        optBoolElemPtr(ctypes.OnOff{}),
				Kern:         singleUint64ValPtr(ctypes.Uint64Elem{Val: 20}),
				Spacing:      singleIntValPtr(ctypes.DecimalNum{Val: 100}),
				Style:        singleStrValPtr(ctypes.CTString{Val: "Heading1"}),
				Position:     singleIntValPtr(ctypes.DecimalNum{Val: 10}),
			},
			expected: `<w:rPr><w:rStyle w:val="Heading1"></w:rStyle><w:b></w:b><w:bCs></w:bCs><w:i></w:i><w:iCs></w:iCs><w:caps></w:caps><w:smallCaps></w:smallCaps><w:strike></w:strike><w:dstrike></w:dstrike><w:outline></w:outline><w:shadow></w:shadow><w:emboss></w:emboss><w:imprint></w:imprint><w:noProof></w:noProof><w:snapToGrid></w:snapToGrid><w:vanish></w:vanish><w:webHidden></w:webHidden><w:spacing w:val="100"></w:spacing><w:kern w:val="20"></w:kern><w:position w:val="10"></w:position><w:rtl></w:rtl><w:cs></w:cs><w:specVanish></w:specVanish><w:oMath></w:oMath></w:rPr>`,
		},
		{
			name: "Only Bold set",
			prop: RunProperty{
				Bold: optBoolElemPtr(ctypes.OnOff{}),
			},
			expected: `<w:rPr><w:b></w:b></w:rPr>`,
		},
		{
			name:     "No attributes set",
			prop:     RunProperty{},
			expected: `<w:rPr></w:rPr>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			e := xml.NewEncoder(&result)
			err := tt.prop.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:rPr"}})
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}
			e.Flush()

			if result.String() != tt.expected {
				t.Errorf("Expected XML:\n%s\nBut got:\n%s", tt.expected, result.String())
			}
		})
	}
}

func TestRunProperty_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name         string
		inputXML     string
		expectedProp RunProperty
	}{
		{
			name:     "All attributes set",
			inputXML: `<w:rPr><w:rStyle w:val="Heading1"></w:rStyle><w:b></w:b><w:bCs></w:bCs><w:i></w:i><w:iCs></w:iCs><w:caps></w:caps><w:smallCaps></w:smallCaps><w:strike></w:strike><w:dstrike></w:dstrike><w:outline></w:outline><w:shadow></w:shadow><w:emboss></w:emboss><w:imprint></w:imprint><w:noProof></w:noProof><w:snapToGrid></w:snapToGrid><w:vanish></w:vanish><w:webHidden></w:webHidden><w:spacing w:val="100"></w:spacing><w:kern w:val="20"></w:kern><w:position w:val="10"></w:position><w:rtl></w:rtl><w:cs></w:cs><w:specVanish></w:specVanish><w:oMath></w:oMath></w:rPr>`,
			expectedProp: RunProperty{
				Bold:         optBoolElemPtr(ctypes.OnOff{}),
				BoldCS:       optBoolElemPtr(ctypes.OnOff{}),
				Italic:       optBoolElemPtr(ctypes.OnOff{}),
				ItalicCS:     optBoolElemPtr(ctypes.OnOff{}),
				Strike:       optBoolElemPtr(ctypes.OnOff{}),
				DoubleStrike: optBoolElemPtr(ctypes.OnOff{}),
				Outline:      optBoolElemPtr(ctypes.OnOff{}),
				Shadow:       optBoolElemPtr(ctypes.OnOff{}),
				Caps:         optBoolElemPtr(ctypes.OnOff{}),
				SmallCaps:    optBoolElemPtr(ctypes.OnOff{}),
				Emboss:       optBoolElemPtr(ctypes.OnOff{}),
				Imprint:      optBoolElemPtr(ctypes.OnOff{}),
				NoGrammar:    optBoolElemPtr(ctypes.OnOff{}),
				SnapToGrid:   optBoolElemPtr(ctypes.OnOff{}),
				Vanish:       optBoolElemPtr(ctypes.OnOff{}),
				WebHidden:    optBoolElemPtr(ctypes.OnOff{}),
				RightToLeft:  optBoolElemPtr(ctypes.OnOff{}),
				CSFormat:     optBoolElemPtr(ctypes.OnOff{}),
				SpecVanish:   optBoolElemPtr(ctypes.OnOff{}),
				OMath:        optBoolElemPtr(ctypes.OnOff{}),
				Kern:         singleUint64ValPtr(ctypes.Uint64Elem{Val: 20}),
				Spacing:      singleIntValPtr(ctypes.DecimalNum{Val: 100}),
				Style:        singleStrValPtr(ctypes.CTString{Val: "Heading1"}),
				Position:     singleIntValPtr(ctypes.DecimalNum{Val: 10}),
			},
		},
		{
			name:     "Only Bold set",
			inputXML: `<w:rPr><w:b/></w:rPr>`,
			expectedProp: RunProperty{
				Bold: optBoolElemPtr(ctypes.OnOff{}),
			},
		},
		{
			name:         "No attributes set",
			inputXML:     `<w:rPr></w:rPr>`,
			expectedProp: RunProperty{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var prop RunProperty
			err := xml.Unmarshal([]byte(tt.inputXML), &prop)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			// Compare expected and actual RunProperty structs
			if !reflect.DeepEqual(prop, tt.expectedProp) {
				t.Errorf("Expected:\n%+v\nBut got:\n%+v", tt.expectedProp, prop)
			}
		})
	}
}
