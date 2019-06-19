// Package currency contains ISO 4217 currency codes and currency names
package currency

import (
	"math"
	"strconv"
)

// Currency type for a currency ISO 4217 code
type Currency int

// String returns a name of currency
func (c Currency) String() string {
	if p, ok := currencyProperties[c]; ok {
		return p.Name
	}
	return "non-ISO 4216 currency"
}

// FormatAmount returns an integer amount formatted depending on the number of decimal places of the currency
func (c Currency) FormatAmount(raw int) string {
	p := currencyProperties[c]
	b := float64(raw) / math.Pow10(int(p.Decimals))
	return strconv.FormatFloat(b, 'f', -1, 64)
}

const (
	AFN Currency = 971
	ALL Currency = 8
	DZD Currency = 12
	ARS Currency = 32
	AUD Currency = 36
	BSD Currency = 44
	BHD Currency = 48
	BDT Currency = 50
	AMD Currency = 51
	BBD Currency = 52
	BMD Currency = 60
	BTN Currency = 64
	BOB Currency = 68
	BWP Currency = 72
	BZD Currency = 84
	SBD Currency = 90
	BND Currency = 96
	MMK Currency = 104
	BIF Currency = 108
	KHR Currency = 116
	CAD Currency = 124
	CVE Currency = 132
	KYD Currency = 136
	LKR Currency = 144
	CLP Currency = 152
	CNY Currency = 156
	COP Currency = 170
	KMF Currency = 174
	CRC Currency = 188
	HRK Currency = 191
	CUP Currency = 192
	CZK Currency = 203
	DKK Currency = 208
	DOP Currency = 214
	SVC Currency = 222
	ETB Currency = 230
	ERN Currency = 232
	FKP Currency = 238
	FJD Currency = 242
	DJF Currency = 262
	GMD Currency = 270
	GIP Currency = 292
	GTQ Currency = 320
	GNF Currency = 324
	GYD Currency = 328
	HTG Currency = 332
	HNL Currency = 340
	HKD Currency = 344
	HUF Currency = 348
	ISK Currency = 352
	INR Currency = 356
	IDR Currency = 360
	IRR Currency = 364
	IQD Currency = 368
	ILS Currency = 376
	JMD Currency = 388
	JPY Currency = 392
	KZT Currency = 398
	JOD Currency = 400
	KES Currency = 404
	KPW Currency = 408
	KRW Currency = 410
	KWD Currency = 414
	KGS Currency = 417
	LAK Currency = 418
	LBP Currency = 422
	LSL Currency = 426
	LRD Currency = 430
	LYD Currency = 434
	MOP Currency = 446
	MWK Currency = 454
	MYR Currency = 458
	MVR Currency = 462
	MUR Currency = 480
	MXN Currency = 484
	MNT Currency = 496
	MDL Currency = 498
	MAD Currency = 504
	OMR Currency = 512
	NAD Currency = 516
	NPR Currency = 524
	ANG Currency = 532
	AWG Currency = 533
	VUV Currency = 548
	NZD Currency = 554
	NIO Currency = 558
	NGN Currency = 566
	NOK Currency = 578
	PKR Currency = 586
	PAB Currency = 590
	PGK Currency = 598
	PYG Currency = 600
	PEN Currency = 604
	PHP Currency = 608
	QAR Currency = 634
	RUB Currency = 643
	RWF Currency = 646
	SHP Currency = 654
	SAR Currency = 682
	SCR Currency = 690
	SLL Currency = 694
	SGD Currency = 702
	VND Currency = 704
	SOS Currency = 706
	ZAR Currency = 710
	SSP Currency = 728
	SZL Currency = 748
	SEK Currency = 752
	CHF Currency = 756
	SYP Currency = 760
	THB Currency = 764
	TOP Currency = 776
	TTD Currency = 780
	AED Currency = 784
	TND Currency = 788
	UGX Currency = 800
	MKD Currency = 807
	EGP Currency = 818
	GBP Currency = 826
	TZS Currency = 834
	USD Currency = 840
	UYU Currency = 858
	UZS Currency = 860
	WST Currency = 882
	YER Currency = 886
	TWD Currency = 901
	UYW Currency = 927
	VES Currency = 928
	MRU Currency = 929
	STN Currency = 930
	CUC Currency = 931
	ZWL Currency = 932
	BYN Currency = 933
	TMT Currency = 934
	GHS Currency = 936
	SDG Currency = 938
	UYI Currency = 940
	RSD Currency = 941
	MZN Currency = 943
	AZN Currency = 944
	RON Currency = 946
	CHE Currency = 947
	CHW Currency = 948
	TRY Currency = 949
	XAF Currency = 950
	XCD Currency = 951
	XOF Currency = 952
	XPF Currency = 953
	XDR Currency = 960
	XUA Currency = 965
	ZMW Currency = 967
	SRD Currency = 968
	MGA Currency = 969
	COU Currency = 970
	TJS Currency = 972
	AOA Currency = 973
	BGN Currency = 975
	CDF Currency = 976
	BAM Currency = 977
	EUR Currency = 978
	MXV Currency = 979
	UAH Currency = 980
	GEL Currency = 981
	BOV Currency = 984
	PLN Currency = 985
	BRL Currency = 986
	CLF Currency = 990
	XSU Currency = 994
	USN Currency = 997
)

