package runcontent

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/internal"
)

func TestTextMarshalXML(t *testing.T) {
	testCases := []struct {
		input    *Text
		expected string
	}{
		{NewText(), `<w:t></w:t>`},
		{TextFromString("Hello, World!"), `<w:t>Hello, World!</w:t>`},
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
		{`<w:t xml:space="preserve">Hello, World!</w:t>`, &Text{
			text:  "Hello, World!",
			space: internal.ToPtr("preserve"),
		}},
		{`<w:t xml:space="preserve">Some text</w:t>`, &Text{
			text:  "Some text",
			space: internal.ToPtr("preserve"),
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			decoder := xml.NewDecoder(strings.NewReader(tc.input))
			var result Text

			startToken, err := decoder.Token()
			if err != nil {
				t.Fatalf("Error getting start token: %v", err)
			}

			err = result.UnmarshalXML(decoder, startToken.(xml.StartElement))
			if err != nil {
				t.Fatalf("Error during UnmarshalXML: %v", err)
			}

			if result.text != tc.expected.text {
				t.Errorf("Expected text %q, but got %q", tc.expected.text, result.text)
			}

			if err := internal.ComparePtr("space", tc.expected.space, result.space); err != nil {
				t.Errorf(err.Error())
			}
		})
	}
}
