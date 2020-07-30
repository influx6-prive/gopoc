package gopoc

import (
	"os"

	"github.com/spf13/afero"
)

type ParserResultHandler interface {
	Handle(FeedHeader, interface{}) error
}

type DataFileSystem interface {
	OpenDir(eamNamespace string, dataFeed string) (afero.File, error)
	OpenFile(eamNamespace string, dataFeed string, fileName string) (afero.File, error)
}

type FileParser interface {
	GetParser(afero.File, os.FileInfo) (FeedParser, error)
}

type ParseProcessor interface {
	CanHandle(parser FeedParser) bool
	Handle(parser FeedParser, handle ParserResultHandler) error
}

type FeedIterator interface {
	HasNext() bool
	Next() interface{}
	Err() error
}

type FeedHeader interface {
	Type() string
}

type FeedParser interface {
	Err() error

	// Header returns a detail which represent the underline
	// data.
	Header() FeedHeader

	// Functions for going through
	// data feed lines, text and tags.
	ByRow() FeedIterator
	HasTag(tag string) bool

	// Tag returns an iterator that cycles through all
	// matching tag
	ByTag(tag string) FeedIterator

	// ByPosition returns value of feed at giving position.
	ByPosition(col int) (interface{}, error)
}
