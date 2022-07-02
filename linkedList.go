package algo

type node struct {
	data any
	next *node
}

func NewNode(value any) *node {
	return &node{
		data: value,
	}
}

type linkedList struct {
	head *node
}

func NewLinkedList(value any) *linkedList {
	return &linkedList{
		head: NewNode(value),
	}
}

// read returns node.data of the linkedList's given index.
func (ll *linkedList) read(index int) any {
	currentNode := ll.head

	for i := 0; i < index; i++ {
		if currentNode.next == nil {
			return nil
		}
		currentNode = currentNode.next
	}

	return currentNode.data
}

// indexOf search's linkedList for the given value and returns the index position.
func (ll *linkedList) indexOf(value any) int {
	currentNode := ll.head

	for i := 0; ; i++ {
		if currentNode.data == value {
			return i
		}
		if currentNode.next == nil {
			return -1
		}
		currentNode = currentNode.next
	}
}

// insertion will insert a new node at the given index.
// It returns false if the given index exceeds the length of linkedList.
func (ll *linkedList) insertAtIndex(index int, value any) bool {
	newNode := NewNode(value)

	if index == 0 {
		newNode.next = ll.head
		ll.head = newNode

		return true
	}

	currentNode := ll.head
	for i := 0; i < index-1; i++ { // Access the node immediately before the index position of the new node.
		if currentNode.next == nil {
			return false
		}
		currentNode = currentNode.next
	}

	newNode.next = currentNode.next
	currentNode.next = newNode

	return true
}

// deleteAtIndex delete a linkedList's node at the given index.
// It returns false if the given index exceeds the length of linkedList
func (ll *linkedList) deleteAtIndex(index int) bool {
	if index == 0 { // Simply set the head as the next node if deleting first node.
		ll.head = ll.head.next

		return true
	}

	currentNode := ll.head
	for i := 0; i < index-1; i++ { // Access the node immediately before the index position of the node to be deleted.
		if currentNode.next == nil || currentNode.next.next == nil {
			return false
		}
		currentNode = currentNode.next
	}

	nodeAfterDeletedNode := currentNode.next.next
	currentNode.next = nodeAfterDeletedNode

	return true
}