package docxrun

import "github.com/gomutex/godocx/wml/ctypes"

// NewRunStyle creates a new RunStyle.
func NewRunStyle(value string) *ctypes.CTString {
	return ctypes.NewCTString(value)
}

// DefaultRunStyle creates the default RunStyle with the value "Normal".
func DefaultRunStyle() *ctypes.CTString {
	return ctypes.NewCTString("Normal")
}
