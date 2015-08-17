package ghtasklists_test

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/bfontaine/ghtasklist/ghtasklists"
	"github.com/stretchr/testify/assert"
)

type eofReader struct{}

func (e eofReader) Read(p []byte) (int, error) { return 0, io.EOF }

type errReader struct{ e error }

func (e errReader) Read(p []byte) (int, error) { return 0, e.e }

func TestTransformEmpty(t *testing.T) {
	s, e := ghtasklists.Transform(eofReader{})

	assert.Nil(t, e)
	assert.Equal(t, "", s)
}

func TestTransformError(t *testing.T) {
	e := errors.New("foo")
	s, err := ghtasklists.Transform(errReader{e: e})

	assert.Equal(t, e, err)
	assert.Equal(t, "", s)
}

func TestTransform(t *testing.T) {
	input := `
		foo
		bar qux
		choucroute`

	output := `- [ ] 		foo
- [ ] 		bar qux
- [ ] 		choucroute
`

	s, e := ghtasklists.Transform(strings.NewReader(input))

	assert.Nil(t, e)
	assert.Equal(t, output, s)
}
