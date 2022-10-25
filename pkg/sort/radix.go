package sort

import (
	utils "demo/pkg/utils"
	"math"
)

func remplir(array []int, digit int) [][]int {
	var digitValue int
	var radArray [][]int = make([][]int, 10)

	for _, element := range array {
		digitValue = element / int(math.Pow10(digit))
		digitValue %= 10
		radArray[digitValue] = append(radArray[digitValue], element)
	}

	return radArray
}

func vider(radArray [][]int) []int {
	var array []int

	for i := 0; i < len(radArray); i++ {
		for len(radArray[i]) > 0 {
			array = append(array, radArray[i][0])
			radArray[i] = radArray[i][1:]
		}
	}

	return array
}

func Radix(array []int) []int {
	var nbDigits int = utils.NumberOfDigits(array)
	var radArray [][]int

	for index := 0; index < nbDigits; index++ {
		radArray = remplir(array, index)
		array = vider(radArray)
	}

	return array
}
