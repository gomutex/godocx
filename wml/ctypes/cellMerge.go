package ctypes

import (
	"encoding/xml"
	"errors"
	"strconv"
)

type CellMerge struct {
	ID         int               `xml:"id,attr"`
	Author     string            `xml:"author,attr"`
	Date       *string           `xml:"date,attr,omitempty"`
	VMerge     *AnnotationVMerge `xml:"vMerge,attr,omitempty"`     //Revised Vertical Merge Setting
	VMergeOrig *AnnotationVMerge `xml:"vMergeOrig,attr,omitempty"` //Vertical Merge Setting Removed by Revision

}

func (t CellMerge) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "w:id"}, Value: strconv.Itoa(t.ID)},
		{Name: xml.Name{Local: "w:author"}, Value: t.Author},
	}

	if t.Date != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:date"}, Value: *t.Date})
	}

	if t.VMerge != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:vMerge"}, Value: string(*t.VMerge)})
	}

	if t.VMergeOrig != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:vMergeOrig"}, Value: string(*t.VMergeOrig)})
	}

	return e.EncodeElement("", start)
}

type AnnotationVMerge string

const (
	// AnnotationVMergeCont represents a vertically merged cell.
	AnnotationVMergeCont AnnotationVMerge = "cont"
	// AnnotationVMergeRest represents a vertically split cell.
	AnnotationVMergeRest AnnotationVMerge = "rest"
)

// AnnotationVMergeFromStr converts a string to AnnotationVMerge type.
func AnnotationVMergeFromStr(value string) (AnnotationVMerge, error) {
	switch value {
	case "cont":
		return AnnotationVMergeCont, nil
	case "rest":
		return AnnotationVMergeRest, nil
	default:
		return "", errors.New("invalid AnnotationVMerge value")
	}
}

// UnmarshalXMLAttr unmarshals XML attribute into AnnotationVMerge.
func (a *AnnotationVMerge) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := AnnotationVMergeFromStr(attr.Value)
	if err != nil {
		return err
	}
	*a = val
	return nil
}
