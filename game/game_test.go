package game

import (
	"testing"
)

func TestGameRender(t *testing.T) {
	window := Window{5, 5}

	tt := []struct {
		name     string
		content  []string
		relx     int
		expected []string
	}{
		{
			name: "simple level",
			content: []string{
				"#####",
				"#   #",
				"#   #",
				"#   #",
				"#####",
			},
			relx: 0,
			expected: []string{
				"#####",
				"#   #",
				"#   #",
				"#   #",
				"#####",
			},
		},
		{
			name: "simple level with offset",
			content: []string{
				"#####",
				"#   #",
				"#   #",
				"#   #",
				"#####",
			},
			relx: 1,
			expected: []string{
				"#####",
				"#   #",
				"#   #",
				"#   #",
				"#####",
			},
		},
		{
			name: "large level",
			content: []string{
				"########",
				"#      #",
				"#      #",
				"#      #",
				"########",
			},
			relx: 0,
			expected: []string{
				"#####",
				"#    ",
				"#    ",
				"#    ",
				"#####",
			},
		},
		{
			name: "large level with offset",
			content: []string{
				"########",
				"#      #",
				"#      #",
				"#      #",
				"########",
			},
			relx: 1,
			expected: []string{
				"#####",
				"     ",
				"     ",
				"     ",
				"#####",
			},
		},
		{
			name: "level smaller than window height",
			content: []string{
				"#####",
				"#   #",
				"#####",
			},
			relx: 0,
			expected: []string{
				"     ",
				"     ",
				"#####",
				"#   #",
				"#####",
			},
		},
		{
			name: "level smaller than window height with offset",
			content: []string{
				"#########",
				"#       #",
				"#########",
			},
			relx: 1,
			expected: []string{
				"     ",
				"     ",
				"#####",
				"     ",
				"#####",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			level := NewLayer(tc.content)
			result := Render(tc.relx, window, level)

			if len(result) != len(tc.expected) {
				t.Fatalf("Expected %d lines, got %d", len(tc.expected), len(result))
			}

			for i := range result {
				if result[i] != tc.expected[i] {
					t.Errorf("Expected line %d to be %q, got %q", i, tc.expected[i], result[i])
				}
			}
		})
	}
}
