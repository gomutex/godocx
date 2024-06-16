package dml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type Anchor struct {
	/// Specifies that this object shall be positioned using the positioning information in the
	/// simplePos child element (§20.4.2.13). This positioning, when specified, positions the
	/// object on the page by placing its top left point at the x-y coordinates specified by that
	/// element.
	/// Reference: http://officeopenxml.com/drwPicFloating-position.php
	SimplePosAttr *int `xml:"simplePos,attr,omitempty"`

	/// Specifies the minimum distance which shall be maintained between the top edge of this drawing object and any subsequent text within the document when this graphical object is displayed within the document's contents.,
	/// The distance shall be measured in EMUs (English Mektric Units).,
	DistTAttr *uint `xml:"distT,attr,omitempty"`
	DistBAttr *uint `xml:"distB,attr,omitempty"`
	DistLAttr *uint `xml:"distL,attr,omitempty"`
	DistRAttr *uint `xml:"distR,attr,omitempty"`

	LayoutInCellAttr   *int `xml:"layoutInCell,attr,omitempty"`
	AllowOverlapAttr   *int `xml:"allowOverlap,attr,omitempty"`
	RelativeHeightAttr *int `xml:"relativeHeight,attr,omitempty"`
	BehindDocAttr      *int `xml:"behindDoc,attr,omitempty"`
	LockedAttr         *int `xml:"locked,attr,omitempty"`

	// Child elements:
	SimplePos         *PositionType              `xml:"simplePos,omitempty"`
	PositionH         *PoistionH                 `xml:"positionH,omitempty"`
	PositionV         *PoistionV                 `xml:"positionV,omitempty"`
	Graphic           *Graphic                   `xml:"graphic,omitempty"`
	Extent            *Extent                    `xml:"extent,omitempty"`
	DocProp           *DocProp                   `xml:"docPr,omitempty"`
	CNvGraphicFramePr *NonVisualGraphicFrameProp `xml:"cNvGraphicFramePr,omitempty"`
	EffectExtent      *EffectExtent              `xml:"effectExtent,omitempty"`
	WrapNone          *WrapNone                  `xml:"wrapNone,omitempty"`
}

func NewAnchor() *Anchor {
	return &Anchor{}
}

