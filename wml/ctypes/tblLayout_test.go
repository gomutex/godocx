package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/gomutex/godocx/internal"
	"github.com/gomutex/godocx/wml/stypes"
)

func TestTableLayout_MarshalXML(t *testing.T) {
	layout := &TableLayout{LayoutType: internal.ToPtr(stypes.TableLayoutFixed)}

	expected := `<w:tblLayout w:type="fixed"></w:tblLayout>`

	var builder strings.Builder
	encoder := xml.NewEncoder(&builder)
	err := encoder.Encode(layout)
	if err != nil {
		t.Fatalf("Error encoding TableLayout: %v", err)
	}

	result := builder.String()
	if result != expected {
		t.Errorf("Unexpected XML. Expected: %s, Got: %s", expected, result)
	}
}

func TestLayout_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name       string
		inputXML   string
		expected   TableLayout
		expectFail bool // Whether unmarshalling is expected to fail
	}{
		{
			name:     "Test with Overlap Value `never`",
			inputXML: `<w:tblLayout w:type="fixed"></w:tblLayout>`,
			expected: TableLayout{LayoutType: internal.ToPtr(stypes.TableLayoutFixed)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result TableLayout
			err := xml.Unmarshal([]byte(tt.inputXML), &result)

			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			err = internal.ComparePtr("Type", tt.expected.LayoutType, result.LayoutType)
			if err != nil {
				t.Error(err)
			}

		})
	}
}
