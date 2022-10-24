package bankinterestrate

var bidvScraper *BankInfo

func init() {
	bidvScraper = &BankInfo{
		GeneralInfo: BankGeneralInfo{
			Code: "BIDV",
			Name: "BIDV",
		},
		RateInfo: BankRateInfo{
			PersonalInterestRate: RateRequestInfo{Url: "https://www.bidv.com.vn/ServicesBIDV/InterestDetailServlet"},
			BusinessInterestRate: RateRequestInfo{Url: "https://www.bidv.com.vn/ServicesBIDV/InterestDetailServlet"},
		},
	}
	var bidvScraperHandler = GetBidvScraperHandler()
	bidvScraper.SetScraperHandler(bidvScraperHandler)
}

func GetBidvScraper() *BankInfo {
	return bidvScraper
}
