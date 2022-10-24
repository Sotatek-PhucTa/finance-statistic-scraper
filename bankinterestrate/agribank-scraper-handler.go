package bankinterestrate

import (
	"fmt"
	"github.com/gocolly/colly"
)

type AgribankScraperHandler struct{}

var agribankScraperHandler *AgribankScraperHandler

func init() {
	agribankScraperHandler = &AgribankScraperHandler{}
}

func GetAgribankScraperHandler() *AgribankScraperHandler {
	return agribankScraperHandler
}

func (handler *AgribankScraperHandler) GetPersonalInterestRate(requestInfo RateRequestInfo) []InterestRate {
	fmt.Println("Get Personal rate")
	c := colly.NewCollector()
	var interestRates []InterestRate
	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String(), "to get personal interest rate")
	})

	c.OnHTML("div#laiSuatCn > table > tbody", func(laiSuatCnData *colly.HTMLElement) {
		fmt.Println("Lai suat CN")
		laiSuatCnData.ForEach("tr", func(_ int, tr *colly.HTMLElement) {
			interestRates = append(interestRates, InterestRate{
				RateType: PersonalRate,
				Duration: tr.ChildText("td:nth-child(1)"),
				Amount:   tr.ChildText("td:nth-child(2)")})
		})
	})

	c.OnResponse(func(response *colly.Response) {
		fmt.Printf("Visited %s with status code %d\n", response.Request.URL.String(), response.StatusCode)
	})

	c.Visit(requestInfo.Url)
	return interestRates
}

func (handler *AgribankScraperHandler) GetBusinessInterestRate(requestInfo RateRequestInfo) []InterestRate {
	fmt.Println("Get Business rate")
	c := colly.NewCollector()
	var interestRates []InterestRate
	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String(), "to get business interest rate")
	})

	c.OnHTML("div#laiSuatDn > table > tbody", func(laiSuatCnData *colly.HTMLElement) {
		fmt.Println("Lai suat DN")
		laiSuatCnData.ForEach("tr", func(_ int, tr *colly.HTMLElement) {
			interestRates = append(interestRates, InterestRate{
				RateType: BusinessRate,
				Duration: tr.ChildText("td:nth-child(1)"),
				Amount:   tr.ChildText("td:nth-child(2)")})
		})
	})

	c.OnResponse(func(response *colly.Response) {
		fmt.Printf("Visited %s with status code %d\n", response.Request.URL.String(), response.StatusCode)
	})

	c.Visit(requestInfo.Url)
	return interestRates
}
