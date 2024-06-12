package dml

import "encoding/xml"

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
