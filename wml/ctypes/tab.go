package ctypes

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/gomutex/godocx/wml/stypes"
)

// Custom Tab Stop
type Tab struct {
	// Tab Stop Type
	Val stypes.CustTabStop `xml:"val,attr,omitempty"`

	//Tab Stop Position
	Position int `xml:"pos,attr,omitempty"`

	//Custom Tab Stop Leader Character
	LeaderChar *stypes.CustLeadChar `xml:"leader,attr,omitempty"`
}

func (t Tab) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:tab"
	start.Attr = []xml.Attr{}

	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"}, Value: string(t.Val)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:pos"}, Value: strconv.Itoa(t.Position)})

	if t.LeaderChar != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:leader"}, Value: string(*t.LeaderChar)})
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

type Tabs struct {
	Tab []Tab `xml:"tab,omitempty"`
}

func (t Tabs) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	if len(t.Tab) == 0 {
		return nil
	}

	// Create the enclosing XML element
	start.Name = xml.Name{Local: "w:tabs"}

	err := e.EncodeToken(start)
	if err != nil {
		return fmt.Errorf("error encoding start element: %v", err)
	}

	for _, tab := range t.Tab {
		if err := tab.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("error encoding tab: %v", err)
		}
	}

	err = e.EncodeToken(start.End())
	if err != nil {
		return fmt.Errorf("error encoding end element: %v", err)
	}

	return nil
}

type PTab struct {
	Alignment  stypes.PTabAlignment  `xml:"alignment,attr,omitempty"`
	RelativeTo stypes.PTabRelativeTo `xml:"relativeTo,attr,omitempty"`
	Leader     stypes.PTabLeader     `xml:"leader,attr,omitempty"`
}

func (t PTab) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "w:alignment"}, Value: string(t.Alignment)},
		{Name: xml.Name{Local: "w:relativeTo"}, Value: string(t.RelativeTo)},
		{Name: xml.Name{Local: "w:leader"}, Value: string(t.Leader)},
	}

	return e.EncodeElement("", start)
}
