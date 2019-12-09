package markdown

import "testing"

func TestDocument(t *testing.T) {
	var tests = []string{
		// Empty document.
		"",
		"[]",

		" ",
		"[]",

		// This shouldn't panic.
		// https://github.com/russross/blackfriday/issues/172
		"[]:<",
		"[{\"type\":\"paragraph\",\"children\":[{\"literal\":\"[]:\\u003c\"}]}]",

		// This shouldn't panic.
		// https://github.com/russross/blackfriday/issues/173
		"   [",
		"[{\"type\":\"paragraph\",\"children\":[{\"literal\":\"[\"}]}]",
	}
	doTests(t, tests)
}
