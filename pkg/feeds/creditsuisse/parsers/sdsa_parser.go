package parsers

import (
	"strconv"
	"strings"
	"time"

	"github.com/beevik/etree"
	"github.com/influx6/npkg/nerror"
	"github.com/shopspring/decimal"

	"github.com/JSchillinger/gopoc"
	"github.com/JSchillinger/gopoc/pkg/datatypes"
)

const (
	CrtnDtTMTimeFormat  = time.RFC3339Nano
	ReportingTimeFormat = "2020-02-01"
)

type SafeKeepingAccountInfo struct {
	Location                  string
	AccountId                 string
	ExtAccountId              string
	InvestmentCurrencyISOCd   string
	InvestmentCurrencyISODesc string
	AccountTpCd               int
	AccountTpDesc             string
	AccountSTS                bool
	AccountClsgSTS            bool
	AccountPldgdInd           bool
	AccountPrtlyPldfdInd      bool
	AccountLmtPldgdInd        bool
	AccountOpeningDate        time.Time
}

type ClientKeyData struct {
	IntRptUnit          string
	IntRptUnitDesc      string
	ClientId            string
	ExtClientId         string
	SafeKeepingAccounts map[string]SafeKeepingAccountInfo
}

type SenderReceiverInfo struct {
	SenderBIC   string
	ReceiverBIC string
}

type FLInfo struct {
	DWHMsgId              string
	LocationISO           string
	CrtnDtTmDate          time.Time
	ReportingDate         time.Time
	Type                  string
	TypeCD                string
	Version               string
	ClientFormat          string
	MessageSequenceNumber int
}

type AccountHeader struct {
	Sender SenderReceiverInfo
	Info   FLInfo
}

type PSNStaticDataSafeKeepingAccountInformation struct {
	Account       AccountHeader
	ClientKeyData []ClientKeyData
}

type PSNOptionContractCtlInfo struct {
	CtrctId                    string
	DealDate                   time.Time
	ExpiryDate                 time.Time
	BaseCurrencyCsCd           string
	BaseCurrencyIsoCd          string
	SettlementDate             time.Time
	BuySettlementDate          time.Time
	SellSettlementDate         time.Time
	SettlementTpInd            int
	BuySellInd                 string
	OptionTpCd                 int
	ExticOptionTpCd            string
	NtnlLeadCurrencyAmount     decimal.Decimal
	NtnlLeadCurrencyIsoCd      string
	NtnlLeadCurrencyIsoDesc    string
	StrikePriceRate            decimal.Decimal
	NtnlCounterCurrencyAmount  decimal.Decimal
	NtnlCounterCurrencyISOCd   string
	NtnlCounterCurrencyISODesc string
	PremiumAmount              decimal.Decimal
	PremiumAccountIsoCd        string          `xml:"PremiumCcyIsoCd"`
	PremiumAccountIsoDesc      string          `xml:"PremiumCcyIsoDesc"`
	PositionDate               time.Time       `xml:"posDt"`
	RevaluationDate            time.Time       `xml:"RevaltnDt"`
	PrtFlId                    string          `xml:"PrtflId"`
	MtmValOrigCurrencyIsoCd    string          `xml:"MtmValOrigCcyIsoCd"`
	MtmValOrigCurrencyAmount   decimal.Decimal `xml:"MtmValOrigCcyAmt"`
	MtmValReportCurrencyISOCd  string          `xml:"MtmValRptCcyIsoCd"`
	MtmValReportCurrencyAmount decimal.Decimal `xml:"MtmValRptCcyAmt"`
	MtmValueDate               time.Time       `xml:"MtmValDt"`
}

type OptionContract struct {
	Account   AccountHeader
	Contracts []PSNOptionContractCtlInfo
}

type StaticDataSafeKeepingAccountInformationParser struct{}

func (ic *StaticDataSafeKeepingAccountInformationParser) CanHandle(parser gopoc.FeedParser) bool {
	var header = parser.Header()
	if header.Type() == datatypes.XMlDataFeed && parser.HasTag("PsNSafekeepingAccountInformation") {
		return true
	}
	return false
}

