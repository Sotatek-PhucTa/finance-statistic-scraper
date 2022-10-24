package bankinterestrate

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type MbbankScraperHandler struct{}

func (m MbbankScraperHandler) GetPersonalInterestRate(requestInfo RateRequestInfo) []InterestRate {
	var interestRates []InterestRate
	client := resty.New()
	resp, _ := client.R().SetHeaders(requestInfo.Header).Get(requestInfo.Url)
	fmt.Println(resp)
	//var res map[string]interface{}
	//json.NewDecoder(resp.Body()).Decode(&res)
	//fmt.Println(res)
	return interestRates
}

func (m MbbankScraperHandler) GetBusinessInterestRate(requestInfo RateRequestInfo) []InterestRate {
	//TODO implement me
	panic("implement me")
}

var mbbankScraperHandler *MbbankScraperHandler

func init() {
	mbbankScraperHandler = &MbbankScraperHandler{}
}

func GetMbbankScraperHandler() *MbbankScraperHandler {
	return mbbankScraperHandler
}
