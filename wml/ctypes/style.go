package ctypes

import (
	"encoding/xml"
	"fmt"

	"github.com/gomutex/godocx/common/constants"
	"github.com/gomutex/godocx/wml/stypes"
)

var defaultStyleNSAttrs = []xml.Attr{
	{Name: xml.Name{Local: "xmlns:w"}, Value: "http://schemas.openxmlformats.org/wordprocessingml/2006/main"},
	{Name: xml.Name{Local: "xmlns:mc"}, Value: "http://schemas.openxmlformats.org/markup-compatibility/2006"},
	{Name: xml.Name{Local: "xmlns:w14"}, Value: "http://schemas.microsoft.com/office/word/2010/wordml"},
	{Name: xml.Name{Local: "mc:Ignorable"}, Value: "w14"},
}

// Style Definitions
type Styles struct {
	RelativePath string `xml:"-"`
	Attr         []xml.Attr

	// Sequence

	//1. Document Default Paragraph and Run Properties
	DocDefaults *DocDefault `xml:"docDefaults,omitempty"`

	//2. Latent Style Information
	LatentStyle *LatentStyle `xml:"latentStyles,omitempty"`

	//3. Style Definition
	StyleList []Style `xml:",any"`
}

func (s *Styles) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:styles"

	if len(s.Attr) == 0 {
		start.Attr = append(start.Attr, defaultStyleNSAttrs...)
	} else {
		start.Attr = s.Attr
	}

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	// 1. Document Default Paragraph and Run Properties
	if s.DocDefaults != nil {
		if err := s.DocDefaults.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:docDefaults"}}); err != nil {
			return fmt.Errorf("docDefaults: %w", err)
		}
	}

	// 2. Latent Style Information
	if s.LatentStyle != nil {
		if err := s.LatentStyle.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:latentStyles"}}); err != nil {
			return fmt.Errorf("latentStyle: %w", err)
		}
	}

	//3. Style Definition
	for _, elem := range s.StyleList {
		if err := elem.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:style"}}); err != nil {
			return fmt.Errorf("style: %w", err)
		}
	}

	return e.EncodeToken(start.End())
}

func (s *Styles) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	s.Attr = make([]xml.Attr, len(start.Attr))

	for _, attr := range start.Attr {
		// Broken Go xml - Handling the Docx namespace
		//TODO: Centralize, generic function for document also
		ns := attr.Name.Space
		if ns != "xmlns" {
			// Convert if it is namespace with http:// not the xmlns attribute
			local, ok := constants.NSToLocal[ns]
			ns = local
			if !ok {
				continue
			}
		}

		s.Attr = append(s.Attr, xml.Attr{
			Name: xml.Name{
				Local: fmt.Sprintf("%s:%s", ns, attr.Name.Local),
			},
			Value: attr.Value,
		})
	}
loop:
	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "docDefaults":
				dd := DocDefault{}
				if err = d.DecodeElement(&dd, &elem); err != nil {
					return err
				}
				s.DocDefaults = &dd
			case "latentStyles":
				ls := LatentStyle{}
				if err = d.DecodeElement(&ls, &elem); err != nil {
					return err
				}
				s.LatentStyle = &ls
			case "style":
				style := Style{}
				if err = d.DecodeElement(&style, &elem); err != nil {
					return err
				}
				s.StyleList = append(s.StyleList, style)
			default:
				if err = d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break loop
		}
	}

	return nil
}

