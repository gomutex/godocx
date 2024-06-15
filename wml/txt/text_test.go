package txt

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"
)

func TestTextMarshalXML(t *testing.T) {
	testCases := []struct {
		input    *Text
		expected string
	}{
		{NewText(), `<w:t></w:t>`},
		{TextFromString("Hello, World!"), `<w:t xml:space="preserve">Hello, World!</w:t>`},
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

func TestTextUnmarshalXML(t *testing.T) {
	testCases := []struct {
		input    string
		expected *Text
	}{
		{`<w:t></w:t>`, NewText()},
		{`<w:t xml:space="preserve">Hello, World!</w:t>`, TextFromString("Hello, World!")},
		{`<w:t xml:space="preserve">Some text</w:t><w:other>Other element</w:other>`, TextFromString("Some text")},
	}

	for _, tc := range testCases {
		decoder := xml.NewDecoder(strings.NewReader(tc.input))
		var result Text

		startToken, err := decoder.Token()
		if err != nil {
			t.Errorf("Error getting start token: %v", err)
		}

		err = result.UnmarshalXML(decoder, startToken.(xml.StartElement))
		if err != nil {
			t.Errorf("Error during UnmarshalXML: %v", err)
		}

		if !reflect.DeepEqual(&result, tc.expected) {
			t.Errorf("Expected %+v, but got %+v", tc.expected, &result)
		}
	}
}
