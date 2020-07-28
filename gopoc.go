package gopoc

import (
	"github.com/spf13/afero"
)

type DataFileSystem interface {
	OpenDir(eamNamespace string, dataFeed string) (afero.File, error)
}

type DataFeedHeader interface {
	Type() string
}

type DataFeedFile interface {
	File(afero.File) FeedParser
}

type FeedIterator interface {
	HasNext() bool
	Next() interface{}
	Err() error
}

type FeedParser interface {
	Err() error
	Header() DataFeedHeader

	// Functions for going through
	// data feed lines, text and tags.
	ByRow() FeedIterator
	ByTags(tag string) FeedIterator
	ByPosition(col int) interface{}
}
