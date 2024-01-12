package main

import (
	"fmt"
)

func main() {
	fmt.Println(Solution("qwerty"))
	fmt.Println(PositiveSum([]int{1, -4, 7, 12}))
	fmt.Println(RepeatStr(3, "Tr"))
	fmt.Println(NoSpace("Tr Tr"))
	fmt.Println(DNAStrand("GTAT"))
	fmt.Println(ToAlternatingCase("hello WORLD 1a2b3c4d5e"))

	fmt.Println(GetSum(-5, -5))
}
