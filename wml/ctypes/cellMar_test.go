package ctypes

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"

	"github.com/gomutex/godocx/wml/stypes"
)

func TestCellMarginsMarshalXML(t *testing.T) {
	testCases := []struct {
		input    CellMargins
		expected string
	}{
		{
			input: CellMargins{
				Top:    NewTableWidth(0, stypes.TableWidthDxa),
				Left:   NewTableWidth(55, stypes.TableWidthDxa),
				Bottom: NewTableWidth(0, stypes.TableWidthDxa),
				Right:  NewTableWidth(55, stypes.TableWidthDxa),
			},
			expected: `<w:tblCellMar><w:top w:w="0" w:type="dxa"></w:top><w:left w:w="55" w:type="dxa"></w:left><w:bottom w:w="0" w:type="dxa"></w:bottom><w:right w:w="55" w:type="dxa"></w:right></w:tblCellMar>`,
		},
		{
			input:    DefaultCellMargins(),
			expected: `<w:tblCellMar></w:tblCellMar>`,
		},
	}

	for _, tc := range testCases {
		var result strings.Builder
		encoder := xml.NewEncoder(&result)

		start := xml.StartElement{Name: xml.Name{Local: "w:tblCellMar"}}
		err := tc.input.MarshalXML(encoder, start)

		if err != nil {
			t.Errorf("Error during MarshalXML: %v", err)
		}

		err = encoder.Flush()
		if err != nil {
			t.Errorf("Error flushing encoder: %v", err)
		}

		if result.String() != tc.expected {
			t.Errorf("Expected XML:\n%s\n\nActual XML:\n%s", tc.expected, result.String())
		}
	}
}

func TestCellMarginsUnmarshalXML(t *testing.T) {
	testCases := []struct {
		input    string
		expected CellMargins
	}{
		{
			input: `<w:tblCellMar xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main"><w:top w:w="0" w:type="dxa"></w:top><w:left w:w="55" w:type="dxa"></w:left><w:bottom w:w="0" w:type="dxa"></w:bottom><w:right w:w="55" w:type="dxa"></w:right></w:tblCellMar>`,
			expected: CellMargins{
				Top:    NewTableWidth(0, stypes.TableWidthDxa),
				Left:   NewTableWidth(55, stypes.TableWidthDxa),
				Bottom: NewTableWidth(0, stypes.TableWidthDxa),
				Right:  NewTableWidth(55, stypes.TableWidthDxa),
			},
		},
	}

	for _, tc := range testCases {
		var result CellMargins

		err := xml.Unmarshal([]byte(tc.input), &result)
		if err != nil {
			t.Fatalf("Error unmarshaling XML: %v", err)
		}

		if result.Top == nil {
			t.Errorf("Got nil value for Top")
		}

		if result.Bottom == nil {
			t.Errorf("Got nil value for Bottom")
		}

		if result.Left == nil {
			t.Errorf("Got nil value for Left")
		}

		if result.Right == nil {
			t.Errorf("Got nil value for Right")
		}

		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %+v, but got %+v", tc.expected, result)
		}
	}
}
