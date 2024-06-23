package shapes

import (
	"encoding/xml"
	"testing"

	"github.com/gomutex/godocx/dml/dmlst"
	"github.com/gomutex/godocx/internal"
)

func TestTile_MarshalXML(t *testing.T) {
	tests := []struct {
		tile     Tile
		expected string
	}{
		{Tile{}, `<a:tile></a:tile>`},                                                            // Empty Tile should produce empty element
		{Tile{Tx: internal.ToPtr(int64(100))}, `<a:tile tx="100"></a:tile>`},                     // Only Tx attribute
		{Tile{Ty: internal.ToPtr(int64(200))}, `<a:tile ty="200"></a:tile>`},                     // Only Ty attribute
		{Tile{Sx: internal.ToPtr(50)}, `<a:tile sx="50"></a:tile>`},                              // Only Sx attribute
		{Tile{Sy: internal.ToPtr(75)}, `<a:tile sy="75"></a:tile>`},                              // Only Sy attribute
		{Tile{Flip: internal.ToPtr(dmlst.TileFlipModeHorizontal)}, `<a:tile flip="x"></a:tile>`}, // Only Flip attribute
		{Tile{Algn: internal.ToPtr(dmlst.RectAlignmentCenter)}, `<a:tile algn="ctr"></a:tile>`},  // Only Algn attribute
		{Tile{
			Tx:   internal.ToPtr(int64(100)),
			Ty:   internal.ToPtr(int64(200)),
			Sx:   internal.ToPtr(50),
			Sy:   internal.ToPtr(75),
			Flip: internal.ToPtr(dmlst.TileFlipModeBoth),
			Algn: internal.ToPtr(dmlst.RectAlignmentBottomRight),
		}, `<a:tile tx="100" ty="200" sx="50" sy="75" flip="xy" algn="br"></a:tile>`}, // All attributes
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			output, err := xml.Marshal(tt.tile)
			if err != nil {
				t.Fatalf("Error marshalling XML: %v", err)
			}

			if string(output) != tt.expected {
				t.Errorf("Expected XML:\n%s\nGot XML:\n%s", tt.expected, output)
			}
		})
	}
}

func TestTile_UnmarshalXML(t *testing.T) {
	tests := []struct {
		xmlInput    string
		expected    Tile
		expectError bool
	}{
		{`<a:tile></a:tile>`, Tile{}, false},
		{`<a:tile tx="100"></a:tile>`, Tile{Tx: internal.ToPtr(int64(100))}, false},
		{`<a:tile ty="200"></a:tile>`, Tile{Ty: internal.ToPtr(int64(200))}, false},
		{`<a:tile sx="50"></a:tile>`, Tile{Sx: internal.ToPtr(50)}, false},
		{`<a:tile sy="75"></a:tile>`, Tile{Sy: internal.ToPtr(75)}, false},
		{`<a:tile flip="x"></a:tile>`, Tile{Flip: internal.ToPtr(dmlst.TileFlipModeHorizontal)}, false},
		{`<a:tile algn="ctr"></a:tile>`, Tile{Algn: internal.ToPtr(dmlst.RectAlignmentCenter)}, false},
		{`<a:tile tx="100" ty="200" sx="50" sy="75" flip="xy" algn="br"></a:tile>`, Tile{
			Tx:   internal.ToPtr(int64(100)),
			Ty:   internal.ToPtr(int64(200)),
			Sx:   internal.ToPtr(50),
			Sy:   internal.ToPtr(75),
			Flip: internal.ToPtr(dmlst.TileFlipModeBoth),
			Algn: internal.ToPtr(dmlst.RectAlignmentBottomRight),
		}, false},
		{`<a:tile tx="invalid"></a:tile>`, Tile{}, true},         // Expect error for invalid Tx attribute
		{`<a:tile unknownAttr="value"></a:tile>`, Tile{}, false}, // Unknown attribute should be ignored
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			var tile Tile
			err := xml.Unmarshal([]byte(tt.xmlInput), &tile)
			if tt.expectError {
				if err == nil {
					t.Fatalf("Expected error but got none")
				}
				// Test passed as error was expected and occurred
				return
			}

			if err != nil {
				t.Fatalf("Error unmarshalling XML: %v", err)
			}

			if err := internal.ComparePtr("Tx", tt.expected.Tx, tile.Tx); err != nil {
				t.Errorf(err.Error())
			}
			if err := internal.ComparePtr("Ty", tt.expected.Ty, tile.Ty); err != nil {
				t.Errorf(err.Error())
			}
			if err := internal.ComparePtr("Sx", tt.expected.Sx, tile.Sx); err != nil {
				t.Errorf(err.Error())
			}
			if err := internal.ComparePtr("Sy", tt.expected.Sy, tile.Sy); err != nil {
				t.Errorf(err.Error())
			}
			if err := internal.ComparePtr("Flip", tt.expected.Flip, tile.Flip); err != nil {
				t.Errorf(err.Error())
			}
			if err := internal.ComparePtr("Algn", tt.expected.Algn, tile.Algn); err != nil {
				t.Errorf(err.Error())
			}
		})
	}
}
