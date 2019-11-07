package markdown

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"
)

type testData struct {
	md   []byte
	html []byte
}

func testDataToStrArray(tests []*testData) []string {
	res := []string{}
	for _, td := range tests {
		res = append(res, string(td.md))
		res = append(res, string(td.html))
	}
	return res
}

func strArrayToTestData(a []string) []*testData {
	if len(a)%2 == 1 {
		panic("must have even number of items in a")
	}
	res := []*testData{}
	for i := 0; i < len(a)/2; i++ {
		j := i * 2
		td := &testData{
			md:   []byte(a[j]),
			html: []byte(a[j+1]),
		}
		res = append(res, td)
	}
	return res
}

func readTestFile2(t *testing.T, fileName string) []string {
	tests := readTestFile(t, fileName)
	return testDataToStrArray(tests)
}

func readTestFile(t *testing.T, fileName string) []*testData {
	path := filepath.Join("testdata", fileName)
	d, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("ioutil.ReadFile('%s') failed with %s", path, err)
	}
	parts := bytes.Split(d, []byte("+++\n"))
	if len(parts)%2 != 0 {
		t.Fatalf("odd test tuples in file %s: %d", path, len(parts))
	}
	res := []*testData{}
	n := len(parts) / 2
	for i := 0; i < n; i++ {
		j := i * 2
		td := &testData{
			md:   parts[j],
			html: parts[j+1],
		}
		res = append(res, td)
	}
	return res
}
