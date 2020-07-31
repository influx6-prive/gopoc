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

var (
	_ gopoc.ParseProcessor = (*StaticDataSafeKeepingAccountInformationXMLParser)(nil)
	_ gopoc.ParseProcessor = (*PsnOptionContractXMLParser)(nil)
)

const (
	CrtnDtTMTimeFormat  = time.RFC3339
	ReportingTimeFormat = "2006-01-02"
)

type SafeKeepingAccountInfo struct {
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

type SafeKeepingClientInfo struct {
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
	CrtnDtTmDate          string
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

type SafeKeepingAccountInformation struct {
	Account         AccountHeader
	SafeKeepingInfo []SafeKeepingClientInfo
}

type OptionContractCtlInfo struct {
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
	Contracts []OptionClientKeyData
}

type OptionClientKeyData struct {
	IntRptUnit     string
	IntRptUnitDesc string
	ClientId       string
	ExtClientId    string
	Contracts      []OptionContractCtlInfo
}

type StaticDataSafeKeepingAccountInformationXMLParser struct{}

func (ic *StaticDataSafeKeepingAccountInformationXMLParser) CanHandle(parser gopoc.FeedParser) (bool, error) {
	var header = parser.Header()
	var hasPSNTag, err = parser.HasTag("//psn:Document/PsNSafekeepingAccountInformation")
	if err != nil {
		return false, nerror.WrapOnly(err)
	}

	if header.Type() == datatypes.XMlDataFeed && hasPSNTag {
		return true, nil
	}
	return false, nil
}

func (ic *StaticDataSafeKeepingAccountInformationXMLParser) Handle(parser gopoc.FeedParser, handler gopoc.Collector) error {
	var header = parser.Header()
	if header.Type() != datatypes.XMlDataFeed {
		return nerror.New("must be an xml feed")
	}

	var xmlParser, notXMLParserType = parser.(*datatypes.XMLParser)
	if !notXMLParserType {
		return nerror.New("expecting datatypes.XMLParser type")
	}

	var psnAccountInfo = xmlParser.ByTag("//psn:Document/PsNSafekeepingAccountInformation").(*datatypes.XMLElementIterator)
	if psnErr := psnAccountInfo.Err(); psnErr != nil {
		return nerror.Wrap(psnErr, "PsNSafekeepingAccountInformation tag not found")
	}

	for psnAccountInfo.HasNext() {
		var currentPSNAccountInfo = psnAccountInfo.NextNode()
		var result, parseErr = parsePsNSafeKeeping(currentPSNAccountInfo)
		if parseErr != nil {
			handler.CollectErr(parser.Header(), nerror.WrapOnly(parseErr))
			continue
		}

		handler.Collect(parser.Header(), result)
	}

	return nil
}

func parsePsNSafeKeeping(node *etree.Element) (SafeKeepingAccountInformation, error) {
	var psn SafeKeepingAccountInformation

	var psnHeaderErr error
	psn.Account, psnHeaderErr = parsePsNHeader(node)
	if psnHeaderErr != nil {
		return psn, nerror.WrapOnly(psnHeaderErr)
	}

	var clientKeyErr error
	psn.SafeKeepingInfo, clientKeyErr = parseClientDataBlocks(node)
	if clientKeyErr != nil {
		return psn, nerror.WrapOnly(clientKeyErr)
	}

	return psn, nil
}

type PsnOptionContractXMLParser struct{}

func (ic *PsnOptionContractXMLParser) CanHandle(parser gopoc.FeedParser) (bool, error) {
	var header = parser.Header()
	var hasTag, err = parser.HasTag("//psn:Document/PsNOptionContract")
	if err != nil {
		return false, nerror.WrapOnly(err)
	}
	if header.Type() == datatypes.XMlDataFeed && hasTag {
		return true, nil
	}
	return false, nil
}

func (ic *PsnOptionContractXMLParser) Handle(parser gopoc.FeedParser, handler gopoc.Collector) error {
	var header = parser.Header()
	if header.Type() != datatypes.XMlDataFeed {
		return nerror.New("must be an xml feed")
	}

	var xmlParser, notXMLParserType = parser.(*datatypes.XMLParser)
	if !notXMLParserType {
		return nerror.New("expecting datatypes.XMLParser type")
	}

	var psnAccountInfo = xmlParser.ByTag("//psn:Document/PsNOptionContract").(*datatypes.XMLElementIterator)
	if psnErr := psnAccountInfo.Err(); psnErr != nil {
		return nerror.Wrap(psnErr, "PsNSafekeepingAccountInformation tag not found")
	}

	for psnAccountInfo.HasNext() {
		var currentPSNAccountInfo = psnAccountInfo.NextNode()
		var result, parseErr = parseOptionContract(currentPSNAccountInfo)
		if parseErr != nil {
			handler.CollectErr(parser.Header(), nerror.WrapOnly(parseErr))
			continue
		}

		if cErr := handler.Collect(parser.Header(), result); cErr != nil {
			return nerror.WrapOnly(cErr)
		}
	}

	return nil
}

func parseOptionContract(node *etree.Element) (OptionContract, error) {
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
			data.CrtnDtTmDate = elem.Text()
		case "rprtngdt":
			var reportErr error
			data.ReportingDate, reportErr = parseReportingDate(elem.Text())
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
		case "msqseqno":
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

func parseClientDataBlocks(node *etree.Element) ([]SafeKeepingClientInfo, error) {
	var clientData []SafeKeepingClientInfo

	var dataBlock = node.SelectElement("Data")
	if dataBlock == nil {
		return clientData, nerror.New("Data tag not found")
	}

	var clientSfkDataList = dataBlock.SelectElements("ClntSfkData")
	// TODO: Should we return an error for an empty ClntSfkData block?
	if len(clientSfkDataList) == 0 {
		return clientData, nerror.New("No ClntSfkData found")
	}

	// initialize adequate space which should suffice for available list of
	// ClientSfkData.
	clientData = make([]SafeKeepingClientInfo, 0, len(clientSfkDataList))

	for _, clientSfkItem := range clientSfkDataList {
		var clientDataBlock, clientDataErr = parseDataBlock(clientSfkItem)
		if clientDataErr != nil {
			return clientData, clientDataErr
		}
		clientData = append(clientData, clientDataBlock)
	}
	return clientData, nil
}

func parseDataBlock(node *etree.Element) (SafeKeepingClientInfo, error) {
	var clientData, clientDataErr = parseDataBlockClientKey(node)
	if clientDataErr != nil {
		return clientData, nerror.Wrap(clientDataErr, "Failed to parse ClientKey")
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

func parseDataBlockClientKey(node *etree.Element) (SafeKeepingClientInfo, error) {
	var clientData SafeKeepingClientInfo

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
		case "acctid":
			data.AccountId = sfkItem.Text()
		case "extacctid":
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
			data.AccountOpeningDate, reportErr = parseReportingDate(sfkItem.Text())
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

func parseClientOptionContractBlocks(node *etree.Element) ([]OptionClientKeyData, error) {
	var contracts []OptionClientKeyData

	var dataBlock = node.SelectElement("Data")
	if dataBlock == nil {
		return contracts, nerror.New("Data tag not found")
	}

	var clientOptCtrctDataList = dataBlock.SelectElements("ClntOptCtrctData")
	if len(clientOptCtrctDataList) == 0 {
		return contracts, nerror.New("No ClntSfkData found")
	}

	contracts = make([]OptionClientKeyData, 0, len(clientOptCtrctDataList))
	for _, dataItem := range clientOptCtrctDataList {
		var clientKeyData, clKeyDataErr = parseClientOptionClientKeyData(dataItem)
		if clKeyDataErr != nil {
			return contracts, nerror.WrapOnly(clKeyDataErr)
		}
		contracts = append(contracts, clientKeyData)
	}
	return contracts, nil
}

func parseClientOptionClientKeyData(node *etree.Element) (OptionClientKeyData, error) {
	var data OptionClientKeyData

	var clientKeyElement = node.SelectElement("ClntKey")
	if clientKeyElement == nil {
		return data, nerror.New("ClntKey not found in ClntSfkData")
	}

	for _, field := range clientKeyElement.ChildElements() {
		switch strings.ToLower(field.Tag) {
		case "intrptunit":
			data.IntRptUnit = field.Text()
		case "intrptunitdesc":
			data.IntRptUnitDesc = field.Text()
		case "clntid":
			data.ClientId = field.Text()
		case "extclntid":
			data.ExtClientId = field.Text()
		}
	}

	var optionCtrlInfoList = node.SelectElements("OptnCtrctInf")
	data.Contracts = make([]OptionContractCtlInfo, 0, len(optionCtrlInfoList))
	for _, optionCtrlInfoItem := range optionCtrlInfoList {
		var contract, contractErr = parseOptionContractCtlInfo(optionCtrlInfoItem)
		if contractErr != nil {
			return data, nerror.WrapOnly(contractErr)
		}
		data.Contracts = append(data.Contracts, contract)
	}

	return data, nil
}

func parseOptionContractCtlInfo(parent *etree.Element) (OptionContractCtlInfo, error) {
	var data OptionContractCtlInfo
	for _, node := range parent.ChildElements() {
		switch strings.ToLower(node.Tag) {
		case "ctrctid":
			data.CtrctId = node.Text()
		case "dealdt":
			var dealDateErr error
			data.DealDate, dealDateErr = parseReportingDate(node.Text())
			if dealDateErr != nil {
				return data, nerror.WrapOnly(dealDateErr)
			}
		case "expirydt":
			var expiryDateErr error
			data.ExpiryDate, expiryDateErr = parseReportingDate(node.Text())
			if expiryDateErr != nil {
				return data, nerror.WrapOnly(expiryDateErr)
			}
		case "baseccycscd":
			data.BaseCurrencyCsCd = node.Text()
		case "baseccyisocd":
			data.BaseCurrencyIsoCd = node.Text()
		case "sttlmdt":
			var dateErr error
			data.SettlementDate, dateErr = parseReportingDate(node.Text())
			if dateErr != nil {
				return data, nerror.WrapOnly(dateErr)
			}
		case "buysttlmdt":
			var dateErr error
			data.BuySettlementDate, dateErr = parseReportingDate(node.Text())
			if dateErr != nil {
				return data, nerror.WrapOnly(dateErr)
			}
		case "sellsttlmdt":
			var dateErr error
			data.SellSettlementDate, dateErr = parseReportingDate(node.Text())
			if dateErr != nil {
				return data, nerror.WrapOnly(dateErr)
			}
		case "sttlmtpind":
			var valErr error
			data.SettlementTpInd, valErr = strconv.Atoi(node.Text())
			if valErr != nil {
				return data, nerror.WrapOnly(valErr)
			}
		case "buysellind":
			data.BuySellInd = node.Text()
		case "optntpcd":
			var valErr error
			data.OptionTpCd, valErr = strconv.Atoi(node.Text())
			if valErr != nil {
				return data, nerror.WrapOnly(valErr)
			}
		case "exticoptntpcd":
			data.ExticOptionTpCd = node.Text()
		case "ntnlleadccyamt":
			var valErr error
			data.NtnlLeadCurrencyAmount, valErr = decimal.NewFromString(node.Text())
			if valErr != nil {
				return data, nerror.WrapOnly(valErr)
			}
		case "ntnlleadccyisocd":
			data.NtnlLeadCurrencyIsoCd = node.Text()
		case "ntnlleadccyisodesc":
			data.NtnlLeadCurrencyIsoDesc = node.Text()
		case "strikeprcrate":
			var valErr error
			data.StrikePriceRate, valErr = decimal.NewFromString(node.Text())
			if valErr != nil {
				return data, nerror.WrapOnly(valErr)
			}
		case "ntnlcounterccyamt":
			var valErr error
			data.NtnlCounterCurrencyAmount, valErr = decimal.NewFromString(node.Text())
			if valErr != nil {
				return data, nerror.WrapOnly(valErr)
			}
		case "ntnlcounterccyisocd":
			data.NtnlCounterCurrencyISOCd = node.Text()
		case "ntnlcounterccyisodesc":
			data.NtnlCounterCurrencyISODesc = node.Text()
		case "premiumamt":
			var valErr error
			data.PremiumAmount, valErr = decimal.NewFromString(node.Text())
			if valErr != nil {
				return data, nerror.WrapOnly(valErr)
			}
		case "premiumccyisocd":
			data.PremiumAccountIsoCd = node.Text()
		case "premiumccyisodesc":
			data.PremiumAccountIsoDesc = node.Text()
		case "posdt":
			var dateErr error
			data.PositionDate, dateErr = parseReportingDate(node.Text())
			if dateErr != nil {
				return data, nerror.WrapOnly(dateErr)
			}
		case "revaltndt":
			var dateErr error
			data.RevaluationDate, dateErr = parseReportingDate(node.Text())
			if dateErr != nil {
				return data, nerror.WrapOnly(dateErr)
			}
		case "prtflid":
			data.PrtFlId = node.Text()
		case "mtmvalorigccyisocd":
			data.MtmValOrigCurrencyIsoCd = node.Text()
		case "mtmvalorigccyamt":
			var valErr error
			data.MtmValOrigCurrencyAmount, valErr = decimal.NewFromString(node.Text())
			if valErr != nil {
				return data, nerror.WrapOnly(valErr)
			}
		case "mtmvalrptccyisocd":
			data.MtmValReportCurrencyISOCd = node.Text()
		case "mtmvalrptccyamt":
			var valErr error
			data.MtmValReportCurrencyAmount, valErr = decimal.NewFromString(node.Text())
			if valErr != nil {
				return data, nerror.WrapOnly(valErr)
			}
		case "mtmvaldt":
			var dateErr error
			data.MtmValueDate, dateErr = parseReportingDate(node.Text())
			if dateErr != nil {
				return data, nerror.WrapOnly(dateErr)
			}
		}
	}
	return data, nil
}

func parseRFC3339Nano(value string) (time.Time, error) {
	return time.Parse(time.RFC3339Nano, value)
}

func parseReportingDate(value string) (time.Time, error) {
	return time.Parse(ReportingTimeFormat, value)
}
