package main

import (
	"finance-statistic-scraper/bankinterestrate"
)

func main() {
	var agribankScraper = bankinterestrate.GetAgribankScraper()
	agribankScraper.GetBusinessInterestRate()
	agribankScraper.GetPersonalInterestRate()
	agribankScraper.SaveInterestRate(true)
	//bankinterestrate.GetVietinbankScraper().GetInterestRate().SaveInterestRate(true)
}
