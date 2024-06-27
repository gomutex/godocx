package docx

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/wml/stypes"
)

func TestDocument_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    Document
		expected []string
	}{
		{
			name: "With Background and Body",
			input: Document{
				Background: &Background{
					Color:      StringPtr("FF0000"),
					ThemeColor: ThemeColorPtr(stypes.ThemeColorAccent1),
					ThemeTint:  StringPtr("500"),
					ThemeShade: StringPtr("200"),
				},
				Body: &Body{},
			},
			expected: []string{
				`xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships"`,
				`<w:background w:color="FF0000" w:themeColor="accent1" w:themeTint="500" w:themeShade="200"></w:background>`,
				`<w:body></w:body>`,
				`xmlns:o="urn:schemas-microsoft-com:office:office"`,
				`xmlns:v="urn:schemas-microsoft-com:vml"`,
				`xmlns:w10="urn:schemas-microsoft-com:office:word"`,
				`xmlns:wp="http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"`,
				`xmlns:wps="http://schemas.microsoft.com/office/word/2010/wordprocessingShape"`,
				`xmlns:wpg="http://schemas.microsoft.com/office/word/2010/wordprocessingGroup"`,
				`xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006"`,
				`xmlns:wp14="http://schemas.microsoft.com/office/word/2010/wordprocessingDrawing"`,
				`xmlns:w14="http://schemas.microsoft.com/office/word/2010/wordml"`,
				`xmlns:w15="http://schemas.microsoft.com/office/word/2012/wordml"`,
				`mc:Ignorable="w14 wp14 w15"`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:document"}}

			err := tt.input.MarshalXML(encoder, start)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			encoder.Flush()
			actual := result.String()

			for _, exp := range tt.expected {
				if !strings.Contains(actual, exp) {
					t.Errorf("Expected XML part not found in actual XML:\nExpected part: %s\nActual XML: %s", exp, actual)
				}
			}
		})
	}
}
