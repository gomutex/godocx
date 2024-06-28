package ctypes

import (
	"bytes"
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/gomutex/godocx/internal"
)

func TestCellMerge_MarshalXML(t *testing.T) {
	// Create a CellMerge instance for testing
	cellMerge := &CellMerge{
		ID:         1,
		Author:     "John Doe",
		Date:       nil,
		VMerge:     internal.ToPtr(AnnotationVMergeCont),
		VMergeOrig: internal.ToPtr(AnnotationVMergeRest),
	}

	// Marshal the CellMerge instance to XML
	xmlData, err := xml.Marshal(cellMerge)
	if err != nil {
		t.Fatalf("Error marshaling CellMerge to XML: %v", err)
	}

	// Define the expected XML string
	expectedXMLString := `<CellMerge w:id="1" w:author="John Doe" w:vMerge="cont" w:vMergeOrig="rest"></CellMerge>`

	// Compare the generated XML with the expected XML string
	if !bytes.Contains(xmlData, []byte(expectedXMLString)) {
		t.Errorf("Expected XML string %s, got %s", expectedXMLString, string(xmlData))
	}
}

func TestCellMerge_UnmarshalXML(t *testing.T) {
	// Define a sample XML data corresponding to CellMerge structure
	xmlData := []byte(`<CellMerge w:id="2" w:author="Jane Smith" w:date="2024-06-25" w:vMerge="rest"></CellMerge>`)

	// Define the expected CellMerge instance after unmarshaling
	expectedCellMerge := &CellMerge{
		ID:         2,
		Author:     "Jane Smith",
		Date:       xmlStrPtr("2024-06-25"), // Helper function to get pointer to string
		VMerge:     internal.ToPtr(AnnotationVMergeRest),
		VMergeOrig: nil,
	}

	// Variable to hold unmarshaled CellMerge instance
	var unmarshaledCellMerge CellMerge

	// Unmarshal the XML into the CellMerge instance
	err := xml.Unmarshal(xmlData, &unmarshaledCellMerge)
	if err != nil {
		t.Fatalf("Error unmarshaling XML to CellMerge: %v", err)
	}

	// Compare the unmarshaled CellMerge instance with the expected instance
	if !reflect.DeepEqual(&unmarshaledCellMerge, expectedCellMerge) {
		t.Errorf("Expected %#v, got %#v", expectedCellMerge, unmarshaledCellMerge)
	}
}

// Helper function to return pointer to string
func xmlStrPtr(s string) *string {
	return &s
}
