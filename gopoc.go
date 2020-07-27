package gopoc

import (
	"github.com/spf13/afero"
)

type DataFileSystem interface {
	OpenDir(eamNamespace string, dataFeed string) (afero.File, error)
}

type DataFeedRow interface {
	Type() string
	Data() interface{}
}

type DataFeedHeader interface {
	Type() string
	Header() interface{}
}

type DataFeedFile interface {
	File(afero.File) FeedParser
}

type FeedParser interface {
	Err() error
	Header() DataFeedHeader

	// Functions for going through
	// data feed lines, text and tags.
	ByRow() DataFeedRow
	ByTag(tag string) DataFeedRow
	ByPosition(col int) DataFeedRow

	Parse(file afero.File)
}
