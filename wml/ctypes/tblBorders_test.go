package ctypes

import (
	"bytes"
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/gomutex/godocx/wml/stypes"
)

func TestTableBorders_MarshalXML(t *testing.T) {
	colorRed := "red"
	themeColorAccent1 := stypes.ThemeColor("accent1")
	themeTint := "80"
	themeShade := "20"
	space := "4"
	shadow := stypes.OnOff("true")
	frame := stypes.OnOff("false")

	borders := &TableBorders{
		Top:     &Border{Val: stypes.BorderStyleSingle, Color: &colorRed, ThemeColor: &themeColorAccent1, ThemeTint: &themeTint, ThemeShade: &themeShade, Space: &space, Shadow: &shadow, Frame: &frame},
		Left:    &Border{Val: stypes.BorderStyleDouble},
		Bottom:  &Border{Val: stypes.BorderStyleDashed},
		Right:   &Border{Val: stypes.BorderStyleDotted},
		InsideH: &Border{Val: stypes.BorderStyleSingle},
		InsideV: &Border{Val: stypes.BorderStyleDouble},
	}

	xmlData, err := xml.Marshal(borders)
	if err != nil {
		t.Fatalf("Error marshaling TableBorders to XML: %v", err)
	}

	expectedXMLString := `<w:tblBorders>` +
		`<w:top w:val="single" w:color="red" w:themeColor="accent1" w:themeTint="80" w:themeShade="20" w:space="4" w:shadow="true" w:frame="false"></w:top>` +
		`<w:left w:val="double"></w:left>` +
		`<w:bottom w:val="dashed"></w:bottom>` +
		`<w:right w:val="dotted"></w:right>` +
		`<w:insideH w:val="single"></w:insideH>` +
		`<w:insideV w:val="double"></w:insideV>` +
		`</w:tblBorders>`

	if !bytes.Contains(xmlData, []byte(expectedXMLString)) {
		t.Errorf("Expected XML string %s, got %s", expectedXMLString, string(xmlData))
	}
}

func TestDefaultTableBorders_MarshalXML(t *testing.T) {
	borders := DefaultTableBorders()

	xmlData, err := xml.Marshal(borders)
	if err != nil {
		t.Fatalf("Error marshaling DefaultTableBorders to XML: %v", err)
	}

	expectedXMLString := `<w:tblBorders></w:tblBorders>`
	if string(xmlData) != expectedXMLString {
		t.Errorf("Expected XML string %s, got %s", expectedXMLString, string(xmlData))
	}
}

func TestTableBorders_UnmarshalXML_Valid(t *testing.T) {
	xmlData := `
	<w:tblBorders xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
		<w:top w:val="single"></w:top>
		<w:left w:val="double"></w:left>
		<w:bottom w:val="dashed"></w:bottom>
		<w:right w:val="dotted"></w:right>
		<w:insideH w:val="single"></w:insideH>
		<w:insideV w:val="double"></w:insideV>
	</w:tblBorders>`

	expectedBorders := &TableBorders{
		Top:     &Border{Val: stypes.BorderStyleSingle},
		Left:    &Border{Val: stypes.BorderStyleDouble},
		Bottom:  &Border{Val: stypes.BorderStyleDashed},
		Right:   &Border{Val: stypes.BorderStyleDotted},
		InsideH: &Border{Val: stypes.BorderStyleSingle},
		InsideV: &Border{Val: stypes.BorderStyleDouble},
	}

	var unmarshaledBorders TableBorders

	err := xml.Unmarshal([]byte(xmlData), &unmarshaledBorders)
	if err != nil {
		t.Fatalf("Error unmarshaling XML to TableBorders: %v", err)
	}

	if !reflect.DeepEqual(&unmarshaledBorders, expectedBorders) {
		t.Errorf("Expected %#v, got %#v", expectedBorders, unmarshaledBorders)
	}
}
