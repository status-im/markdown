package markdown

import (
	"bytes"
	"encoding/json"
	"regexp"
	"strings"
	"testing"

	"github.com/gomarkdown/markdown/parser"
)

type TestParams struct {
	extensions        parser.Extensions
	referenceOverride parser.ReferenceOverrideFunc
}

func runMarkdown(input string, params TestParams) (string, error) {
	parser := parser.NewWithExtensions(params.extensions)
	parser.ReferenceOverride = params.referenceOverride
	result := parser.Parse([]byte(input))
	if result == nil {
		return "", nil
	}
	output, err := json.Marshal(result)
	return string(output), err
}

// doTests runs full document tests using MarkdownCommon configuration.
func doTests(t *testing.T, tests []string) {
	doTestsParam(t, tests, TestParams{
		extensions: parser.CommonExtensions,
	})
}

func doTestsBlock(t *testing.T, tests []string, extensions parser.Extensions) {
	doTestsParam(t, tests, TestParams{
		extensions: extensions,
	})
}

func doTestsParam(t *testing.T, tests []string, params TestParams) {
	for i := 0; i+1 < len(tests); i += 2 {
		input := tests[i]
		expected := strings.TrimRight(tests[i+1], "\n")
		got, err := runMarkdown(input, params)
		if err != nil {
			t.Errorf("Failed to marshal json: %+v\n", err)
		}
		if got != expected {
			t.Errorf("\nInput   [%#v]\nExpected[%#v]\nGot     [%#v]\nInput:\n%s\nExpected:\n%s\nGot:\n%s\n",
				input, expected, got, input, expected, got)
		}
	}
}

func doTestsInline(t *testing.T, tests []string) {
	doTestsInlineParam(t, tests, TestParams{})
}

func doLinkTestsInline(t *testing.T, tests []string) {
	doTestsInline(t, tests)

	prefix := "http://localhost"
	transformTests := transformLinks(tests, prefix)
	doTestsInlineParam(t, transformTests, TestParams{})
	doTestsInlineParam(t, transformTests, TestParams{})
}

func doSafeTestsInline(t *testing.T, tests []string) {
	doTestsInlineParam(t, tests, TestParams{})

	// All the links in this test should not have the prefix appended, so
	// just rerun it with different parameters and the same expectations.
	prefix := "http://localhost"
	transformTests := transformLinks(tests, prefix)
	doTestsInlineParam(t, transformTests, TestParams{})
}

func doTestsInlineParam(t *testing.T, tests []string, params TestParams) {
	params.extensions |= parser.Autolink | parser.Strikethrough
	doTestsParam(t, tests, params)
}

func transformLinks(tests []string, prefix string) []string {
	newTests := make([]string, len(tests))
	anchorRe := regexp.MustCompile(`<a href="/(.*?)"`)
	imgRe := regexp.MustCompile(`<img src="/(.*?)"`)
	for i, test := range tests {
		if i%2 == 1 {
			test = anchorRe.ReplaceAllString(test, `<a href="`+prefix+`/$1"`)
			test = imgRe.ReplaceAllString(test, `<img src="`+prefix+`/$1"`)
		}
		newTests[i] = test
	}
	return newTests
}

func normalizeNewlines(d []byte) []byte {
	// replace CR LF (windows) with LF (unix)
	d = bytes.Replace(d, []byte{13, 10}, []byte{10}, -1)
	// replace CF (mac) with LF (unix)
	d = bytes.Replace(d, []byte{13}, []byte{10}, -1)
	return d
}
