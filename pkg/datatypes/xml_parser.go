package datatypes

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/spf13/afero"

	"github.com/JSchillinger/gopoc"

	"github.com/beevik/etree"
	"github.com/influx6/npkg/nerror"
)

const (
	XMlDataFeed = "datafeed/xml"
)

var _ gopoc.FeedParser = (*XMLParser)(nil)

type XMLHeader struct {
	FileInfo os.FileInfo
}

func (xe XMLHeader) String() string {
	return fmt.Sprintf(`Type: %s,File: %s`, XMlDataFeed, xe.FileInfo.Name())
}

func (xe XMLHeader) Type() string {
	return XMlDataFeed
}

type XMLParser struct {
	err         error
	header      XMLHeader
	doParse     sync.Once
	source      io.Reader
	xmlDocument *etree.Document
}

func NewXMLParser(xmlSource io.Reader, header XMLHeader) *XMLParser {
	var parser XMLParser
	parser.header = header
	parser.source = xmlSource
	parser.xmlDocument = etree.NewDocument()
	return &parser
}

func NewXMLParserFromFile(xmlFile afero.File, xmlFileInfo os.FileInfo) *XMLParser {
	return NewXMLParser(xmlFile, XMLHeader{
		FileInfo: xmlFileInfo,
	})
}

func NewXMLParserFromString(xmlContent string) *XMLParser {
	return NewXMLParser(strings.NewReader(xmlContent), XMLHeader{})
}

func (xm *XMLParser) parseSource() error {
	if xm.err != nil {
		return xm.err
	}

	var err error
	xm.doParse.Do(func() {
		_, err = xm.xmlDocument.ReadFrom(xm.source)
	})
	xm.err = err
	if err != nil {
		return nerror.Wrap(err, "CollectErr to parse data source")
	}
	return nil
}

func (xm *XMLParser) Err() error {
	_ = xm.parseSource()
	return xm.err
}

func (xm *XMLParser) Header() gopoc.FeedHeader {
	return xm.header
}

func (xm *XMLParser) ByPosition(col int) (interface{}, error) {
	return nil, nerror.New("not supported")
}

func (xm *XMLParser) ByRow() gopoc.FeedIterator {
	_ = xm.parseSource()
	var element XMLElementIterator
	element.parent = xm
	element.index = -1

	if xm.err != nil {
		element.err = nerror.WrapOnly(xm.err)
		return &element
	}

	element.nodes = xm.xmlDocument.ChildElements()
	element.totalNodes = len(element.nodes)
	return &element
}

func (xm *XMLParser) HasTag(xselector string) (bool, error) {
	_ = xm.parseSource()
	if xm.err != nil {
		return false, xm.err
	}
	var xPath, xPathErr = etree.CompilePath(xselector)
	if xPathErr != nil {
		return false, xPathErr
	}

	var node = xm.xmlDocument.FindElementPath(xPath)
	return node != nil, nil
}

func (xm *XMLParser) ByTag(selector string) gopoc.FeedIterator {
	_ = xm.parseSource()
	var element XMLElementIterator
	element.parent = xm
	element.index = -1

	if xm.err != nil {
		element.err = nerror.WrapOnly(xm.err)
		return &element
	}

	var xPath, xPathErr = etree.CompilePath(selector)
	if xPathErr != nil {
		element.err = nerror.WrapOnly(xm.err)
		return &element
	}

	element.nodes = xm.xmlDocument.FindElementsPath(xPath)
	element.totalNodes = len(element.nodes)
	return &element
}

type XMLElementIterator struct {
	index      int
	totalNodes int
	err        error
	parent     *XMLParser
	nodes      []*etree.Element
}

func (xe *XMLElementIterator) Err() error {
	return xe.err
}

func (xe *XMLElementIterator) HasNext() bool {
	if xe.err != nil {
		return false
	}
	if len(xe.nodes) == 0 {
		return false
	}
	return xe.index < xe.totalNodes-1
}

func (xe *XMLElementIterator) Next() interface{} {
	return xe.NextNode()
}

func (xe *XMLElementIterator) NextNode() *etree.Element {
	if !xe.HasNext() {
		return nil
	}
	xe.index++
	return xe.nodes[xe.index]
}
