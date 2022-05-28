package structures

type LinkedList[T any] struct {
	Head *Node[T]
}

type Node[T any] struct {
	Data T
	Next *Node[T]
	Prev *Node[T]
}
