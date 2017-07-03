package cprimes

import (
	"testing"
	"github.com/dmitrijivanenko/gokatas/utilhelpers"
)

func TestExplode(t *testing.T) {

	var testData = []struct{
		number int
		result []int
	}{
		{ 123, []int{3,2,1} },
	}

	for _, pair := range testData  {
		if !utilhelpers.CompareSlices(explode(pair.number), pair.result) {
			t.Error("For ", pair.number,
				" expected ", pair.result,
				" got ", explode(pair.number))
		}
	}

}

func TestIsPrime(t *testing.T) {

	var testData = []struct{
		number int
		result bool
	}{
		{ 1, true },
		{ 2, true },
		{ 4, false },
		{ 5, true },
		{ 5, true },
	}

	for _, pair := range testData  {
		if isPrime(pair.number) != pair.result {
			t.Error("For ", pair.number,
				" expected ", pair.result,
				" got ", pair.number)
		}
	}
}

func TestCalculateCircularPrimes(t *testing.T) {

	var testData = []struct{
		number int
		result int
	}{
		{ 1, 1 },
		{ 2, 2 },
		{ 4, 3 },
		{ 5, 4 },
		{ 6, 4 },
	}

	for _, pair := range testData  {
		if CalculateCircularPrimes(pair.number) != pair.result {
			t.Error("For ", pair.number,
				" expected ", pair.result,
				" got ", pair.number)
		}
	}
}
