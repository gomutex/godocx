package dml

type RelativeFromHType string

const (
	RelativeFromHTypeCharacter     RelativeFromHType = "character"
	RelativeFromHTypeColumn        RelativeFromHType = "column"
	RelativeFromHTypeInsideMargin  RelativeFromHType = "insideMargin"
	RelativeFromHTypeLeftMargin    RelativeFromHType = "leftMargin"
	RelativeFromHTypeMargin        RelativeFromHType = "margin"
	RelativeFromHTypeOutsizeMargin RelativeFromHType = "outsizeMargin"
	RelativeFromHTypePage          RelativeFromHType = "page"
	RelativeFromHTypeRightMargin   RelativeFromHType = "rightMargin"
)

type RelativeFromVType string

const (
	RelativeFromVTypeBottomMargin  RelativeFromVType = "bottomMargin"
	RelativeFromVTypeInsideMargin  RelativeFromVType = "insideMargin"
	RelativeFromVTypeLine          RelativeFromVType = "line"
	RelativeFromVTypeMargin        RelativeFromVType = "margin"
	RelativeFromVTypeOutsizeMargin RelativeFromVType = "outsizeMargin"
	RelativeFromVTypePage          RelativeFromVType = "page"
	RelativeFromVTypeParagraph     RelativeFromVType = "paragraph"
	RelativeFromVTypeTopMargin     RelativeFromVType = "topMargin"
)