type Style struct {
	//Sequence:

	//1. Primary Style Name
	Name *CTString `xml:"name,omitempty"`

	//2. Alternate Style Names
	Alias *CTString `xml:"alias,omitempty"`

	//3. Parent Style ID
	BasedOn *CTString `xml:"basedOn,omitempty"`

	//4. Style For Next Paragraph
	Next *CTString `xml:"next,omitempty"`

	//5. Linked Style Reference
	Link *CTString `xml:"link,omitempty"`

	//6.Automatically Merge User Formatting Into Style Definition
	AutoRedefine *OnOff `xml:"autoRedefine,omitempty"`

	//7.Hide Style From User Interface
	Hidden *OnOff `xml:"hidden,omitempty"`

	//8.Optional User Interface Sorting Order
	UIPriority *DecimalNum `xml:"uiPriority,omitempty"`

	// 9. Hide Style From Main User Interface
	SemiHidden *OnOff `xml:"semiHidden,omitempty"`

	// 10. Remove Semi-Hidden Property When Style Is Used
	UnhideWhenUsed *OnOff `xml:"unhideWhenUsed,omitempty"`

	// 11. Primary Style
	QFormat *OnOff `xml:"qFormat,omitempty"`

	// 12. Style Cannot Be Applied
	Locked *OnOff `xml:"locked,omitempty"`

	// 13. E-Mail Message Text Style
	Personal *OnOff `xml:"personal,omitempty"`

	// 14. E-Mail Message Composition Style
	PersonalCompose *OnOff `xml:"personalCompose,omitempty"`

	// 15. E-Mail Message Reply Style
	PersonalReply *OnOff `xml:"personalReply,omitempty"`

	//16. Revision Identifier for Style Definition
	RevID *GenSingleStrVal[stypes.LongHexNum] `xml:"rsid,omitempty"`

	//17. Style Paragraph Properties
	ParaProp *ParagraphProp `xml:"pPr,omitempty"`

	//18. Run Properties
	RunProp *RunProperty `xml:"rPr,omitempty"`

	//19. Style Table Properties
	TableProp *TableProp `xml:"tblPr,omitempty"`

	//20. Style Table Row Properties
	TableRowProp *RowProperty `xml:"trPr,omitempty"`

	//21. Style Table Cell Properties
	TableCellProp *CellProperty `xml:"tcPr,omitempty"`

	//22.Style Conditional Table Formatting Properties
	TableStylePr []TableStyleProp `xml:",any"`

	// Attributes

	//Style Type
	Type *stypes.StyleType `xml:"type,attr,omitempty"`

	//Style ID
	ID *string `xml:"styleId,attr,omitempty"`

	//Default Style
	Default *stypes.OnOff `xml:"default,attr,omitempty"`

	//User-Defined Style
	CustomStyle *stypes.OnOff `xml:"customStyle,attr,omitempty"`
}

