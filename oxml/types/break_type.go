package types

type BreakType string

const (
	Page         BreakType = "page"
	Column       BreakType = "column"
	TextWrapping BreakType = "textWrapping"
	Unsupported  BreakType = "unsupported"
)
