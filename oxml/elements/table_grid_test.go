package elements

import (
	"encoding/xml"
	"strings"
	"testing"
)

func equalGrids(g1, g2 []uint64) bool {
	if len(g1) != len(g2) {
		return false
	}
	for i, v := range g1 {
		if v != g2[i] {
			return false
		}
	}
	return true
}

func TestTableGridMarshalingUnmarshaling(t *testing.T) {
	testGrid := NewTableGrid([]uint64{100, 200})

	xmlData, err := xml.Marshal(testGrid)
	if err != nil {
		t.Fatalf("Error marshaling TableGrid to XML: %v", err)
	}

	var unmarshaledGrid TableGrid
	err = xml.Unmarshal(xmlData, &unmarshaledGrid)
	if err != nil {
		t.Fatalf("Error unmarshaling XML to TableGrid: %v", err)
	}

	if !equalGrids(testGrid.Grid, unmarshaledGrid.Grid) {
		t.Errorf("Expected grid values %v, got %v", testGrid.Grid, unmarshaledGrid.Grid)
	}

	expectedXMLString := `<w:tblGrid><w:gridCol w:w="100" w:type="dxa"></w:gridCol><w:gridCol w:w="200" w:type="dxa"></w:gridCol></w:tblGrid>`
	if !strings.Contains(string(xmlData), expectedXMLString) {
		t.Errorf("Expected XML string %s, got %s", expectedXMLString, string(xmlData))
	}
}
