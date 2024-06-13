package dml

import (
	"encoding/xml"
	"fmt"

	"github.com/gomutex/godocx/ooxml/types"
)

type Anchor struct {
	/// Specifies that this object shall be positioned using the positioning information in the
	/// simplePos child element (§20.4.2.13). This positioning, when specified, positions the
	/// object on the page by placing its top left point at the x-y coordinates specified by that
	/// element.
	/// Reference: http://officeopenxml.com/drwPicFloating-position.php
	SimplePosAttr int

	/// Specifies the minimum distance which shall be maintained between the top edge of this drawing object and any subsequent text within the document when this graphical object is displayed within the document's contents.,
	/// The distance shall be measured in EMUs (English Mektric Units).,
	DistTAttr int
	DistBAttr int
	DistLAttr int
	DistRAttr int

	LayoutInCellAttr   int
	AllowOverlapAttr   int
	RelativeHeightAttr int
	BehindDocAttr      int
	LockedAttr         int

	// Child elements:
	SimplePos         *types.PositionType
	PositionH         *types.PoistionH
	PositionV         *types.PoistionV
	Graphic           *Graphic
	Extent            *Extent
	DocProp           *DocProp
	cNvGraphicFramePr *NonVisualGraphicFrameProp
	EffectExtent      *EffectExtent
	WrapNone          *WrapNone `xml:"wrapNone"`
}

func NewAnchor() *Anchor {
	return &Anchor{}
}

func (a *Anchor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "wp:anchor"
	start.Attr = []xml.Attr{
		// {Name: xml.Name{Local: "xmlns:a"}, Value: constants.DrawingMLMainNS},
		// {Name: xml.Name{Local: "xmlns:pic"}, Value: constants.DrawingMLPicNS},
	}

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

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	// Placement(the order) of these elements is important especially the wrapNone

	if a.SimplePos != nil {
		if err := e.EncodeElement(a.SimplePos, xml.StartElement{Name: xml.Name{Local: "wp:simplePos"}}); err != nil {
			return err
		}
	}

	if a.PositionH != nil {
		if err := e.EncodeElement(a.PositionH, xml.StartElement{Name: xml.Name{Local: "wp:positionH"}}); err != nil {
			return err
		}
	}

	if a.PositionV != nil {
		if err := e.EncodeElement(a.PositionV, xml.StartElement{Name: xml.Name{Local: "wp:positionV"}}); err != nil {
			return err
		}
	}

	if a.Extent != nil {
		if err := e.EncodeElement(a.Extent, xml.StartElement{Name: xml.Name{Local: "wp:extent"}}); err != nil {
			return err
		}
	}

	if a.EffectExtent != nil {
		if err := e.EncodeElement(a.EffectExtent, xml.StartElement{Name: xml.Name{Local: "wp:effectExtent"}}); err != nil {
			return err
		}
	}

	// if a.WrapNone.Valid && a.WrapNone.Bool {
	if a.WrapNone != nil {
		err := e.EncodeElement(a.WrapNone, xml.StartElement{Name: xml.Name{Local: "wp:wrapNone"}})
		if err != nil {
			return err
		}
	}

	if a.DocProp != nil {
		if err := e.EncodeElement(a.DocProp, xml.StartElement{Name: xml.Name{Local: "wp:docPr"}}); err != nil {
			return err
		}
	}
	if a.cNvGraphicFramePr != nil {
		if err := e.EncodeElement(a.cNvGraphicFramePr, xml.StartElement{Name: xml.Name{Local: "wp:cNvGraphicFramePr"}}); err != nil {
			return err
		}
	}

	if a.Graphic != nil {
		if err := e.EncodeElement(a.Graphic, xml.StartElement{Name: xml.Name{Local: "a:graphic"}}); err != nil {
			return err
		}
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
		}
	}

	for {
		token, err := decoder.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "extent":
				a.Extent = &Extent{
					XMLName: "wp:extent",
				}
				if err := decoder.DecodeElement(a.Extent, &elem); err != nil {
					return err
				}
			case "effectExtent":
				a.EffectExtent = &EffectExtent{
					XMLName: "wp:effectExtent",
				}
				if err := decoder.DecodeElement(a.EffectExtent, &elem); err != nil {
					return err
				}
			case "cNvGraphicFramePr":
				a.cNvGraphicFramePr = &NonVisualGraphicFrameProp{}
				if err := decoder.DecodeElement(a.cNvGraphicFramePr, &elem); err != nil {
					return err
				}
			case "docPr":
				a.DocProp = &DocProp{}
				if err := decoder.DecodeElement(a.DocProp, &elem); err != nil {
					return err
				}
			case "simplePos":
				a.SimplePos = &types.PositionType{}
				if err := decoder.DecodeElement(a.SimplePos, &elem); err != nil {
					return err
				}
			case "positionV":
				a.PositionV = &types.PoistionV{}
				if err := decoder.DecodeElement(a.PositionV, &elem); err != nil {
					return err
				}
			case "positionH":
				a.PositionH = &types.PoistionH{}
				if err := decoder.DecodeElement(a.PositionH, &elem); err != nil {
					return err
				}
			case "graphic":
				a.Graphic = &Graphic{}
				if err := decoder.DecodeElement(a.Graphic, &elem); err != nil {
					return err
				}
			case "wrapNone":
				a.WrapNone = &WrapNone{}
			default:
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
