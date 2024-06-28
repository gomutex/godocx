package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/wml/stypes"
)

func TestDocGrid_MarshalXML_AllAttributes(t *testing.T) {
	linePitch := 240
	charSpace := 120

	docGrid := DocGrid{
		Type:      stypes.DocGridLinesAndChars,
		LinePitch: &linePitch,
		CharSpace: &charSpace,
	}

	expectedXML := `<w:docGrid w:type="linesAndChars" w:linePitch="240" w:charSpace="120"></w:docGrid>`

	checkMarshalXML(t, docGrid, expectedXML)
}

func TestDocGrid_MarshalXML_OmitEmptyAttributes(t *testing.T) {
	docGrid := DocGrid{
		Type: stypes.DocGridLines,
	}

	expectedXML := `<w:docGrid w:type="lines"></w:docGrid>`

	checkMarshalXML(t, docGrid, expectedXML)
}

func TestDocGrid_MarshalXML_NoAttributes(t *testing.T) {
	docGrid := DocGrid{}

	expectedXML := `<w:docGrid></w:docGrid>`

	checkMarshalXML(t, docGrid, expectedXML)
}

func checkMarshalXML(t *testing.T, docGrid DocGrid, expectedXML string) {
	t.Helper()

	output, err := xml.Marshal(&docGrid)
	if err != nil {
		t.Fatalf("Error marshaling DocGrid: %v", err)
	}

	result := strings.TrimSpace(string(output))
	expected := strings.TrimSpace(expectedXML)

	if result != expected {
		t.Errorf("Expected XML:\n%s\nBut got:\n%s", expected, result)
	}
}

func TestDocGrid_UnmarshalXML_AllAttributes(t *testing.T) {
	xmlInput := `<w:docGrid xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main" w:type="linesAndChars" w:linePitch="240" w:charSpace="120"></w:docGrid>`

	linePitch := 240
	charSpace := 120
	expectedDocGrid := DocGrid{
		Type:      stypes.DocGridLinesAndChars,
		LinePitch: &linePitch,
		CharSpace: &charSpace,
	}

	checkUnmarshalXML(t, xmlInput, expectedDocGrid)
}

func TestDocGrid_UnmarshalXML_MinimalAttributes(t *testing.T) {
	xmlInput := `<w:docGrid xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main" w:type="lines"></w:docGrid>`

	expectedDocGrid := DocGrid{
		Type: stypes.DocGridLines,
	}

	checkUnmarshalXML(t, xmlInput, expectedDocGrid)
}

func TestDocGrid_UnmarshalXML_NoAttributes(t *testing.T) {
	xmlInput := `<w:docGrid xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main"></w:docGrid>`

	expectedDocGrid := DocGrid{}

	checkUnmarshalXML(t, xmlInput, expectedDocGrid)
}

func checkUnmarshalXML(t *testing.T, xmlInput string, expectedDocGrid DocGrid) {
	t.Helper()

	var docGrid DocGrid

	err := xml.Unmarshal([]byte(xmlInput), &docGrid)
	if err != nil {
		t.Fatalf("Error unmarshaling XML: %v", err)
	}

	if docGrid.Type != expectedDocGrid.Type {
		t.Errorf("Expected Type %s but got %s", expectedDocGrid.Type, docGrid.Type)
	}

	if expectedDocGrid.LinePitch == nil {
		if docGrid.LinePitch != nil {
			t.Errorf("Expected LinePitch to be nil but got %d", *docGrid.LinePitch)
		}
	} else {
		if docGrid.LinePitch == nil {
			t.Errorf("Expected LinePitch %d but got nil", *expectedDocGrid.LinePitch)
		} else if *docGrid.LinePitch != *expectedDocGrid.LinePitch {
			t.Errorf("Expected LinePitch %d but got %d", *expectedDocGrid.LinePitch, *docGrid.LinePitch)
		}
	}

	if expectedDocGrid.CharSpace == nil {
		if docGrid.CharSpace != nil {
			t.Errorf("Expected CharSpace to be nil but got %d", *docGrid.CharSpace)
		}
	} else {
		if docGrid.CharSpace == nil {
			t.Errorf("Expected CharSpace %d but got nil", *expectedDocGrid.CharSpace)
		} else if *docGrid.CharSpace != *expectedDocGrid.CharSpace {
			t.Errorf("Expected CharSpace %d but got %d", *expectedDocGrid.CharSpace, *docGrid.CharSpace)
		}
	}
}
