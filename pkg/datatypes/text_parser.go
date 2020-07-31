package datatypes

import (
	"io"
	"os"

	"github.com/spf13/afero"

	"github.com/JSchillinger/gopoc"
)

const (
	TextDataFeed = "datafeed/text"
)

var _ gopoc.FeedParser = (*TextParser)(nil)

func NewTextParser(reader io.Reader, header TextHeader) *TextParser {
	return &TextParser{
		reader: reader,
		header: header,
	}
}

func NewTextParserFromFile(file afero.File, fileInfo os.FileInfo) *TextParser {
	return NewTextParser(file, TextHeader{FileInfo: fileInfo})
}

type TextHeader struct {
	FileInfo os.FileInfo
}

func (xe TextHeader) Type() string {
	return TextDataFeed
}

type TextParser struct {
	err    error
	reader io.Reader
	header TextHeader
}

func (t TextParser) Err() error {
	panic("implement me")
}

func (t TextParser) Header() gopoc.FeedHeader {
	panic("implement me")
}

func (t TextParser) ByRow() gopoc.FeedIterator {
	panic("implement me")
}

func (t TextParser) HasTag(tag string) (bool, error) {
	panic("implement me")
}

func (t TextParser) ByTag(tag string) gopoc.FeedIterator {
	panic("implement me")
}

func (t TextParser) ByPosition(col int) (interface{}, error) {
	panic("implement me")
}

type TextIterator struct {
	err error
}

func (t TextIterator) HasNext() bool {
	return false
}

func (t *TextIterator) Next() interface{} {
	panic("implement me")
}

func (t *TextIterator) Err() error {
	return t.err
}
