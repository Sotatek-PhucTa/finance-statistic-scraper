package bankinterestrate

import (
	"fmt"
)

type BankGeneralInfo struct {
	Code string
	Name string
}

type RateRequestInfo struct {
	Url    string
	Header map[string]string
}

type BankRateInfo struct {
	PersonalInterestRate RateRequestInfo
	BusinessInterestRate RateRequestInfo
}

type RateType int64

const (
	PersonalRate RateType = iota
	BusinessRate
)

type InterestRate struct {
	BankCode string
	RateType RateType
	Duration string
	Amount   string
}

type BankScraperHandler interface {
	GetPersonalInterestRate(requestInfo RateRequestInfo) []InterestRate
	GetBusinessInterestRate(requestInfo RateRequestInfo) []InterestRate
}

type BankInfo struct {
	GeneralInfo          BankGeneralInfo
	RateInfo             BankRateInfo
	Handler              BankScraperHandler
	PersonalInterestRate []InterestRate
	BusinessInterestRate []InterestRate
}

type BankScraper interface {
	SetScraperHandler(handler *BankScraperHandler)
	GetPersonalInterestRate()
	GetBusinessInterestRate()
	SaveInterestRate(consoleOutput bool)
}

func (bank *BankInfo) SetScraperHandler(handler BankScraperHandler) {
	bank.Handler = handler
}

func (bank *BankInfo) GetPersonalInterestRate() {
	var interestRates = bank.Handler.GetPersonalInterestRate(bank.RateInfo.PersonalInterestRate)
	bank.PersonalInterestRate = standardizeInterestRate(interestRates, bank)
}

func (bank *BankInfo) GetBusinessInterestRate() {
	var interestRates = bank.Handler.GetBusinessInterestRate(bank.RateInfo.BusinessInterestRate)
	bank.BusinessInterestRate = standardizeInterestRate(interestRates, bank)
}

func (bank *BankInfo) SaveInterestRate(consoleOutput bool) {
	if consoleOutput {
		for _, rate := range bank.PersonalInterestRate {
			fmt.Println(rate)
		}
		for _, rate := range bank.BusinessInterestRate {
			fmt.Println(rate)
		}
	}
}

func standardizeInterestRate(rates []InterestRate, bankInfo *BankInfo) []InterestRate {
	var newRates []InterestRate
	for _, rate := range rates {
		rate.BankCode = bankInfo.GeneralInfo.Code
		newRates = append(newRates, rate)
	}
	return newRates
}
