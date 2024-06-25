package docxrun

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"

	"github.com/gomutex/godocx/elemtypes"
	"github.com/gomutex/godocx/wml/ctypes"
)

func optBoolElemPtr(value elemtypes.OptBoolElem) *elemtypes.OptBoolElem {
	return &value
}

func singleUint64ValPtr(value elemtypes.SingleUint64Val) *elemtypes.SingleUint64Val {
	return &value
}

func singleIntValPtr(value ctypes.DecimalNum) *ctypes.DecimalNum {
	return &value
}

func singleStrValPtr(value elemtypes.SingleStrVal) *elemtypes.SingleStrVal {
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
				Bold:         optBoolElemPtr(elemtypes.OptBoolElem{}),
				BoldCS:       optBoolElemPtr(elemtypes.OptBoolElem{}),
				Italic:       optBoolElemPtr(elemtypes.OptBoolElem{}),
				ItalicCS:     optBoolElemPtr(elemtypes.OptBoolElem{}),
				Strike:       optBoolElemPtr(elemtypes.OptBoolElem{}),
				DoubleStrike: optBoolElemPtr(elemtypes.OptBoolElem{}),
				Outline:      optBoolElemPtr(elemtypes.OptBoolElem{}),
				Shadow:       optBoolElemPtr(elemtypes.OptBoolElem{}),
				Caps:         optBoolElemPtr(elemtypes.OptBoolElem{}),
				SmallCaps:    optBoolElemPtr(elemtypes.OptBoolElem{}),
				Emboss:       optBoolElemPtr(elemtypes.OptBoolElem{}),
				Imprint:      optBoolElemPtr(elemtypes.OptBoolElem{}),
				NoGrammar:    optBoolElemPtr(elemtypes.OptBoolElem{}),
				SnapToGrid:   optBoolElemPtr(elemtypes.OptBoolElem{}),
				Vanish:       optBoolElemPtr(elemtypes.OptBoolElem{}),
				WebHidden:    optBoolElemPtr(elemtypes.OptBoolElem{}),
				RightToLeft:  optBoolElemPtr(elemtypes.OptBoolElem{}),
				CSFormat:     optBoolElemPtr(elemtypes.OptBoolElem{}),
				SpecVanish:   optBoolElemPtr(elemtypes.OptBoolElem{}),
				OMath:        optBoolElemPtr(elemtypes.OptBoolElem{}),
				Kern:         singleUint64ValPtr(elemtypes.SingleUint64Val{Val: 20}),
				Spacing:      singleIntValPtr(ctypes.DecimalNum{Val: 100}),
				Style:        singleStrValPtr(elemtypes.SingleStrVal{Val: "Heading1"}),
				Position:     singleIntValPtr(ctypes.DecimalNum{Val: 10}),
			},
			expected: `<w:rPr><w:rStyle w:val="Heading1"></w:rStyle><w:b></w:b><w:bCs></w:bCs><w:i></w:i><w:iCs></w:iCs><w:caps></w:caps><w:smallCaps></w:smallCaps><w:strike></w:strike><w:dstrike></w:dstrike><w:outline></w:outline><w:shadow></w:shadow><w:emboss></w:emboss><w:imprint></w:imprint><w:noProof></w:noProof><w:snapToGrid></w:snapToGrid><w:vanish></w:vanish><w:webHidden></w:webHidden><w:spacing w:val="100"></w:spacing><w:kern w:val="20"></w:kern><w:position w:val="10"></w:position><w:rtl></w:rtl><w:cs></w:cs><w:specVanish></w:specVanish><w:oMath></w:oMath></w:rPr>`,
		},
		{
			name: "Only Bold set",
			prop: RunProperty{
				Bold: optBoolElemPtr(elemtypes.OptBoolElem{}),
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
				Bold:         optBoolElemPtr(elemtypes.OptBoolElem{}),
				BoldCS:       optBoolElemPtr(elemtypes.OptBoolElem{}),
				Italic:       optBoolElemPtr(elemtypes.OptBoolElem{}),
				ItalicCS:     optBoolElemPtr(elemtypes.OptBoolElem{}),
				Strike:       optBoolElemPtr(elemtypes.OptBoolElem{}),
				DoubleStrike: optBoolElemPtr(elemtypes.OptBoolElem{}),
				Outline:      optBoolElemPtr(elemtypes.OptBoolElem{}),
				Shadow:       optBoolElemPtr(elemtypes.OptBoolElem{}),
				Caps:         optBoolElemPtr(elemtypes.OptBoolElem{}),
				SmallCaps:    optBoolElemPtr(elemtypes.OptBoolElem{}),
				Emboss:       optBoolElemPtr(elemtypes.OptBoolElem{}),
				Imprint:      optBoolElemPtr(elemtypes.OptBoolElem{}),
				NoGrammar:    optBoolElemPtr(elemtypes.OptBoolElem{}),
				SnapToGrid:   optBoolElemPtr(elemtypes.OptBoolElem{}),
				Vanish:       optBoolElemPtr(elemtypes.OptBoolElem{}),
				WebHidden:    optBoolElemPtr(elemtypes.OptBoolElem{}),
				RightToLeft:  optBoolElemPtr(elemtypes.OptBoolElem{}),
				CSFormat:     optBoolElemPtr(elemtypes.OptBoolElem{}),
				SpecVanish:   optBoolElemPtr(elemtypes.OptBoolElem{}),
				OMath:        optBoolElemPtr(elemtypes.OptBoolElem{}),
				Kern:         singleUint64ValPtr(elemtypes.SingleUint64Val{Val: 20}),
				Spacing:      singleIntValPtr(ctypes.DecimalNum{Val: 100}),
				Style:        singleStrValPtr(elemtypes.SingleStrVal{Val: "Heading1"}),
				Position:     singleIntValPtr(ctypes.DecimalNum{Val: 10}),
			},
		},
		{
			name:     "Only Bold set",
			inputXML: `<w:rPr><w:b/></w:rPr>`,
			expectedProp: RunProperty{
				Bold: optBoolElemPtr(elemtypes.OptBoolElem{}),
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
