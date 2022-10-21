package bankinterestrate

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

type VietcombankScraperHandler struct{}

func (v VietcombankScraperHandler) GetPersonalInterestRate(url string) []InterestRate {
	c := colly.NewCollector()
	var interestRates []InterestRate
	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String(), "to get personal interest rate")
	})

	c.OnHTML("table#danhsachlaisuat > tbody", func(laiSuatCnData *colly.HTMLElement) {
		fmt.Println("Lai suat CN")
		isLaiSuat := false
		laiSuatCnData.ForEach("tr", func(_ int, tr *colly.HTMLElement) {
			className := tr.Attr("class")
			duration, amount := strings.TrimSpace(tr.ChildText("td:nth-child(1)")), strings.TrimSpace(tr.ChildText("td:nth-child(2)"))
			if className == "imp" {
				if duration == "Tiền gửi có kỳ hạn" {
					isLaiSuat = true
				} else {
					isLaiSuat = false
				}
			}
			if isLaiSuat && className == "" {
				interestRates = append(interestRates, InterestRate{
					RateType: PersonalRate,
					Duration: duration,
					Amount:   amount,
				})
			}
		})
	})

	c.OnError(func(response *colly.Response, err error) {
		fmt.Println(err)
	})

	c.OnResponse(func(response *colly.Response) {
		fmt.Printf("Visited %s with status code %d\n", response.Request.URL.String(), response.StatusCode)
	})

	c.Visit(url)
	return interestRates
}

func (v VietcombankScraperHandler) GetBusinessInterestRate(url string) []InterestRate {
	//TODO implement me
	panic("implement me")
}

var vietcombankScraperHandler *VietcombankScraperHandler

func init() {
	vietcombankScraperHandler = &VietcombankScraperHandler{}
}

func GetVietcombankScraperHandler() *VietcombankScraperHandler {
	return vietcombankScraperHandler
}
