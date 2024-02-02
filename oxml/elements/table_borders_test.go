package elements

import (
	"encoding/xml"
	"testing"
)

func TestTableBorderElementMarshaling(t *testing.T) {
	tblBorder := DefaulTableBorderElement()
	tblBorder.BorderType = BorderTypeSingle
	tblBorder.Color = "000000"
	tblBorder.Position = TableBorderPositionTop
	tblBorder.Space = 0
	tblBorder.Size = 2
	xmlData, err := xml.Marshal(tblBorder)
	if err != nil {
		t.Fatalf("Error marshaling TableBorders to XML: %v", err)
	}

	var unmarshalledTableBorder TableBorderElement
	err = xml.Unmarshal(xmlData, &unmarshalledTableBorder)
	if err != nil {
		t.Fatalf("Error unmarshaling XML to TableBorder: %v", err)
	}

	if unmarshalledTableBorder.Position != tblBorder.Position {
		t.Errorf("Expected TableBorder Position value %s, got %s", tblBorder.Position, unmarshalledTableBorder.Position)
	}

	if unmarshalledTableBorder.BorderType != tblBorder.BorderType {
		t.Errorf("Expected TableBorder BorderType value %s, got %s", tblBorder.BorderType, unmarshalledTableBorder.BorderType)
	}

	if unmarshalledTableBorder.Color != tblBorder.Color {
		t.Errorf("Expected TableBorder Color value %s, got %s", tblBorder.Color, unmarshalledTableBorder.Color)
	}

	if unmarshalledTableBorder.Size != tblBorder.Size {
		t.Errorf("Expected TableBorder Size value %v, got %v", tblBorder.Size, unmarshalledTableBorder.Size)
	}

}

func TestTableBordersMarshaling(t *testing.T) {
	tblBorders := DefaultTableBorders()
	tblBorders.Bottom = DefaulTableBorderElement()
	xmlData, err := xml.Marshal(tblBorders)
	if err != nil {
		t.Fatalf("Error marshaling TableBorders to XML: %v", err)
	}

	var unmarshalledTableBorders TableBorders
	err = xml.Unmarshal(xmlData, &unmarshalledTableBorders)
	if err != nil {
		t.Fatalf("Error unmarshaling XML to TableBorders: %v", err)
	}

	if unmarshalledTableBorders.Bottom == nil {
		t.Fatalf("Error unmarshaling XML to TableBorders: Top got nil value")
	}

	if unmarshalledTableBorders.Bottom.Position != tblBorders.Bottom.Position {
		t.Errorf("Expected TableBorder Position value %s, got %s", tblBorders.Bottom.Position, unmarshalledTableBorders.Bottom.Position)
	}

	if unmarshalledTableBorders.Bottom.BorderType != tblBorders.Bottom.BorderType {
		t.Errorf("Expected TableBorder BorderType value %s, got %s", tblBorders.Bottom.BorderType, unmarshalledTableBorders.Bottom.BorderType)
	}

	if unmarshalledTableBorders.Bottom.Color != tblBorders.Bottom.Color {
		t.Errorf("Expected TableBorder Color value %s, got %s", tblBorders.Bottom.Color, unmarshalledTableBorders.Bottom.Color)
	}

	if unmarshalledTableBorders.Bottom.Size != tblBorders.Bottom.Size {
		t.Errorf("Expected TableBorder Size value %v, got %v", tblBorders.Bottom.Size, unmarshalledTableBorders.Bottom.Size)
	}

}
