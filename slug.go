// Package slug turns titles into URL slugs: "Hello, World!" becomes "hello-world".
package slug

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// plain drops accents: "Café" becomes "Cafe".
var plain = transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)

// Make returns the slug for a title: lower case letters and digits, with a dash
// between words. Punctuation is dropped, and accented letters lose their accents.
func Make(title string) string {
	if s, _, err := transform.String(plain, title); err == nil {
		title = s
	}
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
