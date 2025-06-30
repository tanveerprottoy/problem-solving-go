package linkedlist

type SinglyNode[T any] struct {
	Val  T
	Next *SinglyNode[T]
}

type SinglyLinkedList struct {
}
