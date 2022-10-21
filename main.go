package main

import (
	"finance-statistic-scraper/bankinterestrate"
)

func main() {
	var agribankScraper = bankinterestrate.GetAgribankScraper()
	agribankScraper.GetBusinessInterestRate()
	agribankScraper.GetPersonalInterestRate()
	agribankScraper.SaveInterestRate(true)
	var vietinbankScraper = bankinterestrate.GetVietinbankScraper()
	vietinbankScraper.GetBusinessInterestRate()
	vietinbankScraper.GetPersonalInterestRate()
	vietinbankScraper.SaveInterestRate(true)
}
