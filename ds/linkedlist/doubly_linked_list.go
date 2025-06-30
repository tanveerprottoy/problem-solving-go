package linkedlist

type DoublyNode[T any] struct {
	Val      T
	Previous *DoublyNode[T]
	Next     *DoublyNode[T]
}

type DoublyLinkedList struct {
}
