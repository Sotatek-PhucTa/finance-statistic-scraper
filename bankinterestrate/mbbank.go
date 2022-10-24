package bankinterestrate

var mbbankScraper *BankInfo

func init() {
	mbbankScraper = &BankInfo{
		GeneralInfo: BankGeneralInfo{
			Code: "MBB",
			Name: "MBB",
		},
		RateInfo: BankRateInfo{
			PersonalInterestRate: RateRequestInfo{
				Url: "https://www.mbbank.com.vn/api/GetlistFee",
				Header: map[string]string{
					"MB-XSRF-Token-FormOnline": "OMsZRLIvGAO4CUiMJyf6_SRIUc3d89Cvth_YpqGAvi9zbWrQe_Suzi7Lz4-criJu_0xlYexdUlaDssz5YbSmxlckJQ-R0jCsmRdJBUMVUWE1",
					"Cookie":                   "ASP.NET_SessionId=ronddohi4yfjztb0bl22nnwe; __RequestVerificationToken=3ShijBmlIyDJtLJiwvfs7JIijnooGoobY1TT4WEF3aLedQjI2GsFDYIlbWJFbQ0O2YSCZCIc03ix6KKpeVIZ6NPSOB2E7iNAHsN4Z1mbsJM1; BIGipServerWeb_mbbank_pool_443=404488458.47873.0000; alias_current=",
				},
			},
			BusinessInterestRate: RateRequestInfo{Url: "https://www.mbbank.com.vn/api/GetlistFee"},
		},
	}
	var mbbankScraperHandler = GetMbbankScraperHandler()
	mbbankScraper.SetScraperHandler(mbbankScraperHandler)
}

func GetMbbankScraper() *BankInfo {
	return mbbankScraper
}
