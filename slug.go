// Package slug turns titles into URL slugs: "Hello, World!" becomes "hello-world".
package slug

import (
	"strings"
	"unicode"
)

// Make returns the slug for a title: lower case letters and digits, with a dash
// between words. Punctuation is dropped.
func Make(title string) string {
	var b strings.Builder
	for _, r := range strings.ToLower(title) {
		switch {
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			b.WriteRune(r)
		case r == ' ' || r == '-' || r == '_':
			b.WriteRune('-')
		}
	}
	return strings.Trim(b.String(), "-")
}
