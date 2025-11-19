package linkedLists

import (
	"fmt"
)

func AddTwoNumbers(l1, l2 []int) []int {
	ll1 := NewLinkedList[int]()
	ll2 := NewLinkedList[int]()
	ll3 := NewLinkedList[int]()
	capacity := max(len(l1), len(l2)) + 1
	op := make([]int, 0, capacity)
	for _, item := range l1 {
		ll1.Append(item)
	}
	fmt.Println(ll1)
	for _, item := range l2 {
		ll2.Append(item)
	}
	fmt.Println(ll2)
	carry := 0

	ptr1, ptr2 := ll1.head, ll2.head
	for ptr1 != nil && ptr2 != nil {
		sum := ptr1.val + ptr2.val + carry
		if sum >= 10 {
			carry = int(sum / 10)
		}
		ll3.Append(sum % 10)
		ptr1 = ptr1.next
		ptr2 = ptr2.next
	}

	for ptr1 != nil {
		sum := ptr1.val + carry
		if sum >= 10 {
			carry = int(sum / 10)
		}
		ll3.Append(sum % 10)
		ptr1 = ptr1.next
	}

	for ptr2 != nil {
		sum := ptr2.val + carry
		if sum >= 10 {
			carry = int(sum / 10)
		}
		ll3.Append(sum % 10)
		ptr2 = ptr2.next
	}

	if carry != 0 {
		ll3.Append(carry)
	}

	for ptr := ll3.head; ptr != nil; ptr = ptr.next {
		op = append(op, ptr.val)
	}

	return op

}

func SubtractTwoNumbers(l1, l2 []int) []int {
	if len(l1) < len(l2) {
		return nil
	}

	ll1 := NewLinkedList[int]()
	ll2 := NewLinkedList[int]()
	ll3 := NewLinkedList[int]()
	op := make([]int, 0, len(l1))

	for _, v := range l1 {
		ll1.Append(v)
	}
	for _, v := range l2 {
		ll2.Append(v)
	}

	ptr1 := ll1.head
	ptr2 := ll2.head

	for ptr2 != nil {
		if ptr1.val < ptr2.val {
			ptr1.next.val = ptr1.next.val - 1
			ptr1.val = ptr1.val + 10
		}
		ll3.Append(ptr1.val - ptr2.val)
		ptr1 = ptr1.next
		ptr2 = ptr2.next
	}

	for ptr1 != nil {
		ll3.Append(ptr1.val)
		ptr1 = ptr1.next
	}

	for ptr := ll3.head; ptr != nil; ptr = ptr.next {
		op = append(op, ptr.val)
	}

	return op

}
