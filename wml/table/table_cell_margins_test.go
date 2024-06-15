package table

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"
)

func TestTableCellMarginsMarshalXML(t *testing.T) {
	testCases := []struct {
		input    TableCellMargins
		expected string
	}{
		{
			input: TableCellMargins{
				Top:    NewCellMargin(0, WidthTypeDxa),
				Left:   NewCellMargin(55, WidthTypeDxa),
				Bottom: NewCellMargin(0, WidthTypeDxa),
				Right:  NewCellMargin(55, WidthTypeDxa),
			},
			expected: `<w:tblCellMar><w:top w:w="0" w:type="dxa"></w:top><w:left w:w="55" w:type="dxa"></w:left><w:bottom w:w="0" w:type="dxa"></w:bottom><w:right w:w="55" w:type="dxa"></w:right></w:tblCellMar>`,
		},
		{
			input:    DefaultTableCellMargins(),
			expected: `<w:tblCellMar></w:tblCellMar>`,
		},
	}

	for _, tc := range testCases {
		var result strings.Builder
		encoder := xml.NewEncoder(&result)

		start := xml.StartElement{Name: xml.Name{Local: "fake"}}
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

func TestTableCellMarginsUnmarshalXML(t *testing.T) {
	testCases := []struct {
		input    string
		expected TableCellMargins
	}{
		{
			input: `<w:tblCellMar xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main"><w:top w:w="0" w:type="dxa"></w:top><w:left w:w="55" w:type="dxa"></w:left><w:bottom w:w="0" w:type="dxa"></w:bottom><w:right w:w="55" w:type="dxa"></w:right></w:tblCellMar>`,
			expected: TableCellMargins{
				Top:    NewCellMargin(0, WidthTypeDxa),
				Left:   NewCellMargin(55, WidthTypeDxa),
				Bottom: NewCellMargin(0, WidthTypeDxa),
				Right:  NewCellMargin(55, WidthTypeDxa),
			},
		},
	}

	for _, tc := range testCases {
		decoder := xml.NewDecoder(strings.NewReader(tc.input))
		var result TableCellMargins

		startToken, err := decoder.Token()
		if err != nil {
			t.Errorf("Error getting start token: %v", err)
		}

		err = result.UnmarshalXML(decoder, startToken.(xml.StartElement))
		if err != nil {
			t.Errorf("Error during UnmarshalXML: %v", err)
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
