package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// l7 := &ListNode{Val: 1, Next: nil}
	// l6 := &ListNode{Val: 2, Next: l7}
	// l5 := &ListNode{Val: 0, Next: l6}
	// l4 := &ListNode{Val: 0, Next: l5}
	// l3 := &ListNode{Val: 0, Next: l4}
	// l2 := &ListNode{Val: 0, Next: l3}
	// l1 := &ListNode{Val: 9, Next: l2}

	ll4 := &ListNode{Val: 9, Next: nil}
	ll3 := &ListNode{Val: 9, Next: ll4}
	ll2 := &ListNode{Val: 9, Next: ll3}
	ll1 := &ListNode{Val: 9, Next: ll2}

	l7 := &ListNode{Val: 9, Next: nil}
	l6 := &ListNode{Val: 9, Next: l7}
	l5 := &ListNode{Val: 9, Next: l6}
	l4 := &ListNode{Val: 9, Next: l5}
	l3 := &ListNode{Val: 9, Next: l4}
	l2 := &ListNode{Val: 9, Next: l3}
	l1 := &ListNode{Val: 9, Next: l2}

	// ll3 := &ListNode{Val: 3, Next: nil}
	// ll2 := &ListNode{Val: 4, Next: ll3}
	// ll1 := &ListNode{Val: 2, Next: ll2}

	// l3 := &ListNode{Val: 4, Next: nil}
	// l2 := &ListNode{Val: 6, Next: l3}
	// l1 := &ListNode{Val: 5, Next: l2}

	ptr := addTwoNumbers(l1, ll1)
	for ptr != nil {
		fmt.Printf("%v-", ptr.Val)
		ptr = ptr.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	p1 := l1
	p2 := l2
	var p3 *ListNode
	var p4 *ListNode
	var num int

	for p1 != nil && p2 != nil {
		num = p1.Val + p2.Val + num
		if num > 9 {
			p3 = addNode(p3, num%10)
			num = num / 10
		} else {
			p3 = addNode(p3, num)
			num = 0
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	for p1 != nil {
		if num != 0 {
			num += p1.Val
			if num > 9 {
				p3 = addNode(p3, num%10)
				num = num / 10
			} else {
				p3 = addNode(p3, num)
				num = 0
			}
		} else {
			p3 = addNode(p3, p1.Val)
		}
		p1 = p1.Next
	}

	for p2 != nil {
		if num != 0 {
			num += p2.Val
			if num > 9 {
				p3 = addNode(p3, num%10)
				num = num / 10
			} else {
				p3 = addNode(p3, num)
				num = 0
			}
		} else {
			p3 = addNode(p3, p2.Val)
		}
		p2 = p2.Next
	}

	if num != 0 {
		p3 = addNode(p3, num)
	}

	for p3 != nil {
		p4 = addNode(p4, p3.Val)
		p3 = p3.Next
	}

	return p4
}

func addNode(ptr *ListNode, value int) *ListNode {
	return &ListNode{
		Val:  value,
		Next: ptr,
	}
}
