package stypes

import (
	"encoding/xml"
	"errors"
)

// Shading represents the possible values for Shading Pattern in WordprocessingML.
type Shading string

// List of valid Shading values.
const (
	ShdNil                   Shading = "nil"
	ShdClear                 Shading = "clear"
	ShdSolid                 Shading = "solid"
	ShdHorzStripe            Shading = "horzStripe"
	ShdVertStripe            Shading = "vertStripe"
	ShdReverseDiagStripe     Shading = "reverseDiagStripe"
	ShdDiagStripe            Shading = "diagStripe"
	ShdHorzCross             Shading = "horzCross"
	ShdDiagCross             Shading = "diagCross"
	ShdThinHorzStripe        Shading = "thinHorzStripe"
	ShdThinVertStripe        Shading = "thinVertStripe"
	ShdThinReverseDiagStripe Shading = "thinReverseDiagStripe"
	ShdThinDiagStripe        Shading = "thinDiagStripe"
	ShdThinHorzCross         Shading = "thinHorzCross"
	ShdThinDiagCross         Shading = "thinDiagCross"
	ShdPct5                  Shading = "pct5"
	ShdPct10                 Shading = "pct10"
	ShdPct12                 Shading = "pct12"
	ShdPct15                 Shading = "pct15"
	ShdPct20                 Shading = "pct20"
	ShdPct25                 Shading = "pct25"
	ShdPct30                 Shading = "pct30"
	ShdPct35                 Shading = "pct35"
	ShdPct37                 Shading = "pct37"
	ShdPct40                 Shading = "pct40"
	ShdPct45                 Shading = "pct45"
	ShdPct50                 Shading = "pct50"
	ShdPct55                 Shading = "pct55"
	ShdPct60                 Shading = "pct60"
	ShdPct62                 Shading = "pct62"
	ShdPct65                 Shading = "pct65"
	ShdPct70                 Shading = "pct70"
	ShdPct75                 Shading = "pct75"
	ShdPct80                 Shading = "pct80"
	ShdPct85                 Shading = "pct85"
	ShdPct87                 Shading = "pct87"
	ShdPct90                 Shading = "pct90"
	ShdPct95                 Shading = "pct95"
)

// ShdFromStr converts a string value to Shd.
func ShadingFromStr(value string) (Shading, error) {
	switch value {
	case "nil":
		return ShdNil, nil
	case "clear":
		return ShdClear, nil
	case "solid":
		return ShdSolid, nil
	case "horzStripe":
		return ShdHorzStripe, nil
	case "vertStripe":
		return ShdVertStripe, nil
	case "reverseDiagStripe":
		return ShdReverseDiagStripe, nil
	case "diagStripe":
		return ShdDiagStripe, nil
	case "horzCross":
		return ShdHorzCross, nil
	case "diagCross":
		return ShdDiagCross, nil
	case "thinHorzStripe":
		return ShdThinHorzStripe, nil
	case "thinVertStripe":
		return ShdThinVertStripe, nil
	case "thinReverseDiagStripe":
		return ShdThinReverseDiagStripe, nil
	case "thinDiagStripe":
		return ShdThinDiagStripe, nil
	case "thinHorzCross":
		return ShdThinHorzCross, nil
	case "thinDiagCross":
		return ShdThinDiagCross, nil
	case "pct5":
		return ShdPct5, nil
	case "pct10":
		return ShdPct10, nil
	case "pct12":
		return ShdPct12, nil
	case "pct15":
		return ShdPct15, nil
	case "pct20":
		return ShdPct20, nil
	case "pct25":
		return ShdPct25, nil
	case "pct30":
		return ShdPct30, nil
	case "pct35":
		return ShdPct35, nil
	case "pct37":
		return ShdPct37, nil
	case "pct40":
		return ShdPct40, nil
	case "pct45":
		return ShdPct45, nil
	case "pct50":
		return ShdPct50, nil
	case "pct55":
		return ShdPct55, nil
	case "pct60":
		return ShdPct60, nil
	case "pct62":
		return ShdPct62, nil
	case "pct65":
		return ShdPct65, nil
	case "pct70":
		return ShdPct70, nil
	case "pct75":
		return ShdPct75, nil
	case "pct80":
		return ShdPct80, nil
	case "pct85":
		return ShdPct85, nil
	case "pct87":
		return ShdPct87, nil
	case "pct90":
		return ShdPct90, nil
	case "pct95":
		return ShdPct95, nil
	default:
		return "", errors.New("invalid shading value")
	}
}

// UnmarshalXMLAttr unmarshals an XML attribute into a Shd.
func (s *Shading) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := ShadingFromStr(attr.Value)
	if err != nil {
		return err
	}

	*s = val

	return nil
}
