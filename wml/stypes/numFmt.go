package stypes

import (
	"encoding/xml"
	"errors"
)

type NumFmt string

const (
	NumFmtDecimal                      NumFmt = "decimal"
	NumFmtUpperRoman                   NumFmt = "upperRoman"
	NumFmtLowerRoman                   NumFmt = "lowerRoman"
	NumFmtUpperLetter                  NumFmt = "upperLetter"
	NumFmtLowerLetter                  NumFmt = "lowerLetter"
	NumFmtOrdinal                      NumFmt = "ordinal"
	NumFmtCardinalText                 NumFmt = "cardinalText"
	NumFmtOrdinalText                  NumFmt = "ordinalText"
	NumFmtHex                          NumFmt = "hex"
	NumFmtChicago                      NumFmt = "chicago"
	NumFmtIdeographDigital             NumFmt = "ideographDigital"
	NumFmtJapaneseCounting             NumFmt = "japaneseCounting"
	NumFmtAIUEO                        NumFmt = "aiueo"
	NumFmtIroha                        NumFmt = "iroha"
	NumFmtDecimalFullWidth             NumFmt = "decimalFullWidth"
	NumFmtDecimalHalfWidth             NumFmt = "decimalHalfWidth"
	NumFmtJapaneseLegal                NumFmt = "japaneseLegal"
	NumFmtJapaneseDigitalTenThousand   NumFmt = "japaneseDigitalTenThousand"
	NumFmtDecimalEnclosedCircle        NumFmt = "decimalEnclosedCircle"
	NumFmtDecimalFullWidth2            NumFmt = "decimalFullWidth2"
	NumFmtAIUEOFullWidth               NumFmt = "aiueoFullWidth"
	NumFmtIrohaFullWidth               NumFmt = "irohaFullWidth"
	NumFmtDecimalZero                  NumFmt = "decimalZero"
	NumFmtBullet                       NumFmt = "bullet"
	NumFmtGanada                       NumFmt = "ganada"
	NumFmtChosung                      NumFmt = "chosung"
	NumFmtDecimalEnclosedFullstop      NumFmt = "decimalEnclosedFullstop"
	NumFmtDecimalEnclosedParen         NumFmt = "decimalEnclosedParen"
	NumFmtDecimalEnclosedCircleChinese NumFmt = "decimalEnclosedCircleChinese"
	NumFmtIdeographEnclosedCircle      NumFmt = "ideographEnclosedCircle"
	NumFmtIdeographTraditional         NumFmt = "ideographTraditional"
	NumFmtIdeographZodiac              NumFmt = "ideographZodiac"
	NumFmtIdeographZodiacTraditional   NumFmt = "ideographZodiacTraditional"
	NumFmtTaiwaneseCounting            NumFmt = "taiwaneseCounting"
	NumFmtIdeographLegalTraditional    NumFmt = "ideographLegalTraditional"
	NumFmtTaiwaneseCountingThousand    NumFmt = "taiwaneseCountingThousand"
	NumFmtTaiwaneseDigital             NumFmt = "taiwaneseDigital"
	NumFmtChineseCounting              NumFmt = "chineseCounting"
	NumFmtChineseLegalSimplified       NumFmt = "chineseLegalSimplified"
	NumFmtChineseCountingThousand      NumFmt = "chineseCountingThousand"
	NumFmtKoreanDigital                NumFmt = "koreanDigital"
	NumFmtKoreanCounting               NumFmt = "koreanCounting"
	NumFmtKoreanLegal                  NumFmt = "koreanLegal"
	NumFmtKoreanDigital2               NumFmt = "koreanDigital2"
	NumFmtVietnameseCounting           NumFmt = "vietnameseCounting"
	NumFmtRussianLower                 NumFmt = "russianLower"
	NumFmtRussianUpper                 NumFmt = "russianUpper"
	NumFmtNone                         NumFmt = "none"
	NumFmtNumberInDash                 NumFmt = "numberInDash"
	NumFmtHebrew1                      NumFmt = "hebrew1"
	NumFmtHebrew2                      NumFmt = "hebrew2"
	NumFmtArabicAlpha                  NumFmt = "arabicAlpha"
	NumFmtArabicAbjad                  NumFmt = "arabicAbjad"
	NumFmtHindiVowels                  NumFmt = "hindiVowels"
	NumFmtHindiConsonants              NumFmt = "hindiConsonants"
	NumFmtHindiNumbers                 NumFmt = "hindiNumbers"
	NumFmtHindiCounting                NumFmt = "hindiCounting"
	NumFmtThaiLetters                  NumFmt = "thaiLetters"
	NumFmtThaiNumbers                  NumFmt = "thaiNumbers"
	NumFmtThaiCounting                 NumFmt = "thaiCounting"
)

