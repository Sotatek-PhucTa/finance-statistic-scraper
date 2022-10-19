package main

import (
	"finance-statistic-scraper/bankinterestrate"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	bankinterestrate.GetAgribankScraper().SetCollector(c).GetInterestRate()
}
