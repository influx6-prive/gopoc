package parsers

import (
	"io"
	"strings"
	"sync"

	"github.com/beevik/etree"
	"github.com/influx6/npkg/nerror"
)

const (
	XMlDataFeed = "datafeed/xml"
)

type XMLHeader struct{}

func (xe XMLHeader) Type() string {
	return XMlDataFeed
}

type XMLParser struct {
	doParse     sync.Once
	source      io.Reader
	xmlDocument *etree.Document
}

func NewXMLParser(xmlSource io.Reader) *XMLParser {
	var parser XMLParser
	parser.source = xmlSource
	parser.xmlDocument = etree.NewDocument()
	return &parser
}

func NewXMLParserFromString(xmlContent string) *XMLParser {
	return NewXMLParser(strings.NewReader(xmlContent))
}

func (xm *XMLParser) parseSource() error {
	var err error
	xm.doParse.Do(func() {
		_, err = xm.xmlDocument.ReadFrom(xm.source)
	})
	if err != nil {
		return nerror.Wrap(err, "Failed to parse data source")
	}
	return nil
}

func (xm *XMLParser) ByTags(selector string) goproc.FeedIterator {
	var parseErr = xm.parseSource()
	if parseErr != nil {
		parseErr = nerror.WrapOnly(parseErr)
	}

	var element XMLElementIterator
	element.err = parseErr
	element.parent = xm
	element.index = -1
	element.nodes = xm.xmlDocument.SelectElements(selector)
	element.totalNodes = len(element.nodes)
	return &element
}

type XMLElementIterator struct {
	index      int
	totalNodes int
	err        error
	header     *XMLHeader
	parent     *XMLParser
	nodes      []*etree.Element
}

func (xe *XMLElementIterator) Err() error {
	return xe.err
}

func (xe *XMLElementIterator) HasNext() bool {
	return xe.index < xe.totalNodes
}

func (xe *XMLElementIterator) Next() interface{} {
	if !xe.HasNext() {
		return nil
	}
	xe.index++
	return xe.nodes[xe.index]
}
