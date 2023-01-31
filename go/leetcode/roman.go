package main

import (
	"fmt"
	"sort"
)

/*
I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
*/

// roman: integer mapping
var ri = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
	"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000,
}

// integer: roman mapping
var ir = transpose(ri)

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
	fmt.Println(intToRoman(1994))
}

func romanToInt(s string) int {
	var result int

	for i, c := range s {
		cs := string(c)
		last := false
		// Find if iterator at the last character of the string
		if i == len([]rune(s))-1 {
			last = true
		}
		switch cs {
		case "I":
			if !last && contains([]string{"V", "X"}, string(s[i+1])) {
				result += -(ri[cs])
				continue
			}
		case "X":
			if !last && contains([]string{"L", "C"}, string(s[i+1])) {
				result += -(ri[cs])
				continue
			}
		case "C":
			if !last && contains([]string{"D", "M"}, string(s[i+1])) {
				result += -(ri[cs])
				continue
			}
		}
		result += ri[cs]
	}
	return result
}

func intToRoman(n int) string {
	var result string
	var irn []int

	for k := range ir {
		irn = append(irn, k)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(irn)))

	for n != 0 {
		for _, v := range irn {
			for n/v != 0 {
				result += ir[v]
				n += -(v)
			}
		}
	}

	return result
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func transpose(m map[string]int) map[int]string {
	nm := make(map[int]string, len(m))

	for k, v := range m {
		nm[v] = k
	}
	return nm
}
