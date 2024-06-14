package dml

import "encoding/xml"

// Non-Visual Picture Drawing Properties
type CNvPicPr struct {
	PicLocks *PicLocks
}

func (c *CNvPicPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "pic:cNvPicPr"

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if c.PicLocks != nil {
		if err := e.EncodeElement(c.PicLocks, xml.StartElement{Name: xml.Name{Local: "a:picLocks"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (c *CNvPicPr) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := decoder.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "picLocks":
				c.PicLocks = &PicLocks{}
				if err = decoder.DecodeElement(c.PicLocks, &elem); err != nil {
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

// Non-Visual Picture Properties
type NonVisualPicProp struct {
	CNvPr    *CNvPr
	CNvPicPr *CNvPicPr
}

func (n *NonVisualPicProp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "pic:nvPicPr"

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if n.CNvPr != nil {
		if err := e.EncodeElement(n.CNvPr, xml.StartElement{Name: xml.Name{Local: "pic:cNvPr"}}); err != nil {
			return err
		}
	}

	if n.CNvPicPr != nil {
		if err := e.EncodeElement(n.CNvPicPr, xml.StartElement{Name: xml.Name{Local: "pic:nvPicPr"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (n *NonVisualPicProp) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "cNvPr":
				cnvp := &CNvPr{}
				if err := d.DecodeElement(cnvp, &elem); err != nil {
					return err
				}
				n.CNvPr = cnvp
			case "cNvPicPr":
				cnvpp := &CNvPicPr{}
				if err := d.DecodeElement(cnvpp, &elem); err != nil {
					return err
				}
				n.CNvPicPr = cnvpp
			}
		case xml.EndElement:
			if elem == start.End() {
				return nil
			}
		}
	}
}
