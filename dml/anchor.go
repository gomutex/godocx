package dml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/gomutex/godocx/dml/dmlct"
)

type Anchor struct {
	/// Specifies that this object shall be positioned using the positioning information in the
	/// simplePos child element (§20.4.2.13). This positioning, when specified, positions the
	/// object on the page by placing its top left point at the x-y coordinates specified by that
	/// element.
	/// Reference: http://officeopenxml.com/drwPicFloating-position.php
	//Page Positioning
	SimplePosAttr *int `xml:"simplePos,attr,omitempty"`

	/// Specifies the minimum distance which shall be maintained between the top edge of this drawing object and any subsequent text within the document when this graphical object is displayed within the document's contents.,
	/// The distance shall be measured in EMUs (English Mektric Units).,
	//Distance From Text on Top Edge
	DistT uint `xml:"distT,attr,omitempty"`
	//Distance From Text on Bottom Edge
	DistB uint `xml:"distB,attr,omitempty"`
	//Distance From Text on Left Edge
	DistL uint `xml:"distL,attr,omitempty"`
	//Distance From Text on Right Edge
	DistR uint `xml:"distR,attr,omitempty"`

	//Relative Z-Ordering Position
	RelativeHeight int `xml:"relativeHeight,attr"`

	//Layout In Table Cell
	LayoutInCell int `xml:"layoutInCell,attr"`

	//Display Behind Document Text
	BehindDoc int `xml:"behindDoc,attr"`

	//Lock Anchor
	Locked int `xml:"locked,attr"`

	//Allow Objects to Overlap
	AllowOverlap int `xml:"allowOverlap,attr"`

	Hidden *int `xml:"hidden,attr,omitempty"`

	// Child elements:
	// 1. Simple Positioning Coordinates
	SimplePos dmlct.Point2D `xml:"simplePos"`

	// 2. Horizontal Positioning
	PositionH PoistionH `xml:"positionH"`

	// 3. Vertical Positioning
	PositionV PoistionV `xml:"positionV"`

	// 4. Inline Drawing Object Extents
	Extent dmlct.PSize2D `xml:"extent"`

	// 5. EffectExtent
	EffectExtent *EffectExtent `xml:"effectExtent,omitempty"`

	// 6. Wrapping
	// 6.1 .wrapNone
	WrapNone *WrapNone `xml:"wrapNone,omitempty"`

	// 6.2. wrapSquare
	WrapSquare *WrapSquare `xml:"wrapSquare,omitempty"`

	// 6.3. wrapThrough
	WrapThrough *WrapThrough `xml:"wrapThrough,omitempty"`

	// 6.4. wrapTopAndBottom
	WrapTopBtm *WrapTopBtm `xml:"wrapTopAndBottom,omitempty"`

	// 7. Drawing Object Non-Visual Properties
	DocProp DocProp `xml:"docPr"`

	// 8. Common DrawingML Non-Visual Properties
	CNvGraphicFramePr *NonVisualGraphicFrameProp `xml:"cNvGraphicFramePr,omitempty"`

	// 9. Graphic Object
	Graphic Graphic `xml:"graphic"`
}

func NewAnchor() *Anchor {
	return &Anchor{}
}

func (a Anchor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "wp:anchor"

	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "behindDoc"}, Value: strconv.Itoa(a.BehindDoc)})

	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distT"}, Value: strconv.FormatUint(uint64(a.DistT), 10)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distB"}, Value: strconv.FormatUint(uint64(a.DistB), 10)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distL"}, Value: strconv.FormatUint(uint64(a.DistL), 10)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distR"}, Value: strconv.FormatUint(uint64(a.DistR), 10)})

	if a.SimplePosAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "simplePos"}, Value: strconv.Itoa(*a.SimplePosAttr)})
	}

	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "locked"}, Value: strconv.Itoa(a.Locked)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "layoutInCell"}, Value: strconv.Itoa(a.LayoutInCell)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "allowOverlap"}, Value: strconv.Itoa(a.AllowOverlap)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "relativeHeight"}, Value: strconv.Itoa(a.RelativeHeight)})
	if a.Hidden != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hidden"}, Value: strconv.Itoa(*a.Hidden)})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	// The sequence (order) of these element is important

	// 1. SimplePos
	if err := a.SimplePos.MarshalXML(e, xml.StartElement{
		Name: xml.Name{Local: "wp:simplePos"},
	}); err != nil {
		return fmt.Errorf("simplePos: %v", err)
	}

	// 2. PositionH
	if err := a.PositionH.MarshalXML(e, xml.StartElement{}); err != nil {
		return fmt.Errorf("PositionH: %v", err)
	}

	// 3. PositionH
	if err := a.PositionV.MarshalXML(e, xml.StartElement{}); err != nil {
		return fmt.Errorf("PositionV: %v", err)
	}

	// 4. Extent
	if err := a.Extent.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "wp:extent"}}); err != nil {
		return fmt.Errorf("Extent: %v", err)
	}

	// 5. EffectExtent
	if err := a.EffectExtent.MarshalXML(e, xml.StartElement{}); err != nil {
		return fmt.Errorf("EffectExtent: %v", err)
	}

	// 6. Wrap Choice
	if err := a.MarshalWrap(e); err != nil {
		return err
	}

	// 7. DocProp
	if err := a.DocProp.MarshalXML(e, xml.StartElement{}); err != nil {
		return err
	}

	// 8. CNvGraphicFramePr
	if a.CNvGraphicFramePr != nil {
		if err := a.CNvGraphicFramePr.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	// 9. Graphic
	if err := a.Graphic.MarshalXML(e, xml.StartElement{}); err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (a *Anchor) MarshalWrap(e *xml.Encoder) error {
	if a.WrapNone != nil {
		return a.WrapNone.MarshalXML(e, xml.StartElement{})
	} else if a.WrapSquare != nil {
		return a.WrapSquare.MarshalXML(e, xml.StartElement{})
	} else if a.WrapThrough != nil {
		return a.WrapThrough.MarshalXML(e, xml.StartElement{})
	} else if a.WrapTopBtm != nil {
		return a.WrapTopBtm.MarshalXML(e, xml.StartElement{})
	}
	return nil
}
