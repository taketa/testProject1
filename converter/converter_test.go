package main

import "testing"

type testpair struct{
	num float64
	currency string
	result float64
}
var tests=[]testpair{
	{22.5,"uah",587.813355},
	{50, "eur",42.4802},
	{50, "euR",42.4802},
	{50, "EUR",42.4802},
	{50, "gbp",37.302},

}
func TestConvert(t *testing.T){
	for _, pair := range tests {
		v := convert(pair.num,pair.currency)
		if v!=pair.result{
			t.Error(
				"For", pair.num, pair.currency,
				"expected",pair.result,
				"got", v,
			)
		}
	}
}