package hdrftr

import (
	"encoding/xml"
	"errors"
	"strings"
	"testing"

	"github.com/gomutex/godocx/wml/stypes"
)

func TestTitlePg_MarshalXML(t *testing.T) {
	tests := []struct {
		titlePg     TitlePg
		expectedXML string
	}{
		{
			titlePg: TitlePg{
				Val: stypes.OnOff("on"),
			},
			expectedXML: `<w:titlePg w:val="on"></w:titlePg>`,
		},
		{
			titlePg: TitlePg{
				Val: stypes.OnOff("1"),
			},
			expectedXML: `<w:titlePg w:val="1"></w:titlePg>`,
		},
		{
			titlePg: TitlePg{
				Val: stypes.OnOff(""),
			},
			expectedXML: `<w:titlePg></w:titlePg>`,
		},
	}

	for _, test := range tests {
		t.Run(test.expectedXML, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:titlePg"}}

			err := test.titlePg.MarshalXML(encoder, start)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			err = encoder.Flush()
			if err != nil {
				t.Fatalf("Error flushing encoder: %v", err)
			}

			gotXML := result.String()
			if gotXML != test.expectedXML {
				t.Errorf("MarshalXML() = %s, want %s", gotXML, test.expectedXML)
			}
		})
	}
}

func TestTitlePg_UnmarshalXML(t *testing.T) {
	tests := []struct {
		xmlInput        string
		expectedTitlePg TitlePg
		err             error
	}{
		{
			xmlInput: `<w:titlePg w:val="on"></w:titlePg>`,
			expectedTitlePg: TitlePg{
				Val: stypes.OnOff("on"),
			},
		},
		{
			xmlInput: `<w:titlePg w:val="off"></w:titlePg>`,
			expectedTitlePg: TitlePg{
				Val: stypes.OnOff("off"),
			},
		},
		{
			xmlInput: `<w:titlePg w:val=""></w:titlePg>`,
			expectedTitlePg: TitlePg{
				Val: stypes.OnOff(""),
			},
			err: errors.New("invalid OnOff string"),
		},
		{
			xmlInput: `<w:titlePg></w:titlePg>`,
			expectedTitlePg: TitlePg{
				Val: stypes.OnOff(""),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.xmlInput, func(t *testing.T) {
			var titlePg TitlePg
			reader := strings.NewReader(test.xmlInput)
			decoder := xml.NewDecoder(reader)

			err := decoder.Decode(&titlePg)
			if test.err != nil {
				if err.Error() != test.err.Error() {
					t.Errorf("UnmarshalXML() error = %v, want %v", err, test.err)
				}
				return
			}
			if err != nil {
				t.Fatalf("Error decoding XML: %v", err)
			}

			if titlePg.Val != test.expectedTitlePg.Val {
				t.Errorf("UnmarshalXML() Val = %s, want %s", titlePg.Val, test.expectedTitlePg.Val)
			}
		})
	}
}
