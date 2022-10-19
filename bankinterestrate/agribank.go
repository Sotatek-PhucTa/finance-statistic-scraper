package bankinterestrate

import (
	"fmt"
	"github.com/gocolly/colly"
)

func GetInterestRate(c *colly.Collector) {
	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	c.OnHTML("div#laiSuatCn > table > tbody", func(laiSuatCnData *colly.HTMLElement) {
		fmt.Println("Lai suat CN")
		laiSuatCnData.ForEach("tr", func(_ int, tr *colly.HTMLElement) {
			fmt.Println(tr.ChildText("td:nth-child(1)"), " ", tr.ChildText("td:nth-child(2)"))
		})
	})

	c.OnHTML("div#laiSuatDn > table > tbody", func(laiSuatCnData *colly.HTMLElement) {
		fmt.Println("Lai suat DN")
		laiSuatCnData.ForEach("tr", func(_ int, tr *colly.HTMLElement) {
			fmt.Println(tr.ChildText("td:nth-child(1)"), " ", tr.ChildText("td:nth-child(2)"))
		})
	})

	c.OnResponse(func(response *colly.Response) {
		fmt.Printf("Visited %s with status code %d\n", response.Request.URL.String(), response.StatusCode)
	})

	c.Visit("https://www.agribank.com.vn/vn/lai-suat")
}
