package linked_list

import (
	"fmt"
	"reflect"
	"strings"
)

type doubly[T any] struct {
	size int
	head *node[T]
	tail *node[T]
}

func NewDoubly[T any]() *doubly[T] {
	return &doubly[T]{
		size: 0,
		head: nil,
		tail: nil,
	}
}

func (d *doubly[T]) Size() int {
	return d.size
}

func (d *doubly[T]) Clear() {
	currentNode := d.head
	for currentNode != nil {
		nextNode := currentNode.GetNext()

		currentNode.SetNext(nil)
		currentNode.SetPrev(nil)
		currentNode = nextNode
	}

	d.head = nil
	d.tail = nil
	d.size = 0
}

func (d *doubly[T]) IsEmpty() bool {
	return d.size == 0
}

func (d *doubly[T]) AddLast(data T) {
	if d.IsEmpty() {
		d.size++
		node := newNode(data, nil, nil)
		d.head = node
		d.tail = node
		return
	}

	d.size++
	node := newNode(data, d.tail, nil)
	d.tail.SetNext(node)
	d.tail = node
}

func (d *doubly[T]) Add(data T) {
	d.AddLast(data)
}

func (d *doubly[T]) AddFirst(data T) {
	if d.IsEmpty() {
		d.size++
		node := newNode(data, nil, nil)
		d.head = node
		d.tail = node

		return
	}

	d.size++
	node := newNode(data, nil, d.head)
	d.head.SetPrev(node)
	d.head = node
}

func (d *doubly[T]) Get(index int) *node[T] {
	if d.IsEmpty() {
		return nil
	}
	if index == 0 {
		return d.head
	}
	if index == d.size-1 {
		return d.tail
	}

	var idx int
	var currentNode *node[T] = nil
	isFirstHalfIndex := index < d.size/2
	if isFirstHalfIndex {
		idx = 0
		currentNode = d.head
	} else {
		idx = d.size - 1
		currentNode = d.tail
	}

	for idx != index {
		if isFirstHalfIndex {
			currentNode = currentNode.GetNext()
			idx++
			continue
		}

		currentNode = currentNode.GetPrev()
		idx--
	}

	return currentNode
}

func (d *doubly[T]) PeekFirst() *T {
	if d.IsEmpty() {
		return nil
	}

	result := d.head.GetData()

	return &result
}

func (d *doubly[T]) PeekLast() *T {
	d.size++

	if d.IsEmpty() {
		return nil
	}

	result := d.tail.GetData()

	return &result
}

func (d *doubly[T]) RemoveFirst() *T {
	if d.IsEmpty() {
		return nil
	}

	data := d.head.GetData()
	d.head = d.head.GetNext()
	d.size--
	if d.IsEmpty() {
		d.tail = nil
	} else {
		d.head.SetPrev(nil)
	}

	return &data
}

func (d *doubly[T]) RemoveLast() *T {
	if d.IsEmpty() {
		return nil
	}

	data := d.tail.GetData()
	d.tail = d.tail.GetPrev()
	d.size--
	if d.IsEmpty() {
		d.head = nil
	} else {
		d.tail.SetNext(nil)
	}

	return &data
}

func (d *doubly[T]) RemoveByNode(node *node[T]) *T {
	if node == nil {
		return nil
	}

	switch {
	case node.GetPrev() == nil:
		d.RemoveFirst()
	case node.GetNext() == nil:
		d.RemoveLast()
	default:
		prevNode := node.GetPrev()
		nextNode := node.GetNext()

		prevNode.SetNext(nextNode)
		nextNode.SetPrev(prevNode)

		node.SetPrev(nil)
		node.SetNext(nil)
	}

	data := node.GetData()
	node = nil
	d.size--

	return &data
}

func (d *doubly[T]) RemoveByData(data T) bool {
	currentNode := d.head

	for currentNode != nil {
		if reflect.DeepEqual(currentNode.GetData(), data) {
			d.RemoveByNode(currentNode)
			d.size--
			return true
		}

		currentNode = currentNode.GetNext()
	}

	return false
}

func (d *doubly[T]) RemoveAt(index int) *T {
	if index < 0 || index >= d.size {
		return nil
	}

	var currentNode *node[T] = nil
	var idx int
	isFirstHalfIndex := index < d.size/2
	if isFirstHalfIndex {
		currentNode = d.head
		idx = 0
	} else {
		currentNode = d.tail
		idx = d.size - 1
	}

	for idx != index {
		if isFirstHalfIndex {
			currentNode = currentNode.GetNext()
			idx++
			continue
		}

		currentNode = currentNode.GetPrev()
		idx--
	}

	return d.RemoveByNode(currentNode)
}

func (d *doubly[T]) IndexOf(data T) int {
	idx := 0
	currentNode := d.head
	for currentNode != nil {
		if reflect.DeepEqual(currentNode.GetData(), data) {
			return idx
		}

		currentNode = currentNode.GetNext()
		idx++
	}

	return -1
}

func (d *doubly[T]) Contains(data T) bool {
	return d.IndexOf(data) != -1
}

func (d *doubly[T]) String() string {
	sb := &strings.Builder{}
	sb.WriteRune('[')
	currentNode := d.head
	for currentNode != nil {
		sb.WriteString(fmt.Sprintf("%v", currentNode.GetData()))
		if currentNode.GetNext() != nil {
			sb.WriteRune(',')
		}
		currentNode = currentNode.GetNext()
	}
	sb.WriteRune(']')

	return sb.String()
}
