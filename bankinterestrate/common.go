package bankinterestrate

type BankGeneralInfo struct {
	Code string
	Name string
}

type BankRateInfo struct {
	InterestRate string
}

type BankInfo struct {
	generalInfo BankGeneralInfo
	rateInfo    BankRateInfo
}

type RateType int64

const (
	PersonalRate RateType = iota
	BusinessRate
)

type InterestRate struct {
	rateType RateType
	duration string
	amount   uint32
}

type BankScraper interface {
	GetInterestRate() []InterestRate
	SaveInterestRate(rates []InterestRate)
}
