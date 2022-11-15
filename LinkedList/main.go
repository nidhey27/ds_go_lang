package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
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
	itr.next = itr.next.next
}

func (l *LinkedList) insertAt(n *Node, index int) {
	itr := l.head
	count := 0

	if index == 0 {
		l.insertAtBegining(n)
		return
	}
	if index > l.getLength() {
		fmt.Println("Invalid Index!")
		return
	}

	for count != index-1 {
		itr = itr.next
		count += 1
	}
	n.next = itr.next
	itr.next = n

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

func (l *LinkedList) insertAtBegining(n *Node) {
	node := l.head
	l.head = n
	l.head.next = node
}

func (l *LinkedList) insertAtEnd(n *Node) {
	itr := l.head

	if l.getLength() == 0 {
		l.insertAtBegining(n)
		return
	}

	for itr.next != nil {
		itr = itr.next
	}
	node := n
	itr.next = n
	itr.next.data = node.data
}

func (l LinkedList) printList() {
	if l.getLength() == 0 {
		fmt.Println("Linked List is Empty")
		return
	}
	itr := l.head
	for itr != nil {
		fmt.Printf("%v ", itr.data)
		itr = itr.next
	}
}

func main() {
	fmt.Println("------------ LINKED LIST ------------")

	myLinkedList := LinkedList{}

	myLinkedList.insertAtBegining(&Node{data: -27})
	myLinkedList.insertAtBegining(&Node{data: -23})
	myLinkedList.insertAtBegining(&Node{data: -24})

	myLinkedList.insertAtEnd(&Node{data: 0})
	myLinkedList.insertAtEnd(&Node{data: 1})
	myLinkedList.insertAtEnd(&Node{data: 2})
	myLinkedList.insertAtEnd(&Node{data: 3})

	myLinkedList.insertAt(&Node{data: -50}, 2)
	myLinkedList.insertAt(&Node{data: 99}, 8)
	myLinkedList.insertAt(&Node{data: 100}, 10)

	myLinkedList.deleteNodeWithValue(-23)
	myLinkedList.deleteNodeWithValue(100)

	myLinkedList.printList()
	fmt.Println("\nLength")
	fmt.Println(myLinkedList.getLength())
}
