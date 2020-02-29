package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizeMarkdown(t *testing.T) {
	tests := []struct {
		name     string
		markdown string
		expected string
	}{
		{
			name:     "preserve double spaces",
			markdown: "Lorem ipsum dolor sit amet,  \nconsectetur adipiscing elit",
			expected: "Lorem ipsum dolor sit amet,  \nconsectetur adipiscing elit",
		},
		{
			name:     "remove trailing space",
			markdown: "Lorem ipsum dolor sit amet, \nconsectetur adipiscing elit",
			expected: "Lorem ipsum dolor sit amet,\nconsectetur adipiscing elit",
		},
		{
			name:     "remove blank line with only doubl spaces",
			markdown: "Lorem ipsum dolor sit amet,\n  \nconsectetur adipiscing elit",
			expected: "Lorem ipsum dolor sit amet,\nconsectetur adipiscing elit",
		},
		{
			name:     "remove multiple consecutive blank lines",
			markdown: "Lorem ipsum dolor sit amet,\n\nconsectetur adipiscing elit",
			expected: "Lorem ipsum dolor sit amet,\n\nconsectetur adipiscing elit",
		},
		{
			name:     "remove multiple consecutive blank lines",
			markdown: "Lorem ipsum dolor sit amet,\n\n\nconsectetur adipiscing elit",
			expected: "Lorem ipsum dolor sit amet,\n\nconsectetur adipiscing elit",
		},
		{
			name:     "remove multiple consecutive blank lines",
			markdown: "Lorem ipsum dolor sit amet,\n\n\n\nconsectetur adipiscing elit",
			expected: "Lorem ipsum dolor sit amet,\n\nconsectetur adipiscing elit",
		},
		{
			name:     "remove multiple consecutive blank lines",
			markdown: "Lorem ipsum dolor sit amet,\n\n\n\n\nconsectetur adipiscing elit",
			expected: "Lorem ipsum dolor sit amet,\n\nconsectetur adipiscing elit",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := sanitize(tt.markdown)

			assert.Equal(tt.expected, actual)
		})
	}
}

func TestFenceCodeBlock(t *testing.T) {
	tests := []struct {
		name      string
		code      string
		language  string
		expected  string
		extraline bool
	}{
		{
			name:      "single line",
			code:      "foo",
			language:  "json",
			expected:  "`foo`",
			extraline: false,
		},
		{
			name:      "single line",
			code:      "\"bar\"",
			language:  "hcl",
			expected:  "`\"bar\"`",
			extraline: false,
		},
		{
			name:      "single line",
			code:      "fuzz_buzz",
			language:  "",
			expected:  "`fuzz_buzz`",
			extraline: false,
		},
		{
			name:      "multi lines",
			code:      "[\n  \"foo\",\n  \"bar\",\n  \"baz\"\n]",
			language:  "json",
			expected:  "\n\n```json\n[\n  \"foo\",\n  \"bar\",\n  \"baz\"\n]\n```\n",
			extraline: true,
		},
		{
			name:      "multi lines",
			code:      "variable \"foo\" {\n  default = true\n}",
			language:  "hcl",
			expected:  "\n\n```hcl\nvariable \"foo\" {\n  default = true\n}\n```\n",
			extraline: true,
		},
		{
			name:      "multi lines",
			code:      "Usage:\n\nExample of 'foo_bar' module in `foo_bar.tf`.\n\n- list item 1\n- list item 2",
			language:  "",
			expected:  "\n\n```\nUsage:\n\nExample of 'foo_bar' module in `foo_bar.tf`.\n\n- list item 1\n- list item 2\n```\n",
			extraline: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			actual, extraline := printFencedCodeBlock(tt.code, tt.language)

			assert.Equal(tt.expected, actual)
			assert.Equal(tt.extraline, extraline)
		})
	}
}
