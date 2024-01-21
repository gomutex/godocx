package elements

import (
	"encoding/xml"
	"fmt"

	"github.com/gomutex/godocx/constants"
	"github.com/gomutex/godocx/oxml/types"
)

type Anchor struct {
	BehindDocAttr      int
	DistTAttr          int
	DistBAttr          int
	DistLAttr          int
	DistRAttr          int
	SimplePosAttr      int
	LockedAttr         int
	LayoutInCellAttr   int
	AllowOverlapAttr   int
	RelativeHeightAttr int
	SimplePosXAttr     int
	SimplePosYAttr     int
	SimplePos          *types.PositionType `xml:"simplePos"`
}

func NewAnchor() *Anchor {
	return &Anchor{}
}

type Inline struct {
}

func NewInline() *Inline {
	return &Inline{}
}

func (a *Anchor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "wp:anchor"
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "behindDoc"}, Value: fmt.Sprintf("%d", a.BehindDocAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distT"}, Value: fmt.Sprintf("%d", a.DistTAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distB"}, Value: fmt.Sprintf("%d", a.DistBAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distL"}, Value: fmt.Sprintf("%d", a.DistLAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distR"}, Value: fmt.Sprintf("%d", a.DistRAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "simplePos"}, Value: fmt.Sprintf("%d", a.SimplePosAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "locked"}, Value: fmt.Sprintf("%d", a.LockedAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "layoutInCell"}, Value: fmt.Sprintf("%d", a.LayoutInCellAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "allowOverlap"}, Value: fmt.Sprintf("%d", a.AllowOverlapAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "relativeHeight"}, Value: fmt.Sprintf("%d", a.RelativeHeightAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "simplePosX"}, Value: fmt.Sprintf("%d", a.SimplePosXAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "simplePosY"}, Value: fmt.Sprintf("%d", a.SimplePosYAttr)})

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if a.SimplePos != nil {
		e.EncodeElement(a.SimplePos, xml.StartElement{Name: xml.Name{Local: "wp:simplePos"}})
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (a *Anchor) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "behindDoc":
			fmt.Sscanf(attr.Value, "%d", &a.BehindDocAttr)
		case "distT":
			fmt.Sscanf(attr.Value, "%d", &a.DistTAttr)
		case "distB":
			fmt.Sscanf(attr.Value, "%d", &a.DistBAttr)
		case "distL":
			fmt.Sscanf(attr.Value, "%d", &a.DistLAttr)
		case "distR":
			fmt.Sscanf(attr.Value, "%d", &a.DistRAttr)
		case "simplePos":
			fmt.Sscanf(attr.Value, "%d", &a.SimplePosAttr)
		case "locked":
			fmt.Sscanf(attr.Value, "%d", &a.LockedAttr)
		case "layoutInCell":
			fmt.Sscanf(attr.Value, "%d", &a.LayoutInCellAttr)
		case "allowOverlap":
			fmt.Sscanf(attr.Value, "%d", &a.AllowOverlapAttr)
		case "relativeHeight":
			fmt.Sscanf(attr.Value, "%d", &a.RelativeHeightAttr)
		case "simplePosX":
			fmt.Sscanf(attr.Value, "%d", &a.SimplePosXAttr)
		case "simplePosY":
			fmt.Sscanf(attr.Value, "%d", &a.SimplePosYAttr)
		}
	}

	for {
		token, err := decoder.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name {
			case xml.Name{Space: constants.WMLDrawingNS, Local: "simplePos"}:
				a.SimplePos = &types.PositionType{}
				if err := decoder.DecodeElement(a.SimplePos, &elem); err != nil {
					return err
				}

			default:
				// fmt.Println(elem.Name)
				if err = decoder.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			if elem == start.End() {
				return nil
			}
		}
	}

}

func (i *Inline) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "wp:inline"
	return e.EncodeElement(i, start)
}

func (i *Inline) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := decoder.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
		case xml.EndElement:
			if elem == start.End() {
				return nil
			}
		}
	}
}
