package dml

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/types"
)

func TestMarshalNonVisualPicProp(t *testing.T) {
	tests := []struct {
		prop        *NonVisualPicProp
		expectedXML string
	}{
		{
			prop: &NonVisualPicProp{
				CNvPr: &CNvPr{
					ID:          "1",
					Name:        "Pic1",
					Description: "Description of Pic1",
				},
				CNvPicPr: &CNvPicPr{
					PicLocks: &PicLocks{
						NoChangeAspect: types.NewOptBool(true),
					},
				},
			},
			expectedXML: `<pic:nvPicPr><pic:cNvPr id="1" name="Pic1" descr="Description of Pic1"></pic:cNvPr><pic:cNvPicPr><a:picLocks noChangeAspect="1"></a:picLocks></pic:cNvPicPr></pic:nvPicPr>`,
		},
		{
			prop: &NonVisualPicProp{
				CNvPr: &CNvPr{
					ID:          "2",
					Name:        "Pic2",
					Description: "Description of Pic2",
				},
				CNvPicPr: nil,
			},
			expectedXML: `<pic:nvPicPr><pic:cNvPr id="2" name="Pic2" descr="Description of Pic2"></pic:cNvPr></pic:nvPicPr>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.expectedXML, func(t *testing.T) {
			generatedXML, err := xml.Marshal(tt.prop)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			if strings.TrimSpace(string(generatedXML)) != tt.expectedXML {
				t.Errorf("Expected XML:\n%s\nBut got:\n%s", tt.expectedXML, generatedXML)
			}
		})
	}
}

func TestUnmarshalNonVisualPicProp(t *testing.T) {
	tests := []struct {
		inputXML string
		expected NonVisualPicProp
	}{
		{
			inputXML: `<pic:nvPicPr><pic:cNvPr id="1" name="Pic1" descr="Description of Pic1"></pic:cNvPr><pic:cNvPicPr><a:picLocks noChangeAspect="1"></a:picLocks></pic:cNvPicPr></pic:nvPicPr>`,
			expected: NonVisualPicProp{
				CNvPr: &CNvPr{
					ID:          "1",
					Name:        "Pic1",
					Description: "Description of Pic1",
				},
				CNvPicPr: &CNvPicPr{
					PicLocks: &PicLocks{
						NoChangeAspect: types.NewOptBool(true),
					},
				},
			},
		},
		{
			inputXML: `<pic:nvPicPr><pic:cNvPr id="2" name="Pic2" descr="Description of Pic2"></pic:cNvPr></pic:nvPicPr>`,
			expected: NonVisualPicProp{
				CNvPr: &CNvPr{
					ID:          "2",
					Name:        "Pic2",
					Description: "Description of Pic2",
				},
				CNvPicPr: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.inputXML, func(t *testing.T) {
			var prop NonVisualPicProp

			err := xml.Unmarshal([]byte(tt.inputXML), &prop)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if (prop.CNvPr == nil && tt.expected.CNvPr != nil) || (prop.CNvPr != nil && tt.expected.CNvPr == nil) {
				t.Errorf("Expected CNvPr to be %v, but got %v", tt.expected.CNvPr, prop.CNvPr)
			} else if prop.CNvPr != nil && tt.expected.CNvPr != nil {
				if prop.CNvPr.ID != tt.expected.CNvPr.ID {
					t.Errorf("Expected CNvPr.ID %s, but got %s", tt.expected.CNvPr.ID, prop.CNvPr.ID)
				}
				if prop.CNvPr.Name != tt.expected.CNvPr.Name {
					t.Errorf("Expected CNvPr.Name %s, but got %s", tt.expected.CNvPr.Name, prop.CNvPr.Name)
				}
				if prop.CNvPr.Description != tt.expected.CNvPr.Description {
					t.Errorf("Expected CNvPr.Description %s, but got %s", tt.expected.CNvPr.Description, prop.CNvPr.Description)
				}
			}

			if (prop.CNvPicPr == nil && tt.expected.CNvPicPr != nil) || (prop.CNvPicPr != nil && tt.expected.CNvPicPr == nil) {
				t.Errorf("Expected CNvPicPr to be %v, but got %v", tt.expected.CNvPicPr, prop.CNvPicPr)
			} else if prop.CNvPicPr != nil && tt.expected.CNvPicPr != nil {
				if (prop.CNvPicPr.PicLocks == nil && tt.expected.CNvPicPr.PicLocks != nil) || (prop.CNvPicPr.PicLocks != nil && tt.expected.CNvPicPr.PicLocks == nil) {
					t.Errorf("Expected PicLocks to be %v, but got %v", tt.expected.CNvPicPr.PicLocks, prop.CNvPicPr.PicLocks)
				} else if prop.CNvPicPr.PicLocks != nil && tt.expected.CNvPicPr.PicLocks != nil {
					if prop.CNvPicPr.PicLocks.NoChangeAspect != tt.expected.CNvPicPr.PicLocks.NoChangeAspect {
						t.Errorf("Expected PicLocks.NoChangeAspect %v, but got %v", tt.expected.CNvPicPr.PicLocks.NoChangeAspect, prop.CNvPicPr.PicLocks.NoChangeAspect)
					}
				}
			}
		})
	}
}