func (ic *StaticDataSafeKeepingAccountInformationParser) Parse(parser gopoc.FeedParser, handler gopoc.ParserResultHandler) error {
	var header = parser.Header()
	if header.Type() != datatypes.XMlDataFeed {
		return nerror.New("must be an xml feed")
	}

	var xmlParser, notXMLParserType = parser.(*datatypes.XMLParser)
	if !notXMLParserType {
		return nerror.New("expecting datatypes.XMLParser type")
	}

	var psnAccountInfo = xmlParser.ByTag("PsNSafekeepingAccountInformation").(*datatypes.XMLElementIterator)
	if psnErr := psnAccountInfo.Err(); psnErr != nil {
		return nerror.Wrap(psnErr, "PsNSafekeepingAccountInformation tag not found")
	}

	for psnAccountInfo.HasNext() {
		var currentPSNAccountInfo = psnAccountInfo.NextNode()
		var result, parseErr = parsePsNSafeKeeping(currentPSNAccountInfo)
		if parseErr != nil {
			return nerror.WrapOnly(parseErr)
		}

		handler.Handle(parser.Header(), result)
	}

	return nil
}

func parsePsNSafeKeeping(node *etree.Element) (PSNStaticDataSafeKeepingAccountInformation, error) {
	var psn PSNStaticDataSafeKeepingAccountInformation

	var psnHeaderErr error
	psn.Account, psnHeaderErr = parsePsNHeader(node)
	if psnHeaderErr != nil {
		return psn, nerror.WrapOnly(psnHeaderErr)
	}

	var clientKeyErr error
	psn.ClientKeyData, clientKeyErr = parseClientDataBlocks(node)
	if clientKeyErr != nil {
		return psn, nerror.WrapOnly(clientKeyErr)
	}

	return psn, nil
}

type PsnOptionContractParserParser struct{}

func (ic *PsnOptionContractParserParser) CanHandle(parser gopoc.FeedParser) bool {
	var header = parser.Header()
	if header.Type() == datatypes.XMlDataFeed && parser.HasTag("PsNOptionContract") {
		return true
	}
	return false
}

func (ic *PsnOptionContractParserParser) Parse(parser gopoc.FeedParser, handler gopoc.ParserResultHandler) error {
	var header = parser.Header()
	if header.Type() != datatypes.XMlDataFeed {
		return nerror.New("must be an xml feed")
	}

	var xmlParser, notXMLParserType = parser.(*datatypes.XMLParser)
	if !notXMLParserType {
		return nerror.New("expecting datatypes.XMLParser type")
	}

	var psnAccountInfo = xmlParser.ByTag("PsNOptionContract").(*datatypes.XMLElementIterator)
	if psnErr := psnAccountInfo.Err(); psnErr != nil {
		return nerror.Wrap(psnErr, "PsNSafekeepingAccountInformation tag not found")
	}

	for psnAccountInfo.HasNext() {
		var currentPSNAccountInfo = psnAccountInfo.NextNode()
		var result, parseErr = parsePsNOptionContract(currentPSNAccountInfo)
		if parseErr != nil {
			return nerror.WrapOnly(parseErr)
		}

		handler.Handle(parser.Header(), result)
	}

	return nil
}

func parsePsNOptionContract(node *etree.Element) (OptionContract, error) {
	var psn OptionContract

	var psnHeaderErr error
	psn.Account, psnHeaderErr = parsePsNHeader(node)
	if psnHeaderErr != nil {
		return psn, nerror.WrapOnly(psnHeaderErr)
	}

	var clientKeyErr error
	psn.Contracts, clientKeyErr = parseClientOptionContractBlocks(node)
	if clientKeyErr != nil {
		return psn, nerror.WrapOnly(clientKeyErr)
	}

	return psn, nil
}

func parsePsNHeader(node *etree.Element) (AccountHeader, error) {
	var acctHeader AccountHeader

	var header = node.SelectElement("Hdr:Header")
	if header == nil {
		return acctHeader, nerror.New("Hdr:Header tag not found")
	}

	var senderReceiverInfo, senderRecvErr = parseReceiverSenderInfo(header)
	if senderRecvErr != nil {
		return acctHeader, nerror.WrapOnly(senderRecvErr)
	}

	acctHeader.Sender = senderReceiverInfo

	var flInfo, flInfoErr = parseFlInfo(header)
	if flInfoErr != nil {
		return acctHeader, nerror.WrapOnly(flInfoErr)
	}

	acctHeader.Info = flInfo

	return acctHeader, nil
}

