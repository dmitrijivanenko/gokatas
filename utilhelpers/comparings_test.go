package utilhelpers

import "testing"

type testingPairs struct{
	firsSlice, secondSlice []int
	isEqual bool
}

var testData = []testingPairs{
	{ []int{1,2}, []int{2,1}, false },
	{ nil, nil, true },
	{ []int{1,2}, []int{1,2,3}, false },
	{ []int{1,2}, nil, false },
	{ []int{1,2}, []int{1,2}, true },
}


func TestCompareSlices(t *testing.T) {

	for _, pair := range testData {
		if result := CompareSlices(pair.firsSlice, pair.secondSlice); result != pair.isEqual {
			t.Error("Error")
		}
	}
}
