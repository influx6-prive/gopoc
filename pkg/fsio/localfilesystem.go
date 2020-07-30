package fsio

import (
	"path"

	"github.com/influx6/npkg/nerror"
	"github.com/spf13/afero"

	"github.com/JSchillinger/gopoc"
)

var _ gopoc.DataFileSystem = (*LocalFS)(nil)

type LocalFS struct {
	fs afero.Fs
}

func (l *LocalFS) OpenFile(eamNamespace string, datafeed string, targetPath string) (afero.File, error) {
	l.initFS()
	var eamTargetFilePath = path.Join(datafeed, eamNamespace, targetPath)
	var targetFile, fileErr = l.fs.Open(eamTargetFilePath)
	if fileErr != nil {
		return nil, nerror.WrapOnly(fileErr)
	}
	return targetFile, nil
}

// OpenDir returns giving directory which has EAM datafeed files stored.
func (l *LocalFS) OpenDir(eamNamespace string, datafeed string) (afero.File, error) {
	l.initFS()
	var eamDirectory = path.Join(datafeed, eamNamespace)
	var directoryFile, dirErr = l.fs.Open(eamDirectory)
	if dirErr != nil {
		return nil, nerror.WrapOnly(dirErr)
	}
	var stat, statErr = directoryFile.Stat()
	if statErr != nil {
		return nil, nerror.WrapOnly(statErr)
	}
	if !stat.IsDir() {
		return nil, nerror.New("EAM %s for datafeed %s as no directory", eamNamespace, datafeed)
	}
	return directoryFile, nil
}

func (l *LocalFS) initFS() {
	if l.fs == nil {
		l.fs = afero.NewOsFs()
	}
}
