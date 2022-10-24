package bankinterestrate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type BidvScraperHandler struct{}

func (b BidvScraperHandler) GetPersonalInterestRate(requestInfo RateRequestInfo) []InterestRate {
	var interestRates []InterestRate
	resp, _ := http.Post(requestInfo.Url, "application/json", new(bytes.Buffer))
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	for k, v := range res {
		//Query data in "hanoi" field only
		if k == "hanoi" {
			for k1, v1 := range v.(map[string]interface{}) {
				// Query data in "data" field only
				if k1 == "data" {
					//Cast transformedV1 as array of data
					transformedV1 := v1.([]interface{})
					for _, vTransformedV1 := range transformedV1 {
						transformedV2 := vTransformedV1.(map[string]interface{})
						fmt.Println(transformedV2)
						duration := transformedV2["VND"].(string)
						amount := transformedV2["title_vi"].(string)
						interestRates = append(interestRates, InterestRate{RateType: PersonalRate, Duration: duration, Amount: amount})
					}
				}
			}
		}
	}
	return interestRates
}

func (b BidvScraperHandler) GetBusinessInterestRate(requestInfo RateRequestInfo) []InterestRate {
	//TODO implement me
	panic("Need to extract information from pdf, comeback later")
}

var bidvScraperHandler *BidvScraperHandler

func init() {
	bidvScraperHandler = &BidvScraperHandler{}
}

func GetBidvScraperHandler() *BidvScraperHandler {
	return bidvScraperHandler
}
