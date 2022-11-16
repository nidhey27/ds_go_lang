package main

import "fmt"

type Node struct {
	prev *Node
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) insertAtEnd(n *Node) {
	if l.getLength() == 0 {
		l.insertAtBegining(n)
		return
	}
	itr := l.getLastNode()
	itr.next = n
	n.prev = itr
	n.next = nil

}

func (l *LinkedList) deleteNodeWithValue(value int) {
	if l.getLength() == 0 {
		fmt.Printf("Linked List is Empty\n")
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		return
	}
	itr := l.head
	for itr.next.data != value {
		if itr.next.next == nil {
			fmt.Printf("%v Not Present in LL", value)
			return
		}
		itr = itr.next
	}

	// fmt.Println(itr)
	// itr.next = itr.next.next
	// fmt.Println(itr.next.next)
	if itr.next != nil {
		itr.next.prev = itr.prev
	}
	if itr.prev != nil {
		itr.prev.next = itr.next
	}
	// fmt.Println(itr)
	// fmt.Println(itr.next.next)
	// // itr.next.next = itr.prev
	// // itr.next.next = itr
	// // itr.next.next = itr
	// itr.next = itr.next.next

	// itr.prev = itr.next.prev
	// itr.next = itr.next.next

	// itr.prev = itr.prev.next
	// itr.next.next.prev = itr
}

func (l *LinkedList) insertAtBegining(n *Node) {
	if l.head == nil {
		node := l.head
		l.head = n
		l.head.next = node
	} else {
		n.next = l.head
		n.prev = nil
		l.head.prev = n
		l.head = n
	}
}

func (l LinkedList) printForwardList() {
	if l.getLength() == 0 {
		fmt.Println("Linked List is Empty")
		return
	}
	itr := l.head
	for itr != nil {
		fmt.Printf("%v ", itr)
		itr = itr.next
	}
}

func (l *LinkedList) getLastNode() *Node {
	itr := l.head
	for itr.next != nil {
		itr = itr.next
	}
	return itr
}

func (l *LinkedList) printBackwardList() {
	if l.getLength() == 0 {
		fmt.Println("Linked List is Empty")
		return
	}

	itr := l.getLastNode()
	for itr != nil {
		fmt.Printf("%v ", itr)
		itr = itr.prev
	}
}

func (l LinkedList) getLength() int {
	itr := l.head
	count := 0
	for itr != nil {
		count += 1
		itr = itr.next
	}

	return count
}

func main() {
	fmt.Println("------------DOUBLY LINKED LIST ------------")

	doublyLinkedList := LinkedList{}

	doublyLinkedList.insertAtBegining(&Node{data: 21})
	doublyLinkedList.insertAtBegining(&Node{data: 24})
	doublyLinkedList.insertAtBegining(&Node{data: 26})

	doublyLinkedList.insertAtEnd(&Node{data: -1})
	doublyLinkedList.insertAtEnd(&Node{data: -2})

	doublyLinkedList.deleteNodeWithValue(-2)
	doublyLinkedList.deleteNodeWithValue(24)

	fmt.Println("Forward")
	doublyLinkedList.printForwardList()
	fmt.Println()
	fmt.Println("Backward")
	doublyLinkedList.printBackwardList()
	fmt.Println()
	fmt.Println("Length")
	fmt.Println(doublyLinkedList.getLength())
}
