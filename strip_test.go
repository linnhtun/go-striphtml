package striphtml

import "testing"

func TestStripHtml(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Simple HTML",
			input:    "<p>Hello, <b>world</b>!</p>",
			expected: "Hello, world!",
		},
		{
			name:     "Nested Tags",
			input:    "<div><span>Text</span> with <i>italics</i>.</div>",
			expected: "Text with italics.",
		},
		{
			name:     "Non-breaking spaces",
			input:    "<p>Text&nbsp;with&nbsp;spaces.</p>",
			expected: "Text with spaces.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StripHTML(tt.input)
			if err != nil {
				t.Errorf("StripHTML() error = %v", err)
				return
			}
			if result != tt.expected {
				t.Errorf("StripHTML() = %v, want %v", result, tt.expected)
			}
		})
	}
}
