package main

import (
	"finance-statistic-scraper/bankinterestrate"
)

func main() {
	//var agribankScraper = bankinterestrate.GetAgribankScraper()
	//agribankScraper.GetBusinessInterestRate()
	//agribankScraper.GetPersonalInterestRate()
	//agribankScraper.SaveInterestRate(true)
	//var vietinbankScraper = bankinterestrate.GetVietinbankScraper()
	//vietinbankScraper.GetBusinessInterestRate()
	//vietinbankScraper.GetPersonalInterestRate()
	//vietinbankScraper.SaveInterestRate(true)
	//var vietcombankScraper = bankinterestrate.GetVietcombankScraper()
	//vietcombankScraper.GetPersonalInterestRate()
	//vietcombankScraper.GetBusinessInterestRate()
	//vietcombankScraper.SaveInterestRate(true)
	//var bidvScraper = bankinterestrate.GetBidvScraper()
	//fmt.Println("bidv scraper", bidvScraper)
	//bidvScraper.GetPersonalInterestRate()
	//bidvScraper.SaveInterestRate(true)
	var mbbankScraper = bankinterestrate.GetMbbankScraper()
	mbbankScraper.GetPersonalInterestRate()
	mbbankScraper.SaveInterestRate(true)
}
