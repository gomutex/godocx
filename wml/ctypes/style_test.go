package ctypes

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/gomutex/godocx/internal"
	"github.com/gomutex/godocx/wml/stypes"
)

func tmpOnOffFromStr(value string) *OnOff {
	v, _ := OnOffFromStr(value)
	return v
}

func TestStyle_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		style    *Style
		expected string
	}{
		{
			name: "all properties set",
			style: &Style{
				Name:         &CTString{"StyleName"},
				Alias:        &CTString{"Alias"},
				BasedOn:      &CTString{"BaseStyle"},
				Next:         &CTString{"NextStyle"},
				Link:         &CTString{"LinkedStyle"},
				AutoRedefine: tmpOnOffFromStr("on"),
				Hidden:       tmpOnOffFromStr("off"),
				UIPriority: &DecimalNum{
					Val: 1,
				},
				SemiHidden:      tmpOnOffFromStr("on"),
				UnhideWhenUsed:  tmpOnOffFromStr("on"),
				QFormat:         tmpOnOffFromStr("off"),
				Locked:          tmpOnOffFromStr("on"),
				Personal:        tmpOnOffFromStr("off"),
				PersonalCompose: tmpOnOffFromStr("off"),
				PersonalReply:   tmpOnOffFromStr("off"),
				RevID:           NewGenSingleStrVal(stypes.LongHexNum("FF00")),
				ParaProp:        &ParagraphProp{},
				RunProp:         &RunProperty{},
				TableProp:       &TableProp{},
				TableRowProp:    &RowProperty{},
				TableCellProp:   &CellProperty{},
				TableStylePr: []TableStyleProp{
					{
						Type: stypes.TblStyleOverrideWholeTable,
					},
				},
				Type:        internal.ToPtr(stypes.StyleTypeCharacter),
				ID:          stringPtr("styleID"),
				Default:     internal.ToPtr(stypes.OnOffOn),
				CustomStyle: internal.ToPtr(stypes.OnOffOff),
			},
			expected: `<w:style w:type="character" w:styleId="styleID" w:default="on" w:customStyle="off">` +
				`<w:name w:val="StyleName"></w:name>` +
				`<w:alias w:val="Alias"></w:alias>` +
				`<w:basedOn w:val="BaseStyle"></w:basedOn>` +
				`<w:next w:val="NextStyle"></w:next>` +
				`<w:link w:val="LinkedStyle"></w:link>` +
				`<w:autoRedefine w:val="on"></w:autoRedefine>` +
				`<w:hidden w:val="off"></w:hidden>` +
				`<w:uiPriority w:val="1"></w:uiPriority>` +
				`<w:semiHidden w:val="on"></w:semiHidden>` +
				`<w:unhideWhenUsed w:val="on"></w:unhideWhenUsed>` +
				`<w:qFormat w:val="off"></w:qFormat>` +
				`<w:locked w:val="on"></w:locked>` +
				`<w:personal w:val="off"></w:personal>` +
				`<w:personalCompose w:val="off"></w:personalCompose>` +
				`<w:personalReply w:val="off"></w:personalReply>` +
				`<w:rsid w:val="FF00"></w:rsid>` +
				`<w:pPr></w:pPr>` +
				`<w:rPr></w:rPr>` +
				`<w:tblPr></w:tblPr>` +
				`<w:trPr></w:trPr>` +
				`<w:tcPr></w:tcPr>` +
				`<w:tblStylePr w:type="wholeTable"></w:tblStylePr>` +
				`</w:style>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(tt.style)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got := string(output); got != tt.expected {
				t.Errorf("expected %s, but got %s", tt.expected, got)
			}
		})
	}
}
func TestStyle_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		xmlInput string
		expected *Style
	}{
		{
			name: "all properties set",
			xmlInput: `<w:style w:type="character" w:styleId="styleID" w:default="on" w:customStyle="off">
							<w:name w:val="StyleName"></w:name>
							<w:alias w:val="Alias"></w:alias>
							<w:basedOn w:val="BaseStyle"></w:basedOn>
							<w:next w:val="NextStyle"></w:next>
							<w:link w:val="LinkedStyle"></w:link>
							<w:autoRedefine w:val="on"></w:autoRedefine>
							<w:hidden w:val="off"></w:hidden>
							<w:uiPriority w:val="1"></w:uiPriority>
							<w:semiHidden w:val="on"></w:semiHidden>
							<w:unhideWhenUsed w:val="on"></w:unhideWhenUsed>
							<w:qFormat w:val="off"></w:qFormat>
							<w:locked w:val="on"></w:locked>
							<w:personal w:val="off"></w:personal>
							<w:personalCompose w:val="off"></w:personalCompose>
							<w:personalReply w:val="off"></w:personalReply>
							<w:rsid w:val="FF00"></w:rsid>
							<w:pPr></w:pPr>
							<w:rPr></w:rPr>
							<w:tblPr></w:tblPr>
							<w:trPr></w:trPr>
							<w:tcPr></w:tcPr>
							<w:tblStylePr w:type="wholeTable"></w:tblStylePr>
						</w:style>`,
			expected: &Style{
				Name:            &CTString{Val: "StyleName"},
				Alias:           &CTString{Val: "Alias"},
				BasedOn:         &CTString{Val: "BaseStyle"},
				Next:            &CTString{Val: "NextStyle"},
				Link:            &CTString{Val: "LinkedStyle"},
				AutoRedefine:    tmpOnOffFromStr("on"),
				Hidden:          tmpOnOffFromStr("off"),
				UIPriority:      &DecimalNum{Val: 1},
				SemiHidden:      tmpOnOffFromStr("on"),
				UnhideWhenUsed:  tmpOnOffFromStr("on"),
				QFormat:         tmpOnOffFromStr("off"),
				Locked:          tmpOnOffFromStr("on"),
				Personal:        tmpOnOffFromStr("off"),
				PersonalCompose: tmpOnOffFromStr("off"),
				PersonalReply:   tmpOnOffFromStr("off"),
				RevID:           NewGenSingleStrVal(stypes.LongHexNum("FF00")),
				ParaProp:        &ParagraphProp{},
				RunProp:         &RunProperty{},
				TableProp:       &TableProp{},
				TableRowProp:    &RowProperty{},
				TableCellProp:   &CellProperty{},
				TableStylePr: []TableStyleProp{
					{
						Type: stypes.TblStyleOverrideWholeTable,
					},
				},
				Type:        internal.ToPtr(stypes.StyleTypeCharacter),
				ID:          stringPtr("styleID"),
				Default:     internal.ToPtr(stypes.OnOffOn),
				CustomStyle: internal.ToPtr(stypes.OnOffOff),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var style Style
			err := xml.Unmarshal([]byte(tt.xmlInput), &style)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(&style, tt.expected) {
				t.Errorf("expected %+v, but got %+v", tt.expected, &style)
			}
		})
	}
}
