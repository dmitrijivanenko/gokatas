package rnumerals

import (
	"sort"
)

var table = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

func Convert(number int) string {

	if number < 0 {
		return "Wrong number"
	}

	keys := []int{}
	for k, _ := range table {
		keys = append(keys, k)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	var result string

	for _, decimal := range keys {
		char := table[decimal]
		for number >= decimal {
			result += string(char)
			number -= decimal
		}
	}

	return result
}
