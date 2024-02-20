package elements

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/gomutex/godocx/constants"
	"github.com/gomutex/godocx/oxml/types"
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
	// TODO:
	// EffectExtent
	//
	// implement any other
}

func NewAnchor() *Anchor {
	return &Anchor{}
}

type Inline struct {
	/// Specifies the minimum distance which shall be maintained between the top edge of this drawing object and any subsequent text within the document when this graphical object is displayed within the document's contents.,
	/// The distance shall be measured in EMUs (English Mektric Units).,
	DistTAttr int
	DistBAttr int
	DistLAttr int
	DistRAttr int

	// Child elements:
	Extent            *Extent
	DocProp           *DocProp
	Graphic           *Graphic
	cNvGraphicFramePr *NonVisualGraphicFrameProp
}

func NewInline() *Inline {
	return &Inline{}
}

func (a *Anchor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "wp:anchor"
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "xmlns:a"}, Value: constants.DrawingMLMainNS},
		{Name: xml.Name{Local: "xmlns:pic"}, Value: constants.DrawingMLPicNS},
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

	if a.Extent != nil {
		if err := e.EncodeElement(a.Extent, xml.StartElement{Name: xml.Name{Local: "wp:extent"}}); err != nil {
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

func (i *Inline) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "wp:inline"
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "xmlns:a"}, Value: constants.DrawingMLMainNS},
		{Name: xml.Name{Local: "xmlns:pic"}, Value: constants.DrawingMLPicNS},
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if i.Extent != nil {
		if err := e.EncodeElement(i.Extent, xml.StartElement{Name: xml.Name{Local: "wp:extent"}}); err != nil {
			return err
		}
	}

	if i.DocProp != nil {
		if err := e.EncodeElement(i.DocProp, xml.StartElement{Name: xml.Name{Local: "wp:docPr"}}); err != nil {
			return err
		}
	}

	if i.cNvGraphicFramePr != nil {
		if err := e.EncodeElement(i.cNvGraphicFramePr, xml.StartElement{Name: xml.Name{Local: "wp:cNvGraphicFramePr"}}); err != nil {
			return err
		}
	}

	if i.Graphic != nil {
		if err = i.Graphic.MarshalXML(e, start); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (i *Inline) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := decoder.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "extent":
				i.Extent = &Extent{}
				if err := decoder.DecodeElement(i.Extent, &elem); err != nil {
					return err
				}
			case "docPr":
				i.DocProp = &DocProp{}
				if err := decoder.DecodeElement(i.DocProp, &elem); err != nil {
					return err
				}
			case "cNvGraphicFramePr":
				i.cNvGraphicFramePr = &NonVisualGraphicFrameProp{}
				if err := decoder.DecodeElement(i.cNvGraphicFramePr, &elem); err != nil {
					return err
				}
			case "graphic":
				i.Graphic = DefaultGraphic()
				if err := decoder.DecodeElement(i.Graphic, &elem); err != nil {
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

type Extent struct {
	XMLName string
	Length  uint64 // cx
	Width   uint64 // cy
}

func (x *Extent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// start.Name.Local = x.XMLName
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "cx"}, Value: strconv.FormatUint(x.Length, 10)},
		{Name: xml.Name{Local: "cy"}, Value: strconv.FormatUint(x.Width, 10)},
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (x *Extent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		if a.Name.Local == "cx" {
			cx, err := strconv.ParseUint(a.Value, 10, 32)
			if err != nil {
				return nil
			}
			x.Width = cx
		} else if a.Name.Local == "cy" {
			cy, err := strconv.ParseUint(a.Value, 10, 32)
			if err != nil {
				return nil
			}
			x.Length = cy
		}
	}

	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {

			default:
				if err = d.Skip(); err != nil {
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

type DocProp struct {
	ID   uint64 // cx
	Name string // cy
}

func (d *DocProp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "wp:docPr"
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "id"}, Value: strconv.FormatUint(d.ID, 10)},
		{Name: xml.Name{Local: "name"}, Value: d.Name},
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (d *DocProp) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		if a.Name.Local == "id" {
			id, err := strconv.ParseUint(a.Value, 10, 32)
			if err != nil {
				return nil
			}
			d.ID = id
		} else if a.Name.Local == "name" {
			d.Name = a.Value
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

type NonVisualGraphicFrameProp struct {
	GraphicFrameLocks *GraphicFrameLocks
}

func (n *NonVisualGraphicFrameProp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "wp:cNvGraphicFramePr"
	start.Attr = []xml.Attr{}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if n.GraphicFrameLocks != nil {
		if err := e.EncodeElement(n.GraphicFrameLocks, xml.StartElement{Name: xml.Name{Local: "a:graphicFrameLocks"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (n *NonVisualGraphicFrameProp) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {

	for {
		token, err := decoder.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "graphicFrameLocks":
				n.GraphicFrameLocks = &GraphicFrameLocks{}
				if err := decoder.DecodeElement(n.GraphicFrameLocks, &elem); err != nil {
					return err
				}
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

type GraphicFrameLocks struct {
	//Disallow Aspect Ratio Change
	noChangeAspect bool
}

func (g *GraphicFrameLocks) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "a:graphicFrameLocks"

	noChangeAspect := "0"
	if g.noChangeAspect {
		noChangeAspect = "1"
	}

	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "noChangeAspect"}, Value: noChangeAspect},
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (g *GraphicFrameLocks) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		if a.Name.Local == "noChangeAspect" {
			if a.Value == "1" {
				g.noChangeAspect = true
			} else {
				g.noChangeAspect = false
			}
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
