package bankinterestrate

import "github.com/gocolly/colly"

type BankGeneralInfo struct {
	Code string
	Name string
}

type BankRateInfo struct {
	InterestRate string
}

type BankInfo struct {
	GeneralInfo BankGeneralInfo
	RateInfo    BankRateInfo
}

type RateType int64

const (
	PersonalRate RateType = iota
	BusinessRate
)

type InterestRate struct {
	BankCode string
	RateType RateType
	Duration string
	Amount   string
}

type BankScraper interface {
	SetCollector(c *colly.Collector) BankScraper
	GetInterestRate() BankScraper
	SaveInterestRate(consoleOutput bool)
}

type BankScraperHandler interface {
	SetCollector(c *colly.Collector)
	GetPersonalInterestRate(url string) []InterestRate
	GetBusinessInterestRate(url string) []InterestRate
}
