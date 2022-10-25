package utils

import (
	"math/rand"
	"time"
)

func FillArray(length int) []int {
	s1 := rand.NewSource(time.Now().Unix())
	r1 := rand.New(s1)
	var array []int

	for index := 0; index < length; index++ {
		array = append(array, r1.Intn(length))
	}

	return array
}

func countDigits(n int) int {
	if n/10 == 0 {
		return 1
	}
	return 1 + countDigits(n/10)
}

func NumberOfDigits(array []int) int {
	var max int = array[0]

	for _, element := range array {
		if element > max {
			max = element
		}
	}

	return countDigits(max)
}