func (a *Anchor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "wp:anchor"

	if a.BehindDocAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "behindDoc"}, Value: strconv.Itoa(*a.BehindDocAttr)})
	}

	if a.DistTAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distT"}, Value: strconv.FormatUint(uint64(*a.DistTAttr), 10)})
	}

	if a.DistBAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distB"}, Value: strconv.FormatUint(uint64(*a.DistBAttr), 10)})
	}

	if a.DistLAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distL"}, Value: strconv.FormatUint(uint64(*a.DistLAttr), 10)})
	}

	if a.DistRAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distR"}, Value: strconv.FormatUint(uint64(*a.DistRAttr), 10)})
	}

	if a.SimplePosAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "simplePos"}, Value: strconv.Itoa(*a.SimplePosAttr)})
	}

	if a.LockedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "locked"}, Value: strconv.Itoa(*a.LockedAttr)})
	}

	if a.LayoutInCellAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "layoutInCell"}, Value: strconv.Itoa(*a.LayoutInCellAttr)})
	}

	if a.AllowOverlapAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "allowOverlap"}, Value: strconv.Itoa(*a.AllowOverlapAttr)})
	}

	if a.RelativeHeightAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "relativeHeight"}, Value: strconv.Itoa(*a.RelativeHeightAttr)})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	// Placement(the order) of these elements is important especially the wrapNone

	if a.SimplePos != nil {
		if err := e.EncodeElement(a.SimplePos, xml.StartElement{Name: xml.Name{Local: "wp:simplePos"}}); err != nil {
			return fmt.Errorf("SimplePos: %v", err)
		}
	}

	if a.PositionH != nil {
		if err := a.PositionH.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("PositionH: %v", err)
		}
	}

	if a.PositionV != nil {
		if err := a.PositionV.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("PositionV: %v", err)
		}
	}

	if a.Extent != nil {
		if err := a.Extent.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "wp:extent"}}); err != nil {
			return fmt.Errorf("Extent: %v", err)
		}
	}

	if a.EffectExtent != nil {
		if err := a.EffectExtent.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("EffectExtent: %v", err)
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
	if a.CNvGraphicFramePr != nil {
		if err := e.EncodeElement(a.CNvGraphicFramePr, xml.StartElement{Name: xml.Name{Local: "wp:cNvGraphicFramePr"}}); err != nil {
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

// func (a *Anchor) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
// 	for _, attr := range start.Attr {
// 		switch attr.Name.Local {
// 		case "behindDoc":
// 			fmt.Sscanf(attr.Value, "%d", &a.BehindDocAttr)
// 		case "distT":
// 			fmt.Sscanf(attr.Value, "%d", &a.DistTAttr)
// 		case "distB":
// 			fmt.Sscanf(attr.Value, "%d", &a.DistBAttr)
// 		case "distL":
// 			fmt.Sscanf(attr.Value, "%d", &a.DistLAttr)
// 		case "distR":
// 			fmt.Sscanf(attr.Value, "%d", &a.DistRAttr)
// 		case "simplePos":
// 			fmt.Sscanf(attr.Value, "%d", &a.SimplePosAttr)
// 		case "locked":
// 			fmt.Sscanf(attr.Value, "%d", &a.LockedAttr)
// 		case "layoutInCell":
// 			fmt.Sscanf(attr.Value, "%d", &a.LayoutInCellAttr)
// 		case "allowOverlap":
// 			fmt.Sscanf(attr.Value, "%d", &a.AllowOverlapAttr)
// 		case "relativeHeight":
// 			fmt.Sscanf(attr.Value, "%d", &a.RelativeHeightAttr)
// 		}
// 	}

// 	for {
// 		token, err := decoder.Token()
// 		if err != nil {
// 			return err
// 		}

// 		switch elem := token.(type) {
// 		case xml.StartElement:
// 			switch elem.Name.Local {
// 			case "extent":
// 				a.Extent = &Extent{
// 					XMLName: "wp:extent",
// 				}
// 				if err = decoder.DecodeElement(a.Extent, &elem); err != nil {
// 					return err
// 				}
// 			case "effectExtent":
// 				a.EffectExtent = &EffectExtent{}
// 				if err = decoder.DecodeElement(a.EffectExtent, &elem); err != nil {
// 					return err
// 				}
// 			case "cNvGraphicFramePr":
// 				a.cNvGraphicFramePr = &NonVisualGraphicFrameProp{}
// 				if err = decoder.DecodeElement(a.cNvGraphicFramePr, &elem); err != nil {
// 					return err
// 				}
// 			case "docPr":
// 				a.DocProp = &DocProp{}
// 				if err = decoder.DecodeElement(a.DocProp, &elem); err != nil {
// 					return err
// 				}
// 			case "simplePos":
// 				a.SimplePos = &PositionType{}
// 				if err = decoder.DecodeElement(a.SimplePos, &elem); err != nil {
// 					return err
// 				}
// 			case "positionV":
// 				a.PositionV = &PoistionV{}
// 				if err = decoder.DecodeElement(a.PositionV, &elem); err != nil {
// 					return err
// 				}
// 			case "positionH":
// 				a.PositionH = &PoistionH{}
// 				if err = decoder.DecodeElement(a.PositionH, &elem); err != nil {
// 					return err
// 				}
// 			case "graphic":
// 				a.Graphic = &Graphic{}
// 				if err = decoder.DecodeElement(a.Graphic, &elem); err != nil {
// 					return err
// 				}
// 			case "wrapNone":
// 				a.WrapNone = &WrapNone{}
// 			default:
// 				if err = decoder.Skip(); err != nil {
// 					return err
// 				}
// 			}
// 		case xml.EndElement:
// 			if elem == start.End() {
// 				return nil
// 			}
// 		}
// 	}
// }
