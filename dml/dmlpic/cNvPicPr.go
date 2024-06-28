package dmlpic

import (
	"encoding/xml"
	"fmt"

	"github.com/gomutex/godocx/dml/dmlct"
	"github.com/gomutex/godocx/dml/dmlprops"
)

// Non-Visual Picture Drawing Properties
type CNvPicPr struct {
	//Relative Resize Preferred	- Default value is "true"(i.e when attr not specified).
	PreferRelativeResize *bool `xml:"preferRelativeResize,attr,omitempty"`

	//1. Picture Locks
	PicLocks *dmlprops.PicLocks `xml:"picLocks,omitempty"`

	//TODO:
	// 2. Extension List

}

func NewCNvPicPr() CNvPicPr {
	return CNvPicPr{}
}

func (c CNvPicPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "pic:cNvPicPr"

	if c.PreferRelativeResize != nil {
		if *c.PreferRelativeResize {
			start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "preferRelativeResize"}, Value: "true"})
		} else {
			start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "preferRelativeResize"}, Value: "false"})
		}
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	// 1. PicLocks
	if c.PicLocks != nil {
		if err := e.EncodeElement(c.PicLocks, xml.StartElement{Name: xml.Name{Local: "a:picLocks"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

// Non-Visual Picture Properties
type NonVisualPicProp struct {
	// 1. Non-Visual Drawing Properties
	CNvPr dmlct.CNvPr `xml:"cNvPr,omitempty"`

	// 2.Non-Visual Picture Drawing Properties
	CNvPicPr CNvPicPr `xml:"cNvPicPr,omitempty"`
}

func NewNVPicProp(cNvPr dmlct.CNvPr, cNvPicPr CNvPicPr) NonVisualPicProp {
	return NonVisualPicProp{
		CNvPr:    cNvPr,
		CNvPicPr: cNvPicPr,
	}
}

func DefaultNVPicProp(id uint, name string) NonVisualPicProp {
	cnvPicPr := NewCNvPicPr()
	cnvPicPr.PicLocks = dmlprops.DefaultPicLocks()
	return NonVisualPicProp{
		CNvPr:    *dmlct.NewNonVisProp(id, name),
		CNvPicPr: cnvPicPr,
	}
}

func (n NonVisualPicProp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "pic:nvPicPr"

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	// 1. cNvPr
	if err = n.CNvPr.MarshalXML(e, xml.StartElement{
		Name: xml.Name{Local: "pic:cNvPr"},
	}); err != nil {
		return fmt.Errorf("marshalling cNvPr: %w", err)
	}

	// 2. cNvPicPr
	if err = n.CNvPicPr.MarshalXML(e, xml.StartElement{
		Name: xml.Name{Local: "pic:cNvPicPr"},
	}); err != nil {
		return fmt.Errorf("marshalling cNvPicPr: %w", err)
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
