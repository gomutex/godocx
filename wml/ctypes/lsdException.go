package ctypes

import (
	"encoding/xml"
	"strconv"

	"github.com/gomutex/godocx/wml/stypes"
)

// Latent Style Exception
type LsdException struct {
	Name           string        `xml:"name,attr"`
	Locked         *stypes.OnOff `xml:"locked,attr,omitempty"`
	UIPriority     *int          `xml:"uiPriority,attr,omitempty"`
	SemiHidden     *stypes.OnOff `xml:"semiHidden,attr,omitempty"`
	UnhideWhenUsed *stypes.OnOff `xml:"unhideWhenUsed,attr,omitempty"`
	QFormat        *stypes.OnOff `xml:"qFormat,attr,omitempty"`
}

func (l LsdException) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	start.Name.Local = "w:lsdException"

	start.Attr = append(start.Attr,
		xml.Attr{Name: xml.Name{Local: "w:name"}, Value: l.Name},
	)

	if l.Locked != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:locked"}, Value: string(*l.Locked)})
	}
	if l.UIPriority != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:uiPriority"}, Value: strconv.Itoa(*l.UIPriority)})
	}
	if l.SemiHidden != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:semiHidden"}, Value: string(*l.SemiHidden)})
	}
	if l.UnhideWhenUsed != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:unhideWhenUsed"}, Value: string(*l.UnhideWhenUsed)})
	}
	if l.QFormat != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:qFormat"}, Value: string(*l.QFormat)})
	}

	return e.EncodeElement("", start)
}
