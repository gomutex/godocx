package ctypes

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/gomutex/godocx/wml/stypes"
)

// Latent Style Information
type LatentStyle struct {
	//Default Style Locking Setting
	DefLockedState *stypes.OnOff `xml:"defLockedState,attr,omitempty"`

	//Default User Interface Priority Setting
	DefUIPriority *int `xml:"defUIPriority,attr,omitempty"`

	//Default Semi-Hidden Setting
	DefSemiHidden *stypes.OnOff `xml:"defSemiHidden,attr,omitempty"`

	//Default Hidden Until Used Setting
	DefUnhideWhenUsed *stypes.OnOff `xml:"defUnhideWhenUsed,attr,omitempty"`

	//Default Primary Style Setting
	DefQFormat *stypes.OnOff `xml:"defQFormat,attr,omitempty"`

	//Latent Style Count
	Count *int `xml:"count,attr,omitempty"`

	LsdExceptions []LsdException `xml:",any"`
}

func (l LatentStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:latentStyles"

	if l.DefLockedState != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Space: "", Local: "w:defLockedState"}, Value: string(*l.DefLockedState)})
	}

	if l.DefUIPriority != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Space: "", Local: "w:defUIPriority"}, Value: strconv.Itoa(*l.DefUIPriority)})
	}

	if l.DefSemiHidden != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Space: "", Local: "w:defSemiHidden"}, Value: string(*l.DefSemiHidden)})
	}

	if l.DefUnhideWhenUsed != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Space: "", Local: "w:defUnhideWhenUsed"}, Value: string(*l.DefUnhideWhenUsed)})
	}

	if l.DefQFormat != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Space: "", Local: "w:defQFormat"}, Value: string(*l.DefQFormat)})
	}

	if l.Count != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Space: "", Local: "w:count"}, Value: strconv.Itoa(*l.Count)})
	}

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	for _, elem := range l.LsdExceptions {
		if err := elem.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:lsdException"},
		}); err != nil {
			return fmt.Errorf("LsdException: %w", err)
		}
	}

	return e.EncodeToken(start.End())
}
