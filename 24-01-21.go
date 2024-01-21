package main

import (
	"sort"
	"strings"
)

//https://leetcode.com/problems/minimum-number-of-pushes-to-type-word-i/
func minimumPushes(word string) int {
	// stantard phone keys
	// res := 0
	// keys := []string{"abc","def","ghi","jkl","mno","pqrs","tuv","wxyz"}
	// for _, v := range word {
	// 	for _, k := range keys {
	// 		if strings.Contains(k, string(v)) {
	// 			res += strings.IndexRune(k, v) + 1
	// 			break
	// 		}
	// 	}
	// }
	// return res

	// dynamic smart phone keys
	res := 0
	keys := [8]string{}
	for i, v := range word {
		keys[i%8] += string(v)
	}
	for _, v := range word {
		for _, k := range keys {
			if strings.Contains(k, string(v)) {
				res += strings.IndexRune(k, v) + 1
				break
			}
		}
	}
	return res
}

//https://leetcode.com/problems/minimum-number-of-pushes-to-type-word-ii/
func minimumPushes2(word string) int {
	type runeInt struct{
		r rune //rune
		c int //count
	}
	res := 0
	keys := [8]string{}
	used := ""
	count := []runeInt{}
	for _, v := range word {
		if strings.Contains(used, string(v)) {
			count[strings.IndexRune(used, v)].c += 1
		} else {
			count = append(count, runeInt{v, 1})
			used += string(v)
		}
	}

	sort.Slice(count, func(i, j int) bool {
		return count[i].c > count[j].c
	})
	for i, v := range count {

		keys[i%8] += string(v.r)
	}
	for _, v := range word {
		for _, k := range keys {
			if strings.Contains(k, string(v)) {
				res += strings.IndexRune(k, v) + 1
				break
			}
		}
	}

	return res
}
