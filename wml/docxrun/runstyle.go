package docxrun

import "github.com/gomutex/godocx/elemtypes"

// NewRunStyle creates a new RunStyle.
func NewRunStyle(value string) *elemtypes.SingleStrVal {
	return elemtypes.NewSingleStrVal(value)
}

// DefaultRunStyle creates the default RunStyle with the value "Normal".
func DefaultRunStyle() *elemtypes.SingleStrVal {
	return elemtypes.NewSingleStrVal("Normal")
}
