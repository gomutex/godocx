package txt

import "encoding/xml"

// Custom Tab Stop
type Tab struct {
	Val        *string   // Tab Stop Type
	Pos        *string   // Tab Stop Position
	LeaderChar *LeadChar // Custom Tab Stop Leader Character
}

type LeadChar string

const (
	LeadCharNone       LeadChar = "none"
	LeadCharDot        LeadChar = "dot"
	LeadCharHypen      LeadChar = "hyphen"
	LeadCharUnderScore LeadChar = "underscore"
	LeadCharHeavy      LeadChar = "heavy"
	LeadCharMiddleDot  LeadChar = "middleDot"
	LeadCharInvalid    LeadChar = ""
)

func LeadCharFromStr(val string) (LeadChar, error) {
	switch val {
	case "none":
		return LeadCharNone, nil
	case "dot":
		return LeadCharDot, nil
	case "hyphen":
		return LeadCharHypen, nil
	case "underscore":
		return LeadCharUnderScore, nil
	case "heavy":
		return LeadCharHeavy, nil
	case "middleDot":
		return LeadCharMiddleDot, nil
	default:
		return LeadCharInvalid, nil
	}
}

func (t *Tab) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:tab"
	start.Attr = []xml.Attr{}

	if t.Val != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"}, Value: *t.Val})
	}

	if t.Pos != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "pos"}, Value: *t.Pos})
	}

	if t.LeaderChar != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "leader"}, Value: string(*t.LeaderChar)})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (t *Tab) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		switch a.Name.Local {
		case "val":
			t.Val = &a.Value
		case "pos":
			t.Pos = &a.Value
		case "leader":
			lc, err := LeadCharFromStr(a.Value)
			if err != nil {
				return err
			}
			t.LeaderChar = &lc
		}
	}

	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			}
		case xml.EndElement:
			if elem == start.End() {
				return nil
			}
		}
	}
}
