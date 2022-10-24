package bankinterestrate

import (
	"fmt"
	"github.com/gocolly/colly"
)

type VietinbankScraperHandler struct{}

func (v *VietinbankScraperHandler) GetPersonalInterestRate(requestInfo RateRequestInfo) []InterestRate {
	var interestRates []InterestRate
	c := colly.NewCollector()
	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})
	c.OnHTML("#articles > table > tbody", func(laiSuatData *colly.HTMLElement) {
		laiSuatData.ForEach("tr", func(i int, tr *colly.HTMLElement) {
			if i >= 3 {
				interestRates = append(interestRates, InterestRate{
					BankCode: vietinbankScraper.GeneralInfo.Code,
					RateType: PersonalRate,
					Duration: tr.ChildText("td:nth-child(1)"),
					Amount:   tr.ChildText("td:nth-child(2)"),
				})
			}
		})
	})
	c.OnResponse(func(response *colly.Response) {
		fmt.Printf("Visited %s with status code %d\n", response.Request.URL.String(), response.StatusCode)
	})
	c.Visit(requestInfo.Url)
	return interestRates
}

func (v *VietinbankScraperHandler) GetBusinessInterestRate(requestInfo RateRequestInfo) []InterestRate {
	var interestRates []InterestRate
	c := colly.NewCollector()
	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})
	c.OnHTML("#articles > table > tbody", func(laiSuatData *colly.HTMLElement) {
		laiSuatData.ForEach("tr", func(i int, tr *colly.HTMLElement) {
			if i >= 3 {
				interestRates = append(interestRates, InterestRate{
					BankCode: vietinbankScraper.GeneralInfo.Code,
					RateType: BusinessRate,
					Duration: tr.ChildText("td:nth-child(1)"),
					Amount:   tr.ChildText("td:nth-child(5)"),
				})
			}
		})
	})
	c.OnResponse(func(response *colly.Response) {
		fmt.Printf("Visited %s with status code %d\n", response.Request.URL.String(), response.StatusCode)
	})
	c.Visit(requestInfo.Url)
	return interestRates
}

var vietinbankScraperHandler *VietinbankScraperHandler

func init() {
	vietinbankScraperHandler = &VietinbankScraperHandler{}
}

func GetVietinbankScraperHandler() *VietinbankScraperHandler {
	return vietinbankScraperHandler
}
