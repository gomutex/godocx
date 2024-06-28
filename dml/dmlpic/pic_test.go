package dmlpic

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/dml/dmlct"
	"github.com/gomutex/godocx/dml/dmlprops"
	"github.com/gomutex/godocx/dml/dmlst"
	"github.com/gomutex/godocx/dml/shapes"
)

func TestPicMarshalXML(t *testing.T) {
	p := &Pic{
		NonVisualPicProp: NonVisualPicProp{
			CNvPr: dmlct.CNvPr{
				ID:          1,
				Name:        "Pic 1",
				Description: "Description",
			},
			CNvPicPr: CNvPicPr{
				PicLocks: &dmlprops.PicLocks{
					NoChangeAspect:     dmlst.NewOptBool(true),
					NoChangeArrowheads: dmlst.NewOptBool(true),
				},
			},
		},
		BlipFill: BlipFill{
			Blip: &Blip{
				EmbedID: "rId1",
			},
			FillModeProps: FillModeProps{
				Stretch: &shapes.Stretch{
					FillRect: &dmlct.RelativeRect{},
				},
			},
		},
		PicShapeProp: PicShapeProp{
			TransformGroup: &TransformGroup{
				Offset: &Offset{
					X: 0,
					Y: 0,
				},
				Extent: &dmlct.PSize2D{
					Width:  100000,
					Height: 100000,
				},
			},
			PresetGeometry: &PresetGeometry{
				Preset: "rect",
			},
		},
	}

	// Expected XML output
	expectedXML := `<pic:pic xmlns:pic="http://schemas.openxmlformats.org/drawingml/2006/picture"><pic:nvPicPr><pic:cNvPr id="1" name="Pic 1" descr="Description"></pic:cNvPr><pic:cNvPicPr><a:picLocks noChangeAspect="1" noChangeArrowheads="1"></a:picLocks></pic:cNvPicPr></pic:nvPicPr><pic:blipFill><a:blip r:embed="rId1"></a:blip><a:stretch><a:fillRect></a:fillRect></a:stretch></pic:blipFill><pic:spPr><a:xfrm><a:off x="0" y="0"></a:off><a:ext cx="100000" cy="100000"></a:ext></a:xfrm><a:prstGeom prst="rect"></a:prstGeom></pic:spPr></pic:pic>`

	output, err := xml.Marshal(p)
	if err != nil {
		t.Fatalf("Error marshaling Pic to XML: %v", err)
	}

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedXML) {
		t.Errorf("Generated XML does not match expected XML.\nExpected:\n%s\nGenerated:\n%s", expectedXML, output)
	}
}

func TestPicUnmarshalXML(t *testing.T) {
	xmlData := `<pic:pic xmlns:pic="http://schemas.openxmlformats.org/drawingml/2006/picture">
        	<pic:nvPicPr>
        		<pic:cNvPr id="1" name="Pic 1" descr="Description"></pic:cNvPr>
        		<pic:cNvPicPr>
        			<a:picLocks noChangeAspect="1" noChangeArrowheads="1"></a:picLocks>
        		</pic:cNvPicPr>
        	</pic:nvPicPr>
        	<pic:blipFill>
        		<a:blip r:embed="rId1"></a:blip>
        		<a:stretch>
        			<a:fillRect></a:fillRect>
        		</a:stretch>
        	</pic:blipFill>
        	<pic:spPr>
        		<a:xfrm>
        			<a:off x="0" y="0"></a:off>
        			<a:ext cx="100000" cy="100000"></a:ext>
        		</a:xfrm>
        		<a:prstGeom prst="rect"></a:prstGeom>
        	</pic:spPr>
        </pic:pic>`

	var pic Pic

	err := xml.NewDecoder(strings.NewReader(xmlData)).Decode(&pic)
	if err != nil {
		t.Errorf("Error decoding XML: %v", err)
	}

	checkNotNil := func(fieldName string, fieldValue interface{}) {
		if fieldValue == nil {
			t.Errorf("Expected field '%s' to be unmarshaled, but it was nil", fieldName)
		}
	}

	checkNotNil("NonVisualPicProp", pic.NonVisualPicProp)
	checkNotNil("BlipFill", pic.BlipFill)
	checkNotNil("PicShapeProp", pic.PicShapeProp)
}
