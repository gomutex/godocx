package stypes

import (
	"encoding/xml"
	"errors"
)

type Justification string

const (
	JustificationLeft           Justification = "left"           // Align Left
	JustificationCenter         Justification = "center"         // Align Center
	JustificationRight          Justification = "right"          // Align Right
	JustificationBoth           Justification = "both"           // Justified
	JustificationMediumKashida  Justification = "mediumKashida"  // Medium Kashida Length
	JustificationDistribute     Justification = "distribute"     // Distribute All Characters Equally
	JustificationNumTab         Justification = "numTab"         // Align to List Tab
	JustificationHighKashida    Justification = "highKashida"    // Widest Kashida Length
	JustificationLowKashida     Justification = "lowKashida"     // Low Kashida Length
	JustificationThaiDistribute Justification = "thaiDistribute" // Thai Language Justification
)

func JustificationFromStr(value string) (Justification, error) {
	switch value {
	case "left":
		return JustificationLeft, nil
	case "center":
		return JustificationCenter, nil
	case "right":
		return JustificationRight, nil
	case "both":
		return JustificationBoth, nil
	case "mediumKashida":
		return JustificationMediumKashida, nil
	case "distribute":
		return JustificationDistribute, nil
	case "numTab":
		return JustificationNumTab, nil
	case "highKashida":
		return JustificationHighKashida, nil
	case "lowKashida":
		return JustificationLowKashida, nil
	case "thaiDistribute":
		return JustificationThaiDistribute, nil
	default:
		return "", errors.New("invalid Justification value")
	}
}

func (j *Justification) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := JustificationFromStr(attr.Value)
	if err != nil {
		return err
	}

	*j = val

	return nil
}
