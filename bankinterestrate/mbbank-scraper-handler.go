package bankinterestrate

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MbbankScraperHandler struct{}

func (m MbbankScraperHandler) GetPersonalInterestRate(requestInfo RateRequestInfo) []InterestRate {
	var interestRates []InterestRate
	resp, _ := http.Get(requestInfo.Url)
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	fmt.Println(res)
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
