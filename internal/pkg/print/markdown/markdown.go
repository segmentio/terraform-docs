package markdown

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/segmentio/terraform-docs/internal/pkg/print"

	"mvdan.cc/xurls/v2"
)

// Sanitize cleans a Markdown document to soothe linters.
func Sanitize(markdown string) string {
	result := markdown

	// Preserve double spaces at the end of the line
	result = regexp.MustCompile(` {2}(\r?\n)`).ReplaceAllString(result, "‡‡$1")

	// Remove trailing spaces from the end of lines
	result = regexp.MustCompile(` +(\r?\n)`).ReplaceAllString(result, "$1")
	result = regexp.MustCompile(` +$`).ReplaceAllLiteralString(result, "")

	// Preserve double spaces at the end of the line
	result = regexp.MustCompile(`‡‡(\r?\n)`).ReplaceAllString(result, "  $1")

	// Remove blank line with only double spaces in it
	result = regexp.MustCompile(`(\r?\n)  (\r?\n)`).ReplaceAllString(result, "$1")

	// Remove multiple consecutive blank lines
	result = regexp.MustCompile(`(\r?\n){3,}`).ReplaceAllString(result, "$1$1")
	result = regexp.MustCompile(`(\r?\n){2,}$`).ReplaceAllString(result, "$1")

	return result
}

// SanitizeName escapes underscore character which have special meaning in Markdown.
func SanitizeName(name string, settings *print.Settings) string {
	if settings.EscapeCharacters {
		// Escape underscore
		name = strings.Replace(name, "_", "\\_", -1)
	}
	return name
}

// SanitizeItemForDocument converts passed 'string' to suitable Markdown representation
// for a document. (including line-break, illegal characters, code blocks etc)
func SanitizeItemForDocument(s string, settings *print.Settings) string {
	if s == "" {
		return "n/a"
	}
	result := processSegments(
		s,
		"```",
		func(segment string) string {
			segment = ConvertMultiLineText(segment, false)
			segment = EscapeIllegalCharacters(segment, settings)
			segment = NormalizeURLs(segment, settings)
			return segment
		},
		func(segment string) string {
			lastbreak := ""
			if !strings.HasSuffix(segment, "\n") {
				lastbreak = "\n"
			}
			segment = fmt.Sprintf("```%s%s```", segment, lastbreak)
			return segment
		},
	)
	return strings.Replace(result, "<br>", "\n", -1)
}

// SanitizeItemForTable converts passed 'string' to suitable Markdown representation
// for a table. (including line-break, illegal characters, code blocks etc)
func SanitizeItemForTable(s string, settings *print.Settings) string {
	if s == "" {
		return "n/a"
	}
	result := processSegments(
		s,
		"```",
		func(segment string) string {
			segment = ConvertMultiLineText(segment, true)
			segment = EscapeIllegalCharacters(segment, settings)
			segment = NormalizeURLs(segment, settings)
			return segment
		},
		func(segment string) string {
			segment = strings.TrimSpace(segment)
			segment = strings.Replace(segment, "\n", "<br>", -1)
			segment = strings.Replace(segment, "\r", "", -1)
			segment = fmt.Sprintf("<pre>%s</pre>", segment)
			return segment
		},
	)
	return result
}

// ConvertMultiLineText converts a multi-line text into a suitable Markdown representation.
func ConvertMultiLineText(s string, isTable bool) string {
	if isTable {
		s = strings.TrimSpace(s)
	}

	// Convert double newlines to <br><br>.
	s = strings.Replace(s, "\n\n", "<br><br>", -1)

	// Convert line-break on a non-empty line followed by another line
	// starting with "alphanumeric" word into space-space-newline
	// which is a know convention of Markdown for multi-lines paragprah.
	// This doesn't apply on a markdown list for example, because all the
	// consecutive lines start with hyphen which is a special character.
	s = regexp.MustCompile(`(\S*)(\r?\n)(\w+)`).ReplaceAllString(s, "$1  $2$3")
	s = strings.Replace(s, "    \n", "  \n", -1)

	if isTable {
		// Convert space-space-newline to <br>
		s = strings.Replace(s, "  \n", "<br>", -1)

		// Convert single newline to space.
		s = strings.Replace(s, "\n", " ", -1)
	}

	return s
}

// EscapeIllegalCharacters escapes characters which have special meaning in Markdown into their corresponding literal.
func EscapeIllegalCharacters(s string, settings *print.Settings) string {
	// Escape pipe
	if settings.EscapePipe {
		s = strings.Replace(s, "|", "\\|", -1)
	}

	if settings.EscapeCharacters {
		s = processSegments(
			s,
			"`",
			func(segment string) string {
				escape := func(char string) {
					segment = strings.Replace(segment, char+char, "‡‡", -1)
					segment = strings.Replace(segment, " "+char, " ‡", -1)
					segment = strings.Replace(segment, char+" ", "‡ ", -1)
					segment = strings.Replace(segment, char, "\\"+char, -1)
					segment = strings.Replace(segment, "‡", char, -1)
				}
				// Escape underscore
				escape("_")
				// Escape asterisk
				escape("*")
				return segment
			},
			func(segment string) string {
				segment = fmt.Sprintf("`%s`", segment)
				return segment
			},
		)
	}

	return s
}

// NormalizeURLs runs after escape function and normalizes URL back
// to the original state. For example any underscore in the URL which
// got escaped by 'EscapeIllegalCharacters' will be reverted back.
func NormalizeURLs(s string, settings *print.Settings) string {
	if settings.EscapeCharacters {
		if urls := xurls.Strict().FindAllString(s, -1); len(urls) > 0 {
			for _, url := range urls {
				normalized := strings.Replace(url, "\\", "", -1)
				s = strings.Replace(s, url, normalized, -1)
			}
		}
	}
	return s
}

// GenerateIndentation generates indentation of Markdown headers
// with base level of provided 'settings.MarkdownIndent' plus any
// extra level needed for subsection (e.g. 'Required Inputs' which
// is a subsection of 'Inputs' section)
func GenerateIndentation(extra int, settings *print.Settings) string {
	var base = settings.MarkdownIndent
	if base < 1 || base > 5 {
		base = 2
	}
	var indent string
	for i := 0; i < base+extra; i++ {
		indent += "#"
	}
	return indent
}

// PrintFencedCodeBlock prints codes in fences, it automatically detects if
// the input 'code' contains '\n' it will use multi line fence, otherwise it
// wraps the 'code' inside single-tick block.
// If the fenced is multi-line it also appens an extra '\n` at the end and
// returns true accordingly, otherwise returns false for non-carriage return.
func PrintFencedCodeBlock(code string, language string) (string, bool) {
	if strings.Contains(code, "\n") {
		return fmt.Sprintf("\n\n```%s\n%s\n```\n", language, code), true
	}
	return fmt.Sprintf("`%s`", code), false
}

func processSegments(s string, prefix string, normalFn func(segment string) string, codeFn func(segment string) string) string {
	// Isolate blocks of code. Dont escape anything inside them
	nextIsInCodeBlock := strings.HasPrefix(s, prefix)
	segments := strings.Split(s, prefix)
	buffer := bytes.NewBufferString("")
	for _, segment := range segments {
		if len(segment) == 0 {
			continue
		}
		if !nextIsInCodeBlock {
			segment = normalFn(segment)
		} else {
			segment = codeFn(segment)
		}
		buffer.WriteString(segment)
		nextIsInCodeBlock = !nextIsInCodeBlock
	}
	return buffer.String()
}
