package docxrun

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/wml/simpletypes"
)

func TestRunFontsMarshalXML(t *testing.T) {
	rf := RunFonts{
		Hint:          simpletypes.FontTypeHintDefault,
		Ascii:         "Arial",
		HAnsi:         "Calibri",
		EastAsia:      "SimSun",
		CS:            "Arial",
		AsciiTheme:    simpletypes.ThemeFontMajorAscii,
		HAnsiTheme:    simpletypes.ThemeFontMajorHAnsi,
		EastAsiaTheme: simpletypes.ThemeFontMajorEastAsia,
		CSTheme:       simpletypes.ThemeFontMajorBidi,
	}

	var output strings.Builder
	encoder := xml.NewEncoder(&output)
	err := rf.MarshalXML(encoder, xml.StartElement{})
	if err != nil {
		t.Fatalf("Error marshaling XML: %v", err)
	}
	encoder.Flush()

	expected := `<w:rFonts w:eastAsia="SimSun" w:hint="default" w:ascii="Arial" w:hAnsi="Calibri" w:cs="Arial" w:asciiTheme="majorAscii" w:hAnsiTheme="majorHAnsi" w:eastAsiaTheme="majorEastAsia" w:cstheme="majorBidi"></w:rFonts>`
	if output.String() != expected {
		t.Errorf("Expected %s but got %s", expected, output.String())
	}
}

func TestRunFontsUnmarshalXML(t *testing.T) {
	input := `<w:rFonts w:eastAsia="SimSun" w:hint="default" w:ascii="Arial" w:hAnsi="Calibri" w:cs="Arial" w:asciiTheme="majorAscii" w:hAnsiTheme="majorHAnsi" w:eastAsiaTheme="majorEastAsia" w:cstheme="majorBidi"></w:rFonts>`

	var rf RunFonts
	err := xml.Unmarshal([]byte(input), &rf)
	if err != nil {
		t.Fatalf("Error unmarshaling XML: %v", err)
	}

	expected := RunFonts{
		Hint:          simpletypes.FontTypeHintDefault,
		Ascii:         "Arial",
		HAnsi:         "Calibri",
		EastAsia:      "SimSun",
		CS:            "Arial",
		AsciiTheme:    simpletypes.ThemeFontMajorAscii,
		HAnsiTheme:    simpletypes.ThemeFontMajorHAnsi,
		EastAsiaTheme: simpletypes.ThemeFontMajorEastAsia,
		CSTheme:       simpletypes.ThemeFontMajorBidi,
	}

	if rf != expected {
		t.Errorf("Expected %+v but got %+v", expected, rf)
	}
}
