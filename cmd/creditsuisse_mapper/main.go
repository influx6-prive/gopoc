package main

import (
	"fmt"
	"log"
	"path"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/afero"

	"github.com/JSchillinger/gopoc"
	"github.com/JSchillinger/gopoc/pkg/datatypes"
	"github.com/JSchillinger/gopoc/pkg/feeds/creditsuisse"
	"github.com/JSchillinger/gopoc/pkg/feeds/creditsuisse/parsers"
)

var creditSuisseData = "../../data/CreditSuisse"
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

type Fault struct {
	Header gopoc.FeedHeader
	Err    error
}

type Collector struct {
	Accounts []parsers.SafeKeepingAccountInformation
	Options  []parsers.OptionContract
	Errs     []Fault
}

func (c *Collector) CollectErr(header gopoc.FeedHeader, err error) {
	c.Errs = append(c.Errs, Fault{Header: header, Err: err})
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

func main() {
	var feedControl creditsuisse.DataFeed
	feedControl.FileSystem = &LocalFS{}
	feedControl.FeedParser = &datatypes.BaseFileParser{}
	feedControl.FeedProcessors = []gopoc.ParseProcessor{
		&parsers.StaticDataSafeKeepingAccountInformationXMLParser{},
		&parsers.PsnOptionContractXMLParser{},
	}

	var collector Collector
	if err := feedControl.Process("creditSuisse", &collector); err != nil {
		log.Fatalf("Failed to parse: %#q\n", err)
		return
	}

	if len(collector.Errs) != 0 {
		for _, fault := range collector.Errs {
			log.Printf("Failed parsing of %#q with error: %#s\n", fault.Header, fault.Err)
		}
	}

	log.Printf("---------Credit Suise SafeKeeping Accounts----------\n")
	spew.Dump(collector.Accounts)

	log.Printf("---------Credit Suise Option Contracts--------------\n")
	spew.Dump(collector.Options)
}
