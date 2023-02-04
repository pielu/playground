package addTwoNumbers

import (
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func lnFollow(ln *ListNode) string {
	s := strconv.Itoa(ln.Val)
	p := ln.Next
	for p != nil {
		s += " -> " + strconv.Itoa(p.Val)
		p = p.Next
	}
	return s
}

func lnInsert(s []int) *ListNode {
	ln, p := &ListNode{}, &ListNode{}

	if len(s) == 0 {
		return ln
	}

	ln.Val = s[0]

	if len(s) > 1 {
		ln.Next = p
	}

	for i := 1; i < len(s); i++ {
		p.Val = s[i]
		if i != len(s)-1 {
			p.Next = &ListNode{}
			p = p.Next
		}
	}

	return ln
}

func lnToList(ln *ListNode) []int {
	var sl []int

	for ln.Next != nil {
		sl = append(sl, ln.Val)
		ln = ln.Next
	}
	sl = append(sl, ln.Val)

	return sl
}

func toNumber(s []int) int {
	rs := reverseIntSlice(s)
	return sliceToInt(rs)
}

func toListNode(n int) *ListNode {
	s := intToReverseSlice(n)
	ln := lnInsert(s)
	return ln
}

func intToReverseSlice(i int) []int {
	ns := make([]int, 0)

	for i > 0 {
		r := i % 10
		i = i / 10
		ns = append(ns, r)

	}
	return ns
}

func sliceToInt(s []int) int {
	res := 0
	op := 1
	for i := len(s) - 1; i >= 0; i-- {
		res += s[i] * op
		op *= 10
	}
	return res
}

func reverseIntSlice(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1st, l2st := lnToList(l1), lnToList(l2)
	l1n, l2n := toNumber(l1st), toNumber(l2st)
	l1l2 := l1n + l2n
	r := toListNode(l1l2)
	return r
}
