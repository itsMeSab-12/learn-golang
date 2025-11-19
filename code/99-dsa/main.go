package main

import (
	"dsa/linkedLists"
	"fmt"
)

func linkedListsMain() {

	// ll := linkedLists.NewLinkedList[int]()
	// fmt.Println(ll)
	// ll.Append(1)
	// ll.Prepend(0)
	// ll.Append(2)
	// ll.Append(3)
	// fmt.Println(ll)
	// ll.Find(3)
	// ll.Find(4)
	// ll.Delete(2)
	// ll.Find(2)
	// ll.Prepend(4)
	// ll.Prepend(2)
	// fmt.Println(ll)
	// ll.Shift()
	// ll.Pop()
	// fmt.Println(ll)

	ip1 := []int{2, 3, 4, 9}
	ip2 := []int{9, 9, 9}

	op1 := linkedLists.AddTwoNumbers(ip1, ip2)
	op2 := linkedLists.SubtractTwoNumbers(ip1, ip2)
	fmt.Println(ip1)
	fmt.Println(ip2)
	fmt.Println(op1)
	fmt.Println(op2)
}

func main() {
	//linkedListsMain()
}
