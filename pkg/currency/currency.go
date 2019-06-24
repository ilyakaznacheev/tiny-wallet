// Package currency contains ISO 4217 currency codes and currency names.
//
// Each currency contain it's name and decimal numbers. Since each currency have different decimal numbers, the database and internal application logic processes currency amounts as integers in smallest currency unit.
//
// The package allows to process currency conversions from external (float) format into internal (integer) and vice versa.
//
// For more information about ISO 4217 currency codes see https://www.iso.org/iso-4217-currency-codes.html
package currency

import (
	"fmt"
	"math"
	"strconv"
)

// Currency type for a currency ISO 4217 code
type Currency string

// String returns a name of currency
func (c Currency) String() string {
	if p, ok := currencyProperties[c]; ok {
		return p.Name
	}
	return fmt.Sprintf("non-ISO 4216 currency (%s)", string(c))
}

// FormatAmount returns an integer amount formatted depending on the number of decimal places of the currency
func (c Currency) FormatAmount(raw int) string {
	p := currencyProperties[c]
	b := float64(raw) / math.Pow10(int(p.Decimals))
	return strconv.FormatFloat(b, 'f', -1, 64)
}

// Decimals returns a number of decimal places of a currency
func (c Currency) Decimals() int {
	if p, ok := currencyProperties[c]; ok {
		return int(p.Decimals)
	}
	return 0
}

const (
	AFN Currency = "AFN"
	AED Currency = "AED"
	ALL Currency = "ALL"
	AMD Currency = "AMD"
	ANG Currency = "ANG"
	AOA Currency = "AOA"
	ARS Currency = "ARS"
	AUD Currency = "AUD"
	AWG Currency = "AWG"
	AZN Currency = "AZN"
	BAM Currency = "BAM"
	BBD Currency = "BBD"
	BDT Currency = "BDT"
	BGN Currency = "BGN"
	BHD Currency = "BHD"
	BIF Currency = "BIF"
	BMD Currency = "BMD"
	BND Currency = "BND"
	BOB Currency = "BOB"
	BOV Currency = "BOV"
	BRL Currency = "BRL"
	BSD Currency = "BSD"
	BTN Currency = "BTN"
	BWP Currency = "BWP"
	BYN Currency = "BYN"
	BZD Currency = "BZD"
	CAD Currency = "CAD"
	CDF Currency = "CDF"
	CHE Currency = "CHE"
	CHF Currency = "CHF"
	CHW Currency = "CHW"
	CLF Currency = "CLF"
	CLP Currency = "CLP"
	CNY Currency = "CNY"
	COP Currency = "COP"
	COU Currency = "COU"
	CRC Currency = "CRC"
	CUC Currency = "CUC"
	CUP Currency = "CUP"
	CVE Currency = "CVE"
	CZK Currency = "CZK"
	DJF Currency = "DJF"
	DKK Currency = "DKK"
	DOP Currency = "DOP"
	DZD Currency = "DZD"
	EGP Currency = "EGP"
	ERN Currency = "ERN"
	ETB Currency = "ETB"
	EUR Currency = "EUR"
	FJD Currency = "FJD"
	FKP Currency = "FKP"
	GBP Currency = "GBP"
	GEL Currency = "GEL"
	GHS Currency = "GHS"
	GIP Currency = "GIP"
	GMD Currency = "GMD"
	GNF Currency = "GNF"
	GTQ Currency = "GTQ"
	GYD Currency = "GYD"
	HKD Currency = "HKD"
	HNL Currency = "HNL"
	HRK Currency = "HRK"
	HTG Currency = "HTG"
	HUF Currency = "HUF"
	IDR Currency = "IDR"
	ILS Currency = "ILS"
	INR Currency = "INR"
	IQD Currency = "IQD"
	IRR Currency = "IRR"
	ISK Currency = "ISK"
	JMD Currency = "JMD"
	JOD Currency = "JOD"
	JPY Currency = "JPY"
	KES Currency = "KES"
	KGS Currency = "KGS"
	KHR Currency = "KHR"
	KMF Currency = "KMF"
	KPW Currency = "KPW"
	KRW Currency = "KRW"
	KWD Currency = "KWD"
	KYD Currency = "KYD"
	KZT Currency = "KZT"
	LAK Currency = "LAK"
	LBP Currency = "LBP"
	LKR Currency = "LKR"
	LRD Currency = "LRD"
	LSL Currency = "LSL"
	LYD Currency = "LYD"
	MAD Currency = "MAD"
	MDL Currency = "MDL"
	MGA Currency = "MGA"
	MKD Currency = "MKD"
	MMK Currency = "MMK"
	MNT Currency = "MNT"
	MOP Currency = "MOP"
	MRU Currency = "MRU"
	MUR Currency = "MUR"
	MVR Currency = "MVR"
	MWK Currency = "MWK"
	MXN Currency = "MXN"
	MXV Currency = "MXV"
	MYR Currency = "MYR"
	MZN Currency = "MZN"
	NAD Currency = "NAD"
	NGN Currency = "NGN"
	NIO Currency = "NIO"
	NOK Currency = "NOK"
	NPR Currency = "NPR"
	NZD Currency = "NZD"
	OMR Currency = "OMR"
	PAB Currency = "PAB"
	PEN Currency = "PEN"
	PGK Currency = "PGK"
	PHP Currency = "PHP"
	PKR Currency = "PKR"
	PLN Currency = "PLN"
	PYG Currency = "PYG"
	QAR Currency = "QAR"
	RON Currency = "RON"
	RSD Currency = "RSD"
	RUB Currency = "RUB"
	RWF Currency = "RWF"
	SAR Currency = "SAR"
	SBD Currency = "SBD"
	SCR Currency = "SCR"
	SDG Currency = "SDG"
	SEK Currency = "SEK"
	SGD Currency = "SGD"
	SHP Currency = "SHP"
	SLL Currency = "SLL"
	SOS Currency = "SOS"
	SRD Currency = "SRD"
	SSP Currency = "SSP"
	STN Currency = "STN"
	SVC Currency = "SVC"
	SYP Currency = "SYP"
	SZL Currency = "SZL"
	THB Currency = "THB"
	TJS Currency = "TJS"
	TMT Currency = "TMT"
	TND Currency = "TND"
	TOP Currency = "TOP"
	TRY Currency = "TRY"
	TTD Currency = "TTD"
	TWD Currency = "TWD"
	TZS Currency = "TZS"
	UAH Currency = "UAH"
	UGX Currency = "UGX"
	USD Currency = "USD"
	USN Currency = "USN"
	UYI Currency = "UYI"
	UYU Currency = "UYU"
	UYW Currency = "UYW"
	UZS Currency = "UZS"
	VES Currency = "VES"
	VND Currency = "VND"
	VUV Currency = "VUV"
	WST Currency = "WST"
	XAF Currency = "XAF"
	XCD Currency = "XCD"
	XDR Currency = "XDR"
	XOF Currency = "XOF"
	XPF Currency = "XPF"
	XSU Currency = "XSU"
	XUA Currency = "XUA"
	YER Currency = "YER"
	ZAR Currency = "ZAR"
	ZMW Currency = "ZMW"
	ZWL Currency = "ZWL"
)

