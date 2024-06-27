package ctypes

// NewRunStyle creates a new RunStyle.
func NewRunStyle(value string) *CTString {
	return NewCTString(value)
}

// DefaultRunStyle creates the default RunStyle with the value "Normal".
func DefaultRunStyle() *CTString {
	return NewCTString("Normal")
}
