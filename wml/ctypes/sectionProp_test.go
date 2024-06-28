package ctypes

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"

	"github.com/gomutex/godocx/wml/stypes"
)

func TestSectionProp_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    SectionProp
		expected string
	}{
		{
			name: "All attributes",
			input: SectionProp{
				HeaderReference: &HeaderReference{Type: "default", ID: "rId1"},
				FooterReference: &FooterReference{Type: "default", ID: "rId2"},
				PageSize: &PageSize{
					Width:  uint64Ptr(12240),
					Height: uint64Ptr(15840),
				},
				Type:       NewGenSingleStrVal(stypes.SectionMarkNextPage),
				PageMargin: &PageMargin{Top: intPtr(1440), Bottom: intPtr(1440), Left: intPtr(1440), Right: intPtr(1440)},
				PageNum:    &PageNumbering{Format: stypes.NumFmtDecimal},
				FormProt:   NewGenSingleStrVal(stypes.OnOffTrue),
				TitlePg:    NewGenSingleStrVal(stypes.OnOffTrue),
				TextDir:    NewGenSingleStrVal(stypes.TextDirectionLrTb),
				DocGrid:    &DocGrid{Type: "default", LinePitch: intPtr(360)},
			},
			expected: `<w:sectPr><w:headerReference w:type="default" r:id="rId1"></w:headerReference><w:footerReference w:type="default" r:id="rId2"></w:footerReference><w:type w:val="nextPage"></w:type><w:pgSz w:w="12240" w:h="15840"></w:pgSz><w:pgMar w:left="1440" w:right="1440" w:top="1440" w:bottom="1440"></w:pgMar><w:pgNumType w:fmt="decimal"></w:pgNumType><w:formProt w:val="true"></w:formProt><w:titlePg w:val="true"></w:titlePg><w:textDirection w:val="lrTb"></w:textDirection><w:docGrid w:type="default" w:linePitch="360"></w:docGrid></w:sectPr>`,
		},
		{
			name:     "No attributes",
			input:    SectionProp{},
			expected: `<w:sectPr></w:sectPr>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:sectPr"}}

			err := tt.input.MarshalXML(encoder, start)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			if err = encoder.Flush(); err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			output := result.String()
			if output != tt.expected {
				t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", tt.expected, output)
			}
		})
	}
}

func TestSectionProp_MarshalXML_EmptyFields(t *testing.T) {
	// Test case to ensure empty fields are not serialized
	input := SectionProp{}

	var result strings.Builder
	encoder := xml.NewEncoder(&result)
	start := xml.StartElement{Name: xml.Name{Local: "w:sectPr"}}

	err := input.MarshalXML(encoder, start)
	if err != nil {
		t.Fatalf("Error marshaling XML: %v", err)
	}

	if err = encoder.Flush(); err != nil {
		t.Fatalf("Error marshaling XML: %v", err)
	}

	expected := `<w:sectPr></w:sectPr>`
	output := result.String()
	if output != expected {
		t.Errorf("XML mismatch\nExpected:\n%s\nActual:\n%s", expected, output)
	}
}

func TestSectionProp_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected SectionProp
	}{
		{
			name: "All attributes",
			inputXML: `<w:sectPr>
				<w:headerReference w:type="default" w:id="rId1"></w:headerReference>
				<w:footerReference w:type="default" w:id="rId2"></w:footerReference>
				<w:pgSz w:w="12240" w:h="15840"></w:pgSz>
				<w:type w:val="nextPage"></w:type>
				<w:pgMar w:top="1440" w:bottom="1440" w:left="1440" w:right="1440"></w:pgMar>
				<w:pgNumType w:fmt="decimal"></w:pgNumType>
				<w:formProt w:val="true"></w:formProt>
				<w:titlePg w:val="true"></w:titlePg>
				<w:textDirection w:val="lrTb"></w:textDirection>
				<w:docGrid w:type="default" w:linePitch="360"></w:docGrid>
			</w:sectPr>`,
			expected: SectionProp{
				HeaderReference: &HeaderReference{Type: "default", ID: "rId1"},
				FooterReference: &FooterReference{Type: "default", ID: "rId2"},
				PageSize: &PageSize{
					Width:  uint64Ptr(12240),
					Height: uint64Ptr(15840),
				},
				Type:       NewGenSingleStrVal(stypes.SectionMarkNextPage),
				PageMargin: &PageMargin{Top: intPtr(1440), Bottom: intPtr(1440), Left: intPtr(1440), Right: intPtr(1440)},
				PageNum:    &PageNumbering{Format: stypes.NumFmtDecimal},
				FormProt:   NewGenSingleStrVal(stypes.OnOffTrue),
				TitlePg:    NewGenSingleStrVal(stypes.OnOffTrue),
				TextDir:    NewGenSingleStrVal(stypes.TextDirectionLrTb),
				DocGrid:    &DocGrid{Type: "default", LinePitch: intPtr(360)},
			},
		},
		{
			name:     "No attributes",
			inputXML: `<w:sectPr></w:sectPr>`,
			expected: SectionProp{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result SectionProp

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error during unmarshaling: %v", err)
			}

			// Compare individual fields for equality
			if !reflect.DeepEqual(result.HeaderReference, tt.expected.HeaderReference) {
				t.Errorf("HeaderReference mismatch\nExpected: %#v\nActual:   %#v", tt.expected.HeaderReference, result.HeaderReference)
			}
			if !reflect.DeepEqual(result.FooterReference, tt.expected.FooterReference) {
				t.Errorf("FooterReference mismatch\nExpected: %#v\nActual:   %#v", tt.expected.FooterReference, result.FooterReference)
			}
			if !reflect.DeepEqual(result.PageSize, tt.expected.PageSize) {
				t.Errorf("PageSize mismatch\nExpected: %#v\nActual:   %#v", tt.expected.PageSize, result.PageSize)
			}
			if !reflect.DeepEqual(result.Type, tt.expected.Type) {
				t.Errorf("Type mismatch\nExpected: %#v\nActual:   %#v", tt.expected.Type, result.Type)
			}
			if !reflect.DeepEqual(result.PageMargin, tt.expected.PageMargin) {
				t.Errorf("PageMargin mismatch\nExpected: %#v\nActual:   %#v", tt.expected.PageMargin, result.PageMargin)
			}
			if !reflect.DeepEqual(result.PageNum, tt.expected.PageNum) {
				t.Errorf("PageNum mismatch\nExpected: %#v\nActual:   %#v", tt.expected.PageNum, result.PageNum)
			}
			if !reflect.DeepEqual(result.FormProt, tt.expected.FormProt) {
				t.Errorf("FormProt mismatch\nExpected: %#v\nActual:   %#v", tt.expected.FormProt, result.FormProt)
			}
			if !reflect.DeepEqual(result.TitlePg, tt.expected.TitlePg) {
				t.Errorf("TitlePg mismatch\nExpected: %#v\nActual:   %#v", tt.expected.TitlePg, result.TitlePg)
			}
			if !reflect.DeepEqual(result.TextDir, tt.expected.TextDir) {
				t.Errorf("TextDir mismatch\nExpected: %#v\nActual:   %#v", tt.expected.TextDir, result.TextDir)
			}
			if !reflect.DeepEqual(result.DocGrid, tt.expected.DocGrid) {
				t.Errorf("DocGrid mismatch\nExpected: %#v\nActual:   %#v", tt.expected.DocGrid, result.DocGrid)
			}
		})
	}
}
