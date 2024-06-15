package txt

import (
	"encoding/xml"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestSzMarshalXML(t *testing.T) {
	testCases := []struct {
		input    *Sz
		expected string
	}{
		{NewSz(12), `<w:sz w:val="12"></w:sz>`},
		{NewSz(0), ``}, // Empty because value is 0
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

func TestSzUnmarshalXML(t *testing.T) {
	testCases := []struct {
		input    string
		expected *Sz
	}{
		{`<w:sz w:val="14"></w:sz>`, NewSz(14)},
		{`<w:sz></w:sz>`, NewSz(0)}, // Default value when not specified
	}

	for _, tc := range testCases {
		decoder := xml.NewDecoder(strings.NewReader(tc.input))
		var result Sz

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

	invalidXML := `<w:sz w:val="-1"></w:sz>`
	decoder := xml.NewDecoder(strings.NewReader(invalidXML))
	var result Sz

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
