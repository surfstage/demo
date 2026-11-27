package slug

import "testing"

func TestMake(t *testing.T) {
	cases := map[string]string{
		"Hello, World!":         "hello-world",
		"Release notes for 2.0": "release-notes-for-20",
		"snake_case_title":      "snake-case-title",
		"  Trim me  ":           "trim-me",
		"Hello  World":          "hello-world",
		"Release - notes":       "release-notes",
		"a_-_b":                 "a-b",
		"-- Dashes at ends --":  "dashes-at-ends",
		"Q & A":                 "q-a",
	}
	for in, want := range cases {
		if got := Make(in); got != want {
			t.Errorf("Make(%q) = %q, want %q", in, got, want)
		}
	}
}
