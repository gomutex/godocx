package ctypes

// 单位转换函数
func InchesToTwips(inches float64) uint64 {
	return uint64(inches * 1440)
}

func CentimetersToTwips(cm float64) uint64 {
	return uint64(cm * 567)
}

func MillimetersToTwips(mm float64) uint64 {
	return uint64(mm * 56.7)
}

func PointsToTwips(points float64) int {
	return int(points * 20)
}
