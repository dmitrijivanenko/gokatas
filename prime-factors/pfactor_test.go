package pfactor

import "testing"
import "github.com/dmitrijivanenko/gokatas/utilhelpers"

type testingPair struct {
	number int
	primes []int
}

var testData = []testingPair{
	{1, []int{}},
	{2, []int{2}},
	{100, []int{2, 2, 5, 5}},
}

func TestFactor(t *testing.T) {

	for _, pair := range testData {
		if result := Factor(pair.number); !utilhelpers.CompareSlices(pair.primes, result) {
			t.Error("For ", pair.number, " expected ", pair.primes)
		}
	}
}
