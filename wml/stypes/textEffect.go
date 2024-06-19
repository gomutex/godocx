package stypes

import (
	"encoding/xml"
	"errors"
)

// TextEffect represents the possible values for text animation effects.
type TextEffect string

const (
	TextEffectBlinkBackground TextEffect = "blinkBackground" // Blinking Background Animation
	TextEffectLights          TextEffect = "lights"          // Colored Lights Animation
	TextEffectAntsBlack       TextEffect = "antsBlack"       // Black Dashed Line Animation
	TextEffectAntsRed         TextEffect = "antsRed"         // Marching Red Ants
	TextEffectShimmer         TextEffect = "shimmer"         // Shimmer Animation
	TextEffectSparkle         TextEffect = "sparkle"         // Sparkling Lights Animation
	TextEffectNone            TextEffect = "none"            // No Animation
)

// TextEffectFromStr converts a string to a TextEffect.
func TextEffectFromStr(value string) (TextEffect, error) {
	switch value {
	case "blinkBackground":
		return TextEffectBlinkBackground, nil
	case "lights":
		return TextEffectLights, nil
	case "antsBlack":
		return TextEffectAntsBlack, nil
	case "antsRed":
		return TextEffectAntsRed, nil
	case "shimmer":
		return TextEffectShimmer, nil
	case "sparkle":
		return TextEffectSparkle, nil
	case "none":
		return TextEffectNone, nil
	default:
		return "", errors.New("invalid TextEffect value")
	}
}

// UnmarshalXMLAttr unmarshals an XML attribute into a TextEffect.
func (t *TextEffect) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := TextEffectFromStr(attr.Value)
	if err != nil {
		return err
	}

	*t = val

	return nil
}
