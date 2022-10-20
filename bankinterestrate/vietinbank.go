package bankinterestrate

import (
	"fmt"
	"github.com/gocolly/colly"
)

type VietinbankScraper struct {
	VietinbankInfo BankInfo
	c              *colly.Collector
	InterestRate   []InterestRate
}

var vietinbankScraper VietinbankScraper

func init() {
	var vietinbankInfo = BankInfo{
		GeneralInfo: BankGeneralInfo{
			Code: "VIETINBANK",
			Name: "VIETINBANK",
		},
		RateInfo: BankRateInfo{
			InterestRate: "https://www.vietinbank.vn/web/home/vn/lai-suat",
		},
	}
	vietinbankScraper = VietinbankScraper{VietinbankInfo: vietinbankInfo}
}

func GetVietinbankScraper() VietinbankScraper {
	return vietinbankScraper
}

func (vietinbankScraper VietinbankScraper) SetCollector(c *colly.Collector) VietinbankScraper {
	vietinbankScraper.c = c
	return vietinbankScraper
}

func (vietinbankScraper VietinbankScraper) GetInterestRate() VietinbankScraper {
	var interestRates []InterestRate
	vietinbankScraper.c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})
	vietinbankScraper.c.OnHTML("#articles > table > tbody", func(laiSuatData *colly.HTMLElement) {
		laiSuatData.ForEach("tr", func(i int, tr *colly.HTMLElement) {
			if i >= 3 {
				interestRates = append(interestRates, InterestRate{
					BankCode: vietinbankScraper.VietinbankInfo.GeneralInfo.Code,
					RateType: PersonalRate,
					Duration: tr.ChildText("td:nth-child(1)"),
					Amount:   tr.ChildText("td:nth-child(2)"),
				})
				interestRates = append(interestRates, InterestRate{
					BankCode: vietinbankScraper.VietinbankInfo.GeneralInfo.Code,
					RateType: BusinessRate,
					Duration: tr.ChildText("td:nth-child(1)"),
					Amount:   tr.ChildText("td:nth-child(5)"),
				})
			}
		})
	})
	vietinbankScraper.c.OnResponse(func(response *colly.Response) {
		fmt.Printf("Visited %s with status code %d\n", response.Request.URL.String(), response.StatusCode)
	})
	vietinbankScraper.c.OnError(func(response *colly.Response, err error) {
		fmt.Println(err)
	})
	vietinbankScraper.c.Visit(vietinbankScraper.VietinbankInfo.RateInfo.InterestRate)
	vietinbankScraper.InterestRate = interestRates
	return vietinbankScraper
}

func (vietinbankScraper VietinbankScraper) SaveInterestRate(consoleOutput bool) {
	if consoleOutput {
		for _, rate := range vietinbankScraper.InterestRate {
			fmt.Println(rate)
		}
	}
}