func (s *Style) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:style"

	// Add attributes
	if s.Type != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:type"}, Value: string(*s.Type)})
	}
	if s.ID != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:styleId"}, Value: *s.ID})
	}

	if s.Default != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:default"}, Value: string(*s.Default)})
	}
	if s.CustomStyle != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:customStyle"}, Value: string(*s.CustomStyle)})
	}

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	// 1. Name: Primary Style Name
	if s.Name != nil {
		if err := s.Name.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:name"}}); err != nil {
			return fmt.Errorf("Style Name: %w", err)
		}
	}

	// 2. Alias: Alternate Style Names
	if s.Alias != nil {
		if err := s.Alias.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:alias"}}); err != nil {
			return fmt.Errorf("Style Alias: %w", err)
		}
	}

	// 3. BasedOn: Parent Style ID
	if s.BasedOn != nil {
		if err := s.BasedOn.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:basedOn"}}); err != nil {
			return fmt.Errorf("Style BasedOn: %w", err)
		}
	}

	// 4. Next: Style For Next Paragraph
	if s.Next != nil {
		if err := s.Next.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:next"}}); err != nil {
			return fmt.Errorf("Style Next: %w", err)
		}
	}

	// 5. Link: Linked Style Reference
	if s.Link != nil {
		if err := s.Link.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:link"}}); err != nil {
			return fmt.Errorf("Style Link: %w", err)
		}
	}

	// 6. AutoRedefine: Automatically Merge User Formatting Into Style Definition
	if s.AutoRedefine != nil {
		if err := s.AutoRedefine.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:autoRedefine"}}); err != nil {
			return fmt.Errorf("Style AutoRedefine: %w", err)
		}
	}

	// 7. Hidden: Hide Style From User Interface
	if s.Hidden != nil {
		if err := s.Hidden.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:hidden"}}); err != nil {
			return fmt.Errorf("Style Hidden: %w", err)
		}
	}

	// 8. UIPriority: Optional User Interface Sorting Order
	if s.UIPriority != nil {
		if err := s.UIPriority.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:uiPriority"}}); err != nil {
			return fmt.Errorf("Style UIPriority: %w", err)
		}
	}

	// 9. SemiHidden: Hide Style From Main User Interface
	if s.SemiHidden != nil {
		if err := s.SemiHidden.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:semiHidden"}}); err != nil {
			return fmt.Errorf("Style SemiHidden: %w", err)
		}
	}

	// 10. UnhideWhenUsed: Remove Semi-Hidden Property When Style Is Used
	if s.UnhideWhenUsed != nil {
		if err := s.UnhideWhenUsed.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:unhideWhenUsed"}}); err != nil {
			return fmt.Errorf("Style UnhideWhenUsed: %w", err)
		}
	}

	// 11. QFormat: Primary Style
	if s.QFormat != nil {
		if err := s.QFormat.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:qFormat"}}); err != nil {
			return fmt.Errorf("Style QFormat: %w", err)
		}
	}

	// 12. Locked: Style Cannot Be Applied
	if s.Locked != nil {
		if err := s.Locked.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:locked"}}); err != nil {
			return fmt.Errorf("Style Locked: %w", err)
		}
	}

	// 13. Personal: E-Mail Message Text Style
	if s.Personal != nil {
		if err := s.Personal.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:personal"}}); err != nil {
			return fmt.Errorf("Style Personal: %w", err)
		}
	}

	// 14. PersonalCompose: E-Mail Message Composition Style
	if s.PersonalCompose != nil {
		if err := s.PersonalCompose.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:personalCompose"}}); err != nil {
			return fmt.Errorf("Style PersonalCompose: %w", err)
		}
	}

	// 15. PersonalReply: E-Mail Message Reply Style
	if s.PersonalReply != nil {
		if err := s.PersonalReply.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:personalReply"}}); err != nil {
			return fmt.Errorf("Style PersonalReply: %w", err)
		}
	}

	// 16. RevID: Revision Identifier for Style Definition
	if s.RevID != nil {
		if err := s.RevID.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:rsid"}}); err != nil {
			return fmt.Errorf("Style RevID: %w", err)
		}
	}

	// 17. ParaProp: Style Paragraph Properties
	if s.ParaProp != nil {
		if err := s.ParaProp.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:pPr"}}); err != nil {
			return fmt.Errorf("Style ParaProp: %w", err)
		}
	}

	// 18. RunProp: Run Properties
	if s.RunProp != nil {
		if err := s.RunProp.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:rPr"}}); err != nil {
			return fmt.Errorf("Style RunProp: %w", err)
		}
	}

	// 19. TableProp: Style Table Properties
	if s.TableProp != nil {
		if err := s.TableProp.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:tblPr"}}); err != nil {
			return fmt.Errorf("Style TableProp: %w", err)
		}
	}

	// 20. TableRowProp: Style Table Row Properties
	if s.TableRowProp != nil {
		if err := s.TableRowProp.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:trPr"}}); err != nil {
			return fmt.Errorf("Style TableRowProp: %w", err)
		}
	}

	// 21. TableCellProp: Style Table Cell Properties
	if s.TableCellProp != nil {
		if err := s.TableCellProp.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:tcPr"}}); err != nil {
			return fmt.Errorf("Style TableCellProp: %w", err)
		}
	}

	// 22. TableStylePr: Style Conditional Table Formatting Properties
	for _, tsPr := range s.TableStylePr {
		if err := tsPr.MarshalXML(e, xml.StartElement{Name: xml.Name{Local: "w:tblStylePr"}}); err != nil {
			return fmt.Errorf("Style TableStylePr: %w", err)
		}
	}

	return e.EncodeToken(start.End())
}
