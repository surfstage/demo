// Package slug turns titles into URL slugs: "Hello, World!" becomes "hello-world".
package slug

import (
	"strings"
	"unicode"
)

// Make returns the slug for a title: lower case letters and digits, with a dash
// between words. Punctuation is dropped. A run of separators (spaces, dashes,
// underscores) becomes a single dash, and separators at either end are dropped.
func Make(title string) string {
	var b strings.Builder
	dash := false // a separator was seen since the last letter or digit
	for _, r := range strings.ToLower(title) {
		switch {
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			if dash && b.Len() > 0 {
				b.WriteRune('-')
			}
			dash = false
			b.WriteRune(r)
		case r == ' ' || r == '-' || r == '_':
			dash = true
		}
	}
	return b.String()
}
