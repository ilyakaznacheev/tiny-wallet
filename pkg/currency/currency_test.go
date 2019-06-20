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

func TestAtoc(t *testing.T) {
	tests := []struct {
		name    string
		code    string
		want    Currency
		wantErr bool
	}{
		{"USD", "USD", USD, false},
		{"CLP", "CLP", CLP, false},
		{"JOD", "JOD", JOD, false},
		{"AAA", "AAA", 000, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Atoc(tt.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("wrong error state %v, want %v", err, tt.wantErr)
				return
			}
			if err == nil && *got != tt.want {
				t.Errorf("wrong code %d, want %d", *got, tt.want)
			}

		})
	}
}

func TestCtoa(t *testing.T) {
	tests := []struct {
		name    string
		c       Currency
		want    string
		wantErr bool
	}{
		{"USD", USD, "USD", false},
		{"CLP", CLP, "CLP", false},
		{"JOD", JOD, "JOD", false},
		{"AAA", 000, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Ctoa(tt.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("wrong error state %v, want %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("wrong code %s, want %s", got, tt.want)
			}
		})
	}
}
