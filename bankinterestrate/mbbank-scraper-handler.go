package bankinterestrate

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MbbankScraperHandler struct{}

func (m MbbankScraperHandler) GetPersonalInterestRate(url string) []InterestRate {
	var interestRates []InterestRate
	resp, _ := http.Get(url)
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	fmt.Println(res)
	return interestRates
}

func (m MbbankScraperHandler) GetBusinessInterestRate(url string) []InterestRate {
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
