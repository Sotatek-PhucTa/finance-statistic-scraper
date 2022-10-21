package bankinterestrate

var agribankScraper BankInfo

func init() {
	agribankScraper = BankInfo{
		GeneralInfo: BankGeneralInfo{
			Code: "AGRIBANK",
			Name: "AGRIBANK",
		},
		RateInfo: BankRateInfo{
			InterestRate: "https://www.agribank.com.vn/vn/lai-suat",
		},
	}

	var agribankScraperHandler = GetAgribankScraperHandler()
	agribankScraper.SetScraperHandler(agribankScraperHandler)
}

func GetAgribankScraper() BankInfo {
	return agribankScraper
}
