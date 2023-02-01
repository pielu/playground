package palindrom

import (
	"strconv"
	"strings"
)

func IsPalindromeString(x int) bool {
	var y []string = strings.Split(strconv.Itoa(x), "")
	if y[0] == y[len(y)-1] {
		return true
	}
	return false
}

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	y := 0
	z := x
	for z > 0 {
		remainder := z % 10
		y *= 10
		y += remainder
		z /= 10
	}
	if y == x {
		return true
	}
	return false
}
