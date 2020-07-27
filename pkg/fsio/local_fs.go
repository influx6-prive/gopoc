package fsio

import (
	"fmt"
	"github.com/spf13/afero"
	"github.com/influx6/npkg/nerror"
)

type LocalFS struct {
	fs afero.Fs
}

// OpenDir returns giving directory which has EAM datafeed files stored.
func (l *LocalFS) OpenDir(eamNamespace string, datafeed string) (afero.File, error) {
	l.initFS()
	var eamDirectory = fmt.Sprintf("%s/%s", datafeed, eamNamespace)
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
