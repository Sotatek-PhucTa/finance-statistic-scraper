package bankinterestrate

var vietcombankScraper *BankInfo

func init() {
	vietcombankScraper = &BankInfo{
		GeneralInfo: BankGeneralInfo{
			Code: "VCB",
			Name: "VIETCOMBANK",
		},
		RateInfo: BankRateInfo{
			PersonalInterestRate: RateRequestInfo{Url: "https://portal.vietcombank.com.vn/UserControls/TVPortal.TyGia/pListLaiSuat.aspx?CusstomType=1&BacrhID=1&InrateType=&isEn=False&numAfter=2"},
			BusinessInterestRate: RateRequestInfo{Url: "https://portal.vietcombank.com.vn/UserControls/TVPortal.TyGia/pListLaiSuat.aspx?CusstomType=2&BacrhID=1&InrateType=&isEn=False&numAfter=2"},
		},
	}
	var vietcombankScraperHandler = GetVietcombankScraperHandler()
	vietcombankScraper.SetScraperHandler(vietcombankScraperHandler)
}

func GetVietcombankScraper() *BankInfo {
	return vietcombankScraper
}
