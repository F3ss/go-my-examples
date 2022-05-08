package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l7 := &ListNode{Val: 1, Next: nil}
	l6 := &ListNode{Val: 2, Next: l7}
	l5 := &ListNode{Val: 0, Next: l6}
	l4 := &ListNode{Val: 0, Next: l5}
	l3 := &ListNode{Val: 0, Next: l4}
	l2 := &ListNode{Val: 0, Next: l3}
	l1 := &ListNode{Val: 9, Next: l2}
	// 9-0-1-0-0-2-0
	// 8-0-0-0-1-2-1
	ll4 := &ListNode{Val: 9, Next: nil}
	ll3 := &ListNode{Val: 9, Next: ll4}
	ll2 := &ListNode{Val: 9, Next: ll3}
	ll1 := &ListNode{Val: 9, Next: ll2}

	// l7 := &ListNode{Val: 9, Next: nil}
	// l6 := &ListNode{Val: 9, Next: l7}
	// l5 := &ListNode{Val: 9, Next: l6}
	// l4 := &ListNode{Val: 9, Next: l5}
	// l3 := &ListNode{Val: 9, Next: l4}
	// l2 := &ListNode{Val: 9, Next: l3}
	// l1 := &ListNode{Val: 9, Next: l2}

	// ll4 := &ListNode{Val: 9, Next: nil}
	// ll3 := &ListNode{Val: 9, Next: ll4}
	// ll2 := &ListNode{Val: 9, Next: ll3}
	// ll1 := &ListNode{Val: 9, Next: ll2}

	// l1 := &ListNode{Val: 5, Next: nil}
	// ll1 := &ListNode{Val: 5, Next: nil}

	ptr := addTwoNumbers(l1, ll1)
	for ptr != nil {
		fmt.Printf("%v-", ptr.Val)
		ptr = ptr.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	counter := 1
	result := 0
	p1 := l1
	p2 := l2

	for p1 != nil {
		if p1.Val == 0 {
			result *= 10
		} else {
			result += p1.Val * counter

		}
		fmt.Println(result)
		p1 = p1.Next
		counter *= 10
	}
	fmt.Print("\n\n\n")

	counter = 1

	for p2 != nil {
		if p2.Val == 0 {
			result *= 10
		} else {
			result += p2.Val * counter

		}
		fmt.Println(result)
		p2 = p2.Next
		counter *= 10
	}

	p1 = nil

	for counter <= result {
		counter = counter * 10
	}

	for i := counter / 10; i > 0; i = i / 10 {
		p1 = addNode(p1, (result / i))
		result = result % i
	}

	return p1
}

func addNode(ptr *ListNode, value int) *ListNode {
	return &ListNode{
		Val:  value,
		Next: ptr,
	}
}
