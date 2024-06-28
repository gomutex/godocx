package ctypes

import (
	"encoding/xml"
	"testing"

	"github.com/gomutex/godocx/internal"
)

func TestTrackChangeNum_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    TrackChangeNum
		expected string
	}{
		{
			name: "With all attributes",
			input: TrackChangeNum{
				ID:       123,
				Author:   "John Doe",
				Date:     internal.ToPtr("2023-06-18T12:34:56Z"),
				Original: internal.ToPtr("42"),
			},
			expected: `<TrackChangeNum w:id="123" w:author="John Doe" w:date="2023-06-18T12:34:56Z" w:original="42"></TrackChangeNum>`,
		},
		{
			name: "Without optional attributes",
			input: TrackChangeNum{
				ID:     124,
				Author: "Jane Doe",
			},
			expected: `<TrackChangeNum w:id="124" w:author="Jane Doe"></TrackChangeNum>`,
		},
		{
			name: "With only date attribute",
			input: TrackChangeNum{
				ID:     125,
				Author: "Alice",
				Date:   internal.ToPtr("2024-06-18T12:34:56Z"),
			},
			expected: `<TrackChangeNum w:id="125" w:author="Alice" w:date="2024-06-18T12:34:56Z"></TrackChangeNum>`,
		},
		{
			name: "With only original attribute",
			input: TrackChangeNum{
				ID:       126,
				Author:   "Bob",
				Original: internal.ToPtr("99"),
			},
			expected: `<TrackChangeNum w:id="126" w:author="Bob" w:original="99"></TrackChangeNum>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := xml.Marshal(tt.input)
			if err != nil {
				t.Fatalf("Error marshaling to XML: %v", err)
			}

			if string(output) != tt.expected {
				t.Errorf("Expected XML:\n%s\nGot:\n%s", tt.expected, string(output))
			}
		})
	}
}

func TestTrackChangeNum_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		inputXML string
		expected TrackChangeNum
	}{
		{
			name:     "With all attributes",
			inputXML: `<TrackChangeNum w:id="123" w:author="John Doe" w:date="2023-06-18T12:34:56Z" w:original="42"></TrackChangeNum>`,
			expected: TrackChangeNum{
				ID:       123,
				Author:   "John Doe",
				Date:     internal.ToPtr("2023-06-18T12:34:56Z"),
				Original: internal.ToPtr("42"),
			},
		},
		{
			name:     "Without optional attributes",
			inputXML: `<TrackChangeNum w:id="124" w:author="Jane Doe"></TrackChangeNum>`,
			expected: TrackChangeNum{
				ID:     124,
				Author: "Jane Doe",
			},
		},
		{
			name:     "With only date attribute",
			inputXML: `<TrackChangeNum w:id="125" w:author="Alice" w:date="2024-06-18T12:34:56Z"></TrackChangeNum>`,
			expected: TrackChangeNum{
				ID:     125,
				Author: "Alice",
				Date:   internal.ToPtr("2024-06-18T12:34:56Z"),
			},
		},
		{
			name:     "With only original attribute",
			inputXML: `<TrackChangeNum w:id="126" w:author="Bob" w:original="99"></TrackChangeNum>`,
			expected: TrackChangeNum{
				ID:       126,
				Author:   "Bob",
				Original: internal.ToPtr("99"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result TrackChangeNum

			err := xml.Unmarshal([]byte(tt.inputXML), &result)
			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			// Compare ID
			if result.ID != tt.expected.ID {
				t.Errorf("Expected ID %d but got %d", tt.expected.ID, result.ID)
			}

			// Compare Author
			if result.Author != tt.expected.Author {
				t.Errorf("Expected Author %s but got %s", tt.expected.Author, result.Author)
			}

			// Compare Date
			if tt.expected.Date != nil {
				if result.Date == nil || *result.Date != *tt.expected.Date {
					t.Errorf("Expected Date %s but got %v", *tt.expected.Date, result.Date)
				}
			} else if result.Date != nil {
				t.Errorf("Expected Date nil but got %v", result.Date)
			}

			// Compare Original
			if tt.expected.Original != nil {
				if result.Original == nil || *result.Original != *tt.expected.Original {
					t.Errorf("Expected Original %s but got %v", *tt.expected.Original, result.Original)
				}
			} else if result.Original != nil {
				t.Errorf("Expected Original nil but got %v", result.Original)
			}
		})
	}
}
