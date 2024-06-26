package styles

import (
	"github.com/gomutex/godocx/wml/ctypes"
)

type Style struct {
	//Sequence:

	//1. Primary Style Name
	Name *ctypes.CTString `xml:"name,omitempty"`

	//2. Alternate Style Names
	Alias *ctypes.CTString `xml:"alias,omitempty"`

	//3. Parent Style ID
	BasedOn *ctypes.CTString `xml:"basedOn,omitempty"`

	//4. Style For Next Paragraph
	Next *ctypes.CTString `xml:"next,omitempty"`

	//5. Linked Style Reference
	Link *ctypes.CTString `xml:"link,omitempty"`

	//6.Automatically Merge User Formatting Into Style Definition
	AutoRedefine *ctypes.OnOff `xml:"autoRedefine,omitempty"`

	//7.Hide Style From User Interface
	Hidden *ctypes.OnOff `xml:"hidden,omitempty"`

	//8.Optional User Interface Sorting Order
	UIPriority *ctypes.DecimalNum `xml:"uiPriority,omitempty"`

	// 9. Hide Style From Main User Interface
	SemiHidden *ctypes.OnOff `xml:"semiHidden,omitempty"`

	// 10. Remove Semi-Hidden Property When Style Is Used
	UnhideWhenUsed *ctypes.OnOff `xml:"unhideWhenUsed,omitempty"`

	// 11. Primary Style
	QFormat *ctypes.OnOff `xml:"qFormat,omitempty"`

	// 12. Style Cannot Be Applied
	Locked *ctypes.OnOff `xml:"locked,omitempty"`

	// 13. E-Mail Message Text Style
	Personal *ctypes.OnOff `xml:"personal,omitempty"`

	// 14. E-Mail Message Composition Style
	PersonalCompose *ctypes.OnOff `xml:"personalCompose,omitempty"`

	// 15. E-Mail Message Reply Style
	PersonalReply *ctypes.OnOff `xml:"personalReply,omitempty"`

	//16. Revision Identifier for Style Definition
	RevID *ctypes.LongHexNum `xml:"rsid,omitempty"`
}
