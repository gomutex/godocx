package wml

// TableBorderPosition represents the position of a border in a table.
type TableBorderPosition string

const (
	TableBorderPositionLeft    TableBorderPosition = "w:left"
	TableBorderPositionRight   TableBorderPosition = "w:right"
	TableBorderPositionTop     TableBorderPosition = "w:top"
	TableBorderPositionBottom  TableBorderPosition = "w:bottom"
	TableBorderPositionInsideH TableBorderPosition = "w:insideH"
	TableBorderPositionInsideV TableBorderPosition = "w:insideV"
)

// TableCellBorderPosition represents the position of a border in a table cell.
type TableCellBorderPosition int

const (
	TableCellBorderPositionLeft    TableCellBorderPosition = 1
	TableCellBorderPositionRight   TableCellBorderPosition = 2
	TableCellBorderPositionTop     TableCellBorderPosition = 3
	TableCellBorderPositionBottom  TableCellBorderPosition = 4
	TableCellBorderPositionInsideH TableCellBorderPosition = 5
	TableCellBorderPositionInsideV TableCellBorderPosition = 6
	TableCellBorderPositionTl2br   TableCellBorderPosition = 7
	TableCellBorderPositionTr2bl   TableCellBorderPosition = 8
)

// ParagraphBorderPosition represents the position of a border in a paragraph.
type ParagraphBorderPosition int

const (
	ParagraphBorderPositionLeft    ParagraphBorderPosition = 1
	ParagraphBorderPositionRight   ParagraphBorderPosition = 2
	ParagraphBorderPositionTop     ParagraphBorderPosition = 3
	ParagraphBorderPositionBottom  ParagraphBorderPosition = 4
	ParagraphBorderPositionBetween ParagraphBorderPosition = 5
	ParagraphBorderPositionBar     ParagraphBorderPosition = 6
)
