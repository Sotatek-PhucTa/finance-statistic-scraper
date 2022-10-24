package bankinterestrate

var mbbankScraper *BankInfo

func init() {
	mbbankScraper = &BankInfo{
		GeneralInfo: BankGeneralInfo{
			Code: "MBB",
			Name: "MBB",
		},
		RateInfo: BankRateInfo{
			PersonalInterestRate: "https://www.mbbank.com.vn/api/GetlistFee",
			BusinessInterestRate: "https://www.mbbank.com.vn/api/GetlistFee",
		},
	}
	var mbbankScraperHandler = GetMbbankScraperHandler()
	mbbankScraper.SetScraperHandler(mbbankScraperHandler)
}

func GetMbbankScraper() *BankInfo {
	return mbbankScraper
}
