package main

import (
	sort "demo/pkg/sort"
	utils "demo/pkg/utils"
	"fmt"
)

func main() {
	fmt.Println("Debut de la demo")
	var array []int = utils.FillArray(1000)

	fmt.Println("Array sans trie:", array)

	array = sort.Radix(array)

	fmt.Println("Array avec trie:", array)

	fmt.Println("Fin de la demo")
}
