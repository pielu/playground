package addTwoNumbers

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1sl, l2sl := []int{2, 4, 3}, []int{5, 6, 4}
	l1, l2 := lnInsert(l1sl), lnInsert(l2sl)
	l1a, l2a := lnToList(l1), lnToList(l2)
	// Test slices are what they should be
	if !equal(l1a, l1sl) || !equal(l2a, l2sl) {
		t.Error("Unexpected result:", t)
	}
	l1n, l2n := toNumber(l1a), toNumber(l2a)
	// Test numbers from the slices are what they should be
	if l1n != 342 || l2n != 465 {
		t.Error("Unexpected result:", t)
	}
	// So the sum of two figures has to be what it should
	l1nl2n := l1n + l2n
	if l1nl2n != 807 {
		t.Error("Unexpected result:", t)
	}

	r := addTwoNumbers(l1, l2)
	rstr := lnFollow(r)
	if rstr != "7 -> 0 -> 8" {
		t.Error("Unexpected result:", t)
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
