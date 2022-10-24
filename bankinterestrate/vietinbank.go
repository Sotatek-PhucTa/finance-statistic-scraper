package bankinterestrate

var vietinbankScraper *BankInfo

func init() {
	vietinbankScraper = &BankInfo{
		GeneralInfo: BankGeneralInfo{
			Code: "VIETINBANK",
			Name: "VIETINBANK",
		},
		RateInfo: BankRateInfo{
			PersonalInterestRate: RateRequestInfo{Url: "https://www.vietinbank.vn/web/home/vn/lai-suat"},
			BusinessInterestRate: RateRequestInfo{Url: "https://www.vietinbank.vn/web/home/vn/lai-suat"},
		},
	}
	var vietinbankScraperHandler = GetVietinbankScraperHandler()
	vietinbankScraper.SetScraperHandler(vietinbankScraperHandler)
}

func GetVietinbankScraper() *BankInfo {
	return vietinbankScraper
}
