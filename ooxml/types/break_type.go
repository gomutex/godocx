package types

type BreakType string

const (
	BreakTypePage         BreakType = "page"
	BreakTypeColumn       BreakType = "column"
	BreakTypeTextWrapping BreakType = "textWrapping"
	BreakTypeUnsupported  BreakType = "unsupported"
)