// property of ISO currency
type property struct {
	Code     string
	Name     string
	Decimals uint
}

// currencyProperties ISO currency property
var currencyProperties = map[Currency]property{
	AFN: {"AFN", "Afghani", 2},
	AED: {"AED", "UAE Dirham", 2},
	ALL: {"ALL", "Lek", 2},
	AMD: {"AMD", "Armenian Dram", 2},
	ANG: {"ANG", "Netherlands Antillean Guilder", 2},
	AOA: {"AOA", "Kwanza", 2},
	ARS: {"ARS", "Argentine Peso", 2},
	AUD: {"AUD", "Australian Dollar", 2},
	AWG: {"AWG", "Aruban Florin", 2},
	AZN: {"AZN", "Azerbaijan Manat", 2},
	BAM: {"BAM", "Convertible Mark", 2},
	BBD: {"BBD", "Barbados Dollar", 2},
	BDT: {"BDT", "Taka", 2},
	BGN: {"BGN", "Bulgarian Lev", 2},
	BHD: {"BHD", "Bahraini Dinar", 3},
	BIF: {"BIF", "Burundi Franc", 0},
	BMD: {"BMD", "Bermudian Dollar", 2},
	BND: {"BND", "Brunei Dollar", 2},
	BOB: {"BOB", "Boliviano", 2},
	BOV: {"BOV", "Mvdol", 2},
	BRL: {"BRL", "Brazilian Real", 2},
	BSD: {"BSD", "Bahamian Dollar", 2},
	BTN: {"BTN", "Ngultrum", 2},
	BWP: {"BWP", "Pula", 2},
	BYN: {"BYN", "Belarusian Ruble", 2},
	BZD: {"BZD", "Belize Dollar", 2},
	CAD: {"CAD", "Canadian Dollar", 2},
	CDF: {"CDF", "Congolese Franc", 2},
	CHE: {"CHE", "WIR Euro", 2},
	CHF: {"CHF", "Swiss Franc", 2},
	CHW: {"CHW", "WIR Franc", 2},
	CLF: {"CLF", "Unidad de Fomento", 4},
	CLP: {"CLP", "Chilean Peso", 0},
	CNY: {"CNY", "Yuan Renminbi", 2},
	COP: {"COP", "Colombian Peso", 2},
	COU: {"COU", "Unidad de Valor Real", 2},
	CRC: {"CRC", "Costa Rican Colon", 2},
	CUC: {"CUC", "Peso Convertible", 2},
	CUP: {"CUP", "Cuban Peso", 2},
	CVE: {"CVE", "Cabo Verde Escudo", 2},
	CZK: {"CZK", "Czech Koruna", 2},
	DJF: {"DJF", "Djibouti Franc", 0},
	DKK: {"DKK", "Danish Krone", 2},
	DOP: {"DOP", "Dominican Peso", 2},
	DZD: {"DZD", "Algerian Dinar", 2},
	EGP: {"EGP", "Egyptian Pound", 2},
	ERN: {"ERN", "Nakfa", 2},
	ETB: {"ETB", "Ethiopian Birr", 2},
	EUR: {"EUR", "Euro", 2},
	FJD: {"FJD", "Fiji Dollar", 2},
	FKP: {"FKP", "Falkland Islands Pound", 2},
	GBP: {"GBP", "Pound Sterling", 2},
	GEL: {"GEL", "Lari", 2},
	GHS: {"GHS", "Ghana Cedi", 2},
	GIP: {"GIP", "Gibraltar Pound", 2},
	GMD: {"GMD", "Dalasi", 2},
	GNF: {"GNF", "Guinean Franc", 0},
	GTQ: {"GTQ", "Quetzal", 2},
	GYD: {"GYD", "Guyana Dollar", 2},
	HKD: {"HKD", "Hong Kong Dollar", 2},
	HNL: {"HNL", "Lempira", 2},
	HRK: {"HRK", "Kuna", 2},
	HTG: {"HTG", "Gourde", 2},
	HUF: {"HUF", "Forint", 2},
	IDR: {"IDR", "Rupiah", 2},
	ILS: {"ILS", "New Israeli Sheqel", 2},
	INR: {"INR", "Indian Rupee", 2},
	IQD: {"IQD", "Iraqi Dinar", 3},
	IRR: {"IRR", "Iranian Rial", 2},
	ISK: {"ISK", "Iceland Krona", 0},
	JMD: {"JMD", "Jamaican Dollar", 2},
	JOD: {"JOD", "Jordanian Dinar", 3},
	JPY: {"JPY", "Yen", 0},
	KES: {"KES", "Kenyan Shilling", 2},
	KGS: {"KGS", "Som", 2},
	KHR: {"KHR", "Riel", 2},
	KMF: {"KMF", "Comorian Franc ", 0},
	KPW: {"KPW", "North Korean Won", 2},
	KRW: {"KRW", "Won", 0},
	KWD: {"KWD", "Kuwaiti Dinar", 3},
	KYD: {"KYD", "Cayman Islands Dollar", 2},
	KZT: {"KZT", "Tenge", 2},
	LAK: {"LAK", "Lao Kip", 2},
	LBP: {"LBP", "Lebanese Pound", 2},
	LKR: {"LKR", "Sri Lanka Rupee", 2},
	LRD: {"LRD", "Liberian Dollar", 2},
	LSL: {"LSL", "Loti", 2},
	LYD: {"LYD", "Libyan Dinar", 3},
	MAD: {"MAD", "Moroccan Dirham", 2},
	MDL: {"MDL", "Moldovan Leu", 2},
	MGA: {"MGA", "Malagasy Ariary", 2},
	MKD: {"MKD", "Denar", 2},
	MMK: {"MMK", "Kyat", 2},
	MNT: {"MNT", "Tugrik", 2},
	MOP: {"MOP", "Pataca", 2},
	MRU: {"MRU", "Ouguiya", 2},
	MUR: {"MUR", "Mauritius Rupee", 2},
	MVR: {"MVR", "Rufiyaa", 2},
	MWK: {"MWK", "Malawi Kwacha", 2},
	MXN: {"MXN", "Mexican Peso", 2},
	MXV: {"MXV", "Mexican Unidad de Inversion (UDI)", 2},
	MYR: {"MYR", "Malaysian Ringgit", 2},
	MZN: {"MZN", "Mozambique Metical", 2},
	NAD: {"NAD", "Namibia Dollar", 2},
	NGN: {"NGN", "Naira", 2},
	NIO: {"NIO", "Cordoba Oro", 2},
	NOK: {"NOK", "Norwegian Krone", 2},
	NPR: {"NPR", "Nepalese Rupee", 2},
	NZD: {"NZD", "New Zealand Dollar", 2},
	OMR: {"OMR", "Rial Omani", 3},
	PAB: {"PAB", "Balboa", 2},
	PEN: {"PEN", "Sol", 2},
	PGK: {"PGK", "Kina", 2},
	PHP: {"PHP", "Philippine Peso", 2},
	PKR: {"PKR", "Pakistan Rupee", 2},
	PLN: {"PLN", "Zloty", 2},
	PYG: {"PYG", "Guarani", 0},
	QAR: {"QAR", "Qatari Rial", 2},
	RON: {"RON", "Romanian Leu", 2},
	RSD: {"RSD", "Serbian Dinar", 2},
	RUB: {"RUB", "Russian Ruble", 2},
	RWF: {"RWF", "Rwanda Franc", 0},
	SAR: {"SAR", "Saudi Riyal", 2},
	SBD: {"SBD", "Solomon Islands Dollar", 2},
	SCR: {"SCR", "Seychelles Rupee", 2},
	SDG: {"SDG", "Sudanese Pound", 2},
	SEK: {"SEK", "Swedish Krona", 2},
	SGD: {"SGD", "Singapore Dollar", 2},
	SHP: {"SHP", "Saint Helena Pound", 2},
	SLL: {"SLL", "Leone", 2},
	SOS: {"SOS", "Somali Shilling", 2},
	SRD: {"SRD", "Surinam Dollar", 2},
	SSP: {"SSP", "South Sudanese Pound", 2},
	STN: {"STN", "Dobra", 2},
	SVC: {"SVC", "El Salvador Colon", 2},
	SYP: {"SYP", "Syrian Pound", 2},
	SZL: {"SZL", "Lilangeni", 2},
	THB: {"THB", "Baht", 2},
	TJS: {"TJS", "Somoni", 2},
	TMT: {"TMT", "Turkmenistan New Manat", 2},
	TND: {"TND", "Tunisian Dinar", 3},
	TOP: {"TOP", "Pa’anga", 2},
	TRY: {"TRY", "Turkish Lira", 2},
	TTD: {"TTD", "Trinidad and Tobago Dollar", 2},
	TWD: {"TWD", "New Taiwan Dollar", 2},
	TZS: {"TZS", "Tanzanian Shilling", 2},
	UAH: {"UAH", "Hryvnia", 2},
	UGX: {"UGX", "Uganda Shilling", 0},
	USD: {"USD", "US Dollar", 2},
	USN: {"USN", "US Dollar (Next day)", 2},
	UYI: {"UYI", "Uruguay Peso en Unidades Indexadas (UI)", 0},
	UYU: {"UYU", "Peso Uruguayo", 2},
	UYW: {"UYW", "Unidad Previsional", 4},
	UZS: {"UZS", "Uzbekistan Sum", 2},
	VES: {"VES", "Bolívar Soberano", 2},
	VND: {"VND", "Dong", 0},
	VUV: {"VUV", "Vatu", 0},
	WST: {"WST", "Tala", 2},
	XAF: {"XAF", "CFA Franc BEAC", 0},
	XCD: {"XCD", "East Caribbean Dollar", 2},
	XDR: {"XDR", "SDR (Special Drawing Right)", 0},
	XOF: {"XOF", "CFA Franc BCEAO", 0},
	XPF: {"XPF", "CFP Franc", 0},
	XSU: {"XSU", "Sucre", 0},
	XUA: {"XUA", "ADB Unit of Account", 0},
	YER: {"YER", "Yemeni Rial", 2},
	ZAR: {"ZAR", "Rand", 2},
	ZMW: {"ZMW", "Zambian Kwacha", 2},
	ZWL: {"ZWL", "Zimbabwe Dollar", 2},
}
