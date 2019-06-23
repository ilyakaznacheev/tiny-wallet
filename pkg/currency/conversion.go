package currency

import (
	"fmt"
	"math"
)

// ConvertToInternal converts external floating point currency amount to internal integer in the lowest unit of the currency
// E.g. USD (2): 15.25 -> 1525
func ConvertToInternal(m float64, c Currency) int {
	return int(m * math.Pow10(c.Decimals()))
}

// ConvertToExternal converts internal integer amount in the lowest unit of the currency  to external floating point format
// E.g. USD (2): 1525 -> 15.25
func ConvertToExternal(m int, c Currency) float64 {
	return float64(m) / math.Pow10(c.Decimals())
}

// AtoCurrency converts string to ISO 4216 currency
func AtoCurrency(a string) (*Currency, error) {
	c := Currency(a)
	if _, ok := currencyProperties[c]; ok {
		return &c, nil
	}
	return nil, fmt.Errorf("non-ISO 4216 currency (%s)", a)
}
