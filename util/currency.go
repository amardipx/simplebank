package util

const (
	USD = "USD"
	EUR = "EUR"
	// add more currency if needed
)

// returns true if the input currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR:
		return true
	}
	return false
}
