package linked_list

type node[T any] struct {
	data T
	prev *node[T]
	next *node[T]
}

func newNode[T any](data T, prev *node[T], next *node[T]) *node[T] {
	return &node[T]{
		data: data,
		prev: prev,
		next: next,
	}
}

func (n *node[T]) GetData() T {
	return n.data
}

func (n *node[T]) GetPrev() *node[T] {
	return n.prev
}

func (n *node[T]) GetNext() *node[T] {
	return n.next
}

func (n *node[T]) SetPrev(node *node[T]) {
	n.prev = node
}

func (n *node[T]) SetNext(node *node[T]) {
	n.next = node
}
