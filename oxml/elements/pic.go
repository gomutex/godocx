package elements

type DrawingPositionType int

const (
	DrawingInlinePosition DrawingPositionType = iota
	DrawingAnchorPosition
)

type Pic struct {
	PositionType DrawingPositionType
}
