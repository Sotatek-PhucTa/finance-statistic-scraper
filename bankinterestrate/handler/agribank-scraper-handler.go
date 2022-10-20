package handler

import (
	"finance-statistic-scraper/bankinterestrate"
	"fmt"
	"github.com/gocolly/colly"
)

type AgribankScraperHandler struct {
	c *colly.Collector
}

var agribankScraperHandler AgribankScraperHandler

func init() {
	var c = colly.NewCollector()
	agribankScraperHandler = AgribankScraperHandler{c: c}
}

func GetAgribankScraperHandler() AgribankScraperHandler {
	return agribankScraperHandler
}

func (handler *AgribankScraperHandler) SetCollector(c *colly.Collector) {
	handler.c = c
}

func (handler *AgribankScraperHandler) GetPersonalInterestRate(url string) []bankinterestrate.InterestRate {
	var interestRates []bankinterestrate.InterestRate
	handler.c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String(), "to get personal interest rate")
	})

	handler.c.OnHTML("div#laiSuatCn > table > tbody", func(laiSuatCnData *colly.HTMLElement) {
		fmt.Println("Lai suat CN")
		laiSuatCnData.ForEach("tr", func(_ int, tr *colly.HTMLElement) {
			interestRates = append(interestRates, bankinterestrate.InterestRate{
				RateType: bankinterestrate.PersonalRate,
				Duration: tr.ChildText("td:nth-child(1)"),
				Amount:   tr.ChildText("td:nth-child(2)")})
		})
	})

	handler.c.OnResponse(func(response *colly.Response) {
		fmt.Printf("Visited %s with status code %d\n", response.Request.URL.String(), response.StatusCode)
	})

	handler.c.Visit(url)
	return interestRates
}

func (handler *AgribankScraperHandler) GetBusinessInterestRate(url string) []bankinterestrate.InterestRate {
	var interestRates []bankinterestrate.InterestRate
	handler.c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String(), "to get business interest rate")
	})

	handler.c.OnHTML("div#laiSuatDn > table > tbody", func(laiSuatCnData *colly.HTMLElement) {
		fmt.Println("Lai suat CN")
		laiSuatCnData.ForEach("tr", func(_ int, tr *colly.HTMLElement) {
			interestRates = append(interestRates, bankinterestrate.InterestRate{
				RateType: bankinterestrate.BusinessRate,
				Duration: tr.ChildText("td:nth-child(1)"),
				Amount:   tr.ChildText("td:nth-child(2)")})
		})
	})

	handler.c.OnResponse(func(response *colly.Response) {
		fmt.Printf("Visited %s with status code %d\n", response.Request.URL.String(), response.StatusCode)
	})

	handler.c.Visit(url)
	return interestRates
}
