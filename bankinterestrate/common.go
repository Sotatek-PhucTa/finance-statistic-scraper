package bankinterestrate

import (
	"fmt"
)

type BankGeneralInfo struct {
	Code string
	Name string
}

type BankRateInfo struct {
	InterestRate string
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
	GetPersonalInterestRate(url string) []InterestRate
	GetBusinessInterestRate(url string) []InterestRate
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
	var interestRates = bank.Handler.GetPersonalInterestRate(bank.RateInfo.InterestRate)
	bank.PersonalInterestRate = interestRates
}

func (bank *BankInfo) GetBusinessInterestRate() {
	var interestRates = bank.Handler.GetBusinessInterestRate(bank.RateInfo.InterestRate)
	bank.BusinessInterestRate = interestRates
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
