package cprimes

import "github.com/dmitrijivanenko/gokatas/prime-factors"


func CalculateCircularPrimes(limit int) int {
	var result int

	for i := 1; i <= limit; i++ {

		{
			var tempResult bool = true

			for _,number := range explode(i) {
				if !isPrime(number) {
					tempResult = false
				}
			}

			if tempResult {
				result++
			}
		}
	}

	return result
}

func explode(number int) []int {

	var result = make([]int,0)

	for number > 0  {
		result = append(result, number % 10)
		number /= 10
	}

	return result
}

func isPrime(number int) bool {

	if len(pfactor.Factor(number)) > 1 {
		return false
	}

	return true
}
