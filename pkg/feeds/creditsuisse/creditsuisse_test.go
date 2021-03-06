package creditsuisse_test

import (
	"fmt"
	"path"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"

	"github.com/JSchillinger/gopoc"
	"github.com/JSchillinger/gopoc/pkg/datatypes"
	"github.com/JSchillinger/gopoc/pkg/feeds/creditsuisse"
	"github.com/JSchillinger/gopoc/pkg/feeds/creditsuisse/parsers"
)

var creditSuisseData = "../../../data/CreditSuisse"
var creditSuisseFS = afero.NewOsFs()

type LocalFS struct{}

func (l LocalFS) OpenDir(eamNamespace string, dataFeed string) (afero.File, error) {
	fmt.Printf("Requesting dir for %#q -> %#q\n", eamNamespace, dataFeed)
	return creditSuisseFS.Open(creditSuisseData)
}

func (l LocalFS) OpenFile(eamNamespace string, dataFeed string, fileName string) (afero.File, error) {
	fmt.Printf("Requesting file %#q\n", fileName)
	return creditSuisseFS.Open(path.Join(creditSuisseData, fileName))
}

type Collector struct {
	Accounts []parsers.SafeKeepingAccountInformation
	Options  []parsers.OptionContract
	Err      []error
}

func (c *Collector) CollectErr(header gopoc.FeedHeader, err error) {
	c.Err = append(c.Err, err)
}

func (c *Collector) Collect(header gopoc.FeedHeader, val interface{}) error {
	switch vl := val.(type) {
	case parsers.OptionContract:
		c.Options = append(c.Options, vl)
	case parsers.SafeKeepingAccountInformation:
		c.Accounts = append(c.Accounts, vl)
	default:
		fmt.Printf("Unable to handle type %#v\n", vl)
	}
	return nil
}

func TestCreditSuisseSDSAParsing(t *testing.T) {
	var feedControl creditsuisse.DataFeed
	feedControl.FileSystem = &LocalFS{}
	feedControl.FeedParser = &datatypes.BaseFileParser{}
	feedControl.FeedProcessors = []gopoc.ParseProcessor{
		&parsers.StaticDataSafeKeepingAccountInformationXMLParser{},
		&parsers.PsnOptionContractXMLParser{},
	}

	var collector Collector
	require.NoError(t, feedControl.Process("credo", &collector))
	require.Len(t, collector.Err, 0)
	require.Len(t, collector.Accounts, 1)
	require.Len(t, collector.Options, 1)

	require.Len(t, collector.Accounts[0].SafeKeepingInfo, 9)
	require.Equal(t, collector.Accounts[0].Account.Info.ClientFormat, "XML")

	require.Len(t, collector.Options[0].Contracts, 1)
	require.Equal(t, collector.Options[0].Contracts[0].IntRptUnit, "0973")
}
