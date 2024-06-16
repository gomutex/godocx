package formatting

import "github.com/gomutex/godocx/elemtypes"

func NewBold(value bool) *elemtypes.NullBoolElem {
	return elemtypes.NewNullBoolElem(value)
}
