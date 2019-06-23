package currency

import (
	"testing"
)

func TestConvertToInternal(t *testing.T) {
	tests := []struct {
		name string
		c    Currency
		m    float64
		want int
	}{
		{"USD", USD, 123.45, 12345},
		{"IQD", IQD, 12.345, 12345},
		{"UYW", UYW, 1.2345, 12345},
		{"CLP", CLP, 12345.0, 12345},
		{"AAA", "AAA", 12345.0, 12345},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToInternal(tt.m, tt.c); got != tt.want {
				t.Errorf("wrong result %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertToExternal(t *testing.T) {
	tests := []struct {
		name string
		c    Currency
		m    int
		want float64
	}{
		{"USD", USD, 12345, 123.45},
		{"IQD", IQD, 12345, 12.345},
		{"UYW", UYW, 12345, 1.2345},
		{"CLP", CLP, 12345, 12345.0},
		{"AAA", "AAA", 12345, 12345.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToExternal(tt.m, tt.c); got != tt.want {
				t.Errorf("wrong result %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAtoCurrency(t *testing.T) {
	tests := []struct {
		name    string
		a       string
		want    Currency
		wantErr bool
	}{
		{"USD", "USD", USD, false},
		{"IQD", "IQD", IQD, false},
		{"UYW", "UYW", UYW, false},
		{"AAA", "AAA", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AtoCurrency(tt.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("unexpected error state = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && *got != tt.want {
				t.Errorf("wrong value %v, want %v", got, tt.want)
			}
		})
	}
}
