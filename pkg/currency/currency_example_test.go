package currency_test

import (
	"fmt"

	"github.com/ilyakaznacheev/tiny-wallet/pkg/currency"
)

func ExampleConvertToInternal() {
	amountExtUSD := 1234.50
	amountExtIQD := 145.345
	amountExtISK := 25.0

	// convert external (normal) money representations
	// into internal (integer) format
	fmt.Println(currency.ConvertToInternal(amountExtUSD, currency.USD))
	fmt.Println(currency.ConvertToInternal(amountExtIQD, currency.IQD))
	fmt.Println(currency.ConvertToInternal(amountExtISK, currency.ISK))
	// Output: 123450
	// 145345
	// 25
}

func ExampleConvertToExternal() {
	amountIntUSD := 123450
	amountIntIQD := 145345
	amountIntISK := 25

	// convert external (normal) money representations
	// into internal (integer) format
	fmt.Println(currency.ConvertToExternal(amountIntUSD, currency.USD))
	fmt.Println(currency.ConvertToExternal(amountIntIQD, currency.IQD))
	fmt.Println(currency.ConvertToExternal(amountIntISK, currency.ISK))
	// Output: 1234.5
	// 145.345
	// 25
}

func ExampleAtoCurrency() {
	strUSD := "USD"
	strCLP := "CLP"
	strISK := "ISK"

	mustConvert := func(a string) currency.Currency {
		c, _ := currency.AtoCurrency(a)
		return *c
	}

	usd := mustConvert(strUSD)
	clp := mustConvert(strCLP)
	isk := mustConvert(strISK)

	fmt.Printf("%T: %#v: %s\n", usd, usd, usd)
	fmt.Printf("%T: %#v: %s\n", clp, clp, clp)
	fmt.Printf("%T: %#v: %s\n", isk, isk, isk)
	// Output: currency.Currency: "USD": US Dollar
	// currency.Currency: "CLP": Chilean Peso
	// currency.Currency: "ISK": Iceland Krona
}

func Example() {
	strUSD := "USD"
	strCLP := "CLP"
	strISK := "ISK"

	// convert string codes into Currency codes
	usd, _ := currency.AtoCurrency(strUSD)
	clp, _ := currency.AtoCurrency(strCLP)
	isk, _ := currency.AtoCurrency(strISK)

	// print currencies
	// each currency will print its name
	fmt.Println(usd)
	fmt.Println(clp)
	fmt.Println(isk)
	// Output: US Dollar
	// Chilean Peso
	// Iceland Krona
}