func parseFlInfo(node *etree.Element) (FLInfo, error) {
	var data FLInfo
	var flInfo = node.SelectElement("FlInf")
	if flInfo == nil {
		return data, nerror.New("FlInf tag not found in Hdr:Header")
	}

	var children = flInfo.ChildElements()
	for _, elem := range children {
		switch strings.ToLower(elem.Tag) {
		case "dwhmsgid":
			data.DWHMsgId = elem.Text()
		case "locationiso2cd":
			data.LocationISO = elem.Text()
		case "crtndttm":
			var crtErr error
			data.CrtnDtTmDate, crtErr = time.Parse(CrtnDtTMTimeFormat, elem.Text())
			if crtErr != nil {
				return data, nerror.WrapOnly(crtErr)
			}
		case "rprtngdt":
			var reportErr error
			data.ReportingDate, reportErr = time.Parse(ReportingTimeFormat, elem.Text())
			if reportErr != nil {
				return data, nerror.WrapOnly(reportErr)
			}
		case "type":
			data.Type = elem.Text()
		case "typecd":
			data.TypeCD = elem.Text()
		case "vrsn":
			data.Version = elem.Text()
		case "clntfmt":
			data.ClientFormat = elem.Text()
		case "msgseqno":
			var msgSeq, err = strconv.Atoi(elem.Text())
			if err != nil {
				return data, nerror.WrapOnly(err)
			}
			data.MessageSequenceNumber = msgSeq
		}
	}

	return data, nil
}

func parseReceiverSenderInfo(node *etree.Element) (SenderReceiverInfo, error) {
	var data SenderReceiverInfo

	var senderToReceiverTag = node.SelectElement("SndrToRcvrInf")
	if senderToReceiverTag == nil {
		return data, nerror.New("SndrToRcvrInf tag not found in Hdr:Header")
	}

	var receiverBICTag = senderToReceiverTag.SelectElement("RcvrBIC")
	if receiverBICTag == nil {
		return data, nerror.New("RcvrBIC tag not found")
	}

	var senderBICTag = senderToReceiverTag.SelectElement("SndrBIC")
	if senderBICTag == nil {
		return data, nerror.New("SndrBIC tag not found")
	}

	data.ReceiverBIC = receiverBICTag.Text()
	data.SenderBIC = senderBICTag.Text()
	return data, nil
}

func parseClientDataBlocks(node *etree.Element) ([]ClientKeyData, error) {
	var clientData []ClientKeyData

	var dataBlock = node.SelectElement("Data")
	if dataBlock == nil {
		return clientData, nerror.New("Data tag not found")
	}

	var clientSfkDataList = node.SelectElements("ClntSfkData")
	// TODO: Should we return an error for an empty ClntSfkData block?
	if len(clientSfkDataList) == 0 {
		return clientData, nil
	}

	// initialize adequate space which should suffice for available list of
	// ClientSfkData.
	clientData = make([]ClientKeyData, 0, len(clientSfkDataList))

	for _, clientSfkItem := range clientSfkDataList {
		var clientDataBlock, clientDataErr = parseDataBlock(clientSfkItem)
		if clientDataErr != nil {
			return clientData, clientDataErr
		}
		clientData = append(clientData, clientDataBlock)
	}
	return clientData, nil
}

func parseDataBlock(node *etree.Element) (ClientKeyData, error) {
	var clientData, clientDataErr = parseDataBlockClientKey(node)
	if clientDataErr == nil {
		return clientData, nerror.WrapOnly(clientDataErr)
	}

	var clientSkfInfoList = node.SelectElements("SfkInfo")
	if len(clientSkfInfoList) == 0 {
		return clientData, nerror.New("no SfkInfo found in ClntSfkData")
	}

	var parsedSfkItems = make(map[string]SafeKeepingAccountInfo, len(clientSkfInfoList))
	for _, skfItem := range clientSkfInfoList {
		var parsedSkfItem, parsedSkfErr = parseDataBlockSkfInfo(skfItem)
		if parsedSkfErr != nil {
			return clientData, nerror.WrapOnly(parsedSkfErr)
		}
		parsedSfkItems[parsedSkfItem.AccountId] = parsedSkfItem
	}

	clientData.SafeKeepingAccounts = parsedSfkItems
	return clientData, nil
}

