package dmlprops

import (
	"encoding/xml"
	"testing"

	"github.com/gomutex/godocx/dml/dmlst"
)

func TestMarshalPicLocks(t *testing.T) {
	tests := []struct {
		picLocks    *PicLocks
		expectedXML string
	}{
		{
			picLocks: &PicLocks{
				DisallowShadowGrouping: dmlst.NewOptBool(true),
				NoSelect:               dmlst.NewOptBool(false),
				NoRot:                  dmlst.NewOptBool(true),
				NoChangeAspect:         dmlst.NewOptBool(false),
				NoMove:                 dmlst.NewOptBool(true),
				NoResize:               dmlst.NewOptBool(true),
				NoEditPoints:           dmlst.NewOptBool(false),
				NoAdjustHandles:        dmlst.NewOptBool(true),
				NoChangeArrowheads:     dmlst.NewOptBool(false),
				NoChangeShapeType:      dmlst.NewOptBool(true),
				NoCrop:                 dmlst.NewOptBool(false),
			},
			expectedXML: `<a:picLocks noGrp="1" noSelect="0" noRot="1" noChangeAspect="0" noMove="1" noResize="1" noEditPoints="0" noAdjustHandles="1" noChangeArrowheads="0" noChangeShapeType="1" noCrop="0"></a:picLocks>`,
		},
		{
			picLocks:    &PicLocks{},
			expectedXML: `<a:picLocks></a:picLocks>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.expectedXML, func(t *testing.T) {
			generatedXML, err := xml.Marshal(tt.picLocks)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			if string(generatedXML) != tt.expectedXML {
				t.Errorf("Expected XML:\n%s\nBut got:\n%s", tt.expectedXML, generatedXML)
			}
		})
	}
}

func TestUnmarshalPicLocks(t *testing.T) {
	tests := []struct {
		inputXML      string
		expectedLocks PicLocks
	}{
		{
			inputXML: `<a:picLocks noGrp="1" noSelect="0" noRot="1" noChangeAspect="0" noMove="1" noResize="1" noEditPoints="0" noAdjustHandles="1" noChangeArrowheads="0" noChangeShapeType="1" noCrop="0"></a:picLocks>`,
			expectedLocks: PicLocks{
				DisallowShadowGrouping: dmlst.NewOptBool(true),
				NoSelect:               dmlst.NewOptBool(false),
				NoRot:                  dmlst.NewOptBool(true),
				NoChangeAspect:         dmlst.NewOptBool(false),
				NoMove:                 dmlst.NewOptBool(true),
				NoResize:               dmlst.NewOptBool(true),
				NoEditPoints:           dmlst.NewOptBool(false),
				NoAdjustHandles:        dmlst.NewOptBool(true),
				NoChangeArrowheads:     dmlst.NewOptBool(false),
				NoChangeShapeType:      dmlst.NewOptBool(true),
				NoCrop:                 dmlst.NewOptBool(false),
			},
		},
		{
			inputXML:      `<a:picLocks></a:picLocks>`,
			expectedLocks: PicLocks{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			var picLocks PicLocks

			err := xml.Unmarshal([]byte(tt.inputXML), &picLocks)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			comparePicLocks(t, tt.expectedLocks, picLocks)
		})
	}
}

func comparePicLocks(t *testing.T, expected, actual PicLocks) {
	if expected.DisallowShadowGrouping != actual.DisallowShadowGrouping {
		t.Errorf("Expected DisallowShadowGrouping %v, but got %v", expected.DisallowShadowGrouping, actual.DisallowShadowGrouping)
	}
	if expected.NoSelect != actual.NoSelect {
		t.Errorf("Expected NoSelect %v, but got %v", expected.NoSelect, actual.NoSelect)
	}
	if expected.NoRot != actual.NoRot {
		t.Errorf("Expected NoRot %v, but got %v", expected.NoRot, actual.NoRot)
	}
	if expected.NoChangeAspect != actual.NoChangeAspect {
		t.Errorf("Expected NoChangeAspect %v, but got %v", expected.NoChangeAspect, actual.NoChangeAspect)
	}
	if expected.NoMove != actual.NoMove {
		t.Errorf("Expected NoMove %v, but got %v", expected.NoMove, actual.NoMove)
	}
	if expected.NoResize != actual.NoResize {
		t.Errorf("Expected NoResize %v, but got %v", expected.NoResize, actual.NoResize)
	}
	if expected.NoEditPoints != actual.NoEditPoints {
		t.Errorf("Expected NoEditPoints %v, but got %v", expected.NoEditPoints, actual.NoEditPoints)
	}
	if expected.NoAdjustHandles != actual.NoAdjustHandles {
		t.Errorf("Expected NoAdjustHandles %v, but got %v", expected.NoAdjustHandles, actual.NoAdjustHandles)
	}
	if expected.NoChangeArrowheads != actual.NoChangeArrowheads {
		t.Errorf("Expected NoChangeArrowheads %v, but got %v", expected.NoChangeArrowheads, actual.NoChangeArrowheads)
	}
	if expected.NoChangeShapeType != actual.NoChangeShapeType {
		t.Errorf("Expected NoChangeShapeType %v, but got %v", expected.NoChangeShapeType, actual.NoChangeShapeType)
	}
	if expected.NoCrop != actual.NoCrop {
		t.Errorf("Expected NoCrop %v, but got %v", expected.NoCrop, actual.NoCrop)
	}
}
