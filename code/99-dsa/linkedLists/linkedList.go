package linkedLists

import (
	"fmt"
	"strings"
)

type LinkedList[T comparable] struct {
	head, tail *Node[T]
}

type Node[T comparable] struct {
	next *Node[T]
	val  T
}

// Create New
func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Add Element at tail end O(1)
func (ll *LinkedList[T]) Append(val T) *Node[T] {

	node := &Node[T]{}
	node.val = val

	if ll.head == nil {
		ll.head = node
		ll.tail = node
	} else {
		ll.tail.next = node
		ll.tail = node
	}

	fmt.Printf("Element %v Added at the Tail. \n", val)
	return ll.tail
}

// Add Element at head end O(1)
func (ll *LinkedList[T]) Prepend(val T) *Node[T] {

	node := &Node[T]{}
	node.val = val

	if ll.head == nil {
		ll.head = node
		ll.tail = node
	}

	node.next = ll.head
	ll.head = node

	fmt.Printf("Element %v Added at the Head. \n", val)
	return ll.head
}

// Find Element O(n)
func (ll *LinkedList[T]) Find(val T) *Node[T] {
	var found *Node[T]
	for start := ll.head; start != nil; start = start.next {
		if start.val == val {
			found = start
			break
		}
	}
	if found != nil {
		fmt.Printf("Element %v Found. \n", val)
	} else {
		fmt.Printf("Element %v Not Found. \n", val)
	}
	return found
}

// Delete Element O(n)
func (ll *LinkedList[T]) Delete(val T) *Node[T] {

	if ll.head == nil {
		fmt.Printf("Cannot Delete from Empty List \n")
		return nil
	}

	if ll.head == ll.tail && ll.head.val == val {
		deleted := ll.Shift()
		fmt.Printf("Deleted Element %v \n", deleted.val)
		return deleted
	}

	var prev *Node[T]
	curr := ll.head

	for curr != nil {
		if curr.val == val {
			if prev == nil {
				ll.head = curr.next
				if curr == ll.tail {
					ll.tail = nil
				}
			} else {
				prev.next = curr.next
				if curr == ll.tail {
					ll.tail = prev
				}
			}
			curr.next = nil
			return curr
		}
		prev = curr
		curr = curr.next
	}

	return nil
}

// Delete at head end O(1)
func (ll *LinkedList[T]) Shift() *Node[T] {
	if ll.head == nil {
		fmt.Printf("Cannot Delete from Empty List \n")
		return nil
	}

	if ll.head.next == nil {
		deleted := ll.head
		ll.head = nil
		ll.tail = nil
		fmt.Printf("Deleted Element %v \n", deleted.val)
		return deleted
	}

	curr := ll.head
	ll.head = curr.next
	curr.next = nil
	fmt.Printf("Deleted Element %v \n", curr.val)
	return curr
}

// Delete at tail end O(1)
func (ll *LinkedList[T]) Pop() *Node[T] {
	if ll.tail == nil {
		fmt.Printf("Cannot Delete from Empty List \n")
		return nil
	}

	if ll.head == ll.tail {
		deleted := ll.Shift()
		fmt.Printf("Deleted Element %v \n", deleted.val)
		return deleted
	}

	prev := ll.head
	for prev.next != ll.tail {
		prev = prev.next
	}

	deleted := ll.tail
	ll.tail = prev
	ll.tail.next = nil
	fmt.Printf("Deleted Element %v \n", deleted.val)
	return deleted
}

//Insert After
//Insert Before
//Length
//Insert at Pos

// Print O(n)
func (ll *LinkedList[T]) String() string {
	var str strings.Builder
	for start := ll.head; start != nil; start = start.next {
		fmt.Fprintf(&str, "(%v)", start.val)
		if start.next != nil {
			str.WriteString(" => ")
		}
	}
	return str.String()
}
