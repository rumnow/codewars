package main

import (
	"fmt"
)

func main() {
	fmt.Println(4, maxFrequencyElements([]int{4,2,3,5,7,4,2}))
	fmt.Println(7, maxFrequencyElements([]int{4,2,3,5,7,4,2,2,7}))
	fmt.Println(5, maxFrequencyElements([]int{1,2,3,4,5}))
	fmt.Println(4, maxFrequencyElements([]int{1,2,2,3,1,4}))
	fmt.Println(4, maxFrequencyElements([]int{1,2,2,3,1,4}))
	fmt.Println(findLongestChain([][]int{{4,3},{1,2},{5,6},{7,8}}))
	fmt.Println(minimumPushes2("xyzxyzxyzxyz"))
	fmt.Println(minimumPushes("xycdefghij"))
	fmt.Println(Solution("qwerty"))
	fmt.Println(PositiveSum([]int{1, -4, 7, 12}))
	fmt.Println(RepeatStr(3, "Tr"))
	fmt.Println(NoSpace("Tr Tr"))
	fmt.Println(DNAStrand("GTAT"))
	fmt.Println(ToAlternatingCase("hello WORLD 1a2b3c4d5e"))

	fmt.Println(GetSum(-5, -5))
}