type property struct {
	Name     string
	Decimals uint
}

var currencyProperties = map[Currency]property{
	AFN: {"Afghani", 2},
	AED: {"UAE Dirham", 2},
	ALL: {"Lek", 2},
	AMD: {"Armenian Dram", 2},
	ANG: {"Netherlands Antillean Guilder", 2},
	AOA: {"Kwanza", 2},
	ARS: {"Argentine Peso", 2},
	AUD: {"Australian Dollar", 2},
	AWG: {"Aruban Florin", 2},
	AZN: {"Azerbaijan Manat", 2},
	BAM: {"Convertible Mark", 2},
	BBD: {"Barbados Dollar", 2},
	BDT: {"Taka", 2},
	BGN: {"Bulgarian Lev", 2},
	BHD: {"Bahraini Dinar", 3},
	BIF: {"Burundi Franc", 0},
	BMD: {"Bermudian Dollar", 2},
	BND: {"Brunei Dollar", 2},
	BOB: {"Boliviano", 2},
	BOV: {"Mvdol", 2},
	BRL: {"Brazilian Real", 2},
	BSD: {"Bahamian Dollar", 2},
	BTN: {"Ngultrum", 2},
	BWP: {"Pula", 2},
	BYN: {"Belarusian Ruble", 2},
	BZD: {"Belize Dollar", 2},
	CAD: {"Canadian Dollar", 2},
	CDF: {"Congolese Franc", 2},
	CHE: {"WIR Euro", 2},
	CHF: {"Swiss Franc", 2},
	CHW: {"WIR Franc", 2},
	CLF: {"Unidad de Fomento", 4},
	CLP: {"Chilean Peso", 0},
	CNY: {"Yuan Renminbi", 2},
	COP: {"Colombian Peso", 2},
	COU: {"Unidad de Valor Real", 2},
	CRC: {"Costa Rican Colon", 2},
	CUC: {"Peso Convertible", 2},
	CUP: {"Cuban Peso", 2},
	CVE: {"Cabo Verde Escudo", 2},
	CZK: {"Czech Koruna", 2},
	DJF: {"Djibouti Franc", 0},
	DKK: {"Danish Krone", 2},
	DOP: {"Dominican Peso", 2},
	DZD: {"Algerian Dinar", 2},
	EGP: {"Egyptian Pound", 2},
	ERN: {"Nakfa", 2},
	ETB: {"Ethiopian Birr", 2},
	EUR: {"Euro", 2},
	FJD: {"Fiji Dollar", 2},
	FKP: {"Falkland Islands Pound", 2},
	GBP: {"Pound Sterling", 2},
	GEL: {"Lari", 2},
	GHS: {"Ghana Cedi", 2},
	GIP: {"Gibraltar Pound", 2},
	GMD: {"Dalasi", 2},
	GNF: {"Guinean Franc", 0},
	GTQ: {"Quetzal", 2},
	GYD: {"Guyana Dollar", 2},
	HKD: {"Hong Kong Dollar", 2},
	HNL: {"Lempira", 2},
	HRK: {"Kuna", 2},
	HTG: {"Gourde", 2},
	HUF: {"Forint", 2},
	IDR: {"Rupiah", 2},
	ILS: {"New Israeli Sheqel", 2},
	INR: {"Indian Rupee", 2},
	IQD: {"Iraqi Dinar", 3},
	IRR: {"Iranian Rial", 2},
	ISK: {"Iceland Krona", 0},
	JMD: {"Jamaican Dollar", 2},
	JOD: {"Jordanian Dinar", 3},
	JPY: {"Yen", 0},
	KES: {"Kenyan Shilling", 2},
	KGS: {"Som", 2},
	KHR: {"Riel", 2},
	KMF: {"Comorian Franc ", 0},
	KPW: {"North Korean Won", 2},
	KRW: {"Won", 0},
	KWD: {"Kuwaiti Dinar", 3},
	KYD: {"Cayman Islands Dollar", 2},
	KZT: {"Tenge", 2},
	LAK: {"Lao Kip", 2},
	LBP: {"Lebanese Pound", 2},
	LKR: {"Sri Lanka Rupee", 2},
	LRD: {"Liberian Dollar", 2},
	LSL: {"Loti", 2},
	LYD: {"Libyan Dinar", 3},
	MAD: {"Moroccan Dirham", 2},
	MDL: {"Moldovan Leu", 2},
	MGA: {"Malagasy Ariary", 2},
	MKD: {"Denar", 2},
	MMK: {"Kyat", 2},
	MNT: {"Tugrik", 2},
	MOP: {"Pataca", 2},
	MRU: {"Ouguiya", 2},
	MUR: {"Mauritius Rupee", 2},
	MVR: {"Rufiyaa", 2},
	MWK: {"Malawi Kwacha", 2},
	MXN: {"Mexican Peso", 2},
	MXV: {"Mexican Unidad de Inversion (UDI)", 2},
	MYR: {"Malaysian Ringgit", 2},
	MZN: {"Mozambique Metical", 2},
	NAD: {"Namibia Dollar", 2},
	NGN: {"Naira", 2},
	NIO: {"Cordoba Oro", 2},
	NOK: {"Norwegian Krone", 2},
	NPR: {"Nepalese Rupee", 2},
	NZD: {"New Zealand Dollar", 2},
	OMR: {"Rial Omani", 3},
	PAB: {"Balboa", 2},
	PEN: {"Sol", 2},
	PGK: {"Kina", 2},
	PHP: {"Philippine Peso", 2},
	PKR: {"Pakistan Rupee", 2},
	PLN: {"Zloty", 2},
	PYG: {"Guarani", 0},
	QAR: {"Qatari Rial", 2},
	RON: {"Romanian Leu", 2},
	RSD: {"Serbian Dinar", 2},
	RUB: {"Russian Ruble", 2},
	RWF: {"Rwanda Franc", 0},
	SAR: {"Saudi Riyal", 2},
	SBD: {"Solomon Islands Dollar", 2},
	SCR: {"Seychelles Rupee", 2},
	SDG: {"Sudanese Pound", 2},
	SEK: {"Swedish Krona", 2},
	SGD: {"Singapore Dollar", 2},
	SHP: {"Saint Helena Pound", 2},
	SLL: {"Leone", 2},
	SOS: {"Somali Shilling", 2},
	SRD: {"Surinam Dollar", 2},
	SSP: {"South Sudanese Pound", 2},
	STN: {"Dobra", 2},
	SVC: {"El Salvador Colon", 2},
	SYP: {"Syrian Pound", 2},
	SZL: {"Lilangeni", 2},
	THB: {"Baht", 2},
	TJS: {"Somoni", 2},
	TMT: {"Turkmenistan New Manat", 2},
	TND: {"Tunisian Dinar", 3},
	TOP: {"Pa’anga", 2},
	TRY: {"Turkish Lira", 2},
	TTD: {"Trinidad and Tobago Dollar", 2},
	TWD: {"New Taiwan Dollar", 2},
	TZS: {"Tanzanian Shilling", 2},
	UAH: {"Hryvnia", 2},
	UGX: {"Uganda Shilling", 0},
	USD: {"US Dollar", 2},
	USN: {"US Dollar (Next day)", 2},
	UYI: {"Uruguay Peso en Unidades Indexadas (UI)", 0},
	UYU: {"Peso Uruguayo", 2},
	UYW: {"Unidad Previsional", 4},
	UZS: {"Uzbekistan Sum", 2},
	VES: {"Bolívar Soberano", 2},
	VND: {"Dong", 0},
	VUV: {"Vatu", 0},
	WST: {"Tala", 2},
	XAF: {"CFA Franc BEAC", 0},
	XCD: {"East Caribbean Dollar", 2},
	XDR: {"SDR (Special Drawing Right)", 0},
	XOF: {"CFA Franc BCEAO", 0},
	XPF: {"CFP Franc", 0},
	XSU: {"Sucre", 0},
	XUA: {"ADB Unit of Account", 0},
	YER: {"Yemeni Rial", 2},
	ZAR: {"Rand", 2},
	ZMW: {"Zambian Kwacha", 2},
	ZWL: {"Zimbabwe Dollar", 2},
}
