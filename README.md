# GoPOC
A small library to demonstrate basic go-based ETL for mapping feeds into usable structure. 

## Platform Requirements

- Linux system or VM
- Go (go version go1.14.1 darwin/amd64)

 You can install go if on a linux based system with the following:

```bash
make setup
```

else download for your specific OS here: https://golang.org/dl/

## Dependencies Setup
To setup locally, ensure to first download all modules for project with:

```bash
make deps
```

You can run the following manually (if desired):
```bash
go mod download
```

### How to run the test suite

Project comes with a basic test, which can be executed by running:

```
make test
```

## Easiest Local Run

```bash
make creditsuisse_cli
```

See should be presented with a dump of the extracted data like below:

```bash
Requesting dir for `creditSuisse` -> `Credit Suisse`
Requesting file `15_01_26_TDOPT_530103123_MODIFIED.xml`
Requesting file `15_01_50_SDSA_530423111.xml`
2020/07/31 15:19:54 ---------Credit Suise SafeKeeping Accounts----------
([]parsers.SafeKeepingAccountInformation) (len=1 cap=1) {
 (parsers.SafeKeepingAccountInformation) {
  Account: (parsers.AccountHeader) {
   Sender: (parsers.SenderReceiverInfo) {
    SenderBIC: (string) (len=11) "CSPBSGSGPSN",
    ReceiverBIC: (string) (len=11) "CRESSGSGECO"
   },
   Info: (parsers.FLInfo) {
    DWHMsgId: (string) (len=30) "DWHMSGIDSGSDSA2020022207394717",
    LocationISO: (string) (len=2) "SG",
    CrtnDtTmDate: (string) (len=29) "2020-02-22T08:39:47.485874000",
    ReportingDate: (time.Time) 2020-02-21 00:00:00 +0000 UTC,
    Type: (string) (len=45) "Static Data - Safekeeping Account Information",
    TypeCD: (string) (len=4) "SDSA",
    Version: (string) (len=5) "1.0.0",
    ClientFormat: (string) (len=3) "XML",
    MessageSequenceNumber: (int) 1
   }
  },
  SafeKeepingInfo: ([]parsers.SafeKeepingClientInfo) (len=9 cap=9) {
   (parsers.SafeKeepingClientInfo) {
    IntRptUnit: (string) (len=4) "0973",
    IntRptUnitDesc: (string) (len=35) "CS Private Banking Singapore Branch",
    ClientId: (string) (len=6) "129638",
    ExtClientId: (string) (len=6) "129638",
    SafeKeepingAccounts: (map[string]parsers.SafeKeepingAccountInfo) (len=2) {
     (string) (len=8) "129638-1": (parsers.SafeKeepingAccountInfo) {
      AccountId: (string) (len=8) "129638-1",
      ExtAccountId: (string) (len=8) "129638-1",
      InvestmentCurrencyISOCd: (string) (len=3) "USD",
      InvestmentCurrencyISODesc: (string) (len=10) "US dollars",
      AccountTpCd: (int) 100,
      AccountTpDesc: (string) (len=19) "Safekeeping account",
      AccountSTS: (bool) true,
      AccountClsgSTS: (bool) false,
      AccountPldgdInd: (bool) false,
      AccountPrtlyPldfdInd: (bool) false,
      AccountLmtPldgdInd: (bool) false,
      AccountOpeningDate: (time.Time) 2014-10-28 00:00:00 +0000 UTC
     },
     (string) (len=9) "129638-80": (parsers.SafeKeepingAccountInfo) {
      AccountId: (string) (len=9) "129638-80",
      ExtAccountId: (string) (len=9) "129638-80",
      InvestmentCurrencyISOCd: (string) (len=3) "EUR",
      InvestmentCurrencyISODesc: (string) (len=4) "Euro",
      AccountTpCd: (int) 100,
      AccountTpDesc: (string) (len=19) "Safekeeping account",
      AccountSTS: (bool) false,
      AccountClsgSTS: (bool) true,
      AccountPldgdInd: (bool) false,
      AccountPrtlyPldfdInd: (bool) false,
      AccountLmtPldgdInd: (bool) false,
      AccountOpeningDate: (time.Time) 2014-12-02 00:00:00 +0000 UTC
     }
    }
   },
   (parsers.SafeKeepingClientInfo) {
    IntRptUnit: (string) (len=4) "0973",
    IntRptUnitDesc: (string) (len=35) "CS Private Banking Singapore Branch",
    ClientId: (string) (len=6) "132676",
    ExtClientId: (string) (len=6) "132676",
    SafeKeepingAccounts: (map[string]parsers.SafeKeepingAccountInfo) (len=1) {
     (string) (len=8) "132676-1": (parsers.SafeKeepingAccountInfo) {
      AccountId: (string) (len=8) "132676-1",
      ExtAccountId: (string) (len=8) "132676-1",
      InvestmentCurrencyISOCd: (string) (len=3) "USD",
      InvestmentCurrencyISODesc: (string) (len=10) "US dollars",
      AccountTpCd: (int) 100,
      AccountTpDesc: (string) (len=19) "Safekeeping account",
      AccountSTS: (bool) true,
      AccountClsgSTS: (bool) false,
      AccountPldgdInd: (bool) false,
      AccountPrtlyPldfdInd: (bool) false,
      AccountLmtPldgdInd: (bool) false,
      AccountOpeningDate: (time.Time) 2016-06-28 00:00:00 +0000 UTC
     }
    }
   },
   (parsers.SafeKeepingClientInfo) {
    IntRptUnit: (string) (len=4) "0973",
    IntRptUnitDesc: (string) (len=35) "CS Private Banking Singapore Branch",
    ClientId: (string) (len=5) "31288",
    ExtClientId: (string) (len=5) "31288",
    SafeKeepingAccounts: (map[string]parsers.SafeKeepingAccountInfo) (len=1) {
     (string) (len=7) "31288-1": (parsers.SafeKeepingAccountInfo) {
      AccountId: (string) (len=7) "31288-1",
      ExtAccountId: (string) (len=7) "31288-1",
      InvestmentCurrencyISOCd: (string) (len=3) "USD",
      InvestmentCurrencyISODesc: (string) (len=10) "US dollars",
      AccountTpCd: (int) 100,
      AccountTpDesc: (string) (len=19) "Safekeeping account",
      AccountSTS: (bool) true,
      AccountClsgSTS: (bool) false,
      AccountPldgdInd: (bool) false,
      AccountPrtlyPldfdInd: (bool) false,
      AccountLmtPldgdInd: (bool) false,
      AccountOpeningDate: (time.Time) 2004-07-19 00:00:00 +0000 UTC
     }
    }
   },
   (parsers.SafeKeepingClientInfo) {
    IntRptUnit: (string) (len=4) "0973",
    IntRptUnitDesc: (string) (len=35) "CS Private Banking Singapore Branch",
    ClientId: (string) (len=5) "40737",
    ExtClientId: (string) (len=5) "40737",
    SafeKeepingAccounts: (map[string]parsers.SafeKeepingAccountInfo) (len=1) {
     (string) (len=7) "40737-1": (parsers.SafeKeepingAccountInfo) {
      AccountId: (string) (len=7) "40737-1",
      ExtAccountId: (string) (len=7) "40737-1",
      InvestmentCurrencyISOCd: (string) (len=3) "USD",
      InvestmentCurrencyISODesc: (string) (len=10) "US dollars",
      AccountTpCd: (int) 100,
      AccountTpDesc: (string) (len=19) "Safekeeping account",
      AccountSTS: (bool) true,
      AccountClsgSTS: (bool) false,
      AccountPldgdInd: (bool) false,
      AccountPrtlyPldfdInd: (bool) false,
      AccountLmtPldgdInd: (bool) false,
      AccountOpeningDate: (time.Time) 2016-09-09 00:00:00 +0000 UTC
     }
    }
   },
   (parsers.SafeKeepingClientInfo) {
    IntRptUnit: (string) (len=4) "0973",
    IntRptUnitDesc: (string) (len=35) "CS Private Banking Singapore Branch",
    ClientId: (string) (len=5) "40825",
    ExtClientId: (string) (len=5) "40825",
    SafeKeepingAccounts: (map[string]parsers.SafeKeepingAccountInfo) (len=1) {
     (string) (len=7) "40825-1": (parsers.SafeKeepingAccountInfo) {
      AccountId: (string) (len=7) "40825-1",
      ExtAccountId: (string) (len=7) "40825-1",
      InvestmentCurrencyISOCd: (string) (len=3) "USD",
      InvestmentCurrencyISODesc: (string) (len=10) "US dollars",
      AccountTpCd: (int) 100,
      AccountTpDesc: (string) (len=19) "Safekeeping account",
      AccountSTS: (bool) true,
      AccountClsgSTS: (bool) false,
      AccountPldgdInd: (bool) false,
      AccountPrtlyPldfdInd: (bool) false,
      AccountLmtPldgdInd: (bool) false,
      AccountOpeningDate: (time.Time) 2018-07-23 00:00:00 +0000 UTC
     }
    }
   },
   (parsers.SafeKeepingClientInfo) {
    IntRptUnit: (string) (len=4) "0973",
    IntRptUnitDesc: (string) (len=35) "CS Private Banking Singapore Branch",
    ClientId: (string) (len=5) "41130",
    ExtClientId: (string) (len=5) "41130",
    SafeKeepingAccounts: (map[string]parsers.SafeKeepingAccountInfo) (len=1) {
     (string) (len=7) "41130-1": (parsers.SafeKeepingAccountInfo) {
      AccountId: (string) (len=7) "41130-1",
      ExtAccountId: (string) (len=7) "41130-1",
      InvestmentCurrencyISOCd: (string) (len=3) "USD",
      InvestmentCurrencyISODesc: (string) (len=10) "US dollars",
      AccountTpCd: (int) 100,
      AccountTpDesc: (string) (len=19) "Safekeeping account",
      AccountSTS: (bool) true,
      AccountClsgSTS: (bool) false,
      AccountPldgdInd: (bool) false,
      AccountPrtlyPldfdInd: (bool) false,
      AccountLmtPldgdInd: (bool) false,
      AccountOpeningDate: (time.Time) 2019-01-31 00:00:00 +0000 UTC
     }
    }
   },
   (parsers.SafeKeepingClientInfo) {
    IntRptUnit: (string) (len=4) "0973",
    IntRptUnitDesc: (string) (len=35) "CS Private Banking Singapore Branch",
    ClientId: (string) (len=5) "71323",
    ExtClientId: (string) (len=5) "71323",
    SafeKeepingAccounts: (map[string]parsers.SafeKeepingAccountInfo) (len=2) {
     (string) (len=7) "71323-1": (parsers.SafeKeepingAccountInfo) {
      AccountId: (string) (len=7) "71323-1",
      ExtAccountId: (string) (len=7) "71323-1",
      InvestmentCurrencyISOCd: (string) (len=3) "USD",
      InvestmentCurrencyISODesc: (string) (len=10) "US dollars",
      AccountTpCd: (int) 100,
      AccountTpDesc: (string) (len=19) "Safekeeping account",
      AccountSTS: (bool) true,
      AccountClsgSTS: (bool) false,
      AccountPldgdInd: (bool) false,
      AccountPrtlyPldfdInd: (bool) false,
      AccountLmtPldgdInd: (bool) false,
      AccountOpeningDate: (time.Time) 2018-11-01 00:00:00 +0000 UTC
     },
     (string) (len=8) "71323-59": (parsers.SafeKeepingAccountInfo) {
      AccountId: (string) (len=8) "71323-59",
      ExtAccountId: (string) (len=8) "71323-59",
      InvestmentCurrencyISOCd: (string) (len=3) "USD",
      InvestmentCurrencyISODesc: (string) (len=10) "US dollars",
      AccountTpCd: (int) 100,
      AccountTpDesc: (string) (len=19) "Safekeeping account",
      AccountSTS: (bool) true,
      AccountClsgSTS: (bool) false,
      AccountPldgdInd: (bool) false,
      AccountPrtlyPldfdInd: (bool) false,
      AccountLmtPldgdInd: (bool) false,
      AccountOpeningDate: (time.Time) 2018-11-01 00:00:00 +0000 UTC
     }
    }
   },
   (parsers.SafeKeepingClientInfo) {
    IntRptUnit: (string) (len=4) "0973",
    IntRptUnitDesc: (string) (len=35) "CS Private Banking Singapore Branch",
    ClientId: (string) (len=5) "71788",
    ExtClientId: (string) (len=5) "71788",
    SafeKeepingAccounts: (map[string]parsers.SafeKeepingAccountInfo) (len=1) {
     (string) (len=7) "71788-1": (parsers.SafeKeepingAccountInfo) {
      AccountId: (string) (len=7) "71788-1",
      ExtAccountId: (string) (len=7) "71788-1",
      InvestmentCurrencyISOCd: (string) (len=3) "USD",
      InvestmentCurrencyISODesc: (string) (len=10) "US dollars",
      AccountTpCd: (int) 100,
      AccountTpDesc: (string) (len=19) "Safekeeping account",
      AccountSTS: (bool) true,
      AccountClsgSTS: (bool) false,
      AccountPldgdInd: (bool) false,
      AccountPrtlyPldfdInd: (bool) false,
      AccountLmtPldgdInd: (bool) false,
      AccountOpeningDate: (time.Time) 2019-04-15 00:00:00 +0000 UTC
     }
    }
   },
   (parsers.SafeKeepingClientInfo) {
    IntRptUnit: (string) (len=4) "0973",
    IntRptUnitDesc: (string) (len=35) "CS Private Banking Singapore Branch",
    ClientId: (string) (len=5) "84959",
    ExtClientId: (string) (len=5) "84959",
    SafeKeepingAccounts: (map[string]parsers.SafeKeepingAccountInfo) (len=2) {
     (string) (len=7) "84959-1": (parsers.SafeKeepingAccountInfo) {
      AccountId: (string) (len=7) "84959-1",
      ExtAccountId: (string) (len=7) "84959-1",
      InvestmentCurrencyISOCd: (string) (len=3) "USD",
      InvestmentCurrencyISODesc: (string) (len=10) "US dollars",
      AccountTpCd: (int) 100,
      AccountTpDesc: (string) (len=19) "Safekeeping account",
      AccountSTS: (bool) true,
      AccountClsgSTS: (bool) false,
      AccountPldgdInd: (bool) false,
      AccountPrtlyPldfdInd: (bool) false,
      AccountLmtPldgdInd: (bool) false,
      AccountOpeningDate: (time.Time) 2016-06-30 00:00:00 +0000 UTC
     },
     (string) (len=7) "84959-2": (parsers.SafeKeepingAccountInfo) {
      AccountId: (string) (len=7) "84959-2",
      ExtAccountId: (string) (len=7) "84959-2",
      InvestmentCurrencyISOCd: (string) (len=3) "USD",
      InvestmentCurrencyISODesc: (string) (len=10) "US dollars",
      AccountTpCd: (int) 100,
      AccountTpDesc: (string) (len=19) "Safekeeping account",
      AccountSTS: (bool) false,
      AccountClsgSTS: (bool) true,
      AccountPldgdInd: (bool) false,
      AccountPrtlyPldfdInd: (bool) false,
      AccountLmtPldgdInd: (bool) false,
      AccountOpeningDate: (time.Time) 2016-12-02 00:00:00 +0000 UTC
     }
    }
   }
  }
 }
}
2020/07/31 15:19:54 ---------Credit Suise Option Contracts--------------
([]parsers.OptionContract) (len=1 cap=1) {
 (parsers.OptionContract) {
  Account: (parsers.AccountHeader) {
   Sender: (parsers.SenderReceiverInfo) {
    SenderBIC: (string) (len=11) "CSPBSGSGPSN",
    ReceiverBIC: (string) (len=11) "CRESSGSGECO"
   },
   Info: (parsers.FLInfo) {
    DWHMsgId: (string) (len=32) "DWHMSGIDSGTDOPT20200221005317133",
    LocationISO: (string) (len=2) "SG",
    CrtnDtTmDate: (string) (len=29) "2020-02-21T01:53:17.827223000",
    ReportingDate: (time.Time) 2020-02-20 00:00:00 +0000 UTC,
    Type: (string) (len=37) "Transactional Data - Options Contract",
    TypeCD: (string) (len=5) "TDOPT",
    Version: (string) (len=5) "1.0.0",
    ClientFormat: (string) (len=3) "XML",
    MessageSequenceNumber: (int) 1
   }
  },
  Contracts: ([]parsers.OptionClientKeyData) (len=1 cap=1) {
   (parsers.OptionClientKeyData) {
    IntRptUnit: (string) (len=4) "0973",
    IntRptUnitDesc: (string) (len=35) "CS Private Banking Singapore Branch",
    ClientId: (string) (len=5) "71788",
    ExtClientId: (string) (len=5) "71788",
    Contracts: ([]parsers.OptionContractCtlInfo) (len=2 cap=2) {
     (parsers.OptionContractCtlInfo) {
      CtrctId: (string) (len=15) "DXTRA1924600035",
      DealDate: (time.Time) 2019-09-03 00:00:00 +0000 UTC,
      ExpiryDate: (time.Time) 2020-09-01 00:00:00 +0000 UTC,
      BaseCurrencyCsCd: (string) (len=4) "0010",
      BaseCurrencyIsoCd: (string) (len=3) "CHF",
      SettlementDate: (time.Time) 2020-09-03 00:00:00 +0000 UTC,
      BuySettlementDate: (time.Time) 2020-09-03 00:00:00 +0000 UTC,
      SellSettlementDate: (time.Time) 2020-09-03 00:00:00 +0000 UTC,
      SettlementTpInd: (int) 5,
      BuySellInd: (string) (len=3) "Buy",
      OptionTpCd: (int) 4,
      ExticOptionTpCd: (string) (len=23) "TargetRedemptionForward",
      NtnlLeadCurrencyAmount: (decimal.Decimal) 5200000,
      NtnlLeadCurrencyIsoCd: (string) (len=3) "EUR",
      NtnlLeadCurrencyIsoDesc: (string) (len=4) "Euro",
      StrikePriceRate: (decimal.Decimal) 120.55,
      NtnlCounterCurrencyAmount: (decimal.Decimal) 626860000,
      NtnlCounterCurrencyISOCd: (string) (len=3) "JPY",
      NtnlCounterCurrencyISODesc: (string) (len=3) "Yen",
      PremiumAmount: (decimal.Decimal) 0,
      PremiumAccountIsoCd: (string) (len=3) "USD",
      PremiumAccountIsoDesc: (string) (len=10) "US dollars",
      PositionDate: (time.Time) 2020-02-20 00:00:00 +0000 UTC,
      RevaluationDate: (time.Time) 2020-09-03 00:00:00 +0000 UTC,
      PrtFlId: (string) (len=7) "71788-1",
      MtmValOrigCurrencyIsoCd: (string) (len=3) "EUR",
      MtmValOrigCurrencyAmount: (decimal.Decimal) 5342.981836,
      MtmValReportCurrencyISOCd: (string) (len=3) "USD",
      MtmValReportCurrencyAmount: (decimal.Decimal) 5786.44932842,
      MtmValueDate: (time.Time) 2020-02-18 00:00:00 +0000 UTC
     },
     (parsers.OptionContractCtlInfo) {
      CtrctId: (string) (len=15) "DXTRA1928900047",
      DealDate: (time.Time) 2019-10-16 00:00:00 +0000 UTC,
      ExpiryDate: (time.Time) 2020-10-14 00:00:00 +0000 UTC,
      BaseCurrencyCsCd: (string) (len=4) "0010",
      BaseCurrencyIsoCd: (string) (len=3) "CHF",
      SettlementDate: (time.Time) 2020-10-16 00:00:00 +0000 UTC,
      BuySettlementDate: (time.Time) 2020-10-16 00:00:00 +0000 UTC,
      SellSettlementDate: (time.Time) 2020-10-16 00:00:00 +0000 UTC,
      SettlementTpInd: (int) 5,
      BuySellInd: (string) (len=3) "Buy",
      OptionTpCd: (int) 4,
      ExticOptionTpCd: (string) (len=28) "PivotTargetRedemptionForward",
      NtnlLeadCurrencyAmount: (decimal.Decimal) 3120000,
      NtnlLeadCurrencyIsoCd: (string) (len=3) "NZD",
      NtnlLeadCurrencyIsoDesc: (string) (len=19) "New Zealand dollars",
      StrikePriceRate: (decimal.Decimal) 0.6618,
      NtnlCounterCurrencyAmount: (decimal.Decimal) 2064816,
      NtnlCounterCurrencyISOCd: (string) (len=3) "USD",
      NtnlCounterCurrencyISODesc: (string) (len=10) "US dollars",
      PremiumAmount: (decimal.Decimal) 0,
      PremiumAccountIsoCd: (string) (len=3) "NZD",
      PremiumAccountIsoDesc: (string) (len=19) "New Zealand dollars",
      PositionDate: (time.Time) 2020-02-20 00:00:00 +0000 UTC,
      RevaluationDate: (time.Time) 2020-10-16 00:00:00 +0000 UTC,
      PrtFlId: (string) (len=7) "71788-1",
      MtmValOrigCurrencyIsoCd: (string) (len=3) "NZD",
      MtmValOrigCurrencyAmount: (decimal.Decimal) -2104.313972,
      MtmValReportCurrencyISOCd: (string) (len=3) "USD",
      MtmValReportCurrencyAmount: (decimal.Decimal) -1349.28611859,
      MtmValueDate: (time.Time) 2020-02-18 00:00:00 +0000 UTC
     }
    }
   }
  }
 }
}
03:19:54 darkside@TUFStar creditsuisse_mapper ±|master ✗|→ go run main.go 
Requesting dir for `creditSuisse` -> `Credit Suisse`
Requesting file `15_01_26_TDOPT_530103123_MODIFIED.xml`
Requesting file `15_01_50_SDSA_530423111.xml`
2020/07/31 15:24:27 ---------Credit Suise SafeKeeping Accounts----------
([]parsers.SafeKeepingAccountInformation) (len=1 cap=1) {
 (parsers.SafeKeepingAccountInformation) {
  Account: (parsers.AccountHeader) {
   Sender: (parsers.SenderReceiverInfo) {
    SenderBIC: (string) (len=11) "CSPBSGSGPSN",
    ReceiverBIC: (string) (len=11) "CRESSGSGECO"
   },
   Info: (parsers.FLInfo) {
    DWHMsgId: (string) (len=30) "DWHMSGIDSGSDSA2020022207394717",
    LocationISO: (string) (len=2) "SG",
    CrtnDtTmDate: (string) (len=29) "2020-02-22T08:39:47.485874000",
    ReportingDate: (time.Time) 2020-02-21 00:00:00 +0000 UTC,
    Type: (string) (len=45) "Static Data - Safekeeping Account Information",
    TypeCD: (string) (len=4) "SDSA",
    Version: (string) (len=5) "1.0.0",
    ClientFormat: (string) (len=3) "XML",
    MessageSequenceNumber: (int) 1
   }
  },
  SafeKeepingInfo: ([]parsers.SafeKeepingClientInfo) (len=9 cap=9) {
   (parsers.SafeKeepingClientInfo) {
    IntRptUnit: (string) (len=4) "0973",
    IntRptUnitDesc: (string) (len=35) "CS Private Banking Singapore Branch",
    ClientId: (string) (len=6) "129638",
    ExtClientId: (string) (len=6) "129638",
    SafeKeepingAccounts: (map[string]parsers.SafeKeepingAccountInfo) (len=2) {
     (string) (len=8) "129638-1": (parsers.SafeKeepingAccountInfo) {
      AccountId: (string) (len=8) "129638-1",
      ExtAccountId: (string) (len=8) "129638-1",
      InvestmentCurrencyISOCd: (string) (len=3) "USD",
      InvestmentCurrencyISODesc: (string) (len=10) "US dollars",
      AccountTpCd: (int) 100,
      AccountTpDesc: (string) (len=19) "Safekeeping account",
      AccountSTS: (bool) true,
      AccountClsgSTS: (bool) false,
      AccountPldgdInd: (bool) false,
      AccountPrtlyPldfdInd: (bool) false,
      AccountLmtPldgdInd: (bool) false,
      AccountOpeningDate: (time.Time) 2014-10-28 00:00:00 +0000 UTC
     },
     (string) (len=9) "129638-80": (parsers.SafeKeepingAccountInfo) {
      AccountId: (string) (len=9) "129638-80",
      ExtAccountId: (string) (len=9) "129638-80",
      InvestmentCurrencyISOCd: (string) (len=3) "EUR",
      InvestmentCurrencyISODesc: (string) (len=4) "Euro",
      AccountTpCd: (int) 100,
      AccountTpDesc: (string) (len=19) "Safekeeping account",
      AccountSTS: (bool) false,
      AccountClsgSTS: (bool) true,
      AccountPldgdInd: (bool) false,
      AccountPrtlyPldfdInd: (bool) false,
      AccountLmtPldgdInd: (bool) false,
      AccountOpeningDate: (time.Time) 2014-12-02 00:00:00 +0000 UTC
     }
    }
   },
   (parsers.SafeKeepingClientInfo) {
    IntRptUnit: (string) (len=4) "0973",
    IntRptUnitDesc: (string) (len=35) "CS Private Banking Singapore Branch",
    ClientId: (string) (len=6) "132676",
    ExtClientId: (string) (len=6) "132676",
    SafeKeepingAccounts: (map[string]parsers.SafeKeepingAccountInfo) (len=1) {
     (string) (len=8) "132676-1": (parsers.SafeKeepingAccountInfo) {
      AccountId: (string) (len=8) "132676-1",
      ExtAccountId: (string) (len=8) "132676-1",
      InvestmentCurrencyISOCd: (string) (len=3) "USD",
      InvestmentCurrencyISODesc: (string) (len=10) "US dollars",
      AccountTpCd: (int) 100,
      AccountTpDesc: (string) (len=19) "Safekeeping account",
      AccountSTS: (bool) true,
      AccountClsgSTS: (bool) false,
      AccountPldgdInd: (bool) false,
      AccountPrtlyPldfdInd: (bool) false,
      AccountLmtPldgdInd: (bool) false,
      AccountOpeningDate: (time.Time) 2016-06-28 00:00:00 +0000 UTC
     }
    }
   },
   (parsers.SafeKeepingClientInfo) {
    IntRptUnit: (string) (len=4) "0973",
    IntRptUnitDesc: (string) (len=35) "CS Private Banking Singapore Branch",
    ClientId: (string) (len=5) "31288",
    ExtClientId: (string) (len=5) "31288",
    SafeKeepingAccounts: (map[string]parsers.SafeKeepingAccountInfo) (len=1) {
     (string) (len=7) "31288-1": (parsers.SafeKeepingAccountInfo) {
      AccountId: (string) (len=7) "31288-1",
      ExtAccountId: (string) (len=7) "31288-1",
      InvestmentCurrencyISOCd: (string) (len=3) "USD",
      InvestmentCurrencyISODesc: (string) (len=10) "US dollars",
      AccountTpCd: (int) 100,
      AccountTpDesc: (string) (len=19) "Safekeeping account",
      AccountSTS: (bool) true,
      AccountClsgSTS: (bool) false,
      AccountPldgdInd: (bool) false,
      AccountPrtlyPldfdInd: (bool) false,
      AccountLmtPldgdInd: (bool) false,
      AccountOpeningDate: (time.Time) 2004-07-19 00:00:00 +0000 UTC
     }
    }
   },
   (parsers.SafeKeepingClientInfo) {
    IntRptUnit: (string) (len=4) "0973",
    IntRptUnitDesc: (string) (len=35) "CS Private Banking Singapore Branch",
    ClientId: (string) (len=5) "40737",
    ExtClientId: (string) (len=5) "40737",
    SafeKeepingAccounts: (map[string]parsers.SafeKeepingAccountInfo) (len=1) {
     (string) (len=7) "40737-1": (parsers.SafeKeepingAccountInfo) {
      AccountId: (string) (len=7) "40737-1",
      ExtAccountId: (string) (len=7) "40737-1",
      InvestmentCurrencyISOCd: (string) (len=3) "USD",
      InvestmentCurrencyISODesc: (string) (len=10) "US dollars",
      AccountTpCd: (int) 100,
      AccountTpDesc: (string) (len=19) "Safekeeping account",
      AccountSTS: (bool) true,
      AccountClsgSTS: (bool) false,
      AccountPldgdInd: (bool) false,
      AccountPrtlyPldfdInd: (bool) false,
      AccountLmtPldgdInd: (bool) false,
      AccountOpeningDate: (time.Time) 2016-09-09 00:00:00 +0000 UTC
     }
    }
   },
   (parsers.SafeKeepingClientInfo) {
    IntRptUnit: (string) (len=4) "0973",
    IntRptUnitDesc: (string) (len=35) "CS Private Banking Singapore Branch",
    ClientId: (string) (len=5) "40825",
    ExtClientId: (string) (len=5) "40825",
    SafeKeepingAccounts: (map[string]parsers.SafeKeepingAccountInfo) (len=1) {
     (string) (len=7) "40825-1": (parsers.SafeKeepingAccountInfo) {
      AccountId: (string) (len=7) "40825-1",
      ExtAccountId: (string) (len=7) "40825-1",
      InvestmentCurrencyISOCd: (string) (len=3) "USD",
      InvestmentCurrencyISODesc: (string) (len=10) "US dollars",
      AccountTpCd: (int) 100,
      AccountTpDesc: (string) (len=19) "Safekeeping account",
      AccountSTS: (bool) true,
      AccountClsgSTS: (bool) false,
      AccountPldgdInd: (bool) false,
      AccountPrtlyPldfdInd: (bool) false,
      AccountLmtPldgdInd: (bool) false,
      AccountOpeningDate: (time.Time) 2018-07-23 00:00:00 +0000 UTC
     }
    }
   },
   (parsers.SafeKeepingClientInfo) {
    IntRptUnit: (string) (len=4) "0973",
    IntRptUnitDesc: (string) (len=35) "CS Private Banking Singapore Branch",
    ClientId: (string) (len=5) "41130",
    ExtClientId: (string) (len=5) "41130",
    SafeKeepingAccounts: (map[string]parsers.SafeKeepingAccountInfo) (len=1) {
     (string) (len=7) "41130-1": (parsers.SafeKeepingAccountInfo) {
      AccountId: (string) (len=7) "41130-1",
      ExtAccountId: (string) (len=7) "41130-1",
      InvestmentCurrencyISOCd: (string) (len=3) "USD",
      InvestmentCurrencyISODesc: (string) (len=10) "US dollars",
      AccountTpCd: (int) 100,
      AccountTpDesc: (string) (len=19) "Safekeeping account",
      AccountSTS: (bool) true,
      AccountClsgSTS: (bool) false,
      AccountPldgdInd: (bool) false,
      AccountPrtlyPldfdInd: (bool) false,
      AccountLmtPldgdInd: (bool) false,
      AccountOpeningDate: (time.Time) 2019-01-31 00:00:00 +0000 UTC
     }
    }
   },
   (parsers.SafeKeepingClientInfo) {
    IntRptUnit: (string) (len=4) "0973",
    IntRptUnitDesc: (string) (len=35) "CS Private Banking Singapore Branch",
    ClientId: (string) (len=5) "71323",
    ExtClientId: (string) (len=5) "71323",
    SafeKeepingAccounts: (map[string]parsers.SafeKeepingAccountInfo) (len=2) {
     (string) (len=7) "71323-1": (parsers.SafeKeepingAccountInfo) {
      AccountId: (string) (len=7) "71323-1",
      ExtAccountId: (string) (len=7) "71323-1",
      InvestmentCurrencyISOCd: (string) (len=3) "USD",
      InvestmentCurrencyISODesc: (string) (len=10) "US dollars",
      AccountTpCd: (int) 100,
      AccountTpDesc: (string) (len=19) "Safekeeping account",
      AccountSTS: (bool) true,
      AccountClsgSTS: (bool) false,
      AccountPldgdInd: (bool) false,
      AccountPrtlyPldfdInd: (bool) false,
      AccountLmtPldgdInd: (bool) false,
      AccountOpeningDate: (time.Time) 2018-11-01 00:00:00 +0000 UTC
     },
     (string) (len=8) "71323-59": (parsers.SafeKeepingAccountInfo) {
      AccountId: (string) (len=8) "71323-59",
      ExtAccountId: (string) (len=8) "71323-59",
      InvestmentCurrencyISOCd: (string) (len=3) "USD",
      InvestmentCurrencyISODesc: (string) (len=10) "US dollars",
      AccountTpCd: (int) 100,
      AccountTpDesc: (string) (len=19) "Safekeeping account",
      AccountSTS: (bool) true,
      AccountClsgSTS: (bool) false,
      AccountPldgdInd: (bool) false,
      AccountPrtlyPldfdInd: (bool) false,
      AccountLmtPldgdInd: (bool) false,
      AccountOpeningDate: (time.Time) 2018-11-01 00:00:00 +0000 UTC
     }
    }
   },
   (parsers.SafeKeepingClientInfo) {
    IntRptUnit: (string) (len=4) "0973",
    IntRptUnitDesc: (string) (len=35) "CS Private Banking Singapore Branch",
    ClientId: (string) (len=5) "71788",
    ExtClientId: (string) (len=5) "71788",
    SafeKeepingAccounts: (map[string]parsers.SafeKeepingAccountInfo) (len=1) {
     (string) (len=7) "71788-1": (parsers.SafeKeepingAccountInfo) {
      AccountId: (string) (len=7) "71788-1",
      ExtAccountId: (string) (len=7) "71788-1",
      InvestmentCurrencyISOCd: (string) (len=3) "USD",
      InvestmentCurrencyISODesc: (string) (len=10) "US dollars",
      AccountTpCd: (int) 100,
      AccountTpDesc: (string) (len=19) "Safekeeping account",
      AccountSTS: (bool) true,
      AccountClsgSTS: (bool) false,
      AccountPldgdInd: (bool) false,
      AccountPrtlyPldfdInd: (bool) false,
      AccountLmtPldgdInd: (bool) false,
      AccountOpeningDate: (time.Time) 2019-04-15 00:00:00 +0000 UTC
     }
    }
   },
   (parsers.SafeKeepingClientInfo) {
    IntRptUnit: (string) (len=4) "0973",
    IntRptUnitDesc: (string) (len=35) "CS Private Banking Singapore Branch",
    ClientId: (string) (len=5) "84959",
    ExtClientId: (string) (len=5) "84959",
    SafeKeepingAccounts: (map[string]parsers.SafeKeepingAccountInfo) (len=2) {
     (string) (len=7) "84959-1": (parsers.SafeKeepingAccountInfo) {
      AccountId: (string) (len=7) "84959-1",
      ExtAccountId: (string) (len=7) "84959-1",
      InvestmentCurrencyISOCd: (string) (len=3) "USD",
      InvestmentCurrencyISODesc: (string) (len=10) "US dollars",
      AccountTpCd: (int) 100,
      AccountTpDesc: (string) (len=19) "Safekeeping account",
      AccountSTS: (bool) true,
      AccountClsgSTS: (bool) false,
      AccountPldgdInd: (bool) false,
      AccountPrtlyPldfdInd: (bool) false,
      AccountLmtPldgdInd: (bool) false,
      AccountOpeningDate: (time.Time) 2016-06-30 00:00:00 +0000 UTC
     },
     (string) (len=7) "84959-2": (parsers.SafeKeepingAccountInfo) {
      AccountId: (string) (len=7) "84959-2",
      ExtAccountId: (string) (len=7) "84959-2",
      InvestmentCurrencyISOCd: (string) (len=3) "USD",
      InvestmentCurrencyISODesc: (string) (len=10) "US dollars",
      AccountTpCd: (int) 100,
      AccountTpDesc: (string) (len=19) "Safekeeping account",
      AccountSTS: (bool) false,
      AccountClsgSTS: (bool) true,
      AccountPldgdInd: (bool) false,
      AccountPrtlyPldfdInd: (bool) false,
      AccountLmtPldgdInd: (bool) false,
      AccountOpeningDate: (time.Time) 2016-12-02 00:00:00 +0000 UTC
     }
    }
   }
  }
 }
}
2020/07/31 15:24:27 ---------Credit Suise Option Contracts--------------
([]parsers.OptionContract) (len=1 cap=1) {
 (parsers.OptionContract) {
  Account: (parsers.AccountHeader) {
   Sender: (parsers.SenderReceiverInfo) {
    SenderBIC: (string) (len=11) "CSPBSGSGPSN",
    ReceiverBIC: (string) (len=11) "CRESSGSGECO"
   },
   Info: (parsers.FLInfo) {
    DWHMsgId: (string) (len=32) "DWHMSGIDSGTDOPT20200221005317133",
    LocationISO: (string) (len=2) "SG",
    CrtnDtTmDate: (string) (len=29) "2020-02-21T01:53:17.827223000",
    ReportingDate: (time.Time) 2020-02-20 00:00:00 +0000 UTC,
    Type: (string) (len=37) "Transactional Data - Options Contract",
    TypeCD: (string) (len=5) "TDOPT",
    Version: (string) (len=5) "1.0.0",
    ClientFormat: (string) (len=3) "XML",
    MessageSequenceNumber: (int) 1
   }
  },
  Contracts: ([]parsers.OptionClientKeyData) (len=1 cap=1) {
   (parsers.OptionClientKeyData) {
    IntRptUnit: (string) (len=4) "0973",
    IntRptUnitDesc: (string) (len=35) "CS Private Banking Singapore Branch",
    ClientId: (string) (len=5) "71788",
    ExtClientId: (string) (len=5) "71788",
    Contracts: ([]parsers.OptionContractCtlInfo) (len=2 cap=2) {
     (parsers.OptionContractCtlInfo) {
      CtrctId: (string) (len=15) "DXTRA1924600035",
      DealDate: (time.Time) 2019-09-03 00:00:00 +0000 UTC,
      ExpiryDate: (time.Time) 2020-09-01 00:00:00 +0000 UTC,
      BaseCurrencyCsCd: (string) (len=4) "0010",
      BaseCurrencyIsoCd: (string) (len=3) "CHF",
      SettlementDate: (time.Time) 2020-09-03 00:00:00 +0000 UTC,
      BuySettlementDate: (time.Time) 2020-09-03 00:00:00 +0000 UTC,
      SellSettlementDate: (time.Time) 2020-09-03 00:00:00 +0000 UTC,
      SettlementTpInd: (int) 5,
      BuySellInd: (string) (len=3) "Buy",
      OptionTpCd: (int) 4,
      ExticOptionTpCd: (string) (len=23) "TargetRedemptionForward",
      NtnlLeadCurrencyAmount: (decimal.Decimal) 5200000,
      NtnlLeadCurrencyIsoCd: (string) (len=3) "EUR",
      NtnlLeadCurrencyIsoDesc: (string) (len=4) "Euro",
      StrikePriceRate: (decimal.Decimal) 120.55,
      NtnlCounterCurrencyAmount: (decimal.Decimal) 626860000,
      NtnlCounterCurrencyISOCd: (string) (len=3) "JPY",
      NtnlCounterCurrencyISODesc: (string) (len=3) "Yen",
      PremiumAmount: (decimal.Decimal) 0,
      PremiumAccountIsoCd: (string) (len=3) "USD",
      PremiumAccountIsoDesc: (string) (len=10) "US dollars",
      PositionDate: (time.Time) 2020-02-20 00:00:00 +0000 UTC,
      RevaluationDate: (time.Time) 2020-09-03 00:00:00 +0000 UTC,
      PrtFlId: (string) (len=7) "71788-1",
      MtmValOrigCurrencyIsoCd: (string) (len=3) "EUR",
      MtmValOrigCurrencyAmount: (decimal.Decimal) 5342.981836,
      MtmValReportCurrencyISOCd: (string) (len=3) "USD",
      MtmValReportCurrencyAmount: (decimal.Decimal) 5786.44932842,
      MtmValueDate: (time.Time) 2020-02-18 00:00:00 +0000 UTC
     },
     (parsers.OptionContractCtlInfo) {
      CtrctId: (string) (len=15) "DXTRA1928900047",
      DealDate: (time.Time) 2019-10-16 00:00:00 +0000 UTC,
      ExpiryDate: (time.Time) 2020-10-14 00:00:00 +0000 UTC,
      BaseCurrencyCsCd: (string) (len=4) "0010",
      BaseCurrencyIsoCd: (string) (len=3) "CHF",
      SettlementDate: (time.Time) 2020-10-16 00:00:00 +0000 UTC,
      BuySettlementDate: (time.Time) 2020-10-16 00:00:00 +0000 UTC,
      SellSettlementDate: (time.Time) 2020-10-16 00:00:00 +0000 UTC,
      SettlementTpInd: (int) 5,
      BuySellInd: (string) (len=3) "Buy",
      OptionTpCd: (int) 4,
      ExticOptionTpCd: (string) (len=28) "PivotTargetRedemptionForward",
      NtnlLeadCurrencyAmount: (decimal.Decimal) 3120000,
      NtnlLeadCurrencyIsoCd: (string) (len=3) "NZD",
      NtnlLeadCurrencyIsoDesc: (string) (len=19) "New Zealand dollars",
      StrikePriceRate: (decimal.Decimal) 0.6618,
      NtnlCounterCurrencyAmount: (decimal.Decimal) 2064816,
      NtnlCounterCurrencyISOCd: (string) (len=3) "USD",
      NtnlCounterCurrencyISODesc: (string) (len=10) "US dollars",
      PremiumAmount: (decimal.Decimal) 0,
      PremiumAccountIsoCd: (string) (len=3) "NZD",
      PremiumAccountIsoDesc: (string) (len=19) "New Zealand dollars",
      PositionDate: (time.Time) 2020-02-20 00:00:00 +0000 UTC,
      RevaluationDate: (time.Time) 2020-10-16 00:00:00 +0000 UTC,
      PrtFlId: (string) (len=7) "71788-1",
      MtmValOrigCurrencyIsoCd: (string) (len=3) "NZD",
      MtmValOrigCurrencyAmount: (decimal.Decimal) -2104.313972,
      MtmValReportCurrencyISOCd: (string) (len=3) "USD",
      MtmValReportCurrencyAmount: (decimal.Decimal) -1349.28611859,
      MtmValueDate: (time.Time) 2020-02-18 00:00:00 +0000 UTC
     }
    }
   }
  }
 }
}

```


