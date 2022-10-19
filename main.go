package main

import (
	bankinterestrate "finance-statistic-scraper/bankinterestrate"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	bankinterestrate.GetInterestRate(c)
}
