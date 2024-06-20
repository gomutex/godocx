package dmlct

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"
)

func TestNewPoint2D(t *testing.T) {
	x := uint64(100)
	y := uint64(200)
	start := NewPoint2D(x, y)

	if start.XAxis != x {
		t.Errorf("NewPoint2D() failed: Expected XAxis %d, got %d", x, start.XAxis)
	}

	if start.YAxis != y {
		t.Errorf("NewPoint2D() failed: Expected YAxis %d, got %d", y, start.YAxis)
	}
}

func TestMarshalPoint2D(t *testing.T) {
	start := &Point2D{
		XAxis: 100,
		YAxis: 200,
	}

	var result strings.Builder
	encoder := xml.NewEncoder(&result)

	startElement := xml.StartElement{Name: xml.Name{Local: "wp:start"}}
	err := start.MarshalXML(encoder, startElement)
	if err != nil {
		t.Fatalf("MarshalXML() failed: %v", err)
	}

	err = encoder.Flush()
	if err != nil {
		t.Fatalf("Flush() failed: %v", err)
	}

	expectedXML := `<wp:start x="100" y="200"></wp:start>`
	if result.String() != expectedXML {
		t.Errorf("MarshalXML() failed: Expected XML:\n%s\nBut got:\n%s", expectedXML, result.String())
	}
}

func TestUnmarshalPoint2D(t *testing.T) {
	tests := []struct {
		inputXML        string
		expectedPoint2D Point2D
	}{
		{
			inputXML: `<wp:start x="100" y="200"></wp:start>`,
			expectedPoint2D: Point2D{
				XAxis: 100,
				YAxis: 200,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			var start Point2D

			err := xml.Unmarshal([]byte(tt.inputXML), &start)
			if err != nil {
				t.Fatalf("Unmarshal() failed: %v", err)
			}

			if !reflect.DeepEqual(start, tt.expectedPoint2D) {
				t.Errorf("Unmarshal() failed: Expected %+v, got %+v", tt.expectedPoint2D, start)
			}
		})
	}
}
