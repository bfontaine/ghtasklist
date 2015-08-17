// Package ghtasklists provides utilities to work with GitHub's task lists
package ghtasklists

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

var linePrefix = []byte("- [ ] ")
var eol = []byte("\n")

// Transform reads its text from the given reader and return the transformed
// text. It does not close the reader.
//
// The resulting text will be the same as the input except that each line will
// be prefixed with "- [ ] " in order to be parsed by GitHub's Markdown parser
// as a task list item.
func Transform(r io.Reader) (string, error) {
	var buf bytes.Buffer

	sc := bufio.NewScanner(r)

	for sc.Scan() {
		buf.Write(linePrefix)
		buf.Write(sc.Bytes())
		buf.Write(eol)
	}

	return buf.String(), sc.Err()
}

// TransformString is like Transform except that it takes a string
func TransformString(s string) string {
	sn, _ := Transform(strings.NewReader(s))
	return sn
}
