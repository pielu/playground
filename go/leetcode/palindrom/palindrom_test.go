package palindrom

import "testing"

func TestIsPalindrom(t *testing.T) {
	x := -121
	ux := 121

	iPsx := IsPalindromeString(x)
	iPx := IsPalindrome(x)
	if iPx == true || iPsx == true {
		t.Error("Unexpected result:", t)
	}
	iPsux := IsPalindromeString(ux)
	iPux := IsPalindrome(ux)
	if iPux != true || iPsux != true {
		t.Error("Unexpected result:", t)
	}
}
