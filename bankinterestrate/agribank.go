package bankinterestrate

import (
	"fmt"
	"github.com/gocolly/colly"
)

type AgribankScraper struct {
	AgribankInfo BankInfo
	c            *colly.Collector
}

var agribankScraper *AgribankScraper

func init() {
	var agribankInfo = BankInfo{
		GeneralInfo: BankGeneralInfo{
			Code: "AGRIBANK",
			Name: "AGRIBANK",
		},
		RateInfo: BankRateInfo{
			InterestRate: "https://www.agribank.com.vn/vn/lai-suat",
		},
	}
	agribankScraper = &AgribankScraper{AgribankInfo: agribankInfo}
}

func GetAgribankScraper() *AgribankScraper {
	return agribankScraper
}

func (agribankScraper AgribankScraper) SetCollector(c *colly.Collector) AgribankScraper {
	agribankScraper.c = c
	return agribankScraper
}

func (agribankScraper AgribankScraper) GetInterestRate() []InterestRate {
	var interestRates []InterestRate
	agribankScraper.c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	agribankScraper.c.OnHTML("div#laiSuatCn > table > tbody", func(laiSuatCnData *colly.HTMLElement) {
		fmt.Println("Lai suat CN")
		laiSuatCnData.ForEach("tr", func(_ int, tr *colly.HTMLElement) {
			interestRates = append(interestRates, InterestRate{
				RateType: PersonalRate,
				Duration: tr.ChildText("td:nth-child(1)"),
				Amount:   tr.ChildText("td:nth-child(2)")})
		})
	})

	agribankScraper.c.OnHTML("div#laiSuatDn > table > tbody", func(laiSuatCnData *colly.HTMLElement) {
		fmt.Println("Lai suat DN")
		laiSuatCnData.ForEach("tr", func(_ int, tr *colly.HTMLElement) {
			interestRates = append(interestRates, InterestRate{
				RateType: PersonalRate,
				Duration: tr.ChildText("td:nth-child(1)"),
				Amount:   tr.ChildText("td:nth-child(2)")})
		})
	})

	agribankScraper.c.OnResponse(func(response *colly.Response) {
		fmt.Printf("Visited %s with status code %d\n", response.Request.URL.String(), response.StatusCode)
	})

	agribankScraper.c.Visit(agribankScraper.AgribankInfo.RateInfo.InterestRate)
	return interestRates
}