func parseDataBlockClientKey(node *etree.Element) (ClientKeyData, error) {
	var clientData ClientKeyData

	var clientKeyElement = node.SelectElement("ClntKey")
	if clientKeyElement == nil {
		return clientData, nerror.New("ClntKey not found in ClntSfkData")
	}

	for _, field := range clientKeyElement.ChildElements() {
		switch strings.ToLower(field.Tag) {
		case "intrptunit":
			clientData.IntRptUnit = field.Text()
		case "intrptunitdesc":
			clientData.IntRptUnitDesc = field.Text()
		case "clntid":
			clientData.ClientId = field.Text()
		case "extclntid":
			clientData.ExtClientId = field.Text()
		}
	}

	return clientData, nil
}

func parseDataBlockSkfInfo(node *etree.Element) (SafeKeepingAccountInfo, error) {
	var data SafeKeepingAccountInfo

	for _, sfkItem := range node.ChildElements() {
		switch strings.ToLower(sfkItem.Tag) {
		case "acctId":
			data.AccountId = sfkItem.Text()
		case "extacctId":
			data.ExtAccountId = sfkItem.Text()
		case "invstmtccyisocd":
			data.InvestmentCurrencyISOCd = sfkItem.Text()
		case "invstmtccyisodesc":
			data.InvestmentCurrencyISODesc = sfkItem.Text()
		case "accttpcd":
			var acctTp, err = strconv.Atoi(sfkItem.Text())
			if err != nil {
				return data, nerror.WrapOnly(err)
			}
			data.AccountTpCd = acctTp
		case "accttpdesc":
			data.AccountTpDesc = sfkItem.Text()
		case "acctopngdt":
			var reportErr error
			data.AccountOpeningDate, reportErr = time.Parse(ReportingTimeFormat, sfkItem.Text())
			if reportErr != nil {
				return data, nerror.WrapOnly(reportErr)
			}
		case "acctclsgsts":
			var status = strings.ToLower(strings.TrimSpace(sfkItem.Text()))
			if status == "no" {
				data.AccountClsgSTS = false
			}
			if status == "yes" {
				data.AccountClsgSTS = true
			}
		case "acctsts":
			var status = strings.ToLower(strings.TrimSpace(sfkItem.Text()))
			if status == "no" {
				data.AccountSTS = false
			}
			if status == "yes" {
				data.AccountSTS = true
			}
		case "acctpldgdind":
			var status = strings.ToLower(strings.TrimSpace(sfkItem.Text()))
			if status == "no" {
				data.AccountPldgdInd = false
			}
			if status == "yes" {
				data.AccountPldgdInd = true
			}
		case "acctprtlypldgdind":
			var status = strings.ToLower(strings.TrimSpace(sfkItem.Text()))
			if status == "no" {
				data.AccountPrtlyPldfdInd = false
			}
			if status == "yes" {
				data.AccountPrtlyPldfdInd = true
			}
		case "acctlmtpldgdind":
			var status = strings.ToLower(strings.TrimSpace(sfkItem.Text()))
			if status == "no" {
				data.AccountLmtPldgdInd = false
			}
			if status == "yes" {
				data.AccountLmtPldgdInd = true
			}
		}
	}
	return data, nil
}

func parseClientOptionContractBlocks(node *etree.Element) ([]PSNOptionContractCtlInfo, error) {
	var contracts []PSNOptionContractCtlInfo
	return contracts, nil
}

