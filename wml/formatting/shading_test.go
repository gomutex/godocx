package formatting

import (
	"encoding/xml"
	"fmt"
	"strings"
	"testing"
)

func TestShadingMarshalXML(t *testing.T) {
	shading := DefaultShading()

	expectedEmpty := fmt.Sprintf(`<w:shd w:fill="%s" w:val="%s"></w:shd>`, shading.Fill, shading.ShdType)
	testMarshalXML(t, shading, expectedEmpty)

	shading.SetColor("red").
		SetFill("00FF00").
		SetShadingType("DiagCross")
	expectedFilled := `<w:shd w:fill="00FF00" w:val="DiagCross"></w:shd>`
	testMarshalXML(t, shading, expectedFilled)
}

func testMarshalXML(t *testing.T, shading *Shading, expected string) {
	var result strings.Builder
	encoder := xml.NewEncoder(&result)

	start := xml.StartElement{Name: xml.Name{Local: "fake"}}
	if err := shading.MarshalXML(encoder, start); err != nil {
		t.Errorf("Error during MarshalXML: %v", err)
	}

	err := encoder.Flush()
	if err != nil {
		t.Errorf("Error flushing encoder: %v", err)
	}

	if result.String() != expected {
		t.Errorf("Expected XML:\n%s\n\nActual XML:\n%s", expected, result.String())
	}
}
