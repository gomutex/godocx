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
}
