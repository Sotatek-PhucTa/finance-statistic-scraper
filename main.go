package main

import (
	"finance-statistic-scraper/bankinterestrate"
)

func main() {
	bankinterestrate.GetAgribankScraper().GetInterestRate().SaveInterestRate(true)
	bankinterestrate.GetVietinbankScraper().GetInterestRate().SaveInterestRate(true)
}
