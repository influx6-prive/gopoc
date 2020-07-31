package creditsuisse

import (
	"github.com/influx6/npkg/nerror"

	"github.com/JSchillinger/gopoc"
)

const (
	DataFeedName = "Credit Suisse"
)

type DataFeed struct {
	FeedParser     gopoc.FileParser
	FileSystem     gopoc.DataFileSystem
	FeedProcessors []gopoc.ParseProcessor
}

func (cs *DataFeed) Process(eamNamespace string, collector gopoc.Collector) error {
	var directory, err = cs.FileSystem.OpenDir(eamNamespace, DataFeedName)
	if err != nil {
		return nerror.WrapOnly(err)
	}

	var directoryFiles, readDirErr = directory.Readdir(-1)
	if readDirErr != nil {
		return nerror.WrapOnly(readDirErr)
	}

	for _, targetFileInfo := range directoryFiles {
		if targetFileInfo.IsDir() {
			continue
		}

		var targetFile, fileOpenErr = cs.FileSystem.OpenFile(eamNamespace, DataFeedName, targetFileInfo.Name())
		if fileOpenErr != nil {
			return nerror.WrapOnly(fileOpenErr)
		}

		var targetFileParser, parserErr = cs.FeedParser.GetParser(targetFile, targetFileInfo)
		if parserErr != nil {
			return nerror.WrapOnly(parserErr)
		}

		for _, proc := range cs.FeedProcessors {
			var canHandle, handleErr = proc.CanHandle(targetFileParser)
			if handleErr != nil {
				return nerror.WrapOnly(handleErr)
			}

			if !canHandle {
				continue
			}

			if procErr := proc.Handle(targetFileParser, collector); procErr != nil {
				return nerror.WrapOnly(procErr)
			}
		}
	}

	return nil
}