// NumFmtFromStr converts a string value to NumFmt type.
func NumFmtFromStr(value string) (NumFmt, error) {
	switch value {
	case "decimal":
		return NumFmtDecimal, nil
	case "upperRoman":
		return NumFmtUpperRoman, nil
	case "lowerRoman":
		return NumFmtLowerRoman, nil
	case "upperLetter":
		return NumFmtUpperLetter, nil
	case "lowerLetter":
		return NumFmtLowerLetter, nil
	case "ordinal":
		return NumFmtOrdinal, nil
	case "cardinalText":
		return NumFmtCardinalText, nil
	case "ordinalText":
		return NumFmtOrdinalText, nil
	case "hex":
		return NumFmtHex, nil
	case "chicago":
		return NumFmtChicago, nil
	case "ideographDigital":
		return NumFmtIdeographDigital, nil
	case "japaneseCounting":
		return NumFmtJapaneseCounting, nil
	case "aiueo":
		return NumFmtAIUEO, nil
	case "iroha":
		return NumFmtIroha, nil
	case "decimalFullWidth":
		return NumFmtDecimalFullWidth, nil
	case "decimalHalfWidth":
		return NumFmtDecimalHalfWidth, nil
	case "japaneseLegal":
		return NumFmtJapaneseLegal, nil
	case "japaneseDigitalTenThousand":
		return NumFmtJapaneseDigitalTenThousand, nil
	case "decimalEnclosedCircle":
		return NumFmtDecimalEnclosedCircle, nil
	case "decimalFullWidth2":
		return NumFmtDecimalFullWidth2, nil
	case "aiueoFullWidth":
		return NumFmtAIUEOFullWidth, nil
	case "irohaFullWidth":
		return NumFmtIrohaFullWidth, nil
	case "decimalZero":
		return NumFmtDecimalZero, nil
	case "bullet":
		return NumFmtBullet, nil
	case "ganada":
		return NumFmtGanada, nil
	case "chosung":
		return NumFmtChosung, nil
	case "decimalEnclosedFullstop":
		return NumFmtDecimalEnclosedFullstop, nil
	case "decimalEnclosedParen":
		return NumFmtDecimalEnclosedParen, nil
	case "decimalEnclosedCircleChinese":
		return NumFmtDecimalEnclosedCircleChinese, nil
	case "ideographEnclosedCircle":
		return NumFmtIdeographEnclosedCircle, nil
	case "ideographTraditional":
		return NumFmtIdeographTraditional, nil
	case "ideographZodiac":
		return NumFmtIdeographZodiac, nil
	case "ideographZodiacTraditional":
		return NumFmtIdeographZodiacTraditional, nil
	case "taiwaneseCounting":
		return NumFmtTaiwaneseCounting, nil
	case "ideographLegalTraditional":
		return NumFmtIdeographLegalTraditional, nil
	case "taiwaneseCountingThousand":
		return NumFmtTaiwaneseCountingThousand, nil
	case "taiwaneseDigital":
		return NumFmtTaiwaneseDigital, nil
	case "chineseCounting":
		return NumFmtChineseCounting, nil
	case "chineseLegalSimplified":
		return NumFmtChineseLegalSimplified, nil
	case "chineseCountingThousand":
		return NumFmtChineseCountingThousand, nil
	case "koreanDigital":
		return NumFmtKoreanDigital, nil
	case "koreanCounting":
		return NumFmtKoreanCounting, nil
	case "koreanLegal":
		return NumFmtKoreanLegal, nil
	case "koreanDigital2":
		return NumFmtKoreanDigital2, nil
	case "vietnameseCounting":
		return NumFmtVietnameseCounting, nil
	case "russianLower":
		return NumFmtRussianLower, nil
	case "russianUpper":
		return NumFmtRussianUpper, nil
	case "none":
		return NumFmtNone, nil
	case "numberInDash":
		return NumFmtNumberInDash, nil
	case "hebrew1":
		return NumFmtHebrew1, nil
	case "hebrew2":
		return NumFmtHebrew2, nil
	case "arabicAlpha":
		return NumFmtArabicAlpha, nil
	case "arabicAbjad":
		return NumFmtArabicAbjad, nil
	case "hindiVowels":
		return NumFmtHindiVowels, nil
	case "hindiConsonants":
		return NumFmtHindiConsonants, nil
	case "hindiNumbers":
		return NumFmtHindiNumbers, nil
	case "hindiCounting":
		return NumFmtHindiCounting, nil
	case "thaiLetters":
		return NumFmtThaiLetters, nil
	case "thaiNumbers":
		return NumFmtThaiNumbers, nil
	case "thaiCounting":
		return NumFmtThaiCounting, nil
	default:
		return "", errors.New("Invalid Numbering Format")
	}
}

func (d *NumFmt) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := NumFmtFromStr(attr.Value)
	if err != nil {
		return err
	}

	*d = val

	return nil

}
