package main

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {

	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type element[T any] struct {
	next *element[T]
	val  T
}

type List[T any] struct {
	head, tail *element[T]
}

func (list *List[T]) push(v T) {
	if list.head == nil {
		list.head = &element[T]{val: v}
		list.tail = list.head
	} else {
		list.tail.next = &element[T]{val: v}
		list.tail = list.tail.next
	}
}

func (list *List[T]) getAll() []T {
	var values []T
	for node := list.head; node != nil; node = node.next {
		values = append(values, node.val)
	}
	return values
}

func buildList() {
	var list = List[int]{}
	list.push(20)
	list.push(30)
	list.push(10)
	fmt.Println("list:", list.getAll())
}
