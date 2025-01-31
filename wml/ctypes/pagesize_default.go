package ctypes

import "github.com/gomutex/godocx/wml/stypes"

/*
code  |  name    |  size
---------------------------
1     |  A4      |  210 × 297 mm
2     |  Letter  |  8.5 × 11 inches
3     |  Legal   |  8.5 × 14 inches
4     |  A3      |  297 × 420 mm
5     |  B4      |  250 × 353 mm
6     |  B5      |  176 × 250 mm
9     |  A5      |  148 × 210 mm
11    |  A6      |  105 × 148 mm
*/

var (
	// A1
	A1Width  = uint64(23386) // 594mm in twips
	A1Height = uint64(33110) // 841mm in twips
	A1Code   = 4
	A1       = &PageSize{
		Width:  &A1Width,
		Height: &A1Height,
		Orient: stypes.PageOrientPortrait,
		Code:   &A1Code,
	}

	// A2
	A2Width  = uint64(16535) // 420mm in twips
	A2Height = uint64(23386) // 594mm in twips
	A2Code   = 4
	A2       = &PageSize{
		Width:  &A2Width,
		Height: &A2Height,
		Orient: stypes.PageOrientPortrait,
		Code:   &A2Code,
	}

	// A3
	A3Width  = uint64(11693) // 297mm in twips
	A3Height = uint64(16535) // 420mm in twips
	A3Code   = 4
	A3       = &PageSize{
		Width:  &A3Width,
		Height: &A3Height,
		Orient: stypes.PageOrientPortrait,
		Code:   &A3Code,
	}

	// A4
	A4Width  = uint64(11906) // 210mm in twips
	A4Height = uint64(16838) // 297mm in twips
	A4Code   = 1
	A4       = &PageSize{
		Width:  &A4Width,
		Height: &A4Height,
		Orient: stypes.PageOrientPortrait,
		Code:   &A4Code,
	}

	// A5
	A5Width  = uint64(8268)  // 148mm in twips
	A5Height = uint64(11693) // 210mm in twips
	A5Code   = 1
	A5       = &PageSize{
		Width:  &A5Width,
		Height: &A5Height,
		Orient: stypes.PageOrientPortrait,
		Code:   &A5Code,
	}
)
