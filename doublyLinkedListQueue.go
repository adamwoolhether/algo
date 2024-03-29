package algo

import (
	"fmt"
	"io"
)

// doubleNode represents a node within a doubly-linked list. Not concurrency safe as is.
type doubleNode struct {
	data any
	next *doubleNode
	prev *doubleNode
}

func newDoubleNode(value any) *doubleNode {
	return &doubleNode{
		data: value,
	}
}

type doublyLinkedList struct {
	head *doubleNode
	tail *doubleNode
}

func NewDoublyLinkedList() *doublyLinkedList {
	return &doublyLinkedList{}
}

func (ll *doublyLinkedList) insertAtEnd(value any) {
	newNode := newDoubleNode(value)

	if ll.head == nil { // If there are no elements in the linked list.
		ll.head = newNode
		ll.tail = newNode
	} else { // If the linked list has at least one node.
		newNode.prev = ll.tail
		ll.tail.next = newNode
		ll.tail = newNode
	}
}

func (ll *doublyLinkedList) removeFromFront() *doubleNode {
	removedNode := ll.head
	ll.head = ll.head.next

	return removedNode
}

func (ll *doublyLinkedList) iterateReverse(writer io.Writer) {
	currentNode := ll.tail

	for currentNode != nil {
		fmt.Fprintf(writer, "%v\n", currentNode.data)
		currentNode = currentNode.prev
	}
}

// queue uses a doubly linked list as it's datastore.
type queue struct {
	data *doublyLinkedList
}

func NewLinkedListQueue() queue {
	return queue{data: NewDoublyLinkedList()}
}

func (q queue) enqueue(element any) {
	q.data.insertAtEnd(element)
}

func (q queue) dequeue() any {
	if q.data.head == nil {
		return nil
	}
	removedNode := q.data.removeFromFront()

	return removedNode.data
}

func (q queue) read() any {
	if q.data.head == nil {
		return nil
	}

	return q.data.head.data
}