func parseOptionContractCtlInfo(node *etree.Element) (PSNOptionContractCtlInfo, error) {
	var data PSNOptionContractCtlInfo
	switch strings.ToLower(node.Tag) {
	case "ctrctid":
		data.CtrctId = node.Text()
	case "dealdate":
		var dealDateErr error
		data.DealDate, dealDateErr = time.Parse(ReportingTimeFormat, node.Text())
		if dealDateErr != nil {
			return data, nerror.WrapOnly(dealDateErr)
		}
	case "expirydate":
		var expiryDateErr error
		data.ExpiryDate, expiryDateErr = time.Parse(ReportingTimeFormat, node.Text())
		if expiryDateErr != nil {
			return data, nerror.WrapOnly(expiryDateErr)
		}
	case "basecurrencycscd":
		data.BaseCurrencyCsCd = node.Text()
	case "basecurrencyisocd":
		data.BaseCurrencyIsoCd = node.Text()
	case "settlementdate":
		var dateErr error
		data.SettlementDate, dateErr = time.Parse(ReportingTimeFormat, node.Text())
		if dateErr != nil {
			return data, nerror.WrapOnly(dateErr)
		}
	case "buysettlementdate":
		var dateErr error
		data.BuySettlementDate, dateErr = time.Parse(ReportingTimeFormat, node.Text())
		if dateErr != nil {
			return data, nerror.WrapOnly(dateErr)
		}
	case "sellsettlementdate":
		var dateErr error
		data.SellSettlementDate, dateErr = time.Parse(ReportingTimeFormat, node.Text())
		if dateErr != nil {
			return data, nerror.WrapOnly(dateErr)
		}
	case "settlementtpind":
		var valErr error
		data.SettlementTpInd, valErr = strconv.Atoi(node.Text())
		if valErr != nil {
			return data, nerror.WrapOnly(valErr)
		}
	case "buysellind":
		data.BuySellInd = node.Text()
	case "optiontpcd":
		var valErr error
		data.OptionTpCd, valErr = strconv.Atoi(node.Text())
		if valErr != nil {
			return data, nerror.WrapOnly(valErr)
		}
	case "exticoptiontpcd":
		data.ExticOptionTpCd = node.Text()
	case "ntnlleadcurrencyamount":
		var valErr error
		data.NtnlLeadCurrencyAmount, valErr = decimal.NewFromString(node.Text())
		if valErr != nil {
			return data, nerror.WrapOnly(valErr)
		}
	case "ntnlleadcurrencyisocd":
		data.NtnlLeadCurrencyIsoCd = node.Text()
	case "ntnlleadcurrencyisodesc":
		data.NtnlLeadCurrencyIsoDesc = node.Text()
	case "strikepricerate":
		var valErr error
		data.StrikePriceRate, valErr = decimal.NewFromString(node.Text())
		if valErr != nil {
			return data, nerror.WrapOnly(valErr)
		}
	case "ntnlcountercurrencyamount":
		var valErr error
		data.NtnlCounterCurrencyAmount, valErr = decimal.NewFromString(node.Text())
		if valErr != nil {
			return data, nerror.WrapOnly(valErr)
		}
	case "ntnlcountercurrencyisocd":
		data.NtnlCounterCurrencyISOCd = node.Text()
	case "ntnlcountercurrencyisodesc":
		data.NtnlCounterCurrencyISODesc = node.Text()
	case "premiumamount":
		var valErr error
		data.PremiumAmount, valErr = decimal.NewFromString(node.Text())
		if valErr != nil {
			return data, nerror.WrapOnly(valErr)
		}
	case "premiumaccountisocd":
		data.PremiumAccountIsoCd = node.Text()
	case "premiumaccountisodesc":
		data.PremiumAccountIsoDesc = node.Text()
	case "positiondate":
		var dateErr error
		data.PositionDate, dateErr = time.Parse(ReportingTimeFormat, node.Text())
		if dateErr != nil {
			return data, nerror.WrapOnly(dateErr)
		}
	case "revaluationdate":
		var dateErr error
		data.RevaluationDate, dateErr = time.Parse(ReportingTimeFormat, node.Text())
		if dateErr != nil {
			return data, nerror.WrapOnly(dateErr)
		}
	case "prtflid":
		data.PrtFlId = node.Text()
	case "mtmvalorigcurrencyisocd":
		data.MtmValOrigCurrencyIsoCd = node.Text()
	case "mtmvalorigcurrencyamount":
		var valErr error
		data.MtmValOrigCurrencyAmount, valErr = decimal.NewFromString(node.Text())
		if valErr != nil {
			return data, nerror.WrapOnly(valErr)
		}
	case "mtmvalreportcurrencyisocd":
		data.MtmValReportCurrencyISOCd = node.Text()
	case "mtmvalreportcurrencyamount":
		var valErr error
		data.MtmValReportCurrencyAmount, valErr = decimal.NewFromString(node.Text())
		if valErr != nil {
			return data, nerror.WrapOnly(valErr)
		}
	case "mtmvaluedate":
		var dateErr error
		data.MtmValueDate, dateErr = time.Parse(ReportingTimeFormat, node.Text())
		if dateErr != nil {
			return data, nerror.WrapOnly(dateErr)
		}
	}
	return data, nil
}
