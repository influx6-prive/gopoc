package datatypes

import (
	"os"
	"path"
	"strings"

	"github.com/JSchillinger/gopoc"
	"github.com/influx6/npkg/nerror"
	"github.com/spf13/afero"
)

var _ gopoc.FileParser = (*FileParser)(nil)

type FileParser struct{}

func (f *FileParser) GetParser(file afero.File, info os.FileInfo) (gopoc.FeedParser, error) {
	var fileName = file.Name()
	var fileExtension = strings.ToLower(path.Ext(fileName))

	//TODO: Add more parsers here.
	switch fileExtension {
	case ".xml":
		return NewXMLParserFromFile(file, info), nil
	case ".txt", ".text":
		return NewTextParserFromFile(file, info), nil
	default:
	}
	return nil, nerror.New("not yet supported")
}
