package dml

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/common/units"
)

func TestNewExtent(t *testing.T) {
	width := units.Emu(100)
	height := units.Emu(200)
	extent := NewExtent(width, height)

	if extent.Width != uint64(width) {
		t.Errorf("Width does not match. Expected %d, got %d", width, extent.Width)
	}

	if extent.Height != uint64(height) {
		t.Errorf("Height does not match. Expected %d, got %d", height, extent.Height)
	}
}

func TestMarshalExtent(t *testing.T) {
	tests := []struct {
		extent      *Extent
		expectedXML string
		xmlName     string
	}{
		{
			extent:      NewExtent(units.Emu(100), units.Emu(200)),
			expectedXML: `<w:extent cx="100" cy="200"></w:extent>`,
			xmlName:     "w:extent",
		},
		{
			extent:      NewExtent(units.Emu(150), units.Emu(250)),
			expectedXML: `<a:ext cx="150" cy="250"></a:ext>`,
			xmlName:     "a:ext",
		},
	}

	for _, tt := range tests {
		t.Run(tt.xmlName, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)

			start := xml.StartElement{Name: xml.Name{Local: tt.xmlName}}
			err := tt.extent.MarshalXML(encoder, start)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			err = encoder.Flush()
			if err != nil {
				t.Errorf("Error flushing encoder: %v", err)
			}

			if result.String() != tt.expectedXML {
				t.Errorf("Expected XML:\n%s\nBut got:\n%s", tt.expectedXML, result.String())
			}
		})
	}
}

func TestUnmarshalExtent(t *testing.T) {
	tests := []struct {
		inputXML    string
		expectedExt Extent
	}{
		{
			inputXML: `<w:extent cx="100" cy="200"></w:extent>`,
			expectedExt: Extent{
				Width:  100,
				Height: 200,
			},
		},
		{
			inputXML: `<a:extent cx="150" cy="250"></a:extent>`,
			expectedExt: Extent{
				Width:  150,
				Height: 250,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			var extent Extent

			err := xml.Unmarshal([]byte(tt.inputXML), &extent)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if extent.Width != tt.expectedExt.Width {
				t.Errorf("Expected width %d, but got %d", tt.expectedExt.Width, extent.Width)
			}
			if extent.Height != tt.expectedExt.Height {
				t.Errorf("Expected height %d, but got %d", tt.expectedExt.Height, extent.Height)
			}
		})
	}
}
