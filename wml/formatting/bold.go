package formatting

import "github.com/gomutex/godocx/elemtypes"

func DefaultBold() *elemtypes.NullBoolElem {
	return elemtypes.DefaultNullBoolElem("w:b")
}

func NewBold(value bool) *elemtypes.NullBoolElem {
	return elemtypes.NewNullBoolElem("w:b", value)
}
