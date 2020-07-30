package parsers

import (
	"time"

	"github.com/JSchillinger/gopoc"
	"github.com/JSchillinger/gopoc/pkg/datatypes"

	"github.com/shopspring/decimal"
)

type ImportCashTransaction struct {
	ValueDate       time.Time
	AccountNumber   string
	Amount          decimal.Decimal
	AmountInBaseCCY decimal.Decimal
}

/**
*
*	ImportCashTransactionParser is responsible to parse the following
* 	data from a base gopoc.FeedParser.
*
*	<ImportTransactionCashflow>
*	<accountNumber>3883389</accountNumber>
*	<amount>590.71</amount>
*	<amountInBaseCurrency>448.6442431691287</amountInBaseCurrency>
*	<valueDate>2016-04-11 00:00:00.0 UTC</valueDate>
*	</ImportTransactionCashflow>
*
 */
type ImportCashTransactionParser struct {
}

func (ic *ImportCashTransactionParser) CanHandle(parser gopoc.FeedParser) bool {
	var header = parser.Header()
	if header.Type() == datatypes.XMlDataFeed {
		return true
	}
	return false
}

func (ic *ImportCashTransactionParser) Parse(parser gopoc.FeedParser, results gopoc.ParserResultHandler) error {
	var header = parser.Header()
	switch header.Type() {
	case datatypes.XMlDataFeed:
		return ic.parseFeedFromXML(parser)
	default:
	}
	return nil
}

func (ic *ImportCashTransactionParser) parseFeedFromXML(parser gopoc.FeedParser) error {
	return nil
}
