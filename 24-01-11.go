package main

import (
	"strings"
)

// Complete the solution so that it reverses the string passed into it.
// 'world'  =>  'dlrow'
// 'word'   =>  'drow'
func Solution(word string) string {
	result := make([]rune, 0, len(word))
	for _, v := range word {
		result = append([]rune{v}, result...)
	}
	return string(result)
}

// You get an array of numbers, return the sum of all of the positives ones.
// Example [1,-4,7,12] => 1 + 7 + 12 = 20
// Note: if there is nothing to sum, the sum is default to 0.
func PositiveSum(numbers []int) int {
	result := 0
	for _, v := range numbers {
		if v > 0 {
			result += v
		}
	}
	return result
}

// Write a function that accepts an integer n and a string s as parameters, and returns a string of s repeated exactly n times.
// Examples (input -> output)
// 6, "I"     -> "IIIIII"
// 5, "Hello" -> "HelloHelloHelloHelloHello"
func RepeatStr(repetitions int, value string) string {
	return strings.Repeat(value, repetitions)
	//return strings.ReplaceAll(word, " ", "")
	//return regexp.MustCompile(`\s+`).ReplaceAllString(word, "")
}

//return string without Space
func NoSpace(word string) string {
	return strings.Join(strings.Split(word, " "), "")
}

// In DNA strings, symbols "A" and "T" are complements of each other, as "C" and "G". Your function receives one side of the DNA; you need to return the other complementary side. DNA strand is never empty or there is no DNA at all .
// "ATTGC" --> "TAACG"
// "GTAT" --> "CATA"
func DNAStrand(dna string) string {
	// your code here
	result := []rune{}
	for _, v := range dna {
		switch v {
		case 'A':
			result = append(result, 'T')
		case 'T':
			result = append(result, 'A')
		case 'C':
			result = append(result, 'G')
		case 'G':
			result = append(result, 'C')
		}
	}
	return string(result)
	// replacer := strings.NewReplacer("A", "T", "T", "A", "G", "C", "C", "G")
	// return(replacer.Replace(dna))
}

// Define String.prototype.toAlternatingCase (or a similar function/method such as to_alternating_case/toAlternatingCase/ToAlternatingCase in your selected language; see the initial solution for details) such that each lowercase letter becomes uppercase and each uppercase letter becomes lowercase. For example:
// ToAlternatingCase("hello world"); // => "HELLO WORLD"
// ToAlternatingCase("HELLO WORLD"); // => "hello world"
// ToAlternatingCase("hello WORLD"); // => "HELLO world"
// ToAlternatingCase("HeLLo WoRLD"); // => "hEllO wOrld"
// ToAlternativeCase("12345"); // => "12345" (Non-alphabetical characters are unaffected)
// ToAlternativeCase("1a2b3c4d5e"); // => "1A2B3C4D5E"
// ToAlternativeCase("String.prototype.toAlternatingCase"); // => "sTRING.PROTOTYPE.TOaLTERNATINGcASE"
func ToAlternatingCase(str string) string {
	//unicode.IsUpper(ch)
	result := ""
	for _, v := range str {
		if v > 64 && v < 91 {
			result += strings.ToLower(string(v))
		} else
		if v > 96 && v < 123 {
			result += strings.ToUpper(string(v))
		} else {
			result += string(v)
		}
	}
	return result
  }
