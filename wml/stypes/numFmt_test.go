package stypes

import (
	"encoding/xml"
	"testing"
)

func TestNumFmtFromStr_ValidValues(t *testing.T) {
	tests := []struct {
		input    string
		expected NumFmt
	}{
		{"decimal", NumFmtDecimal},
		{"upperRoman", NumFmtUpperRoman},
		{"lowerRoman", NumFmtLowerRoman},
		{"upperLetter", NumFmtUpperLetter},
		{"lowerLetter", NumFmtLowerLetter},
		{"ordinal", NumFmtOrdinal},
		{"cardinalText", NumFmtCardinalText},
		{"ordinalText", NumFmtOrdinalText},
		{"hex", NumFmtHex},
		{"chicago", NumFmtChicago},
		{"ideographDigital", NumFmtIdeographDigital},
		{"japaneseCounting", NumFmtJapaneseCounting},
		{"aiueo", NumFmtAIUEO},
		{"iroha", NumFmtIroha},
		{"decimalFullWidth", NumFmtDecimalFullWidth},
		{"decimalHalfWidth", NumFmtDecimalHalfWidth},
		{"japaneseLegal", NumFmtJapaneseLegal},
		{"japaneseDigitalTenThousand", NumFmtJapaneseDigitalTenThousand},
		{"decimalEnclosedCircle", NumFmtDecimalEnclosedCircle},
		{"decimalFullWidth2", NumFmtDecimalFullWidth2},
		{"aiueoFullWidth", NumFmtAIUEOFullWidth},
		{"irohaFullWidth", NumFmtIrohaFullWidth},
		{"decimalZero", NumFmtDecimalZero},
		{"bullet", NumFmtBullet},
		{"ganada", NumFmtGanada},
		{"chosung", NumFmtChosung},
		{"decimalEnclosedFullstop", NumFmtDecimalEnclosedFullstop},
		{"decimalEnclosedParen", NumFmtDecimalEnclosedParen},
		{"decimalEnclosedCircleChinese", NumFmtDecimalEnclosedCircleChinese},
		{"ideographEnclosedCircle", NumFmtIdeographEnclosedCircle},
		{"ideographTraditional", NumFmtIdeographTraditional},
		{"ideographZodiac", NumFmtIdeographZodiac},
		{"ideographZodiacTraditional", NumFmtIdeographZodiacTraditional},
		{"taiwaneseCounting", NumFmtTaiwaneseCounting},
		{"ideographLegalTraditional", NumFmtIdeographLegalTraditional},
		{"taiwaneseCountingThousand", NumFmtTaiwaneseCountingThousand},
		{"taiwaneseDigital", NumFmtTaiwaneseDigital},
		{"chineseCounting", NumFmtChineseCounting},
		{"chineseLegalSimplified", NumFmtChineseLegalSimplified},
		{"chineseCountingThousand", NumFmtChineseCountingThousand},
		{"koreanDigital", NumFmtKoreanDigital},
		{"koreanCounting", NumFmtKoreanCounting},
		{"koreanLegal", NumFmtKoreanLegal},
		{"koreanDigital2", NumFmtKoreanDigital2},
		{"vietnameseCounting", NumFmtVietnameseCounting},
		{"russianLower", NumFmtRussianLower},
		{"russianUpper", NumFmtRussianUpper},
		{"none", NumFmtNone},
		{"numberInDash", NumFmtNumberInDash},
		{"hebrew1", NumFmtHebrew1},
		{"hebrew2", NumFmtHebrew2},
		{"arabicAlpha", NumFmtArabicAlpha},
		{"arabicAbjad", NumFmtArabicAbjad},
		{"hindiVowels", NumFmtHindiVowels},
		{"hindiConsonants", NumFmtHindiConsonants},
		{"hindiNumbers", NumFmtHindiNumbers},
		{"hindiCounting", NumFmtHindiCounting},
		{"thaiLetters", NumFmtThaiLetters},
		{"thaiNumbers", NumFmtThaiNumbers},
		{"thaiCounting", NumFmtThaiCounting},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := NumFmtFromStr(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, result)
			}
		})
	}
}

func TestNumFmtFromStr_InvalidValue(t *testing.T) {
	input := "invalidValue"

	result, err := NumFmtFromStr(input)

	if err == nil {
		t.Fatalf("Expected error for invalid value %s, but got none. Result: %s", input, result)
	}

	expectedError := "Invalid Numbering Format"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}

func TestNumFmt_UnmarshalXMLAttr_ValidValues(t *testing.T) {
	tests := []struct {
		inputXML string
		expected NumFmt
	}{
		{`<element numFmt="decimal"></element>`, NumFmtDecimal},
		{`<element numFmt="upperRoman"></element>`, NumFmtUpperRoman},
		{`<element numFmt="lowerRoman"></element>`, NumFmtLowerRoman},
		{`<element numFmt="upperLetter"></element>`, NumFmtUpperLetter},
		{`<element numFmt="lowerLetter"></element>`, NumFmtLowerLetter},
		// Add more valid values as needed
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			type Element struct {
				XMLName xml.Name `xml:"element"`
				NumFmt  NumFmt   `xml:"numFmt,attr"`
			}

			var elem Element

			err := xml.Unmarshal([]byte(tt.inputXML), &elem)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if elem.NumFmt != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, elem.NumFmt)
			}
		})
	}
}

func TestNumFmt_UnmarshalXMLAttr_InvalidValue(t *testing.T) {
	inputXML := `<element numFmt="invalidValue"></element>`

	type Element struct {
		XMLName xml.Name `xml:"element"`
		NumFmt  NumFmt   `xml:"numFmt,attr"`
	}

	var elem Element

	err := xml.Unmarshal([]byte(inputXML), &elem)

	if err == nil {
		t.Fatalf("Expected error for invalid value, but got none")
	}

	expectedError := "Invalid Numbering Format"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s' but got '%s'", expectedError, err.Error())
	}
}
