package algo

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

// queue uses a doubly linked list as it's datastore.
type queue struct {
	data *doublyLinkedList
}

func newQueue() queue {
	return queue{data: NewDoublyLinkedList()}
}

func (q queue) enqueue(element any) {
	q.data.insertAtEnd(element)
}

func (q queue) dequeue(element any) any {
	removedNode := q.data.removeFromFront()
	return removedNode.data
}

func (q queue) read() any {
	if q.data.head == nil {
		return nil
	}

	return q.data.head.data
}
