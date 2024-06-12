package wml

import (
	"encoding/xml"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestSzCsMarshalXML(t *testing.T) {
	testCases := []struct {
		input    *SzCs
		expected string
	}{
		{NewSzCs(12), `<w:szCs w:val="12"></w:szCs>`},
		{NewSzCs(0), ``}, // Empty because value is 0
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

func TestSzCsUnmarshalXML(t *testing.T) {
	testCases := []struct {
		input    string
		expected *SzCs
	}{
		{`<w:szCs w:val="14"></w:szCs>`, NewSzCs(14)},
		{`<w:szCs></w:szCs>`, NewSzCs(0)}, // Default value when not specified
	}

	for _, tc := range testCases {
		decoder := xml.NewDecoder(strings.NewReader(tc.input))
		var result SzCs

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

	invalidXML := `<w:szCs w:val="-1"></w:szCs>`
	decoder := xml.NewDecoder(strings.NewReader(invalidXML))
	var result SzCs

	startToken, err := decoder.Token()
	if err != nil {
		t.Errorf("Error getting start token: %v", err)
	}

	err = result.UnmarshalXML(decoder, startToken.(xml.StartElement))
	if err != nil {
		if numErr, ok := err.(*strconv.NumError); ok {
			if numErr.Err == strconv.ErrSyntax && strings.Contains(numErr.Num, "-") {
				t.Logf("Parsing error: Cannot parse negative value: %v", numErr.Num)
			} else {
				t.Errorf("Error during UnmarshalXML: %v", err)
			}

		} else {
			t.Errorf("Error during UnmarshalXML: %v", err)
		}
	} else {
		t.Errorf("Expected error for InvalidXML: %v", invalidXML)
	}
}
