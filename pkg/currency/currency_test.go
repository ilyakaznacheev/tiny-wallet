// Package currency contains ISO 4217 currency codes and currency names
package currency

import (
	"fmt"
	"testing"
)

func TestCurrencyString(t *testing.T) {
	tests := []struct {
		name string
		c    Currency
		want string
	}{
		{"AFN", AFN, "Afghani"},
		{"BWP", BWP, "Pula"},
		{"NOK", NOK, "Norwegian Krone"},
		{"UYI", UYI, "Uruguay Peso en Unidades Indexadas (UI)"},
		{"000", 000, "non-ISO 4216 currency"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fmt.Sprint(tt.c); got != tt.want {
				t.Errorf("Wrong currency name %s, want %s", got, tt.want)
			}
		})
	}
}

func TestCurrencyFormatAmount(t *testing.T) {
	amount := 123456789
	tests := []struct {
		name string
		c    Currency
		want string
	}{
		{"AFN", AFN, "1234567.89"},
		{"BHD", BHD, "123456.789"},
		{"CLP", CLP, "123456789"},
		{"UYW", UYW, "12345.6789"},
		{"JOD", JOD, "123456.789"},
		{"SDG", SDG, "1234567.89"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.FormatAmount(amount); got != tt.want {
				t.Errorf("wrong formatting %v, want %v", got, tt.want)
			}
		})
	}
}
