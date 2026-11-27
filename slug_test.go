package slug

import "testing"

func TestMake(t *testing.T) {
	cases := map[string]string{
		"Hello, World!":         "hello-world",
		"Release notes for 2.0": "release-notes-for-20",
		"snake_case_title":      "snake-case-title",
		"  Trim me  ":           "trim-me",
	}
	for in, want := range cases {
		if got := Make(in); got != want {
			t.Errorf("Make(%q) = %q, want %q", in, got, want)
		}
	}
}
