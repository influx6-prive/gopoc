package datatypes

import (
	"os"
	"path"
	"strings"

	"github.com/JSchillinger/gopoc"
	"github.com/influx6/npkg/nerror"
	"github.com/spf13/afero"
)

var _ gopoc.FileParser = (*BaseFileParser)(nil)

type BaseFileParser struct{}

func (f *BaseFileParser) GetParser(file afero.File, info os.FileInfo) (gopoc.FeedParser, error) {
	var fileName = file.Name()
	var fileExtension = strings.ToLower(path.Ext(fileName))

	//TODO: Add more parsers here.
	var parser gopoc.FeedParser
	switch fileExtension {
	case ".xml":
		parser = NewXMLParserFromFile(file, info)
	case ".txt", ".text":
		parser = NewTextParserFromFile(file, info)
	default:
		return nil, nerror.New("not yet supported")
	}

	if parserErr := parser.Err(); parserErr != nil {
		return parser, nerror.WrapOnly(parserErr)
	}
	return parser, nil
}
