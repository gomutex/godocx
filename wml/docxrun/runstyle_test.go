package docxrun

import (
	"testing"

	"github.com/gomutex/godocx/elemtypes"
)

func TestNewRunStyle(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected *elemtypes.SingleStrVal
	}{
		{
			name:     "Custom RunStyle",
			value:    "CustomStyle",
			expected: elemtypes.NewSingleStrVal("CustomStyle"),
		},
		{
			name:     "Empty RunStyle",
			value:    "",
			expected: elemtypes.NewSingleStrVal(""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewRunStyle(tt.value)

			if result.Val != tt.expected.Val {
				t.Errorf("Expected Val %s but got %s", tt.expected.Val, result.Val)
			}
		})
	}
}

func TestDefaultRunStyle(t *testing.T) {
	expected := elemtypes.NewSingleStrVal("Normal")
	result := DefaultRunStyle()

	if result.Val != expected.Val {
		t.Errorf("Expected Val %s but got %s", expected.Val, result.Val)
	}
}
