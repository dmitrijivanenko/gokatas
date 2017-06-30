package rnumerals

import "testing"

type testPairs struct {
	number       int
	romainNumber string
}

var testData = []testPairs{
	{-1, "Wrong number"},
	{1, "I"},
	{4, "IV"},
	{5, "V"},
	{11, "XI"},
	{20, "XX"},
}

func TestConvert(t *testing.T) {
	for _, pair := range testData {
		if Convert(pair.number) != pair.romainNumber {
			t.Error("For ", pair.number,
				" expected ", pair.romainNumber,
				" but gotten ", Convert(pair.number))
		}
	}
}
